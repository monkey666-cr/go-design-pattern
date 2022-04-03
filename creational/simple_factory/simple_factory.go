package simple_factory

type API interface {
	Say(name string) string
}

func NewAPI(str string) API {
	if str == "en" {
		return &English{}
	} else if str == "cn" {
		return &Chinese{}
	} else if str == "ja" {
		return &Japanese{}
	} else {
		return nil
	}
}

type Chinese struct {
}

func (c *Chinese) Say(name string) string {
	return "你好" + name
}

type English struct {
}

func (e *English) Say(name string) string {
	return "hello" + name
}

type Japanese struct {
}

func (j *Japanese) Say(name string) string {
	return "你好, 鬼子" + name
}
