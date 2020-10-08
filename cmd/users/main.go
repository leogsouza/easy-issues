package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/leogsouza/easy-issues/domain"
	"github.com/leogsouza/easy-issues/persistence/memory"
	"github.com/leogsouza/easy-issues/service"
	"github.com/leogsouza/easy-issues/web/controller"
)

func main() {
	userRepo := memory.NewUserRepository()

	userService := service.NewUserService(userRepo)

	userController := controller.UserController{
		UserService: userService,
	}

	for i := 0; i < 10; i++ {
		userService.Create(&domain.User{Name: fmt.Sprintf("User_%d", i)})
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/", userController.List)

	server := &http.Server{
		Addr:           ":8091",
		Handler:        mux,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())

}
