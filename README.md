## Files needed for running the project

- `config/config.yml`
- `.env`

## Running the project

```bash
$ go mod tidy
```

```bash
$ go run main.go
```

## Migrations

### Create migrations:

```bash
migrate create -ext sql -dir ./migrations -seq init
```

### Run up migration

```bash
migrate -database "postgresql://{USERNAME}:{PASSWORD}@{HOST}:{PORT}/{DBNAME}?sslmode=disable" -path ./migrations up
```

### Run down migration

```bash
migrate -database "postgresql://{USERNAME}:{PASSWORD}@{HOST}:{PORT}/{DBNAME}?sslmode=disable" -path ./migrations down
```
