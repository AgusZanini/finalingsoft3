package repositories

import (
	"fmt"
	"operations/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

type OperationClientImpl struct {
	Db *gorm.DB
}

func NewOperationClientImpl(DBUser string, DBPass string, DBHost string, DBPort int, DBName string) *OperationClientImpl {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", DBUser, DBPass, DBHost, DBPort, DBName))
	if err != nil {
		log.Error(fmt.Sprintf("Error initializing database: %v", err))
		return nil
	}
	return &OperationClientImpl{
		Db: db,
	}
}

func (s *OperationClientImpl) StartDbEngine() {
	s.Db.AutoMigrate(&model.Operation{})
	log.Info("Finishing Migration Tables")
}

func (o *OperationClientImpl) GetOperationById(id int) (model.Operation, error) {
	var operation model.Operation

	result := o.Db.Where("Id = ?", id).First(&operation)
	if result.Error != nil {
		log.Error("Error getting operation by id from database")
		return model.Operation{}, result.Error
	}

	return operation, nil
}

func (o *OperationClientImpl) InsertOperation(operation model.Operation) (float64, error) {
	result := o.Db.Create(&operation)
	if result.Error != nil {
		log.Error("Error inserting operation in database")
		return 0, result.Error
	}

	log.Debug("Operation inserted", operation)
	return operation.Result, nil
}
