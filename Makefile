build:
	go build -o ./bin/price_checker ./cmd/price_checker

run:
	go run ./cmd/price_checker --config ./config.yml --check-all


compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go

all: build
