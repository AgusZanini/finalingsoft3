package service

import (
	"errors"
	"operations/dto"
	"operations/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

// crear de un cliente mock

type MockClient struct{}

func NewMockClient() *MockClient {
	return &MockClient{}
}

func (m *MockClient) GetOperationById(id int) (model.Operation, error) {
	if id == 1 {
		return model.Operation{
			Id:      id,
			Number1: 5,
			Number2: 5,
			Result:  5 * 5,
		}, nil
	} else {
		return model.Operation{}, errors.New("error")
	}
}

func (m *MockClient) InsertOperation(operation model.Operation) (float64, error) {
	return operation.Result, nil
}

// instanciar el cliente mock
var mockclient = NewMockClient()

// instanciar servicio de prueba
var testservice = NewOperationServiceImpl(mockclient)

func TestGetOperationById(t *testing.T) {

	// caso de prueba id existente
	id := 1

	operation1, err := testservice.GetOperationById(id)

	assert.NoError(t, err)
	assert.Equal(t, id, operation1.Id)

	// caso de prueba id inexistente
	id = 2

	operation2, err := testservice.GetOperationById(id)

	assert.Error(t, err)
	assert.NotEqual(t, id, operation2.Id)
}

func TestInsertOperation(t *testing.T) {

	// declaracion de las operaciones de prueba

	var testoperationdto1 = dto.OperationDto{
		Number1: 5,
		Number2: 4,
	}

	var testoperationdto2 = dto.OperationDto{
		Number1: -5,
		Number2: 4,
	}

	var testoperationdto3 = dto.OperationDto{
		Number1: -5,
		Number2: -4,
	}

	var testoperationdto4 = dto.OperationDto{
		Number1: 0,
		Number2: 4,
	}

	var testoperationdto5 = dto.OperationDto{
		Number1: 0,
		Number2: 0,
	}

	var testoperationdto6 = dto.OperationDto{
		Number1: 5.5349,
		Number2: 7.2918,
	}

	// caso de prueba dos numeros enteros positivos
	expectedresult1 := 20.0

	result1, err := testservice.InsertOperation(testoperationdto1)

	assert.NoError(t, err)
	assert.Equal(t, expectedresult1, result1)

	// caso de prueba un numero entero positivo y un numero entero negativo
	expectedresult2 := -20.0

	result2, err := testservice.InsertOperation(testoperationdto2)

	assert.NoError(t, err)
	assert.Equal(t, expectedresult2, result2)

	// caso de prueba dos numeros enteros negativos
	expectedresult3 := 20.0

	result3, err := testservice.InsertOperation(testoperationdto3)

	assert.NoError(t, err)
	assert.Equal(t, expectedresult3, result3)

	// caso de prueba un numero entero positivo y cero
	expectedresult4 := 0.0

	result4, err := testservice.InsertOperation(testoperationdto4)

	assert.NoError(t, err)
	assert.Equal(t, expectedresult4, result4)

	// caso de prueba cero y cero
	expectedresult5 := 0.0

	result5, err := testservice.InsertOperation(testoperationdto5)

	assert.NoError(t, err)
	assert.Equal(t, expectedresult5, result5)

	// caso de prueba dos numeros decimales
	expectedresult6 := 40.359383820000005

	result6, err := testservice.InsertOperation(testoperationdto6)

	assert.NoError(t, err)
	assert.Equal(t, expectedresult6, result6)

}
