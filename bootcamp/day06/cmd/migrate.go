package main

import (
	"log"

	"github.com/azeek21/blog/models"
	"github.com/azeek21/blog/pkg/repository"
	"github.com/azeek21/blog/pkg/service"
	"github.com/azeek21/blog/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func main() {
	err := utils.InitConfig(gin.Mode())
	utils.Must(err)

	dbConf := repository.PostgresConnectionConfig{}
	dbConf, err = utils.LoadConfig(dbConf)
	utils.Must(err)
	db, err := repository.CreateDb(dbConf)
	utils.Must(err)

	err = db.Migrator().DropTable(&models.User{}, &models.Role{}, &models.Article{}) // NOTE: do'nt drop tables

	utils.Must(err)
	err = db.AutoMigrate(&models.User{}, &models.Role{}, &models.Article{})
	utils.Must(err)

	repo := repository.NewRepositroy(db)
	roles := viper.GetStringSlice("ROLES")
	log.Println("ROLES: ", roles)
	for _, role := range roles {
		_, err := repo.GetRoleByRoleCode(role)
		if err == gorm.ErrRecordNotFound {
			_, err = repo.CreateRole(&models.Role{
				Code: role,
			})
			utils.Must(err)
		}
		utils.Must(err)
	}
	passWordSerice := service.NewPasswordSerice()
	pwHash, err := passWordSerice.CreateHash(viper.GetString("SU_PWD"))

	utils.Must(err)

	superUser := &models.User{
		Email:    viper.GetString("SU_EMAIL"),
		FullName: viper.GetString("SU_FULL_NAME"),
		Username: viper.GetString("SU_USERNAME"),
		Password: pwHash,
		RoleCode: "admin",
	}

	usr, err := repo.GetUserByEmail(superUser.Email)
	if err == gorm.ErrRecordNotFound {
		_, err = repo.CreateUser(superUser)
		utils.Must(err)
	} else {
		superUser.Model = usr.Model
		usr, err = repo.UpdateUser(superUser)
	}

	utils.Must(err)
	log.Printf("Migration SUCCESS\nSuper user:\nName: %v, Email: %v, Role: %v\n", superUser.FullName, superUser.Email, superUser.RoleCode)
}
