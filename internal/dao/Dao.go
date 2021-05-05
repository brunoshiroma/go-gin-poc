package dao

import (
	"math/rand"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"fmt"
	"log"
	"os"
	"reflect"
)

type Dao interface {
	StartDB() bool
	Create(entity interface{})
	Retrieve(entity interface{}) []interface{}
	Update(entity interface{})
	Delete(entity interface{})
	GetORM() *gorm.DB
}

type SimpleDao struct {
	db *gorm.DB
	id int64
}

func (s *SimpleDao) StartDB() bool {

	if s.db == nil {
		s.id = rand.Int63()
		dbPort, _ := os.LookupEnv("DB_PORT")
		dbHost, _ := os.LookupEnv("DB_HOST")
		dbUser, _ := os.LookupEnv("DB_USER")
		dbPass, _ := os.LookupEnv("DB_PASS")
		dbName, _ := os.LookupEnv("DB_NAME")

		var err error

		dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%v", dbHost, dbUser, dbPass, dbPort, dbName)
		s.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

		if err != nil {
			log.Printf("Error starting connection with DB, %v", err.Error())
			return false
		}
	}

	return true
}

func (s *SimpleDao) Create(entity interface{}) {
	s.db.Model(entity).Create(entity).Commit()
}

func (s *SimpleDao) Retrieve(e interface{}) []interface{} {

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
	}

	return nil
}

func (s *SimpleDao) Update(entity interface{}) {
	s.db.Save(entity)
}

func (s *SimpleDao) Delete(entity interface{}) {
	s.db.Delete(entity)
}

func (s *SimpleDao) GetORM() *gorm.DB {
	return s.db
}
