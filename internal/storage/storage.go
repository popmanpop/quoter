package storage

import (
	"errors"
	"math/rand"
	"quoter/internal/models"
)

var (
	ErrNoQuotes error = errors.New("quote not found")
)

type Storage struct {
	quotes map[int64]*models.Quote
	pastID int64
}

func New() *Storage {
	return &Storage{
		quotes: make(map[int64]*models.Quote),
		pastID: 0,
	}
}

func (s *Storage) Save(quote string, author string) int64 {
	id := s.pastID + 1
	s.quotes[id] = &models.Quote{Text: quote, Author: author, ID: id}
	s.pastID++

	return id
}

func (s *Storage) ListQuotes() []models.Quote {
	res := make([]models.Quote, 0, len(s.quotes))

	for _, quote := range s.quotes {
		res = append(res, *quote)
	}

	return res
}

func (s *Storage) RandomQuote() (*models.Quote) {
	quotes := make([]models.Quote, 0, len(s.quotes))

	for _, quote := range s.quotes {
		quotes = append(quotes, *quote)
	}

	res := selectRandom(quotes)
	if res == nil {
		return &models.Quote{}
	}

	return res
}

func selectRandom[T any](arr []T) *T {
	if len(arr) == 0 {
		return nil
	}
    i := rand.Intn(len(arr))
    return &arr[i]
}

func (s *Storage) ListQuotesAuthor(author string) []models.Quote {
	quotes := make([]models.Quote, 0, len(s.quotes))

	for _, quote := range s.quotes {
		if quote.Author == author {
			quotes = append(quotes, *quote)
		}
	}

	return quotes
}

func (s *Storage) Delete(id int64) error {
	if _, ok := s.quotes[id]; !ok {
		return ErrNoQuotes
	}

	delete(s.quotes, id)
	return nil
}