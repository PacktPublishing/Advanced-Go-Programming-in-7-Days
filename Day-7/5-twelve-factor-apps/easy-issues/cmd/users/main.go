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
	"net"
	"bufio"
	"io"
	"strings"
	"encoding/json"
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

	go func(us application.UserService) {
		tcpAddr, err := net.ResolveTCPAddr("tcp", ":9091")
		if err != nil {
			log.Fatal(err)
		}

		listener, err := net.ListenTCP("tcp", tcpAddr)
		if err != nil {
			log.Fatal(err)
		}

		for {
			conn, err := listener.Accept()

			if err != nil {
				continue
			}

			go handle(conn, us)
		}
	}(userService)

	log.Fatal(server.ListenAndServe())
}

func handle(conn net.Conn, us application.UserService)  {
	defer conn.Close()

	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))


	for {
		rw.WriteString("$users> ")
		rw.Flush()

		cmd, err := rw.ReadString('\n')
		switch {
		case err == io.EOF:
			fmt.Println("EOF")
		    return
		case err != nil:
			fmt.Println(err)
		return
		}

		cmd = strings.Trim(cmd, "\n ")

		// handle the command
		users, _ := us.UsersRepository.All()
		json.NewEncoder(rw).Encode(users)
	}
}

