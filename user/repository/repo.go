package repository

import (
	"belajar-interface/user/model"
	"database/sql"
	"log"
	"time"
)

type Repository interface {
	CreateUser(model.UserData) (model.UserData, error)
	GetUser(username string) (model.UserData, error)
	UpdateLastLogin(string)
}

type repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{DB: db}
}

func (r *repository) CreateUser(newUser model.UserData) (model.UserData, error) {

	query := "INSERT INTO tb_user (uid, username, email, password, create_date, last_login,role) VALUES (?,?,?,?,?,?,?)"
	_, err := r.DB.Exec(query, newUser.UID, newUser.Username, newUser.Email, newUser.Password, time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339), "user")

	if err != nil {
		log.Fatal(err)
	}
	return newUser, err
}

func (r *repository) GetUser(username string) (model.UserData, error) {
	var user model.UserData

	err := r.DB.QueryRow("SELECT uid, role, username, email, password FROM tb_user WHERE username = ?", username).
		Scan(&user.UID, &user.Role, &user.Username, &user.Email, &user.Password)
	if err != nil {
		log.Fatal(err)
	}
	return user, nil
}

func (r *repository) UpdateLastLogin(uid string) {

	query := "UPDATE tb_user SET last_login = ? WHERE uid = ?"
	_, err := r.DB.Exec(query, time.Now().Format(time.RFC3339), uid)
	if err != nil {
		log.Fatal(err)
	}
}
