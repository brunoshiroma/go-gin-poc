# imagem utilizada para o build
FROM golang:alpine as build

# atualiza os repos do alpine
RUN apk update
# instala o make para poder rodar o make build
RUN apk add make
# instal o git, necessario pelo golang baixar as dependencias
RUN apk add git

# diretorio de trabalho, o golang tem a convenção de ficar nesse local GOPATH + nome do modulo
WORKDIR /go/github.com/brunoshiroma/go-gin-poc

# copia apenas os diretorios e arquivos necessarios para o build
COPY ./internal internal/
COPY ./cmd cmd/
COPY go.mod .
COPY Makefile .

# roda o make build
RUN make build

# imagem para o nosso "runtime", utilizado o alpine "puro"
FROM alpine as runtime

WORKDIR /app

# copia o binario gerado no estagio de build, para o nosso estagio de runtime
COPY --from=build /go/github.com/brunoshiroma/go-gin-poc/go-gin-poc go-gin-poc

# o que vai ser executado quando rodar o docker run ...
ENTRYPOINT [ "/app/go-gin-poc" ]