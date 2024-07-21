package repository

import (
	"log"

	"github.com/azeek21/blog/models"
	"github.com/azeek21/blog/pkg/utils"
	"gorm.io/gorm"
)

type ArticleRepo struct {
	db *gorm.DB
}

func NewArticleRepositroy(db *gorm.DB) ArticleRepository {
	return ArticleRepo{
		db: db,
	}
}

func (r ArticleRepo) GetArticleById(articleId uint) (*models.Article, error) {
	res := &models.Article{
		Model: gorm.Model{ID: articleId},
	}
	err := r.db.First(res).Error
	return res, err
}

func (r ArticleRepo) CreateArticle(article *models.Article) (*models.Article, error) {
	err := r.db.Create(article).Error
	return article, err
}

func (r ArticleRepo) Update(article *models.Article) (*models.Article, error) {
	log.Println("TO BE UPDATED: ", article.GetImage())
	err := r.db.Save(article).Error
	return article, err
}

func (r ArticleRepo) Delete(id uint) (bool, error) {
	article, err := r.GetArticleById(id)
	if err != nil {
		return false, err
	}

	err = r.db.Delete(article).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
func (r ArticleRepo) GetArticles(paging models.PagingIncoming) ([]models.Article, error) {
	var articles []models.Article

	err := r.db.Scopes(utils.Paginate(paging)).Find(&articles).Error
	return articles, err
}
