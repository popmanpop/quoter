package save

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	resp "quoter/internal/http-server/response"
)

type Request struct {
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

type Saver interface {
	Save(quote string, author string) int64
}

func New(log *slog.Logger, storage Saver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http-server.handlers.save.New"

		log = log.With(
			slog.String("op", op),
		)

		var req Request

		w.Header().Set("Content-Type", "application/json")

		data, err := io.ReadAll(r.Body)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Error("request body is empty")

				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(
					resp.Response{
					Status: resp.StatusError,
					Error: "empty request",
				})

            	return
			}

			log.Error("failed read request body", "err", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(
				resp.Response{
				Status: resp.StatusError,
				Error: "internal error",
			})

			return
		}

		err = json.Unmarshal(data, &req)
		if err != nil {
			log.Error("failed unmarshal request", "err", err)

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(
				resp.Response{
					Status: resp.StatusError,
					Error: "bad request",
				},
			)

			return
		}

		log.Info("saved new quote")
		_ = storage.Save(req.Quote, req.Author)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(
			resp.Response{
				Status: resp.StatusOK,
			},
		)
	}
}