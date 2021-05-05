package service

import (
	"github.com/brunoshiroma/go-gin-poc/internal/dao"
	"github.com/brunoshiroma/go-gin-poc/internal/entity"
)

func CreateClient(client *entity.Client) {
	var dao dao.Dao = &dao.SimpleDao{}

	if dao.StartDB() {
		dao.Create(client)
	}
}

func RetrieveAllClient() []entity.Client {
	var dao dao.Dao = &dao.SimpleDao{}

	if dao.StartDB() {
		result := make([]entity.Client, 1)

		dao.GetORM().Find(&result)

		return result
	}

	return nil
}

func DeleteClient(client *entity.Client) {
	var dao dao.Dao = &dao.SimpleDao{}

	if dao.StartDB() {
		dao.Delete(client)
	}
}

func UpdateClient(client *entity.Client) {
	var dao dao.Dao = &dao.SimpleDao{}

	if dao.StartDB() {
		dao.Update(client)
	}
}
