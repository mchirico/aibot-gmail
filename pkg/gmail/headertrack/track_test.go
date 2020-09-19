package headertrack

import (
	"testing"
)

func TestSM_Found(t *testing.T) {
	h := map[string]string{}
	h["Snippet"] = "snip"
	ht := NewSM()
	ht.Found(h)
}
