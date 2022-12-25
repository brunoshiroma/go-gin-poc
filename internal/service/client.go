package service

import (
	"github.com/brunoshiroma/go-gin-poc/internal/dao"
	"github.com/brunoshiroma/go-gin-poc/internal/entity"
	"github.com/brunoshiroma/go-gin-poc/internal/responses"
)

type ClientService struct {
	dao dao.Dao[*entity.Client]
}

func NewClientService(dao dao.Dao[*entity.Client]) ClientService {
	return ClientService{
		dao: dao,
	}
}

func (s *ClientService) CreateClient(client *entity.Client) (*responses.Client, error) {
	err := s.dao.Create(client)
	if err != nil {
		return nil, err
	}
	return mapEntityToResponseWithoutDeleteDate(client), nil
}

func (s *ClientService) RetrieveAllClient() ([]responses.Client, error) {
	resultFromDao, err := s.dao.Retrieve(&entity.Client{})

	result := make([]responses.Client, len(resultFromDao))

	for index, entity := range resultFromDao {
		result[index] = *mapEntityToResponseWithoutDeleteDate(entity)
	}

	return result, err
}

func (s *ClientService) DeleteClient(client *entity.Client) error {
	return s.dao.Delete(client)
}

func (s *ClientService) UpdateClient(client *entity.Client) (*responses.Client, error) {
	err := s.dao.Update(client)
	if err != nil {
		return nil, err
	}
	return mapEntityToResponseWithoutDeleteDate(client), nil
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
