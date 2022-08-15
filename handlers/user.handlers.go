package userhandlers

import (
	"encoding/json"
	"fmt"
	"go/api/models/user"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Create(w http.ResponseWriter, r *http.Request) {

	var model user.User

	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := user.PostUserService(model)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   500,
			"Message": fmt.Sprintf("Error on insert user: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   201,
			"Message": fmt.Sprintf("User inserted successfully, ID: %v", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error converting ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var model user.User

	err = json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := user.UpdateUserService(int64(id), model)
	if err != nil {
		log.Printf("Error on update user: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Many users was update: %v", rows)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	resp := map[string]any{
		"Error":   200,
		"Message": fmt.Sprintf("User updated successfully, ID: %v", id),
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error converting ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var model user.User

	err = json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := user.DeleteUserService(int64(id), model)
	if err != nil {
		log.Printf("Error on delete user: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Many users was deleted: %v", rows)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	resp := map[string]any{
		"Error":   200,
		"Message": fmt.Sprintf("User deleted successfully, ID: %v", id),
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func List(w http.ResponseWriter, r *http.Request) {
	users, err := user.GetAllUsersService(0)
	if err != nil {
		log.Printf("Error on get all users: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
