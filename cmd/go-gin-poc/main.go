package main

import (
	"log"
	"os"

	"github.com/brunoshiroma/go-gin-poc/internal/controller"
	"github.com/brunoshiroma/go-gin-poc/internal/dao"

	"github.com/gin-gonic/gin"

	_ "github.com/joho/godotenv/autoload"

	_ "github.com/brunoshiroma/go-gin-poc/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Go Gin POC API
// @version 1.0
// @description POC utilizando Golang + Gin
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host go-gin-poc.herokuapp.com
// @BasePath /api/v1

func main() {

	var dao dao.Dao = &dao.SimpleDao{}
	err := dao.StartDB()
	if err != nil {
		log.Panic(err)
	}

	if os.Getenv("DO_GORM_AUTOMIGRATE") == "true" {
		err = dao.AutoMigrate()
		if err != nil {
			log.Panic(err)
		}
	}

	// Exemplos do setup do swagger tirado de https://github.com/swaggo/swag/blob/master/example/celler/main.go
	r := gin.Default()

	clientController := controller.NewClientController(dao)

	//agrupa a api com o basepath /api/v1
	v1 := r.Group("/api/v1")
	{
		//agrupa endpoints de client
		client := v1.Group("/client")
		{
			client.GET("", clientController.RetriveAllClient)
			client.POST("", clientController.CreateClient)
			client.PUT("", clientController.UpdateClient)
			client.DELETE("", clientController.DeleteClient)
		}
	}

	//configura o swagger.
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Roda o servidor do Gin, utilizando a porta 8080 padrão ou o definida na ENV PORT
	// E usa o ENV GIN_MODE=release para rodar como "produção"
	r.Run()
}
