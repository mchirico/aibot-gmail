package filters

import "testing"

func TestIgnoreEmail(t *testing.T) {
	s := "job-search@indeed.com"
	if !IgnoreEmail(s) {
		t.Fatalf("should be false")
	}
	s = "good@email.com"
	if IgnoreEmail(s) {
		t.Fatalf("should be true")
	}
}

func TestIsEmail(t *testing.T) {
	e := "bozo@gmail.com"
	if !IsEmail(e) {
		t.Fatalf("should be valid")
	}
}
