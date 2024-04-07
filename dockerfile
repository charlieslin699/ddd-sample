# build
FROM golang:1.22 as build
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

ARG APP_CMD

WORKDIR /go/src/ddd-sample

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN mkdir -p /go/bin/ddd-sample
RUN go build -v -o /go/bin/ddd-sample/app ./cmd/apiserver/${APP_CMD}/main.go

# run
FROM golang:1.22.1-alpine3.19

# 由於程式內取設定檔邏輯的原因，config需在/go/src底下
ADD config /go/src/ddd-sample/config
COPY --from=build /go/bin/ddd-sample/app /go/src/ddd-sample/app
