package main

import (
	"net/http"
	"time"
	"log"
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-7/4-Message-Queues/easy-issues/application"
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-7/4-Message-Queues/easy-issues/domain"
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-7/4-Message-Queues/easy-issues/web"
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-7/4-Message-Queues/easy-issues/persistence/db"
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-6/5-Authentication/easy-issues"
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-6/5-Authentication/easy-issues/web/handler"
)

func main()  {
	authRepo := db.NewAuthRepository()
	authService := application.AuthService{
		AuthRepository: authRepo,
	}
	authController := web.AuthController{
		AuthService: authService,
		EventsProducer: application.NewKafkaSyncProducer(),
		Secret: application.Secret,
	}

	pwhash ,_  := easy_issues.HashPassword("opensessame")

	authRepo.Create(&domain.UserRegistration{
		Email: "aladdin@mail.com",
		Uuid: "1234",
		PasswordHash: pwhash,
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/auth/login", authController.Login)
	mux.HandleFunc("/auth/verify", handler.JWTAuthHandler(authController.Verify))
	mux.HandleFunc("/auth/register", authController.Register)

	server := &http.Server{
		Addr:           ":8090",
		Handler:        mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}