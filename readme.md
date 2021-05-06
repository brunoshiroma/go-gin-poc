# Projeto de teste da biblioca Gin para api rest do Golang
Projeto simples, apenas para testar e validar a biblioteca [Gin](https://github.com/gin-gonic/gin)

## Variaveis de ambiente
Utilizado o godotenv para o desenvolvimento, o repositorio tem o [.env_example](.env_example) como exemplo das variaveis utilizadas

## Documentação
Projeto inclui [swagger](http://localhost:60080/swagger/index.html) (o link usa a porta 60080 != da 8080 que é o padrão do Gin, então caso você mude a porta, precisa adequar no link)    

**Criar os arquivos necessario para o swagger:**
```bash
go get -u github.com/swaggo/swag/cmd/swag
swag init -g cmd/go-gin-poc/main.go #precisa passar o argumento -g pelo arquivo main.go não estar no diretorio raiz
```
## Banco de dados
Para esse projeto foi utilizado o Postgres, o ***docker-compose*** está disponivel no projeto, e foi utilizado para realizar todo o desenvolvimento e testes.    

### Bibliotecas utilizadas
 * [Gin](https://github.com/gin-gonic/gin) Framework para API Rest
 * [godotenv](github.com/joho/godotenv) biblioteca para uso do 'padrão' dotenv
 * [gorm](https://gorm.io/index.html) Biblioteca de ORM
 * [swag](https://github.com/swaggo/gin-swagger) Swagger para o golang