package list_test

import (
	"net/http"
	"net/http/httptest"
	"quoter/internal/http-server/handlers/list"
	"quoter/internal/storage"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	
	handler := list.New(storage.New())

	req, err := http.NewRequest(http.MethodGet, "/quotes", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	require.Equal(t, rr.Code, http.StatusOK)
}