# ILCS To-Do-List Backend Test 
## Pre-requisites
1. Go version 1.18+
2. PostgreSQL 
3. Goose
4. JWT
5. Echo
6. Sqlx

## How to run
1. Copy the .env.example file to .env
2. Open the .env file and fill in the required credentials (e.g., database details).
3. Run the app using:
```
go run cmd/main.go
```
## Run migrations
1. Install Goose:
```
go install github.com/pressly/goose/v3/cmd/goose@latest
```
2. Run the migrations:
```
goose -dir migrations postgres "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" up
```

## Linter
```
golangci-lint run
```
or
```
golangci-lint run path/to/your/file.go
```

## Test
```
go test ./internal/usecase/task -v
```