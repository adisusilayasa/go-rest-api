package model

import "github.com/golang-jwt/jwt/v4"

type SendData struct {
	jwt.StandardClaims
	UserData LeakUserData `json:"user_data"`
	// Password string `json:"Password"`
}

type UserData struct {
	UID      string `json:"uid"`
	Role     string `json:"role"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UsersLogin struct {
	jwt.StandardClaims
	Users []UserData
}

type UserDatabase struct {
	User []UserData
}

type LeakUserData struct {
	UID      string `json:"uid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
