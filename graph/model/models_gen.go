// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateLinkInput struct {
	Title   string `json:"title"`
	Address string `json:"address"`
}

type Link struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Address string `json:"address"`
	User    *User  `json:"user"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type UpdateLinkInput struct {
	Title   string `json:"title"`
	Address string `json:"address"`
	UserID  string `json:"userId"`
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
