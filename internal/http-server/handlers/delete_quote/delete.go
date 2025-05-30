package deletequote

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"quoter/internal/http-server/response"
	"strconv"
	"strings"
)

type Deteler interface {
	Delete(id int64) error
}

func New(deleter Deteler, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http-server.handlers.delete"

		log = log.With(
			slog.String("op", op),
		)

		w.Header().Set("Content-Type", "application/json")

		id, err := getID(r.URL.String())
		if err != nil {
			log.Error("failed get id", "err", err)

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.Response{Status: response.StatusError, Error: "invalid id"})

			return
		}

		err = deleter.Delete(id)
		if err != nil {
			log.Error("failed delete", "err", err)

			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response.Response{Status: response.StatusError, Error: "quote not found"})

			return
		}

		log.Info("quote deleted", "id", id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.Response{Status: response.StatusOK})
	}
}

func getID(url string) (int64, error) {
	path := strings.Split(url, "/")

	idStr := path[len(path)-1]
	
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return -1, err
	}

	return id, nil
}