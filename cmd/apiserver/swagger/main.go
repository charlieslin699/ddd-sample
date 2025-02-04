package main

import (
	_ "ddd-sample/cmd/apiserver/swagger/docs"
	"ddd-sample/pkg/httpserver"
	"ddd-sample/userinterface/api/swagger"
	"fmt"
	"net/http"
	"time"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
func main() {
	server := newServer()

	// windows測試
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func newServer() *http.Server {
	config := getConfig()
	server := httpserver.NewHTTPServer()

	swagger.InitRouter(server)

	return &http.Server{
		Addr:           fmt.Sprintf("%s:%d", config.host, config.port),
		Handler:        server.Engine(),
		ReadTimeout:    config.readTimeout,
		WriteTimeout:   config.writeTimeout,
		MaxHeaderBytes: config.maxHeaderBytes,
	}
}

type apiConfig struct {
	host           string
	port           int
	readTimeout    time.Duration
	writeTimeout   time.Duration
	maxHeaderBytes int
}

// TODO: 測試用
func getConfig() apiConfig {
	return apiConfig{
		port: 8080,
	}
}
