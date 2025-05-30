package getrandom

import (
	"encoding/json"
	"net/http"
	"quoter/internal/http-server/response"
	"quoter/internal/models"
)

type QuoteProvider interface {
	RandomQuote() (*models.Quote)
}

type Response struct {
	response.Response
	Quote models.Quote
}

func New(provider QuoteProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		quote := provider.RandomQuote()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(
			Response{
				Response: response.Response{Status: response.StatusOK},
				Quote: *quote,
			})
	}
}

