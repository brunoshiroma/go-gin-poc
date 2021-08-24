#baixa e roda o swag para gerar a pasta docs, usado posteriormente pelo swagger
swag:
	go get github.com/swaggo/swag/cmd/swag
	swag init -g cmd/go-gin-poc/main.go

#baixa e atualiza as dependencias do go mod
dep:
	go mod tidy

#faz o build (do jeito que está o binario que é gerado precisa ser renomeado com o .exe no windows...)
build: swag dep
	go build -o go-gin-poc cmd/go-gin-poc/main.go

#faz o build da imagem docker e faz o push para o repository docker do heroku
build-heroku:
	heroku container:push web -a go-gin-poc

#faz o release/deploy da versão mais atual da imagem
deploy-heroku: build-heroku
	heroku container:release web -a go-gin-poc