package delivery

import (
	"belajar-interface/article/helper"
	"belajar-interface/article/model"
	"belajar-interface/article/service"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type DataHandler struct {
	article service.Service
}

func NewHandler(articleService service.Service) *DataHandler {
	return &DataHandler{articleService}

}

func (h *DataHandler) GetAllArticleHandler(w http.ResponseWriter, r *http.Request) {
	article, err := h.article.GetAllArticle()
	if err != nil {
		response := helper.APIResponse("Error getting all articles", http.StatusBadRequest, "error", nil)
		json.NewEncoder(w).Encode(response)
		return
	}
	response := helper.APIResponse("List of articles", http.StatusOK, "success", article)
	json.NewEncoder(w).Encode(response)
	return
}

func (h *DataHandler) GetArticleById(id int, w http.ResponseWriter, r *http.Request) {

	article, err := h.article.GetArticleById(id)
	if err != nil {
		response := helper.APIResponse("Error getting all articles", http.StatusBadRequest, "error", nil)
		json.NewEncoder(w).Encode(response)
		return
	}
	response := helper.APIResponse("List of articles", http.StatusOK, "success", article)
	json.NewEncoder(w).Encode(response)
	return

}

func (h *DataHandler) CreateArticleData(w http.ResponseWriter, r *http.Request) {
	var article model.CreateArticle
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(reqBody, &article)

	article, err = h.article.CreateArticle(article)
	if err != nil {
		response := helper.APIResponse("Error creating article", http.StatusBadRequest, "error", nil)
		json.NewEncoder(w).Encode(response)
		return
	}
	response := helper.APIResponse("List of articles", http.StatusOK, "success", article)
	json.NewEncoder(w).Encode(response)
	return
}
