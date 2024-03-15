FROM alpine:edge as builder
LABEL stage=go-builder
WORKDIR /app/
RUN apk add --no-cache bash curl gcc git go musl-dev
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o fanqie-novel-api

FROM alpine:edge
LABEL authors="techoc"
WORKDIR /opt/fanqie/
COPY --from=builder /app ./
RUN apk update && \
    apk upgrade --no-cache && \
    apk add --no-cache bash ca-certificates su-exec tzdata; \
    chmod +x ./fanqie-novel-api&& \
    rm -rf /var/cache/apk/*
ENV PUID=0 PGID=0 UMASK=022
EXPOSE 8000
CMD [ "./fanqie-novel-api" ]