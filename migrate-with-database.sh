dburl=$(go run cmd/print-postgres-connection-url/main.go)

migrate -database "$dburl" -source "file://migrations" "$1"
