package repository

import (
	"belajar-interface/article/model"
	"database/sql"
	"fmt"
	"log"
)

type Repository interface {
	CreateArticle(model.CreateArticle) (model.CreateArticle, error)
	FindByID(id int) (model.Article, error)
	FindAll() ([]model.Article, error)
}

type repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{DB: db}
}

func (r *repository) CreateArticle(article model.CreateArticle) (model.CreateArticle, error) {
	query := "INSERT INTO tb_article VALUES (?, ?, ?, ?)"
	_, err := r.DB.Exec(query, article.Id, article.Title, article.Desc, article.Content)

	if err != nil {
		panic(err)
	}
	return article, err
}

func (r *repository) FindAll() ([]model.Article, error) {
	var articles []model.Article

	res, err := r.DB.Query("SELECT id, title, 'desc', content FROM tb_article")
	if err != nil {
		log.Print(err)
	}
	for res.Next() {
		var article model.Article
		err := res.Scan(&article.Id, &article.Title, &article.Desc, &article.Content)
		if err != nil {
			log.Print(err)
		}
		articles = append(articles, article)
	}

	defer res.Close()
	return articles, nil

}

func (r *repository) FindByID(id int) (model.Article, error) {
	var article model.Article
	err := r.DB.
		QueryRow("SELECT id, title, 'desc', content FROM tb_article WHERE id = ?", id).

		Scan(&article.Id, &article.Title, &article.Desc, &article.Content)

	if err != nil {
		fmt.Println(err.Error())
		return article, err
	}
	return article, nil

}

func (r *repository) CreateArticle(article model.CreateArticle) (model.CreateArticle, error) {
	query := "INSERT INTO tb_article VALUES (?, ?, ?, ?)"
	_, err := r.DB.Exec(query, article.Id, article.Title, article.Desc, article.Content)

	if err != nil {
		panic(err)
	}
	return article, err
}
