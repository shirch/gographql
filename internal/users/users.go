package users

import (
	"github.com/jinzhu/gorm"
	"github.com/shirch/graphql/graph/model"
	"golang.org/x/crypto/bcrypt"
)

type Result struct {
	Username string
	Password string
	Id       string
}

func GetUserIdByUsername(username string, DB *gorm.DB) (string, error) {
	var result Result
	DB.Raw("select id from `users` WHERE name = ?", username).Scan(&result)
	return result.Id, nil
}

func GetUserById(id string, DB *gorm.DB) (model.User, error) {
	var result Result
	DB.Raw("select name, password from `users` WHERE id = ?", id).Scan(&result)
	user := model.User{
		ID:       id,
		Name:     result.Username,
		Password: result.Password,
	}
	return user, nil
}

func Authenticate(user *model.User, DB *gorm.DB) bool {
	var result Result
	DB.Raw("select password from `users` WHERE name = ?", user.Name).Scan(&result)
	return CheckPasswordHash(user.Password, result.Password)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
