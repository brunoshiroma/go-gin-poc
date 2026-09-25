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
	@TEST_PKGS="$$(go list -f '{{if or (gt (len .TestGoFiles) 0) (gt (len .XTestGoFiles) 0)}}{{.ImportPath}}{{end}}' ./... | grep .)"; \
	NO_TEST_PKGS="$$(go list -f '{{if and (eq (len .TestGoFiles) 0) (eq (len .XTestGoFiles) 0)}}{{.ImportPath}}{{end}}' ./... | grep .)"; \
	if [ -n "$$NO_TEST_PKGS" ]; then go test $$NO_TEST_PKGS; fi; \
	if [ -n "$$TEST_PKGS" ]; then go test -coverpkg=./... -coverprofile=coverage.out $$TEST_PKGS; fi

test-with-report: test
	go tool cover -html=coverage.out

coverage: test