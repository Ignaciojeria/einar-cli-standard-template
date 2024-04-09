package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"archetype/app/configuration"
	"archetype/app/infrastructure/server"

	"github.com/labstack/echo/v4"
)

func TestNewTemplateGet(t *testing.T) {
	e := echo.New()
	wrapper := server.NewEchoWrapper(e, configuration.Conf{PORT: "8080"})

	newTemplateGet(wrapper)

	req := httptest.NewRequest(http.MethodGet, "/insert-your-custom-pattern-here", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	var resp struct {
		Message string `json:"message"`
	}

	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	expectedMessage := "Unimplemented"
	if resp.Message != expectedMessage {
		t.Errorf("expected message %q, got %q", expectedMessage, resp.Message)
	}
}