package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"yk/internal/app/platform"
	"yk/internal/app/platform/router"
	"yk/internal/pkg/infra"
	"yk/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// LoadConfig config
	config := infra.LoadConfig("configs/platform.yaml")

	// Register router
	engine := gin.Default()
	engine.Use(middleware.Logger())
	router.Register(engine)

	var port string
	flag.StringVar(&port, "p", config.App.Port, "运行端口")
	flag.Parse()

	// Create server
	server := &http.Server{
		Addr:    ":" + port,
		Handler: engine,
	}

	// Start server
	go func() {
		// Service connections
		_, _ = fmt.Fprintf(os.Stdout, "server is running in port %s\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// Stop server
	httpServerStop(server)
}

func httpServerStop(srv *http.Server) {
	platform.Logger().Info("Shutdown Server ...")

	// Get context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown server
	if err := srv.Shutdown(ctx); err != nil {
		platform.Logger().Fatal("Server Shutdown:" + err.Error())
	}
	platform.Logger().Info("Server exiting")
}
