package web

import (
	"net/http"
	"encoding/json"
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-7/1-Task-Overview/easy-issues/domain"
)

// Controller for Issue model
type IssueController struct {
	IssueService domain.IssueService
}

func (c IssueController) List(w http.ResponseWriter, r *http.Request) {
	options := ParseQuery(r.URL.Path)
	resp, err := c.IssueService.Issues(options)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	listResponseJson, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(listResponseJson)
}