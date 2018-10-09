package main

import (
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-6/3-Domain-Driver-Design/easy-issues/persistence/memory"
	"fmt"
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-6/3-Domain-Driver-Design/easy-issues/application"
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-6/3-Domain-Driver-Design/easy-issues/domain"
	"net/http"
	"time"
	"log"
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-6/3-Domain-Driver-Design/easy-issues/web/controller"
)

func main()  {
	userRepo := memory.NewUserRepository()

	userService := application.UserService{
		UsersRepository: userRepo,
	}

	userController := controller.UserController{
		UserService: userService,
	}

	for i:=0;i<10;i+=1 {
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
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
