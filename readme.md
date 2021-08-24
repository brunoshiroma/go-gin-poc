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

### Deploy
Utilizando o [Heroku](https://www.heroku.com/), é possivel hospedar a aplicação e o banco de forma gratuita.  
Tutorial utilizado para fazer o deploy via docker com imagem sendo buildada localmente e rodando no heroku, a partir da imagem gerada [Heroku docker](https://devcenter.heroku.com/articles/container-registry-and-runtime)  
Após configurar o heroku cli, fazer o login, e também o login no container 
registry como explicado no tutorial do heroku docker, para publicar basta rodar o comando:
```bash
make deploy-heroku
```

##### Comandos para fazer o login do registry do heroku, e o deploy manuamente
```bash
#loga no heroku, vai abrir uma pagina web, ai basta logar na sua conta, e ele vai autorizar o cli
heroku login
#loga no repository de containers
heroku container:login
#depois daqui o comando make deploy-heroku já faz
#faz o build usando o heroku.yml + Dockerfile
heroku container:push web
#faz o release, publicando a imagem previamente buildada e pushada no repo
heroku container:release web
```


### Bibliotecas utilizadas
 * [Gin](https://github.com/gin-gonic/gin) Framework para API Rest
 * [godotenv](github.com/joho/godotenv) biblioteca para uso do 'padrão' dotenv
 * [gorm](https://gorm.io/index.html) Biblioteca de ORM
 * [swag](https://github.com/swaggo/gin-swagger) Swagger para o golang