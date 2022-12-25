package controller

import (
	"github.com/brunoshiroma/go-gin-poc/internal/dao"
	"github.com/brunoshiroma/go-gin-poc/internal/entity"
)

type Controller[E interface{}] struct {
	dao dao.Dao[E]
}

func NewClientController[E *entity.Client](dao dao.Dao[E]) *Controller[E] {
	return &Controller[E]{
		dao: dao,
	}
}
