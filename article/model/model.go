package model

type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type CreateArticle struct {
	Id      string `json:"id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}
