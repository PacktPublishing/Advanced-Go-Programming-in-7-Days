package web

import (
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"time"
	"encoding/json"
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-6/5-Authentication/easy-issues"
	"log"
	"github.com/Shopify/sarama"
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-7/4-Message-Queues/easy-issues/domain"
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-7/4-Message-Queues/easy-issues/application"
	"github.com/google/uuid"
)

// Controller for Auth model
type AuthController struct {
	AuthService domain.AuthService
	EventsProducer sarama.SyncProducer
	Secret string
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

	// Code to Register new User account -> TODO
	event := domain.NewCreateUserRegistrationEvent(uuid.New().String(), "alex.tomas@live.com")
	err := application.SubmitEvent(c.EventsProducer, event)

	if err != nil {
		log.Fatalf("%v.SubmitEvent(_) = _, %v: ", c.EventsProducer, err)
	}
}
