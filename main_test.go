package main

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostAuth(t *testing.T) {
	app := InitFiber()

	msg := `{"data":"true"}`

	req, _ := http.NewRequest(
		"POST",
		"/auth",
		strings.NewReader(msg),
	)

	res, err := app.Test(req)
	assert.Equal(t, nil, err)

	body, err := io.ReadAll(res.Body)
	assert.Equal(t, nil, err)

	println(string(body))
}
