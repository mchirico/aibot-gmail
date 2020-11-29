package httputils

import (
	"errors"
	"io/ioutil"
	"net/http"
	"sync"
)

var debug = false

// SetDebug sets the debug status
// Setting this to true causes the panics to be thrown and logged onto the console.
// Setting this to false causes the errors to be saved in the Error field in the returned struct.
func SetDebug(d bool) {
	debug = d
}

// Header sets a new HTTP header
func (h *HTTP) Header(n string, v string) {
	h.Mutex.Lock()
	defer h.Mutex.Unlock()
	h.Headers[n] = v
}

func (h *HTTP) Cookie(n string, v string) {
	h.Mutex.Lock()
	defer h.Mutex.Unlock()
	h.Cookies[n] = v
}

func (h *HTTP) Get(url string) ([]byte, error) {
	h.Mutex.Lock()
	h.Mutex.Unlock()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		if debug {
			panic("Couldn't perform GET request to " + url)
		}
		return nil, errors.New("couldn't perform GET request to " + url)
	}
	// Set headers
	for hName, hValue := range h.Headers {
		req.Header.Set(hName, hValue)
	}
	// Set cookies
	for cName, cValue := range h.Cookies {
		req.AddCookie(&http.Cookie{
			Name:  cName,
			Value: cValue,
		})
	}
	// Perform request
	resp, err := h.client.Do(req)
	if err != nil {
		if debug {
			panic("Couldn't perform GET request to " + url)
		}
		return nil, errors.New("couldn't perform GET request to " + url)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if debug {
			panic("Unable to read the response body")
		}
		return nil, errors.New("unable to read the response body")
	}
	return bytes, nil
}

type HTTP struct {
	client  *http.Client
	Headers map[string]string
	Cookies map[string]string
	sync.Mutex
}

func NewHTTP() *HTTP {
	h := &HTTP{}
	h.client = &http.Client{}
	h.Headers = map[string]string{}
	h.Cookies = map[string]string{}
	return h
}
