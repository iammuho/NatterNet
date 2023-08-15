package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/iammuho/natternet/cmd/app/config"
	"github.com/iammuho/natternet/pkg/http"
	"github.com/iammuho/natternet/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	// Add the logger
	l, err := logger.NewLogger(
		logger.WithLoggerLevel(config.Config.Logger.Level),
		logger.WithLoggerName(config.Config.Logger.Name),
	)

	if err != nil {
		panic(err)
	}

	// Create a new http server
	l.Info("Creating HTTP Server")
	httpServer := http.NewHTTPServer(
		http.WithHTTPServerAddress(config.Config.HTTPServer.ListenAddress),
		http.WithHTTPServerPort(config.Config.HTTPServer.ListenPort),
		http.WithHTTPServerTLSEnabled(config.Config.HTTPServer.TLSEnabled),
	)

	// Start the http server
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		l.Info("Starting HTTP Server")
		if err := httpServer.App.Listen(fmt.Sprintf("%s:%d", httpServer.Options.Address, httpServer.Options.Port)); err != nil {
			l.Panic("HTTP Server Listen failed: %v", zap.Error(err))
		}
	}()

	// Graceful Shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	l.Info("Shutting down the server...")

	if err := httpServer.App.Shutdown(); err != nil {
		l.Panic("HTTP Server Shutdown failed: %v", zap.Error(err))
	}

	wg.Wait()

	l.Info("Server exited")
}
