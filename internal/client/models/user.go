// User model
package models

type User struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	AuthToken string `json:"token"`
	Password  string `json:"-"`
}
