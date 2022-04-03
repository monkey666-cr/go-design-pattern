package singleton

import "testing"

func TestNew(t *testing.T) {
	s := New()
	s["this"] = "that"

	s2 := New()

	if s["this"] != s2["this"] {
		t.Fatal("create singleton instance failed.")
	}
}
