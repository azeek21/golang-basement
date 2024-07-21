package utils

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func RenderTempl(c *gin.Context, status int, template templ.Component) error {
	return template.Render(c, c.Writer)
}

func RenderMdToHTML(input string) templ.Component {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse([]byte(input))

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	res := markdown.Render(doc, renderer)
	return templ.Raw(string(res), nil)
}
