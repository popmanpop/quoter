package list

import (
	"encoding/json"
	"net/http"
	"quoter/internal/http-server/response"
	"quoter/internal/models"
)

type QuoteProvider interface {
	ListQuotes() []models.Quote
	ListQuotesAuthor(author string) []models.Quote
}

type Response struct {
	response.Response
	Quotes []models.Quote
}

func New(provider QuoteProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		query := r.URL.Query()

		author := query.Get("author")

		var quotes []models.Quote
		if author == "" {
			quotes = provider.ListQuotes()
		} else {
			quotes = provider.ListQuotesAuthor(author)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(
			Response{
				Response: response.Response{Status: response.StatusOK},
				Quotes: quotes,
			})
	}
}