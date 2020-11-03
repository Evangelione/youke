package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
	"yk/internal/app/enterprise"
	"yk/internal/app/enterprise/router"
	"yk/internal/pkg/infra"
	"yk/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// LoadConfig config
	config := infra.LoadConfig("configs/enterprise.yaml")
	infra.ConnMysql(config.Mysql)
	infra.ConnRedis(config.Redis)

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
			enterprise.Logger().Fatal("listen:" + err.Error())
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
	enterprise.Logger().Info("Shutdown Server ...")

	// Get context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown server
	if err := srv.Shutdown(ctx); err != nil {
		enterprise.Logger().Fatal("Server Shutdown:" + err.Error())
	}
	enterprise.Logger().Info("Server exiting")
}
