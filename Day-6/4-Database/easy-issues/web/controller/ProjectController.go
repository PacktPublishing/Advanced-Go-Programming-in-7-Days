package controller

import (
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-6/4-Database/easy-issues/domain"
	"net/http"
)

// Controller for Project model
type ProjectController struct {
	ProjectService domain.ProjectService
}

func (c ProjectController) List(w http.ResponseWriter, r *http.Request) {

}

func (c ProjectController) Show(w http.ResponseWriter, r *http.Request) {

}

func (c ProjectController) Create(w http.ResponseWriter, r *http.Request) {

}

func (c ProjectController) Delete(w http.ResponseWriter, r *http.Request) {

}