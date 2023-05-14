FROM golang:latest

ENV  URL="https://api.openweathermap.org/data/2.5/"
ENV  APIKEY="4a79f62e388436d841b3ffc1992e6b6d"

WORKDIR /app 

COPY ./handler ./handler
COPY ./proto ./proto
COPY go.* ./
RUN go mod download
COPY *.go ./
RUN go build -o ./main
# EXPOSE 8181

# ENTRYPOINT ["./main"]