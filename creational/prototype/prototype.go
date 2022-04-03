package prototype

// 用原型实例指定创建对象的种类，并且通过拷贝这些原型创建新的对象(clone浅拷贝，深拷贝)

type Cloneable interface {
	Clone() Cloneable
}

type Manager struct {
	prototypes map[string]Cloneable
}

func NewManager() *Manager {
	return &Manager{
		prototypes: make(map[string]Cloneable),
	}
}

func (m *Manager) Get(name string) Cloneable {
	return m.prototypes[name].Clone()
}

func (m *Manager) Set(name string, prototype Cloneable) {
	m.prototypes[name] = prototype
}
