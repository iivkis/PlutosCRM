package httpctrl1

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHelloWorld_Welcome(t *testing.T) {
	e, w := engine(), httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/api/v1/welcome", nil)

	e.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)
}
