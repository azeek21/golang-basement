package handler

import (
	"log"

	"github.com/azeek21/blog/models"
	"github.com/azeek21/blog/pkg/utils"
	"github.com/azeek21/blog/views/components"
	"github.com/azeek21/blog/views/layouts"
	"github.com/gin-gonic/gin"
)

func (h Handler) IndexPage(ctx *gin.Context) {
	paging, err := utils.GetPagingIncomingFromContext(ctx)
	if err != nil {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, err.Error()))
		ctx.Abort()
		return
	}

	articles, err := h.service.ArticleService.GetArticles(*paging)
	if err != nil {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, err.Error()))
		ctx.Abort()
		return
	}

	totalArticles := h.service.CountingService.Count(&models.Article{})
	log.Println("TOTAL ARTICLES: ", totalArticles)

	utils.RenderTempl(ctx, 200, layouts.IndexPage(articles, paging.TransformToOutgoing(totalArticles)))
}
