package httputils

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	url := "https://tasks.cwxstat.io/gmail"

	key := "X-API-Key"
	value := "1vtlJSvzaaB6bTjJKzyakYnjnxrRzM22Ex3j2SDR"

	h := NewHTTP()
	h.Header(key, value)

	h.Header("Email", "bozo@b.io")
	h.Header("Value", "SomeStuff")

	r, err := h.Get(url)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}

	fmt.Println(string(r))

}
