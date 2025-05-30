run:
	go run cmd/main.go

build:
	go get ./cmd/
	go build  -o main cmd/main.go

migrate:
	go run app/exe/migrate/*.go
	go run cmd/migrate/main.go