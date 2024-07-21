package handler

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"dario.cat/mergo"
	"github.com/azeek21/blog/models"
	"github.com/azeek21/blog/pkg/utils"
	"github.com/azeek21/blog/views/components"
	"github.com/azeek21/blog/views/layouts"
	"github.com/gin-gonic/gin"
)

var ERR_NOT_HAVE_PERMISSON = errors.New("You don't have a permission to perform this action")

func (h Handler) ArticleByIdPage(ctx *gin.Context) {
	_articleId := ctx.Param("id")
	log.Println("ARTICLE ID: ", _articleId)

	articleId, err := strconv.ParseUint(_articleId, 0, 64)
	if err != nil {
		ctx.String(400, err.Error())
		ctx.Abort()
		return
	}

	article, err := h.service.GetArticleById(uint(articleId))
	if err != nil {
		ctx.String(400, err.Error())
		ctx.Abort()
		return
	}

	utils.RenderTempl(ctx, 200, layouts.ArticleByIdPage(*article))
}

func (h Handler) EditArticlePage(ctx *gin.Context) {
	_articleId := ctx.Param("id")
	log.Println("ARTICLE ID: ", _articleId)

	articleId, err := strconv.ParseUint(_articleId, 0, 64)
	if err != nil {
		ctx.String(400, err.Error())
		ctx.Abort()
		return
	}

	article, err := h.service.GetArticleById(uint(articleId))
	if err != nil {
		ctx.String(400, err.Error())
		ctx.Abort()
		return
	}

	utils.RenderTempl(ctx, 200, layouts.EditArticlePage(*article))
}

func (h Handler) NewArticlePage(ctx *gin.Context) {
	utils.RenderTempl(ctx, 200, layouts.NewArticlePage())
}
func (h Handler) GetAllArticles(ctx *gin.Context) {
	utils.RenderTempl(ctx, 200, layouts.ArticlesPage())
}

func (h Handler) GetArticleById(ctx *gin.Context) {
	imageUrl := "/public/logo.png"
	utils.RenderTempl(ctx, 200, layouts.ArticleByIdPage(models.Article{Title: "Article title", Content: "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet. Nisi anim cupidatat excepteur officia. Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident. Nostrud officia pariatur ut officia. Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate. Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod. Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim. Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis. Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet. Nisi anim cupidatat excepteur officia. Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident. Nostrud officia pariatur ut officia. Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate. Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod. Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim. Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis.", ImageUrl: &imageUrl}))
}

func (h Handler) CreateArticle(ctx *gin.Context) {
	newArticle := &models.Article{}
	err := ctx.ShouldBind(&newArticle)
	if err != nil {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, err.Error()))
		ctx.Abort()
		return
	}

	author_user, err := utils.GetUser(ctx)
	if err != nil {

		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, "500 something went wrong in the server"))
		ctx.Abort()
		return
	}
	_, err = h.service.CreateArticle(newArticle, author_user.ID)
	if err != nil {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, err.Error()))
		ctx.Abort()
		return
	}

	ctx.Header("HX-Redirect", fmt.Sprintf("/articles/%v", newArticle.ID))
}

func (h Handler) UpdateArticle(ctx *gin.Context) {
	_articleId := ctx.Param("id")
	articleId, err := utils.StringToUint(_articleId)
	if err != nil {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, err.Error()))
		ctx.Abort()
		return
	}

	article, err := h.service.GetArticleById(articleId)
	if err != nil {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, err.Error()))
		ctx.Abort()
		return
	}

	incomingArticle := &models.Article{}
	err = ctx.ShouldBind(&incomingArticle)
	if err != nil {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, err.Error()))
		ctx.Abort()
		return
	}

	author_user, err := utils.GetUser(ctx)
	if err != nil {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, "500 something went wrong in the server"))
		ctx.Abort()
		return
	}

	if article.AuthorID != author_user.ID {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, ERR_NOT_HAVE_PERMISSON.Error()))
		ctx.Abort()
		return
	}

	if err := mergo.Merge(article, incomingArticle, mergo.WithOverride); err != nil {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, "500 something went wrong in the server"))
		ctx.Abort()
		return
	}

	_, err = h.service.ArticleService.UpdateArticle(article)
	if err != nil {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, "500 something went wrong in the server"))
		ctx.Abort()
		return
	}

	ctx.Header("HX-Redirect", fmt.Sprintf("/articles/%v", article.ID))
}

func (h Handler) DeleteArticle(ctx *gin.Context) {
	_articleId := ctx.Param("id")
	articleId, err := utils.StringToUint(_articleId)
	if err != nil {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, err.Error()))
		ctx.Abort()
		return
	}
	article, err := h.service.ArticleService.GetArticleById(articleId)
	if err != nil {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, err.Error()))
		ctx.Abort()
		return
	}

	current_user, err := utils.GetUser(ctx)
	if err != nil {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, err.Error()))
		ctx.Abort()
		return
	}

	if article.AuthorID != current_user.ID {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, ERR_NOT_HAVE_PERMISSON.Error()))
		ctx.Abort()
		return
	}

	isDeleteted, err := h.service.DeleteArticle(articleId)
	if err != nil {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, err.Error()))
		ctx.Abort()
		return
	}

	if !isDeleteted {
		utils.RenderTempl(ctx, 200, components.AlertsContainer(models.ALERT_LEVELS.ERROR, "500 something went terribly wrong"))
		ctx.Abort()
		return
	}

	ctx.Header("HX-Redirect", "/")
}
