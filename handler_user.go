package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/vishnuavm5/go-rssagg/internal/database"
)

func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON:%s", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user:%s", err))
		return
	}
	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	// apikey, err := auth.GetApiKey(r.Header)
	// if err != nil {
	// 	respondWithError(w, 403, fmt.Sprintf("Auth error:%v", err))
	// 	return
	// }
	// user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apikey)
	// if err != nil {
	// 	respondWithError(w, 400, fmt.Sprintf("Couldnt get the user:%v", err))
	// 	return
	// }
	respondWithJSON(w, 200, databaseUserToUser(user))
}
