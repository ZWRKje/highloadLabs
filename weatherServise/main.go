package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type weatherData struct {
	Name string `json:"name`
	Main struct {
		Temp float64 `json:"temp`
	} `json:main`
}

type wetherResp struct {
	City        string  `json:"city"`
	Unit        string  `json:"unit"`
	Temperature float64 `json:"temperature"`
}

func main() {
	port := os.Getenv("LISTEN_PORT")
	http.HandleFunc("/v1/current/", func(w http.ResponseWriter, req *http.Request) {
		city := req.URL.Query().Get("city")
		url := os.Getenv("URL")
		url = fmt.Sprintf("%sweather?q=%s&appid=4a79f62e388436d841b3ffc1992e6b6d&units=metric", url, city)
		data, err := current(url)
		resp := wetherResp{data.Name, "celsius", data.Main.Temp}
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
	})

	fmt.Printf("%s \n", port)
	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:"+port, nil)
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
