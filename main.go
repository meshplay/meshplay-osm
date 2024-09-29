package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kumarabd/gokit/logger"
	"github.com/khulnasoft/meshplay-adapter-template/adaptor"
	"github.com/khulnasoft/meshplay-adapter-template/api/grpc"
	"github.com/khulnasoft/meshplay-adapter-template/internal/config"
	"github.com/khulnasoft/meshplay-adapter-template/internal/tracing"
)

var (
	configProvider = "local"
)

// main is the entrypoint of the adaptor
func main() {
	// Initialize application specific configs and dependencies
	// App and request config
	cfg, err := config.New(configProvider)
	if err != nil {
		fmt.Println("Config Init Failed", err.Error())
		os.Exit(1)
	}
	service := &grpc.Service{}
	_ = cfg.Server(&service)

	// Initialize Logger instance
	log, err := logger.New(service.Name)
	if err != nil {
		fmt.Println("Logger Init Failed", err.Error())
		os.Exit(1)
	}

	// Initialize Tracing instance
	traceProvider, err := tracing.New(service.Name, service.TraceURL)
	if err != nil {
		fmt.Println("Tracing Init Failed", err.Error())
		os.Exit(1)
	}

	// Initialize Handler intance
	handler := adaptor.New(cfg, log)
	service.Handler = handler
	service.StartedAt = time.Now()

	// Server Initialization
	log.Info("Adaptor Started")
	err = grpc.Start(service, traceProvider)
	if err != nil {
		log.Err("Adaptor crashed!!", err.Error())
		os.Exit(1)
	}
}
