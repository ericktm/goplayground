package views

import (
	"api/model"
	"encoding/json"
	"net/http"
)

func GetTags(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(model.Tag{"123", "Erick"})
}
func GetTag(w http.ResponseWriter, r *http.Request)    {}
func CreateTag(w http.ResponseWriter, r *http.Request) {}
func DeleteTag(w http.ResponseWriter, r *http.Request) {}
