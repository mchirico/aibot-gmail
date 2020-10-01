package analysis

import (
	"fmt"
	"strings"
	"testing"
)

func TestAnalysis(t *testing.T) {
	r, err := Analysis(40, "TRASH")
	if err != nil {
		t.Fatal(err, r)
	}
	for idx, email := range r {

		for k, v := range email {
			if 1 == 2 {
				fmt.Println(idx)
			}
			//fmt.Println(k,v)
			lower := strings.ToLower(k)
			if strings.Contains(lower, "authenticat") {
				if strings.Contains(v, "dkim=") {
					fmt.Println(k, v)
				}

			}
		}
	}
}
