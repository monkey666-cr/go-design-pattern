package proxy

import "testing"

func TestProxy_Do(t *testing.T) {
	var sub Subject
	sub = &Proxy{}

	res := sub.Do()

	if res != "pre:real:after" {
		t.Fatal()
	}
}
