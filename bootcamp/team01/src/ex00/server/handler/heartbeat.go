package handler

import (
	"net/http"
	"replication/config"
	"replication/models"
	"replication/utils"
	"slices"

	"github.com/gin-gonic/gin"
)

func (h handler) HealthCheck(ctx *gin.Context) {

	heartBeat := models.HeartBeat{}
	err := ctx.ShouldBind(&heartBeat)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			models.NewBadRequestResponse(utils.WithPrefix("/health-check", err)),
		)
		return
	}

	serviceHostIndex := slices.Index(heartBeat.Nodes, config.CONFIG.ADDRESS)
	if serviceHostIndex != -1 {
		heartBeat.Nodes = append(heartBeat.Nodes, config.CONFIG.ADDRESS)
	}

	ctx.JSON(http.StatusOK, heartBeat)
}
