package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	urlM "net/url"
	"os"
	"strconv"
	"time"
	pb "weather/proto"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

type weatherMain struct {
	Temp float64 `json:"temp"`
}
type weatherData struct {
	Name string      `json:"name"`
	Main weatherMain `json:"main"`
}

type forecastWeatherData struct {
	List []struct {
		Main weatherMain `json:"main"`
	} `json:"list"`
	City struct {
		Name string `json:"name"`
	} `json:"city"`
}

type wetherResp struct {
	City        string  `json:"city"`
	Unit        string  `json:"unit"`
	Temperature float64 `json:"temperature"`
}

// mb global
var apiKey = os.Getenv("APIKEY")
var port = os.Getenv("LISTEN_PORT")
var urlP = os.Getenv("URL")

type Handler struct {
	client *redis.Client
	Port   string
}

func NewHandler() *Handler {
	rClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &Handler{rClient, port}
}

func (h *Handler) CurrentWeather(w http.ResponseWriter, req *http.Request) {
	res := checkAuth(context.Background(), req)
	if !res {
		fmt.Print("invalid user name \n")
		http.Error(w, "", http.StatusForbidden)
		return
	}
	city := req.URL.Query().Get("city")
	url := os.Getenv("URL")
	url = fmt.Sprintf("%sweather?q=%s&appid=%s&units=metric", url, city, apiKey)
	val := h.client.Get(req.Context(), city)
	var resp wetherResp
	if val.Err() == redis.Nil {
		data, err := current(url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Print("value not found \n")
		h.client.Set(req.Context(), data.Name, data.Main.Temp, 20*time.Second)
		resp = wetherResp{data.Name, "celsius", data.Main.Temp}
	} else {
		temp, err := strconv.ParseFloat(val.Val(), 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resp = wetherResp{city, "celsius", temp}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World from %s", port)
}

func (h *Handler) ForecastWeather(w http.ResponseWriter, req *http.Request) {
	res := checkAuth(context.Background(), req)
	if !res {
		fmt.Print("invalid user name \n")
		http.Error(w, "", http.StatusForbidden)
		return
	}
	city := req.URL.Query().Get("city")
	dt := req.URL.Query().Get("dt")
	i, err := strconv.ParseInt(dt, 10, 64)
	if err != nil {
		panic(err)
	}
	rKey := fmt.Sprintf("%s,%s", city, dt)
	val := h.client.Get(req.Context(), rKey)
	var resp wetherResp
	if val.Err() == redis.Nil {
		urlN := fmt.Sprintf("%sforecast/?q=%s&appid=%s&units=metric&cnt=%s", urlP, city, apiKey, dt)
		data, err := forecast(urlN, i)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		h.client.Set(req.Context(), rKey, data.List[i-1].Main.Temp, 20*time.Second)
		resp = wetherResp{data.City.Name, "celsius", data.List[i-1].Main.Temp}
	} else {
		temp, err := strconv.ParseFloat(val.Val(), 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resp = wetherResp{city, "celsius", temp}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) SaveWeatherData(w http.ResponseWriter, req *http.Request) {
	if req.Method == "PUT" {
		var weather wetherResp
		err := json.NewDecoder(req.Body).Decode(&weather)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		temp := strconv.FormatFloat(weather.Temperature, 'f', 2, 64)
		h.client.Set(req.Context(), weather.City, temp, 20*time.Second)
		w.WriteHeader(200)
		fmt.Fprintf(w, "You set data: %+v", weather)
		return
	}
	http.Error(w, "", http.StatusInternalServerError)
}

func current(url string) (weatherData, error) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	var d weatherData

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}
	return d, nil
}

func forecast(url string, ts int64) (forecastWeatherData, error) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	var d forecastWeatherData

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return forecastWeatherData{}, err
	}

	return d, nil
}

func checkAuth(ctx context.Context, req *http.Request) bool {
	name := req.Header.Get("Own-Auth-Username")
	decodedName, err := urlM.QueryUnescape(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(decodedName)
	conn, err := grpc.Dial("auth:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewAuthClient(conn)
	res, err := client.IsAuth(ctx, &pb.UserInfo{Login: decodedName})
	if err != nil {
		panic(err)
	}
	fmt.Println("res: ", res.Reply)
	return res.Reply
}
