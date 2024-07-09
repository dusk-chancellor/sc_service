package main

// Микросервисное приложение на golang:
// Сервис сбора статистики

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dusk-chancellor/sc_service/configs"
	"github.com/dusk-chancellor/sc_service/internal/app"
)


func main() {
	cfg := configs.ReadConfig() // читаем конфиг

	ctx := context.TODO() // TODO: добавить контекст

	app := app.NewApp(ctx, cfg) // инициализация приложения

	log.Printf("starting server on %s", cfg.Server.Host + ":" + cfg.Server.Port)

	go func() { // запустим в горутине
		app.Run()
	}()
	// когда приложение остановится, то выполнится простенький GracefulShutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	app.Shutdown(ctx)
}
