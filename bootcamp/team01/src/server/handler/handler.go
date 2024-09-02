package handler

import (
	"net/http"
	"replication/models"
	"replication/server/middleware"
	"replication/server/repository"
	"replication/utils"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	RegisterDefaultRoutes(*gin.RouterGroup) error
}

type handler struct {
	repo repository.Repository
}

func NewHandler(repo repository.Repository) Handler {
	return &handler{
		repo: repo,
	}
}

func (h handler) RegisterDefaultRoutes(root *gin.RouterGroup) error {

	storageRoute := root.Group("/storage")
	{
		idValidatorMiddleware := middleware.ParamValidatorMiddleware("key", utils.IsValidUUID)
		storageRoute.GET("/:key", idValidatorMiddleware, h.GetFromStorage)
		storageRoute.POST("/", h.SetToStorage) // NOTE: this is both create and update a.k.a UPSERT
		storageRoute.DELETE("/:key", idValidatorMiddleware, h.DeleteFromStorage)
	}

	return nil
}

func (h handler) GetFromStorage(ctx *gin.Context) {
	key := ctx.Param("key")
	value, err := h.repo.Get(key)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewBadRequestResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, models.StorageItem{
		Key:   key,
		Value: value,
	})
}

func (h handler) SetToStorage(ctx *gin.Context) {
	res := models.StorageItem{}
	err := ctx.ShouldBind(&res)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewBadRequestResponse(err))
		return
	}

	err = h.repo.Set(res.Key, res.Value)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewBadRequestResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h handler) DeleteFromStorage(ctx *gin.Context) {
	key := ctx.Param("key")
	err := h.repo.Delete(key)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewBadRequestResponse(err))
	}
	ctx.Status(http.StatusOK)
}
