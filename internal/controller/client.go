package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/brunoshiroma/go-gin-poc/internal/entity"
	"github.com/brunoshiroma/go-gin-poc/internal/requests"
	"github.com/brunoshiroma/go-gin-poc/internal/service"
)

// CreateClient godoc
// @Summary Criar um cliente
// @Description cria um cliente
// @Tags client
// @Accept  json
// @Produce  json
// @Param message body requests.Client true "Client"
// @Success 200 {object} responses.Client
// @Failure 400 {string} httputil.HTTPError
// @Failure 404 {string} httputil.HTTPError
// @Failure 500 {string} httputil.HTTPError
// @Router /client/ [post]
func (c *Controller) CreateClient(ctx *gin.Context) {
	newClient := requests.Client{}
	clientService := service.NewClientService(c.dao)

	if err := ctx.ShouldBindJSON(&newClient); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	} else {
		entityClient := entity.Client{Name: newClient.Name, Email: newClient.Email}
		clientService.CreateClient(&entityClient)
		ctx.JSON(200, gin.H{
			"message": fmt.Sprintf("new Client with name %v Created, Id %v", newClient.Name, entityClient.Id),
		})
	}
}

// CreateClient godoc
// @Summary Lista todos os clientes
// @Description lista todos os clientes que não foram deletados
// @Tags client
// @Accept  json
// @Produce  json
// @Success 200 {array} responses.Client
// @Failure 400 {string} httputil.HTTPError
// @Failure 404 {string} httputil.HTTPError
// @Failure 500 {string} httputil.HTTPError
// @Router /client/ [get]
func (c *Controller) RetriveAllClient(ctx *gin.Context) {
	clientService := service.NewClientService(c.dao)

	result, err := clientService.RetrieveAllClient()
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"data": result,
		})
	}

}

// UpdateClient godoc
// @Summary Atualiza um cliente
// @Description atualiza um cliente
// @Tags client
// @Accept  json
// @Produce  json
// @Param message body requests.Client true "Client"
// @Success 200 {object} responses.Client
// @Failure 400 {string} httputil.HTTPError
// @Failure 404 {string} httputil.HTTPError
// @Failure 500 {string} httputil.HTTPError
// @Router /client/ [put]
func (c *Controller) UpdateClient(ctx *gin.Context) {
	clientService := service.NewClientService(c.dao)
	updateClient := requests.Client{}

	if err := ctx.ShouldBindJSON(&updateClient); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	} else {
		entityClient := entity.Client{Name: updateClient.Name, Email: updateClient.Email, Id: updateClient.Id}
		_, err := clientService.UpdateClient(&entityClient)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(200, gin.H{
				"message": fmt.Sprintf("Client with Id %v updated", entityClient.Id),
			})
		}
	}

}

// DeleteClient godoc
// @Summary Deleta um cliente
// @Description Deleta um cliente
// @Tags client
// @Accept  json
// @Produce  json
// @Param message body requests.Client true "Client"
// @Success 200 {object} responses.Client
// @Failure 400 {string} httputil.HTTPError
// @Failure 404 {string} httputil.HTTPError
// @Failure 500 {string} httputil.HTTPError
// @Router /client/ [delete]
func (c *Controller) DeleteClient(ctx *gin.Context) {
	clientService := service.NewClientService(c.dao)
	deleteClient := requests.Client{}

	if err := ctx.ShouldBindJSON(&deleteClient); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	} else {
		entityClient := entity.Client{Name: deleteClient.Name, Email: deleteClient.Email, Id: deleteClient.Id}
		err := clientService.DeleteClient(&entityClient)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(200, gin.H{
				"message": fmt.Sprintf("Client with Id %v deleted", entityClient.Id),
			})
		}

	}

}
