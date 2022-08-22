package route

import (
	articleDelivery "belajar-interface/article/delivery"
	articleService "belajar-interface/article/service"
	"belajar-interface/middleware"
	userDelivery "belajar-interface/user/delivery"
	userService "belajar-interface/user/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func NewHandler(articleService articleService.Service, userService userService.Service) {
	mx := new(middleware.CustomMux)

	handlerArticle := articleDelivery.NewHandler(articleService)
	handlerUser := userDelivery.NewHandler(userService)

	mx.RegisterMiddleware(handlerUser.MiddlewareJWTAuthorization)

	mx.HandleFunc("/article/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "GET" {
			GetId := strings.TrimPrefix(r.URL.Path, "/article/")
			if GetId != "" {
				id, err := strconv.Atoi(GetId)
				if err != nil {
					panic(err)
				}
				handlerArticle.GetArticleById(id, w, r)
				return
			}
			handlerArticle.GetAllArticleHandler(w, r)
		}
		if r.Method == "POST" {
			handlerArticle.CreateArticleData(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	})

	mx.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "POST" {
			http.ServeFile(w, r, "form.html")
			return
		}

		if r.Method == "POST" {
			handlerUser.LoginHandler(w, r)
			return
		}
		return
	})

	mx.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "POST" {
			handlerUser.CreateUser(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	})

	mx.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "POST" {
			handlerUser.LogoutSession(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	})
	mx.HandleFunc("/account/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "GET" {
			handlerUser.UserProfile(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	})

	mx.HandleFunc("/show", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "GET" {
			data, err := r.Cookie("token")
			json.NewEncoder(w).Encode(err)
			json.NewEncoder(w).Encode(data)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	})

	server := &http.Server{
		Addr:    ":8081",
		Handler: mx,
	}

	err := server.ListenAndServe()

	if err != nil {

		panic(err)
	}
}
