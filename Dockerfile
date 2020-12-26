FROM golang:1.15.6 as builder

WORKDIR /app

COPY ./ /app

RUN go get github.com/go-delve/delve/cmd/dlv \
    && go get github.com/ramya-rao-a/go-outline \
    && go get github.com/labstack/echo

