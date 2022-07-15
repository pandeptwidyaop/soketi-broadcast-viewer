package main

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostAuth(t *testing.T) {
	app := InitFiber()

	req, _ := http.NewRequest(
		"POST",
		"/auth",
		nil,
	)

	res, err := app.Test(req)
	assert.Equal(t, nil, err)

	body, err := io.ReadAll(res.Body)
	assert.Equal(t, nil, err)

	println(string(body))
}
