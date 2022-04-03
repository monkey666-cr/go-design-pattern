package builder

type Builder interface {
	Part1()
	Part2()
	Part3()
}

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) Construct() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}

// CarBuilder 汽车构造者
type CarBuilder struct {
	result string
}

func (c *CarBuilder) Part1() {
	c.result += "1"
}

func (c *CarBuilder) Part2() {
	c.result += "2"
}

func (c *CarBuilder) Part3() {
	c.result += "3"
}

func (c *CarBuilder) GetResult() string {
	return c.result
}

// HouseBuilder 房屋建造者
type HouseBuilder struct {
	result string
}

func (h *HouseBuilder) Part1() {
	h.result += "A"
}

func (h *HouseBuilder) Part2() {
	h.result += "B"
}

func (h *HouseBuilder) Part3() {
	h.result += "C"
}

func (h *HouseBuilder) GetResult() string {
	return h.result
}
