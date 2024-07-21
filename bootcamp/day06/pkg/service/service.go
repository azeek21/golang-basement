package service

import (
	"github.com/azeek21/blog/models"
	"github.com/azeek21/blog/pkg/repository"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	GetUserById(id uint) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(user *models.User) (bool, error)
	SetRole(user *models.User, role string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type ArticleService interface {
	GetArticles(paging models.PagingIncoming) ([]models.Article, error)
	GetArticleById(articleId uint) (*models.Article, error)
	CreateArticle(article *models.Article, authorId uint) (*models.Article, error)
	UpdateArticle(article *models.Article) (*models.Article, error)
	DeleteArticle(id uint) (bool, error)
}

type JwtService interface {
	CreateJwt(creds *models.User) (string, error)
	// returns user id and err
	VerifyJwt(token string) (uint, error)
}

type PasswordService interface {
	CreateHash(pwd string) (string, error)
	VerifyPasswordAgainstHash(pwd string, hash string) (bool, error)
}

type CountingService interface {
	Count(interface{}) int
}

type Service struct {
	UserService
	ArticleService
	JwtService
	PasswordService
	CountingService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserService:     NewUserService(repo.UserRepository),
		ArticleService:  newArticleService(repo.ArticleRepository),
		JwtService:      NewJwtSerice(),
		PasswordService: NewPasswordSerice(),
		CountingService: NewCountService(repo.CountRepository),
	}
}
