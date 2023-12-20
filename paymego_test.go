package paymego

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCardsCheck_Success(t *testing.T) {

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"jsonrpc":"2.0","id":"123","result":{"receipt":{"_id":"abc123"},"card":{"number":"****"}}}`))
	}))
	defer mockServer.Close()

	api := SubscribeAPI{
		baseURL:    mockServer.URL,
		timeout:    5 * time.Second,
		httpClient: http.Client{},
	}

	resp, err := api.CardsCheck(context.Background(), "123", "token123")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "abc123", resp.Result.Receipt.ID)
	assert.Equal(t, "****", resp.Result.Card.Number)
}

func TestCardsRemove_Success(t *testing.T) {

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"jsonrpc":"2.0","id":"123","result":{"receipt":{"_id":"abc123"},"card":{"number":"****"}}}`))
	}))
	defer mockServer.Close()

	api := SubscribeAPI{
		baseURL:    mockServer.URL,
		timeout:    5 * time.Second,
		httpClient: http.Client{},
	}

	resp, err := api.CardsRemove(context.Background(), "123", "token123")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "abc123", resp.Result.Receipt.ID)
	assert.Equal(t, "****", resp.Result.Card.Number)
}
