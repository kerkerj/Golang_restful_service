package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bmizerany/assert"
)

var (
	server *httptest.Server
	api    Api
)

func setupMartini() {
	api = &ApiMartini{}
	server = httptest.NewServer(api.Handler())
}

func tearDown() {
	server.Close()
}

func makeRequest(method, url string, body io.Reader) (resp *http.Response, err error) {
	setupMartini()
	defer tearDown()

	url = fmt.Sprintf("%s%s", server.URL, url)

	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("API-KEY", "secret123")
	resp, err = http.DefaultClient.Do(req)

	return
}

func TestPageGetALL(t *testing.T) {
	method, url := "GET", fmt.Sprintf("/page/")
	resp, err := makeRequest(method, url, nil)

	assert.Equal(t, nil, err)
	assert.Equal(t, 200, resp.StatusCode)

	var m map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&m)

	fmt.Println("\nResponse:")
	fmt.Printf("%v\n", m)
	fmt.Printf("\n")
}

func TestPageGet(t *testing.T) {
	num := 1
	method, url := "GET", fmt.Sprintf("/page/%v", num)
	resp, err := makeRequest(method, url, nil)

	assert.Equal(t, nil, err)
	assert.Equal(t, 200, resp.StatusCode)

	var m map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&m)

	fmt.Println("\nResponse:")
	fmt.Printf("%v\n", m)
	fmt.Printf("\n")
}

func TestPagePost(t *testing.T) {

	// Test data
	mcPostBody := map[string]interface{}{
		"test1": "test_data1",
		"test2": "test_data2",
		"test3": "test_data3",
	}
	body, _ := json.Marshal(mcPostBody)

	// Make request
	resp, err := makeRequest("POST", "/page/", bytes.NewReader(body))

	// Check
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, resp.StatusCode)

	var m map[string]interface{}
	//hah, _ := ioutil.ReadAll(resp.Body)
	err = json.NewDecoder(resp.Body).Decode(&m)

	fmt.Println("\nResponse:")
	fmt.Printf("%v", m)
	fmt.Println("\n")
}

func TestPagePut(t *testing.T) {

	// Test data
	mcPostBody := map[string]interface{}{
		"test1": "update",
		"test2": "update1",
		"test3": "update2",
	}
	body, _ := json.Marshal(mcPostBody)

	// Make request
	num := 1
	url := fmt.Sprintf("/page/%v", num)
	resp, err := makeRequest("PUT", url, bytes.NewReader(body))

	// Check
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, resp.StatusCode)

	var m map[string]interface{}
	//hah, _ := ioutil.ReadAll(resp.Body)
	err = json.NewDecoder(resp.Body).Decode(&m)

	fmt.Println("\nResponse:")
	fmt.Printf("%v", m)
	fmt.Println("\n")
}

func TestPageDelete(t *testing.T) {

	// Make request
	num := 1
	url := fmt.Sprintf("/page/%v", num)
	resp, err := makeRequest("DELETE", url, nil)

	// Check
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, resp.StatusCode)

	var m map[string]interface{}
	//hah, _ := ioutil.ReadAll(resp.Body)
	err = json.NewDecoder(resp.Body).Decode(&m)

	fmt.Println("\nResponse:")
	fmt.Printf("%v", m)
	fmt.Println("\n")
}

func TestPageInvalid(t *testing.T) {
	resp, err := makeRequest("GET", "/invalid", nil)

	assert.Equal(t, nil, err)
	assert.Equal(t, 404, resp.StatusCode)
}
