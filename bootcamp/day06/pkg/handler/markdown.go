package handler

import (
	"log"

	"github.com/azeek21/blog/pkg/utils"
	"github.com/azeek21/blog/views/components"
	"github.com/gin-gonic/gin"
)

func (h Handler) ShowMarkdownPreview(ctx *gin.Context) {
	content := ctx.Request.FormValue("content")
	log.Println("Show preview")
	utils.RenderTempl(ctx, 200, components.MarkdownEditor(content, content, true))
}

func (h Handler) ShowMarkdownEditor(ctx *gin.Context) {
	content := ctx.Request.FormValue("content")
	log.Println("Show editor")
	utils.RenderTempl(ctx, 200, components.MarkdownEditor(content, content, false))
}
