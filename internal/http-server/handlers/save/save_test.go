package save_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	save "quoter/internal/http-server/handlers/save"
	"quoter/internal/http-server/response"
	"quoter/internal/storage"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeleteHandler(t *testing.T) {

	cases := []struct {
		name string
		author string
		quote string
	} {
		{
			name: "Fail",
			author: "Confucius",
			quote: "Life is simple, but we insist on making it complicated.",
		},
		{
			name: "Fail #2",
			author: "Gerodot",
			quote: "We insit on it.",
		},
	}

for _, tc := range cases {
	t.Run(tc.name, func(t *testing.T){
		handler := save.New(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})), storage.New())

		input := fmt.Sprintf(`{"author":"%s", "quote":"%s"}`, tc.author, tc.quote)

		req, err := http.NewRequest(http.MethodPost, "/quotes", bytes.NewReader([]byte(input)))
        require.NoError(t, err)

		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		require.Equal(t, rr.Code, http.StatusOK)

		body := rr.Body.String()

        var resp response.Response
        require.NoError(t, json.Unmarshal([]byte(body), &resp))

		require.NotEmpty(t, resp)
	})
}
}