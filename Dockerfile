FROM golang:1.22.1
LABEL authors="techoc"
ENV GOPROXY https://goproxy.cn,direct

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .

RUN go build -o fanqie-novel-api

CMD ["./fanqie-novel-api"]
EXPOSE 8000