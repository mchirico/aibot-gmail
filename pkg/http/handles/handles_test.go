package handles

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func Test_RootGET(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	BaseRoot(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	if !strings.Contains(string(body), "version:") {
		t.Fatalf("GET on root failed")
	}

}

func Test_RootPUT(t *testing.T) {
	req := httptest.NewRequest("PUT", "/", nil)
	w := httptest.NewRecorder()
	BaseRoot(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	if resp.StatusCode != 200 {
		t.Fatalf("PUT is causing error")
	}

}

func Test_RootPOST(t *testing.T) {
	req := httptest.NewRequest("POST", "/", nil)
	w := httptest.NewRecorder()
	BaseRoot(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	if string(body) != "post" {
		t.Log(string(body))
		t.Fatalf("post response not what expected")
	}

}

func Test_SubScribe(t *testing.T) {

	os.Setenv("PUBSUBTOKEN","1a3a")
	token := os.Getenv("PUBSUBTOKEN")

	req := httptest.NewRequest("GET", "/subscript?key="+token, nil)
	w := httptest.NewRecorder()
	Subscript(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)


	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))


	if !strings.Contains(string(body), "bad key") {
		t.Fatalf("GET on SubScript failed")
	}

}