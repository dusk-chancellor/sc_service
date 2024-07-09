# Сервис сбора статистики

  - [Как запускать](#как-запускать)
  - [Структура проекта](#структура-проекта)
  - [Примеры запросов на эндпоинты](#примеры-запросов-на-эндпоинты)
  - [Связь](#связь)

### Как запускать

 - Открываем конфиг файл в `~/configs/config.yaml` и
 подставляем нужные значения
 - Запускаем команду `go run ./cmd/sc/main.go`
 - Тестируем на эндпоинтах:
   - `POST /orderbook` - реализация `SaveOrderBook(exchange_name, pair string, orderBook []*DepthOrder) error`
   - `POST /order` - реализация `SaveOrder(client *Client, order *HistoryOrder) error`
   - `GET /orderbook` - реализация `GetOrderBook(exchange_name, pair string) ([]*DepthOrder, error)`
   - `GET /order` - реализация `GetOrderHistory(client *Client) ([]*HistoryOrder, error)`


### Структура проекта
 1. Приложение запускается `~/cmd/sc/main.go`
 2. Все компоненты собираются в `~/internal/app/app.go`
 3. Пользователь обращается к хендлерам в `~/internal/delivery/http/handler.go`
 4. Бизнес логика в `~/internal/service/service.go`
 5. Взаимодействие с базой данных в `~/internal/repository/postgres/`


### Примеры запросов на эндпоинты

  #### POST /orderbook
  Query параметры:
  - `exchange_name`=`test0`
  - `pair`=`test1`

  Body:
  ```json
    [
	    {
		    "price": 100.0,
		    "base_qty": 10.0
	    },
	    {
		    "price": 101.5,
		    "base_qty": 15.0
	    }
    ]
  ```

  #### POST /order
  Body:
  ```json
    {
	    "client_name":"test0",
	    "exchange_name":"test1",
	    "label":"test2",
	    "pair":"test3",
	    "side":"",
	    "type":"",
	    "base_qty":0,
	    "price":0,
	    "algorithm_name_placed":"",
	    "lowest_sell_prc":0,
	    "highest_buy_prc":0,
	    "commission_quote_qty":0
    }
  ```

  #### GET /orderbook
  Query параметры:
  - `exchange_name`=`test0`
  - `pair`=`test1`

  #### GET /order
  Body:
  ```json
    {
	    "client_name":"test0",
	    "exchange_name":"test1",
	    "label":"test2",
	    "pair":"test3"
    }
  ```


### Связь
  - Telegram: [@duskchancellor](https://t.me/duskchancellor);
  - Email: duskchancellor@gmail.com;
