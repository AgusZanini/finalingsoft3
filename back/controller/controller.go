package controller

import (
	"net/http"
	"operations/dto"
	"operations/service"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type OperationController struct {
	operationservice service.OperationService
}

func NewOperationController(operationservice service.OperationService) *OperationController {
	return &OperationController{
		operationservice: operationservice,
	}
}

func (ctrl *OperationController) GetOperationById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error("Error converting id to integer")
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	operationdto, er := ctrl.operationservice.GetOperationById(id)
	if er != nil {
		log.Error("Error getting operation by id")
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, operationdto)
}

func (ctrl *OperationController) InsertOperation(c *gin.Context) {
	var operationdto dto.OperationDto
	err := c.BindJSON(&operationdto)
	if err != nil {
		log.Error("Error binding json to dto")
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	opresult, er := ctrl.operationservice.InsertOperation(operationdto)
	if er != nil {
		log.Error("Error inserting operation")
		c.JSON(http.StatusBadRequest, er.Error())
		return
	}

	c.JSON(http.StatusOK, opresult)
}
