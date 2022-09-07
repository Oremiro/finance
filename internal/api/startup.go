package api

import (
	"finance/internal/api/http/v1"
	"finance/internal/api/http/v1/controller"
	"finance/internal/core/tinkoff"
	"finance/internal/infra"
	"finance/pkg/httpserver"
	"finance/pkg/postgres"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(config *Config) {

	// Adapter
	pg, err := postgres.NewContext(config.DbOptions.ConnectionString)
	defer pg.Dispose()

	// Storages
	storage := infra.NewFinanceStorage(pg)

	// Services //TODO mediator pattern
	tinkoffService := tinkoff.NewService(storage)

	// Controllers
	var controllers = []v1.IController{
		controller.NewTinkoffController(tinkoffService),
	}

	// HTTP Server
	handler := gin.New()
	v1.UseRouter(handler, controllers)
	httpServer := httpserver.NewServer(handler)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Fatal(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
