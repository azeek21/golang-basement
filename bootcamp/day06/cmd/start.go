package main

import (
	"github.com/azeek21/blog/pkg/handler"
	"github.com/azeek21/blog/pkg/repository"
	"github.com/azeek21/blog/pkg/service"
	"github.com/azeek21/blog/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const (
	HtmlFormat = "html"
	JsonFormat = "json"
	FormatKey  = "format"
	IdKey      = "id"
)

func main() {
	utils.InitConfig(gin.Mode())
	engine := gin.Default()
	engine.LoadHTMLGlob(viper.GetString("VIEWS_PATH"))
	dbConfig := repository.PostgresConnectionConfig{}
	dbConfig, err := utils.LoadConfig(dbConfig)
	utils.Must(err)
	db, err := repository.CreateDb(dbConfig)
	utils.Must(err)

	repo := repository.NewRepositroy(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)
	handler.RegisterHandlers(engine, viper.GetString("STATIC_PATH"))
	engine.Run(viper.GetString("SRV_PORT"))

}
