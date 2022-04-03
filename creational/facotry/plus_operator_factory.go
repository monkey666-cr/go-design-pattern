package facotry

// PlusOperatorFactory 加法工厂
type PlusOperatorFactory struct {
}

func (p *PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		&OperatorBase{},
	}
}

// PlusOperator 加法操作
type PlusOperator struct {
	*OperatorBase
}

func (p *PlusOperator) Result() int {
	return p.a + p.b
}
