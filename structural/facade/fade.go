package facade

import "fmt"

func NewAPI() API {
	return &apiImpl{
		NewAModuleImpl(),
		NewBModuleAPI(),
	}
}

type API interface {
	Test() string
}

type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (a *apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

type AModuleAPI interface {
	TestA() string
}

type aModuleImpl struct {
}

func NewAModuleImpl() AModuleAPI {
	return &aModuleImpl{}
}

func (a *aModuleImpl) TestA() string {
	return "A module running"
}

type BModuleAPI interface {
	TestB() string
}

type bModuleImpl struct {
}

func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

func (b *bModuleImpl) TestB() string {
	return "B module running"
}
