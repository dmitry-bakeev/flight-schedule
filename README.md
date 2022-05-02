# Flight schedule

This is the solution of a test task for the position of Go developer to Medesk

## API

### Method `POST` to `api/flight` - add data about flights

Example request:

```json
// time format is RFC3339
{
    "objects": [
        {
            "number_flight": 125,
            "from_city": "Moscow",
            "time_from_city": "2022-04-29T14:00:00Z",
            "to_city": "Rostov",
            "time_to_city": "2022-04-29T16:00:00Z"
        },
        {
            "number_flight": 120,
            "from_city": "Rostov",
            "time_from_city": "2022-04-30T12:00:00Z",
            "to_city": "Moscow",
            "time_to_city": "2022-04-30T14:00:00Z"
        },
    ]
}
```

Example response:

```json
{
    "status": "ok",
    "message": {
        "added": "done"
    }
}
```

### Method `GET` to `api/flight` - show filghts

Takes the following GET params:

- `filter_from_city` - allows to filter by specified city for field `from_city`
- `filter_to_city` - allows to filter by specified city for field `to_city`
- `order_desc` - using for reverse sort for next params
- `order_number_flight` - allows to order by field `number_flight`
- `order_from_city` - allows to order by field `from_city`
- `order_to_city` - allows to order by field `to_city`
- `order_time_from_city` - allows to order by field `time_from_city`
- `order_time_to_city` - allows to order by field `time_from_city`

Exxample request:

```bash
curl http://localhost:8000/api/flight?order_to_city=1&order_desc=1
```

Example response:

```json
// for GET params ?order_number_flight=1&order_desc=1
{
    "status": "ok",
    "message": {
        "objects": [
            {
            "number_flight": 125,
            "from_city": "Moscow",
            "time_from_city": "2022-04-29T14:00:00Z",
            "to_city": "Rostov",
            "time_to_city": "2022-04-29T16:00:00Z"
        },
        {
            "number_flight": 120,
            "from_city": "Rostov",
            "time_from_city": "2022-04-30T12:00:00Z",
            "to_city": "Moscow",
            "time_to_city": "2022-04-30T14:00:00Z"
        },
        ]
    }
}
```

## How to run

### Requirements

- Go 1.18 or above
- [Migrate](https://github.com/golang-migrate/migrate)
- Docker 20.10.14 or above

1 Create `.env` file in this project dir

```bash
# Example
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=qwerty
POSTGRES_DATABASE=postgres
POSTGRES_SSL_MODE=disable

RUN_PORT=8000
```

2 Run postgres db in docker

```bash
# Example
docker run -dp 5432:5432 -e POSTGRES_PASSWORD=qwerty postgres:13.3-alpine
```

3 Run migrations

```bash
# Example
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' -verbose up
```

4 Run app

```bash
# Example
go run cmd/api/main.go 
```
