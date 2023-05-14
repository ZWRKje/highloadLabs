package main

import (
	"fmt"
	"net/http"

	"weather/handler"
)

func main() {
	h := handler.NewHandler()
	http.HandleFunc("/", h.HelloServer)
	http.HandleFunc("/v1/current/", h.CurrentWeather)
	http.HandleFunc("/v1/forecast/", h.ForecastWeather)
	http.HandleFunc("/v1/save/", h.SaveWeatherData)
	fmt.Printf("%s \n", h.Port)
	fmt.Println("Server is listening at localhost:" + h.Port)
	err := http.ListenAndServe("0.0.0.0:"+h.Port, nil)
	if err != nil {
		fmt.Printf("Server is dead %s", err)
		return
	}
}
