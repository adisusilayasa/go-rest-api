package delivery

import (
	"belajar-interface/user/helper"
	"belajar-interface/user/model"
	"belajar-interface/user/service"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type M map[string]interface{}

var APPLICATION_NAME = " Template Rest-API"
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("the secret of meehh")
var Role = ""

type DataHandler struct {
	user service.Service
}

func NewHandler(userService service.Service) *DataHandler {
	return &DataHandler{userService}

}

func (h *DataHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser model.UserData
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(reqBody, &newUser)

	newUser, err = h.user.CreateUser(newUser)
	log.Print("Delivery", newUser)
	if err != nil {
		response := helper.APIResponse("Error creating User", http.StatusBadRequest, "error", nil)
		json.NewEncoder(w).Encode(response)
		return
	}
	response := helper.APIResponse("List of User", http.StatusOK, "success", newUser)
	json.NewEncoder(w).Encode(response)
	return
}
func (h *DataHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {

	username, password, ok := r.BasicAuth()
	password, err := h.user.PasswordHashing(password)
	if err != nil {
		log.Panic(err)
	}

	ok, userInfo := h.user.AuthenticateUser(username, password)
	if !ok {
		response := helper.APIResponse("Invalid username or password", http.StatusBadRequest, "error", nil)
		json.NewEncoder(w).Encode(response)
		return
	}
	expirationTime := time.Now().Add(59 * time.Minute)
	claims := model.SendData{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: expirationTime.Unix(),
		},
		UserData: model.LeakUserData{

			UID:      userInfo.UID,
			Role:     userInfo.Role,
			Username: userInfo.Username,
			Email:    userInfo.Email,
		},
	}
	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {

		response := helper.APIResponse("Wrong signed token", http.StatusBadRequest, "error", nil)
		json.NewEncoder(w).Encode(response)
		return
	}

	Role = userInfo.Role

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   signedToken,
		Expires: expirationTime,
	})
}

func (h *DataHandler) LogoutSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil || cookie.Value == "" {
		return
	}

	cookie = &http.Cookie{
		Name:     "token",
		Path:     "/",
		HttpOnly: false,
		Secure:   false,
		Domain:   "localhost",
		Expires:  time.Now(),
		MaxAge:   -1}

	http.SetCookie(w, cookie)

}

func (h *DataHandler) UserProfile(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokenString := c.Value
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SIGNATURE_KEY), nil
	})
	json.NewEncoder(w).Encode(token.Claims)
}

func (h *DataHandler) MiddlewareJWTAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
			return
		}
		if r.URL.Path == "/article/" {
			next.ServeHTTP(w, r)
			return
		}
		if r.URL.Path == "/register" {
			next.ServeHTTP(w, r)
			return
		}

		c, err := r.Cookie("token")
		if err != nil {
			log.Print(err.Error())
			next.ServeHTTP(w, r)
			return
		}
		tokenString := c.Value

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Signing method invalid")
			} else if method != JWT_SIGNING_METHOD {
				return nil, fmt.Errorf("Signing method invalid")
			}

			return JWT_SIGNATURE_KEY, nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(context.Background(), "userInfo", claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
