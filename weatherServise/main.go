package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	urlM "net/url"
	"os"
	"strconv"

	pb "weather/proto"

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

func main() {

	http.HandleFunc("/", helloServer)
	http.HandleFunc("/v1/current/", currentWeather)
	http.HandleFunc("/v1/forecast/", forecastWeather)
	fmt.Printf("%s \n", port)
	fmt.Println("Server is listening at localhost:" + port)
	err := http.ListenAndServe("127.0.0.1:"+port, nil)
	if err != nil {
		fmt.Printf("Server is dead %s", err)
		return
	}
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

func helloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func currentWeather(w http.ResponseWriter, req *http.Request) {
	res := checkAuth(context.Background(), req)
	if !res {
		fmt.Print("invalid user name \n")
		http.Error(w, "", http.StatusForbidden)
		return
	}
	city := req.URL.Query().Get("city")
	url := os.Getenv("URL")
	url = fmt.Sprintf("%sweather?q=%s&appid=%s&units=metric", url, city, apiKey)
	data, err := current(url)
	resp := wetherResp{data.Name, "celsius", data.Main.Temp}
	if err != nil {
		w.WriteHeader(403)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(resp)
}

func forecastWeather(w http.ResponseWriter, req *http.Request) {
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
	urlN := fmt.Sprintf("%sforecast/?q=%s&appid=%s&units=metric&cnt=%s", urlP, city, apiKey, dt)
	data, err := forecast(urlN, i)
	resp := wetherResp{data.City.Name, "celsius", data.List[i-1].Main.Temp}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(resp)
}

func checkAuth(ctx context.Context, req *http.Request) bool {
	name := req.Header.Get("Own-Auth-Username")
	decodedName, err := urlM.QueryUnescape(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(decodedName)
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
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
