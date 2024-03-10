FROM golang:1.22.1
LABEL authors="techoc"

RUN go build -o fanqie-novel-api && ./fanqie-novel-api

EXPOSE 8000