package headertrack

import (
	"fmt"
	"testing"
)

func TestSM_Found(t *testing.T) {
	h := map[string]string{}
	h["Snippet"] = "snip"
	ht := NewSM()
	ht.Found(h)
	r, err := ht.GetR()
	fmt.Println(r, err)
}
