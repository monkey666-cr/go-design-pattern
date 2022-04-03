package decorator

type Component interface {
	Calc() int
}

type ConcreteComponent struct {
}

func (*ConcreteComponent) Calc() int {
	return 0
}

type MulDecorator struct {
	Component
	num int
}

func WrapMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		c,
		num,
	}
}

func (m *MulDecorator) Calc() int {
	return m.Component.Calc() * m.num
}

type AddDecorator struct {
	Component
	num int
}

func WrapAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		c,
		num,
	}
}

func (a *AddDecorator) Calc() int {
	return a.Component.Calc() + a.num
}
