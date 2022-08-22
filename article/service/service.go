package service

import (
	"belajar-interface/article/model"
	"belajar-interface/article/repository"
)

type M map[string]interface{}

type Service interface {
	CreateArticle(article model.CreateArticle) (model.CreateArticle, error)
	GetAllArticle() ([]model.Article, error)
	GetArticleById(int) (model.Article, error)
}
type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) CreateArticle(article model.CreateArticle) (model.CreateArticle, error) {
	newProduct := model.CreateArticle{}
	newProduct.Id = article.Id
	newProduct.Title = article.Title
	newProduct.Desc = article.Desc
	newProduct.Content = article.Content
	data, err := s.repository.CreateArticle(newProduct)
	if err != nil {
		panic(err)
	}
	return data, nil
}

func (s *service) GetArticleById(Id int) (model.Article, error) {

	article, err := s.repository.FindByID(Id)
	if err != nil {
		return article, err
	}
	return article, err
}

func (s *service) GetAllArticle() ([]model.Article, error) {
	articles, err := s.repository.FindAll()
	if err != nil {
		return articles, err
	}
	return articles, err
}

func (s *service) CreateArticle(article model.CreateArticle) (model.CreateArticle, error) {
	newProduct := model.CreateArticle{}
	newProduct.Id = article.Id
	newProduct.Title = article.Title
	newProduct.Desc = article.Desc
	newProduct.Content = article.Content
	data, err := s.repository.CreateArticle(newProduct)
	if err != nil {
		panic(err)
	}
	return data, nil
}
