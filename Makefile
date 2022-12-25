#baixa e roda o swag para gerar a pasta docs, usado posteriormente pelo swagger
swag:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init -g cmd/go-gin-poc/main.go

#baixa e atualiza as dependencias do go mod
dep:
	go mod tidy

#faz o build (do jeito que está o binario que é gerado precisa ser renomeado com o .exe no windows...)
build: swag dep
	go build -o go-gin-poc cmd/go-gin-poc/main.go

test:
	go test -cover -coverprofile=coverage.out ./...

test-with-report: test
	go tool cover -html=coverage.out
