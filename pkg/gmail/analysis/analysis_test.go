package analysis

import (
	"fmt"
	"testing"
)

func TestAnalysis(t *testing.T) {
	r, err := Analysis()
	if err != nil {
		t.Fatal(err,r)
	}
	for i,v := range r {
		fmt.Println(i,v)
	}
}