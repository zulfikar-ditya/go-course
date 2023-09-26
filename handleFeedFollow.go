package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/zulfikar-ditya/go-course/internal/database"
)

func (APIConfig *APIConfig) handleCreateNewFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	defer r.Body.Close();

	type FeedFollow struct {
		FeedId uuid.UUID `json:"FeedId"`
	}

	decoder := json.NewDecoder(r.Body)
	
	params := FeedFollow{}
	err := decoder.Decode(&params)

	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	feedFollow, err := APIConfig.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		UserID : user.ID,
		FeedID : params.FeedId,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseWithJson(w, http.StatusOK, database.FeedFollow(feedFollow))
}

func (APIConfig *APIConfig) handleGetUserFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	
	defer r.Body.Close();

	feedFollow, err := APIConfig.DB.GetUserFeedFollow(r.Context(), user.ID)

	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseWithJson(w, http.StatusOK, feedFollow)
}

func (APIConfig *APIConfig) handleDeleteUserFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	defer r.Body.Close();

	id := chi.URLParam(r, "feed_follow_id")
	uuid, err := uuid.Parse(id)

	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = APIConfig.DB.DeleteUserFeedFollow(r.Context(), database.DeleteUserFeedFollowParams{
		ID: uuid,
		UserID: user.ID,
	})

	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseWithJson(w, http.StatusOK, "success")
}