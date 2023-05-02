all: clean docker

docker:
	docker build -t auth-server_img .
	docker run -d -p 0.0.0.0:50051:50051 --name auth-server auth-server_img

clean:
	docker container stop auth-server
	docker container rm auth-server
	docker image rm auth-server_img
