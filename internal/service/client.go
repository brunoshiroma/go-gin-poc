package service

import (
	"github.com/brunoshiroma/go-gin-poc/internal/dao"
	"github.com/brunoshiroma/go-gin-poc/internal/entity"
	"github.com/brunoshiroma/go-gin-poc/internal/responses"
)

func CreateClient(client *entity.Client) *responses.Client {
	var dao dao.Dao = &dao.SimpleDao{}

	if dao.StartDB() {
		dao.Create(client)
		return mapEntityToResponseWithoutDeleteDate(client)
	}

	return nil
}

func RetrieveAllClient() []responses.Client {
	var dao dao.Dao = &dao.SimpleDao{}

	if dao.StartDB() {
		resultFromDao := make([]entity.Client, 1)

		//busca todos os clientes.
		dao.GetORM().Find(&resultFromDao)

		result := make([]responses.Client, len(resultFromDao))

		for index, entity := range resultFromDao {
			result[index] = *mapEntityToResponseWithoutDeleteDate(&entity)
		}

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

func UpdateClient(client *entity.Client) *responses.Client {
	var dao dao.Dao = &dao.SimpleDao{}

	if dao.StartDB() {
		dao.Update(client)
		return mapEntityToResponseWithoutDeleteDate(client)
	}

	return nil
}

func mapEntityToResponseWithoutDeleteDate(client *entity.Client) *responses.Client {
	return &responses.Client{
		Id:        client.Id,
		Name:      client.Name,
		Email:     client.Email,
		CreatedAt: client.CreatedAt,
		UpdatedAt: client.UpdatedAt,
	}
}
