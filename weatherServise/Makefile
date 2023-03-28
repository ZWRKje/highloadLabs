all: clean docker

docker:
	docker build -t weather-app_img .
	docker run -d -p 0.0.0.0:8181:8181 --name weather_app weather-app_img

clean:
	docker container stop weather_app
	docker container rm weather_app
	docker image rm weather-app_img
