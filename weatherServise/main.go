package main

import (
	"fmt"
	"net/http"

	"weather/handler"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	middlewarestd "github.com/slok/go-http-metrics/middleware/std"
)

func main() {

	mdlw := middleware.New(middleware.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})
	h := handler.NewHandler()
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", h.HelloServer)

	hFuncCurent := http.HandlerFunc(h.CurrentWeather)
	hCurent := middlewarestd.Handler("my_mytrics", mdlw, hFuncCurent)
	http.Handle("/v1/current/", hCurent)

	hFuncForecast := http.HandlerFunc(h.ForecastWeather)
	hForecast := middlewarestd.Handler("my_mytrics", mdlw, hFuncForecast)
	http.Handle("/v1/forecast/", hForecast)

	hFuncSave := http.HandlerFunc(h.SaveWeatherData)
	hSave := middlewarestd.Handler("my_mytrics", mdlw, hFuncSave)
	http.Handle("/v1/save/", hSave)
	fmt.Printf("%s \n", h.Port)
	fmt.Println("Server is listening at localhost:" + h.Port)
	err := http.ListenAndServe("0.0.0.0:"+h.Port, nil)
	if err != nil {
		fmt.Printf("Server is dead %s", err)
		return
	}
}
