package replies

import (
	"fmt"
	"testing"
)

func TestReplies_Get(t *testing.T) {
	r := &Replies{}
	result := r.Get(0)
	if !(len(result) > 3) {
		t.Fatalf("Bad text")
	}
}

func TestReplies_Next(t *testing.T) {
	r := &Replies{}
	fmt.Println(r.Next())
	fmt.Println(r.Next())
	fmt.Println(r.Next())
	fmt.Println(r.Next())
}
