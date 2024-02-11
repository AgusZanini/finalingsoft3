package repositories

import (
	"operations/model"
)

type OperationClient interface {
	GetOperationById(id int) (model.Operation, error)
	InsertOperation(operation model.Operation) (float64, error)
}
