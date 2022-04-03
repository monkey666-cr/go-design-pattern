package builder

import "testing"

func TestCarBuilder(t *testing.T) {
	builder := &CarBuilder{}
	director := NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	if res != "123" {
		t.Fatalf("CarBuilder fail expect 123 acture %s", res)
	}
}

func TestHouseBuilder(t *testing.T) {
	builder := &HouseBuilder{}
	director := NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	if res != "ABC" {
		t.Fatalf("CarBuilder fail expect 123 acture %s", res)
	}
}
