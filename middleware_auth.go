package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zulfikar-ditya/go-course/internal/auth"
	"github.com/zulfikar-ditya/go-course/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (APIConfig *APIConfig) authMiddleware(handler authHandler) http.HandlerFunc {

	return func (w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(&r.Header)

		if err != nil {
			log.Println(err)
			responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %s", err))
			return
		}

		getUser, err := APIConfig.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			responseWithError(w, http.StatusNotFound, "Couldn't get user")
			return
		}

		user := database.User{
			ID:       getUser.ID,
			Name:     getUser.Name,
			CreatedAt: getUser.CreatedAt,
			UpdatedAt: getUser.UpdatedAt,
		}

		handler(w, r, user)
}
}