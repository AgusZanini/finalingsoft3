package service

import (
	"operations/dto"
)

type OperationService interface {
	GetOperationById(id int) (dto.OperationDto, error)
	InsertOperation(operationdto dto.OperationDto) (float32, error)
}
