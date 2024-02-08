package app

import (
	"operations/config"
	"operations/controller"
	"operations/service"
	"operations/service/repositories"
)

type Dependencies struct {
	OperationController controller.OperationController
}

func NewDependencies() *Dependencies {
	operationclient := repositories.NewOperationClientImpl(config.DBUSER, config.DBPASS, config.DBHOST, config.DBPORT, config.DBNAME)
	operationclient.StartDbEngine()

	service := service.NewOperationServiceImpl(operationclient)

	operationcontroller := controller.NewOperationController(service)

	return &Dependencies{
		OperationController: *operationcontroller,
	}
}
