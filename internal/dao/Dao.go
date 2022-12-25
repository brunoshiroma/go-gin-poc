package dao

import (
	config "github.com/brunoshiroma/go-gin-poc/internal"
	"github.com/brunoshiroma/go-gin-poc/internal/entity"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"fmt"
	"log"
)

type Dao[E interface{}] interface {
	AutoMigrate() error
	StartDB() error
	Create(entity E) error
	Retrieve(entity E, filters ...interface{}) ([]E, error)
	Update(entity E) error
	Delete(entity E) error
	GetORM() *gorm.DB
}

type SimpleDao[E interface{}] struct {
	db *gorm.DB
}

func (s *SimpleDao[E]) AutoMigrate() error {
	return s.db.AutoMigrate(&entity.Client{})
}

func (s *SimpleDao[E]) StartDB() error {

	if s.db == nil {
		var err error

		if config.Env.DbFile != "" {
			s.db, err = gorm.Open(sqlite.Open(config.Env.DbFile), &gorm.Config{
				Logger: logger.Default.LogMode(logger.Info),
			})
		} else {
			dsn := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%v",
				config.Env.DbHost,
				config.Env.DbUser,
				config.Env.DbPass,
				config.Env.DbPort,
				config.Env.DbName,
			)
			s.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
				Logger: logger.Default.LogMode(logger.Info),
			})
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (s *SimpleDao[E]) Create(e E) error {
	var err error = nil
	s.db.Transaction(func(tx *gorm.DB) error {
		tx.Create(e)
		if tx.Error != nil {
			err = tx.Error
		}
		return err
	})
	return err
}

func (s *SimpleDao[E]) Retrieve(e E, filter ...interface{}) ([]E, error) {
	var result = make([]E, 0)

	dbResult := s.db.Model(e).Find(&result, filter...)

	if dbResult.Error != nil {
		log.Printf("Error on retrieve all %v", dbResult.Error)
		return nil, dbResult.Error
	}

	return result, nil
}

func (s *SimpleDao[E]) Update(entity E) error {
	var err error = nil
	s.db.Transaction(func(tx *gorm.DB) error {
		tx.Save(entity)
		if tx.Error != nil {
			err = tx.Error
		}
		return err
	})
	return err
}

func (s *SimpleDao[E]) Delete(entity E) error {
	var err error = nil
	s.db.Transaction(func(tx *gorm.DB) error {
		s.db.Delete(entity)
		if tx.Error != nil {
			err = tx.Error
		}
		return err
	})
	return err
}

func (s *SimpleDao[E]) GetORM() *gorm.DB {
	return s.db
}
