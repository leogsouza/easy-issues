package controller

import (
	"encoding/json"
	"net/http"

	"github.com/leogsouza/easy-issues/domain"
)

// UserController is a controller for User model
type UserController struct {
	UserService domain.UserService
}

// List returns all users
func (c UserController) List(w http.ResponseWriter, r *http.Request) {
	users, err := c.UserService.Users()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userJSON, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(userJSON)
}
