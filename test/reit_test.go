package test

import (
	"../route"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouterReitWithSuccess(t *testing.T) {
	r := route.Init()
	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/api/reit")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 200, res.StatusCode, "OK response is expected")
}
