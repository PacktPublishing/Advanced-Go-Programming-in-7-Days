package controller

import (
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-6/3-Domain-Driver-Design/easy-issues/domain"
	"net/http"
)

// Controller for Issue model
type IssueController struct {
	IssueService domain.IssueService
}

func (c IssueController) List(w http.ResponseWriter, r *http.Request) {
}

func (c IssueController) Show(w http.ResponseWriter, r *http.Request) {

}

func (c IssueController) Create(w http.ResponseWriter, r *http.Request) {

}

func (c IssueController) Delete(w http.ResponseWriter, r *http.Request) {

}