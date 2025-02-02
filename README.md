# dowhile.uz back-end

## How to create migrations

```
migrate create -dir migrations -ext sql create_<table-name>_table
migrate create -dir migrations -ext sql add_<column-name>_for_<table-name>_table
```

## How to run migrations:

Get database connection:

```
go run cmd/print-connection-urls/main.go
```

Migrate:

```
migrate -database "<database-url>" -source "file://migrations" up
```
