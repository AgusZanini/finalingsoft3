package service

import (
	"operations/dto"
	"operations/model"
	"operations/service/repositories"
)

type OperationServiceImpl struct {
	operationclient repositories.OperationClient
}

func NewOperationServiceImpl(operationclient repositories.OperationClient) *OperationServiceImpl {
	return &OperationServiceImpl{
		operationclient: operationclient,
	}
}

func (s *OperationServiceImpl) InsertOperation(operationdto dto.OperationDto) (float64, error) {
	var operation model.Operation

	operation.Number1 = operationdto.Number1
	operation.Number2 = operationdto.Number2

	operation.Result = operationdto.Number1 / operationdto.Number2

	opresult, err := s.operationclient.InsertOperation(operation)
	if err != nil {
		return 0, err
	}

	return opresult, nil
}

func (s *OperationServiceImpl) GetOperationById(id int) (dto.OperationDto, error) {
	var operationdto dto.OperationDto

	operation, err := s.operationclient.GetOperationById(id)
	if err != nil {
		return dto.OperationDto{}, err
	}

	operationdto.Id = operation.Id
	operationdto.Number1 = operation.Number1
	operationdto.Number2 = operation.Number2
	operationdto.Result = operation.Result

	return operationdto, nil
}
