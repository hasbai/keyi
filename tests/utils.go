package tests

import (
	"bytes"
	"fmt"
	"github.com/goccy/go-json"
	"io/ioutil"
	"keyi/app"
	"keyi/auth"
	. "keyi/models"
	"net/http"
	"strings"
	"sync"
	"testing"

	"github.com/hetiansu5/urlquery"
	"github.com/stretchr/testify/assert"
)

type JsonData interface {
	Map | []Map
}

var App = app.Create()

type Token struct {
	sync.Mutex
	Token string
}

var token Token

func init() {
	generateToken(1)
}

func generateToken(id int) {
	accessToken, _, err := auth.GenerateTokens(&auth.User{
		BaseModel: BaseModel{
			ID: id,
		},
		Permission:   auth.PUser,
		TenantID:     1,
		TenantAreaID: 0,
	})
	if err != nil {
		panic("generate token error, " + err.Error())
	}

	token.Token = "Bearer " + accessToken
}

func setToken(id int) {
	token.Lock()
	generateToken(id)
}

func resetToken() {
	generateToken(1)
	token.Unlock()
}

// testCommon tests status code and returns response body in bytes
func testCommon(t *testing.T, method string, route string, statusCode int, data ...Map) []byte {
	var requestData []byte
	var err error

	if len(data) > 0 && data[0] != nil { // data[0] is request data
		requestData, err = json.Marshal(data[0])
		assert.Nilf(t, err, "encode request body")
	}
	req, err := http.NewRequest(
		strings.ToUpper(method),
		route,
		bytes.NewBuffer(requestData),
	)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token.Token)
	assert.Nilf(t, err, "constructs http request")

	res, err := App.Test(req, -1)
	assert.Nilf(t, err, "perform request")

	responseBody, err := ioutil.ReadAll(res.Body)
	assert.Nilf(t, err, "decode response")
	fmt.Println(string(responseBody))

	assert.Equalf(t, statusCode, res.StatusCode, "status code")

	return responseBody
}

// testCommonQuery tests status code and returns response body in bytes
func testCommonQuery(t *testing.T, method string, route string, statusCode int, data ...Map) []byte {
	var err error
	req, err := http.NewRequest(
		strings.ToUpper(method),
		route,
		nil,
	)
	if len(data) > 0 && data[0] != nil { // data[0] is query data
		queryData, err := urlquery.Marshal(data[0])
		req.URL.RawQuery = string(queryData)
		assert.Nilf(t, err, "encode request body")
	}

	req.Header.Add("Content-Type", "application/json")
	assert.Nilf(t, err, "constructs http request")

	res, err := App.Test(req, -1)
	assert.Nilf(t, err, "perform request")
	assert.Equalf(t, statusCode, res.StatusCode, "status code")

	responseBody, err := ioutil.ReadAll(res.Body)
	assert.Nilf(t, err, "decode response")

	return responseBody
}

// testAPIGeneric inherits testCommon, decodes response body to json, tests whether it's expected
func testAPIGeneric[T JsonData](t *testing.T, method string, route string, statusCode int, data ...Map) T {
	responseBody := testCommon(t, method, route, statusCode, data...)

	if statusCode == 204 { // no content
		return nil
	}
	var responseData T
	err := json.Unmarshal(responseBody, &responseData)
	assert.Nilf(t, err, "decode response")

	if len(data) > 1 { // data[1] is response data
		assert.Equalf(t, data[1], responseData, "response data")
	}

	return responseData
}

// testAPI returns a Map
func testAPI(t *testing.T, method string, route string, statusCode int, data ...Map) Map {
	return testAPIGeneric[Map](t, method, route, statusCode, data...)
}

// testAPIArray returns []Map
func testAPIArray(t *testing.T, method string, route string, statusCode int, data ...Map) []Map {
	return testAPIGeneric[[]Map](t, method, route, statusCode, data...)
}

func testAPIModel[T any](t *testing.T, method string, route string, statusCode int, data ...Map) T {
	var model T
	responseBytes := testCommon(t, method, route, statusCode, data...)
	err := json.Unmarshal(responseBytes, &model)
	assert.Nilf(t, err, "unmarshal response")
	return model
}

func testAPIModelWithQuery[T any](t *testing.T, method string, route string, statusCode int, obj *T, data ...Map) {
	responseBytes := testCommonQuery(t, method, route, statusCode, data...)
	err := json.Unmarshal(responseBytes, obj)
	assert.Nilf(t, err, "unmarshal response")
}
