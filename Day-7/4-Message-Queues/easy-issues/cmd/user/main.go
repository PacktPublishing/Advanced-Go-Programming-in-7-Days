package main

import (
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-7/4-Message-Queues/easy-issues/application"
	"net/http"
	"time"
	"log"
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-7/4-Message-Queues/easy-issues/persistence/memory"
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-7/4-Message-Queues/easy-issues/web"
	"github.com/Shopify/sarama"
)

func main()  {
	userRepo := memory.NewUserRepository()

	userService := application.UserService{
		UsersRepository: userRepo,
		EventsConsumer: application.NewKafkaConsumer(),
	}

	userController := web.UserController{
		UserService: userService,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("/", userController.List)

	go func() {
		defer userService.EventsConsumer.Close()
		consumer, err :=
			userService.EventsConsumer.ConsumePartition(application.Topic, application.Partition, sarama.OffsetOldest)
		if err != nil {
			log.Fatalf("Kafka error: %s\n", err)
		}

		application.ConsumeUserRegistrationEvents(consumer, userService.OnEventTypeHandle)
	}()

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

