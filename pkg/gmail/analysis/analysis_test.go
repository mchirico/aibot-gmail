package analysis

import (
	"fmt"
	"strings"
	"testing"
	"time"
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

func TestAnalysisFirebase(t *testing.T) {
	now := time.Now()
	m, _ := AnalysisFirebase(now.Add(-24 * 2 * time.Hour))
	s, _ := Filter2orMore(m)
	good, _ := FilterJunkString(s)
	t.Logf("result: %v\n", good)
}
