package facotry

// SubOperatorFactory 减法工厂
type SubOperatorFactory struct {
}

func (m *SubOperatorFactory) Create() Operator {
	return &SubOperator{
		&OperatorBase{},
	}
}

// SubOperator 减法操作
type SubOperator struct {
	*OperatorBase
}

func (m *SubOperator) Result() int {
	return m.a - m.b
}
