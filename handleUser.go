package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/zulfikar-ditya/go-course/internal/database"
)

func (APIConfig *APIConfig) handleCreateNewUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	type person struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := person{}
	err := decoder.Decode(&params)

	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload", err))
		return
	}

	user, err := APIConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		Name: params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		responseWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to create user: %s", err))
		return
	}

	responseWithJson(w, http.StatusOK, user)
}

func (cfg *APIConfig) handlerUsersGet(w http.ResponseWriter, r *http.Request, user database.User) {
	responseWithJson(w, http.StatusOK, database.User(user))
}