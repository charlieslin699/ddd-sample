package main

import (
	"ddd-sample/pkg/httpserver"
	"ddd-sample/userinterface/api/auth"
	"fmt"
	"net/http"
	"time"
)

func main() {
	server := newServer()

	// windows測試
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

	//nolint:gocritic // 優雅重啟
	// if err := gracehttp.Serve(server); err != nil {
	// 	panic(err)
	// }
}

func newServer() *http.Server {
	config := getConfig()
	server := httpserver.NewHTTPServer()

	auth.InitRouter(server)

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
		port: 7000,
	}
}
