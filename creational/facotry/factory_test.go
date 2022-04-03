package facotry

import "testing"

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestPlusOperatorFactory_Create(t *testing.T) {
	factory := &PlusOperatorFactory{}
	if compute(factory, 1, 2) != 3 {
		t.Fatal("error with factory method pattern")
	}
}

func TestSubOperatorFactory_Create(t *testing.T) {
	factory := &SubOperatorFactory{}
	if compute(factory, 2, 1) != 1 {
		t.Fatal("error with factory method pattern")
	}
}
