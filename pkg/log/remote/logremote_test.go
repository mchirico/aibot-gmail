package remote

import "testing"

func TestLog(t *testing.T) {
	r, err := Log("Test@log", "test")
	if err != nil {
		t.Logf("err: %v\n", err)
	}
	t.Logf("r: %v\n", r)
}
