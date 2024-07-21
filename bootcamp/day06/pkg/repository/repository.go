package repository

import (
	"fmt"
	"log"

	"github.com/azeek21/blog/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	GetUserById(id uint) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(user *models.User) (bool, error)
	SetRole(user *models.User, role string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type RoleRepository interface {
	GetRoleByRoleCode(role string) (*models.Role, error)
	UpdateRole(role *models.Role) (*models.Role, error)
	GetAllRoles() ([]models.Role, error)
	DeleteRole(role *models.Role) (bool, error)
	CreateRole(role *models.Role) (*models.Role, error)
}

type ArticleRepository interface {
	GetArticles(paging models.PagingIncoming) ([]models.Article, error)
	GetArticleById(articleId uint) (*models.Article, error)
	CreateArticle(*models.Article) (*models.Article, error)
	Update(*models.Article) (*models.Article, error)
	Delete(id uint) (bool, error)
}

type CountRepository interface {
	Count(model interface{}) int64
}

type Repository struct {
	UserRepository
	RoleRepository
	ArticleRepository
	CountRepository
}

func NewRepositroy(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository:    NewUserRepository(db),
		RoleRepository:    NewRolesRepository(db),
		ArticleRepository: NewArticleRepositroy(db),
		CountRepository:   NewCountRepository(db),
	}
}

type PostgresConnectionConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PWD"`
	Dbname   string `mapstructure:"DB_NAME"`
	Port     string `mapstructure:"DB_PORT"`
}

func connectDb(config PostgresConnectionConfig) (*gorm.DB, error) {
	connectionString := fmt.Sprintf(`host=%v user=%v password=%v dbname=%v port=%v sslmode=disable`,
		config.Host, config.User, config.Password, config.Dbname, config.Port)

	return gorm.Open(postgres.Open(connectionString), &gorm.Config{})
}

func CreateDb(config PostgresConnectionConfig) (*gorm.DB, error) {
	log.Printf("User postgres config: %+v", config)
	return connectDb(config)
}
