package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/AndreAppolariFilho/go_tiny_url/internal/database"
	"github.com/go-chi/chi/v5"
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
	originalUrl := params.Url
	if !strings.HasPrefix(originalUrl, "http") || !strings.HasPrefix(originalUrl, "https"){
		originalUrl = fmt.Sprintf("https://%s", originalUrl)
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
			OriginalUrl: originalUrl,
			TinyUrl: generatedURL,
		},
	)
	respondWithJSON(w, 200, databaseUrlToUrl(url))
}


func (apiCfg *apiConfig) handlerRedirectToUrl(w http.ResponseWriter, r *http.Request){
	tinyUrl := fmt.Sprintf("%s/%s", r.Host, chi.URLParam(r, "url"))
	url, err := apiCfg.DB.GetUrlByTyniUrl(context.Background(), tinyUrl)
	if err != nil{
		respondWithError(w, 404, fmt.Sprintf("This url doesn't exists, %v", err))
		return
	}
	http.Redirect(w, r, url.OriginalUrl, http.StatusFound)

}