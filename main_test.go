package main

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsicsList(t *testing.T) {
	app := Setup()
	assert := assert.New(t)

	req, _ := http.NewRequest("GET", "/api/v1/asics/", nil)
	res, err := app.Test(req, -1)

	assert.Equalf(200, res.StatusCode, "")
	body, err := io.ReadAll(res.Body)
	assert.Nilf(err, "Response has an error")

	assert.JSONEq(`{"asics":["asics1","asics2","asics3"],"success":true}`, string(body), "Response is different from expected")
}

func TestStatistics(t *testing.T) {
	app := Setup()
	assert := assert.New(t)

	req, _ := http.NewRequest("GET", "/api/v1/statistics/", nil)
	res, err := app.Test(req, -1)

	assert.Equalf(200, res.StatusCode, "")
	body, err := io.ReadAll(res.Body)
	assert.Nilf(err, "Response has an error")

	assert.JSONEq(`{"statistics":["1","2","3"],"success":true}`, string(body), "Response is different from expected")
}
