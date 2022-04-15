package handler

import (
	"net/http"
	"world-heritage-scrape/app/worldheritage/usecase"
)

func PutWorldHeritageHandler(w http.ResponseWriter, r *http.Request) {
	usecase.PutWorldHeritage()
}
