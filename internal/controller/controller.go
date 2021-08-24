package controller

import "github.com/brunoshiroma/go-gin-poc/internal/dao"

type Controller struct {
	dao dao.Dao
}

func NewClientController(dao dao.Dao) *Controller {
	return &Controller{
		dao: dao,
	}
}
