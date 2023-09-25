package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/zulfikar-ditya/go-course/internal/database"
)

func (APIConfig *APIConfig) handleCreateNewFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	defer r.Body.Close()

	type CreateFeed struct {
		UserId uuid.UUID `json:"user_id"`
		Url  string `json:"url"`
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := CreateFeed{}
	err := decoder.Decode(&params)

	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	feed, err := APIConfig.DB.CreateNewFeed(r.Context(), database.CreateNewFeedParams{
		ID: uuid.New(),
		UserID: user.ID,
		Url: params.Url,
		Name: params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Failed to create feed")
		return
	}

	responseWithJson(w, http.StatusCreated, feed)
}

func (APIConfig *APIConfig) handleGetFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	defer r.Body.Close()

	feeds, err := APIConfig.DB.GetFeeds(r.Context())

	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Failed to get feeds")
		return
	}

	responseWithJson(w, http.StatusOK, feeds)
}