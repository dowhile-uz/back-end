# dowhile.uz back-end

## Where to get GitHub App credentials?

- Go to https://github.com/settings/apps
- Create New GitHub App
- In callback url field insert `http://127.0.0.1:8000/v1/github-auth/complete`
- Select `Expire user authorization tokens` and `Request user authorization (OAuth) during installation` checkboxes
- Copy-paste app-id, client-id, and client-secret (you need to generate this first) into `configs/override.yaml` (as base config you can rely on `configs/base.yaml` and override only fields you want, the rest will be defaulted)
- Generate private key (the section Private keys) and place it inside `configs/github-app.pem`

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
