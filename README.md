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
