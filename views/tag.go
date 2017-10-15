package views

import (
	"net/http"
	"encoding/json"
	"api/model"
)

func GetTags(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(model.Tag{"123","Erick"})
}
func GetTag(w http.ResponseWriter, r *http.Request) {}
func CreateTag(w http.ResponseWriter, r *http.Request) {}
func DeleteTag(w http.ResponseWriter, r *http.Request) {}