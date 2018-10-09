package web

import (
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-6/5-Authentication/easy-issues/domain"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"time"
	"encoding/json"
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-6/5-Authentication/easy-issues"
	pb "github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-7/3-RPCs/easy-issues/protocol"
	"context"
	"log"
)

// Controller for Issue model
type AuthController struct {
	AuthService domain.AuthService
	Secret string
	UserClient pb.UserClient
}

type LoginResponse struct {
	Token string `json:"token"`
}

type JWTData struct {
	jwt.StandardClaims
	CustomClaims map[string]string `json:"custom,omitempty"`
}

func (c AuthController) Verify(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(http.StatusOK)
}

func (c AuthController) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}

	email, _, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
		return
	}

	userRegistration, err := c.AuthService.GetRegistrationByEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if userRegistration.Status == domain.RegistrationStatusDeleted {
		if err != nil {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}
	}

	ok = easy_issues.CheckPasswordHash("secret", userRegistration.PasswordHash)
	if !ok {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	claims := JWTData{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			Issuer: "auth.service",
		},

		CustomClaims: map[string]string{
			"userId": userRegistration.Uuid,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(c.Secret))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	response := LoginResponse{
		Token: tokenString,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}


func (c AuthController) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}

	// TODO code for registering in Auth database
	ctx, cancel := context.WithTimeout(context.Background(), 10* time.Second)
	defer cancel()
	newUserResponse, err := c.UserClient.SubmitNewUser(ctx, &pb.NewUserRequest{
		Email: "alex@hotmail.com",
		Uuid: "12345",
		Status: pb.UserStatus_ACTIVE,
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(newUserResponse)
}