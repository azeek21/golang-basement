package service

import (
	"github.com/azeek21/blog/models"
	"github.com/azeek21/blog/pkg/repository"
)

type ArticleSer struct {
	repo repository.ArticleRepository
}

func newArticleService(repo repository.ArticleRepository) ArticleService {
	return ArticleSer{
		repo: repo,
	}
}

func (s ArticleSer) GetArticleById(articleId uint) (*models.Article, error) {
	return s.repo.GetArticleById(articleId)
}
func (s ArticleSer) CreateArticle(article *models.Article, authorId uint) (*models.Article, error) {
	article.AuthorID = authorId
	return s.repo.CreateArticle(article)
}

func (s ArticleSer) UpdateArticle(article *models.Article) (*models.Article, error) {
	return s.repo.Update(article)
}

func (s ArticleSer) DeleteArticle(id uint) (bool, error) {
	return s.repo.Delete(id)
}

func (s ArticleSer) GetArticles(paging models.PagingIncoming) ([]models.Article, error) {
	return s.repo.GetArticles(paging)
}
