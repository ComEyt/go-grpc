// Package api defines the user model
package api

type User struct {
	// User's name.
	// Required: true
	Name     string `json:"name"`

	// User's nickname.
	// Required: true
	NickName string `json:"nickname"`

	// User's address.
	Address  string `json:"address"`

	// User's email.
	Email    string `json:"email"`

}
