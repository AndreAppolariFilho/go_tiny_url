package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/AndreAppolariFilho/go_tiny_url/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateNewUrl(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Url string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil{
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON:%v", err))
		return
	}
	if params.Url == ""{
		respondWithError(w, 400, fmt.Sprintf("Error should inform the url"))
		return
	}
	lengthOfTheCharacters := 7
	generatedURL := fmt.Sprintf("%s/%s",r.Host, TinyUrl(lengthOfTheCharacters))
	for ;urlExists(apiCfg.DB, generatedURL);{
		generatedURL, err = url.JoinPath(r.URL.Host, TinyUrl(lengthOfTheCharacters))
		if err != nil{
			respondWithError(w, 500, fmt.Sprintf("Error mounting the url: %v", err))
			return
		}
	}
	
	url, err := apiCfg.DB.CreateURL(
		r.Context(),
		database.CreateURLParams{
			ID: uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			OriginalUrl: params.Url,
			TinyUrl: generatedURL,
		},
	)
	respondWithJSON(w, 200, databaseUrlToUrl(url))
}