package facotry

type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) SetA(number int) {
	o.a = number
}

func (o *OperatorBase) SetB(number int) {
	o.b = number
}
