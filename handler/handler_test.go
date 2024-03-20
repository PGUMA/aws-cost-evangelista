package handler

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleRequest(t *testing.T) {
	ctx := context.Background()

	// Test for the default environment
	resp, err := HandleRequest(ctx)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "Hello, this is the default environment.", resp.Body)

	t.Setenv("APP_ENV", "prd")
	resp, err = HandleRequest(ctx)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "Hello, this is the prd environment.", resp.Body)
}
