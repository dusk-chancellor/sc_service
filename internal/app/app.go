package app

// пакет инициализации приложения
// здесь все компоненты собираются воедино для удобного запуска с cmd

import (
	"context"
	"net/http"

	"github.com/dusk-chancellor/sc_service/configs"
	handlers "github.com/dusk-chancellor/sc_service/internal/delivery/http"
	"github.com/dusk-chancellor/sc_service/internal/repository/postgres"
	"github.com/dusk-chancellor/sc_service/internal/service"
)

type App struct { // Структура приложения в общем
	HttpServer http.Server
	// ...
}

func NewApp(ctx context.Context, cfg *configs.Config) *App {
	db, err := postgres.ConnectDB(cfg) // подключение к постгресу
	if err != nil {
		panic(err)
	}

	appService := service.NewService(db, ctx) // инициализация сервисных методов

	appHandlers := handlers.NewHandlers(appService) // инициализация хендлеров

	mux := http.NewServeMux() // REST API хендлеры
	mux.HandleFunc("GET /orderbook", appHandlers.GetOrderBookHandler())
	mux.HandleFunc("GET /order", appHandlers.GetOrderHistoryHandler())
	mux.HandleFunc("POST /orderbook", appHandlers.SaveOrderBookHandler())
	mux.HandleFunc("POST /order", appHandlers.SaveOrderHandler())


	app := &App{
		HttpServer: http.Server{
			Addr:    cfg.Server.Host + ":" + cfg.Server.Port,
			Handler: mux,
		//	ReadTimeout: 101 * time.Millisecond,
		//	WriteTimeout: 201 * time.Millisecond,
		},
	}

	return app
}

func (a *App) Run() { // сразу паникуем при ошибке
	if err := a.HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

func (a *App) Shutdown(ctx context.Context) { // тут тоже без возврата ошибок
	if err := a.HttpServer.Shutdown(ctx); err != nil {
		panic(err)
	}
}
