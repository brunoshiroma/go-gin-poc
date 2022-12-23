package dao

import (
	"github.com/brunoshiroma/go-gin-poc/internal/entity"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"fmt"
	"log"
	"os"
	"reflect"
)

type Dao interface {
	AutoMigrate() error
	StartDB() error
	Create(entity interface{}) error
	Retrieve(entity interface{}) ([]interface{}, error)
	Update(entity interface{}) error
	Delete(entity interface{}) error
	GetORM() *gorm.DB
}

type SimpleDao struct {
	db *gorm.DB
}

func (s *SimpleDao) AutoMigrate() error {
	return s.db.AutoMigrate(&entity.Client{})
}

func (s *SimpleDao) StartDB() error {

	if s.db == nil {
		dbPort, _ := os.LookupEnv("DB_PORT")
		dbHost, _ := os.LookupEnv("DB_HOST")
		dbUser, _ := os.LookupEnv("DB_USER")
		dbPass, _ := os.LookupEnv("DB_PASS")
		dbName, _ := os.LookupEnv("DB_NAME")
		dbFile, _ := os.LookupEnv("DB_FILE")

		var err error

		dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%v", dbHost, dbUser, dbPass, dbPort, dbName)

		if dbFile != "" {
			s.db, err = gorm.Open(sqlite.Open(dbFile), &gorm.Config{
				Logger: logger.Default.LogMode(logger.Info),
			})
		} else {
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

func (s *SimpleDao) Create(entity interface{}) error {
	var err error = nil
	s.db.Transaction(func(tx *gorm.DB) error {
		tx.Create(entity)
		if tx.Error != nil {
			err = tx.Error
		}
		return err
	})
	return err
}

func (s *SimpleDao) Retrieve(e interface{}) ([]interface{}, error) {

	entityType := reflect.TypeOf(e)

	arrayType := reflect.ArrayOf(1, entityType)
	arrayInstance := reflect.New(arrayType)
	arrayElem := arrayInstance.Elem()
	arrayInterface := arrayElem.Interface()

	dbResult := s.db.Model(e).Find(&arrayElem)

	log.Printf("I %v", arrayInterface)

	for index := 0; index < arrayElem.Len(); index++ {
		value := arrayElem.Index(index)
		log.Printf("Value %v", value)
	}

	if dbResult.Error != nil {
		log.Printf("Error on retrieve all %v", dbResult.Error)
		return nil, dbResult.Error
	}
	return nil, nil
}

func (s *SimpleDao) Update(entity interface{}) error {
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

func (s *SimpleDao) Delete(entity interface{}) error {
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

func (s *SimpleDao) GetORM() *gorm.DB {
	return s.db
}
