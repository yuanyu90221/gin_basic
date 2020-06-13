package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuanyu90221/gin_basic/router"
)

func TestIndexGetRouter(t *testing.T) {
	ginRouter := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	ginRouter.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message":"hello gin"}`, w.Body.String())
}
