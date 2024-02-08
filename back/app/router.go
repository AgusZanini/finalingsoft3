package app

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func MapUrls(router *gin.Engine, dependencies *Dependencies) {
	router.GET("/operation/:id", dependencies.OperationController.GetOperationById)
	router.POST("/operation", dependencies.OperationController.InsertOperation)

	log.Info("finishing url mappings...")
}
