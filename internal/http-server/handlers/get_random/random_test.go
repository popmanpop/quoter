package getrandom_test

import (
	"net/http"
	"net/http/httptest"
	getrandom "quoter/internal/http-server/handlers/get_random"
	"quoter/internal/storage"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	
	handler := getrandom.New(storage.New())

	req, err := http.NewRequest(http.MethodGet, "/quotes/random", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	require.Equal(t, rr.Code, http.StatusOK)
}