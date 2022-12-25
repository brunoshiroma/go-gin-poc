package dao

import (
	"testing"
	"time"

	config "github.com/brunoshiroma/go-gin-poc/internal"
	"github.com/brunoshiroma/go-gin-poc/internal/entity"
	"github.com/stretchr/testify/assert"
)

var dao *SimpleDao[*entity.Client] = nil

func setupORM() {
	config.Env.DbFile = ":memory:"

	if dao == nil {
		dao = &SimpleDao[*entity.Client]{}
		dao.StartDB()
		dao.AutoMigrate()
	}
}

func TestAutoMigrate(t *testing.T) {
	setupORM()
}

func TestCreate(t *testing.T) {
	setupORM()
	clientEntity := entity.Client{
		Name:      "TEST",
		Email:     "test@email.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	dao.Create(&clientEntity)
	assert.Greater(t, clientEntity.Id, uint64(0))
}

func TestDelete(t *testing.T) {
	TestCreate(t)

	dao.Delete(&entity.Client{
		Id: 1,
	})
}

func TestUpdate(t *testing.T) {
	TestCreate(t)

	dao.Update(&entity.Client{
		Id:   1,
		Name: "UPDATED",
	})
}

func TestRetriveById(t *testing.T) {
	TestCreate(t)

	entities, err := dao.Retrieve(&entity.Client{}, 1)

	assert.NotNil(t, entities)
	assert.Equal(t, uint64(1), entities[0].Id)
	assert.NoError(t, err)
}

func TestRetriveAll(t *testing.T) {
	TestCreate(t)
	TestCreate(t)

	entities, err := dao.Retrieve(&entity.Client{})

	assert.NotNil(t, entities)
	assert.Greater(t, len(entities), 1)
	assert.NoError(t, err)
}
