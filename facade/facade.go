package facade

import "fmt"

// 外觀模式
// API為facade模塊的外觀接口，大部分代碼使用此接口簡化對facade類的訪問。
// facade模塊同時暴露了a和b兩個Module的NewXXX和interface，其他代碼如果需要使用細節功能時可以直接調用。

// https://www.cnblogs.com/jing99/p/12602696.html
// https://medium.com/@ckpattern35/ck-patt-%E8%A8%AD%E8%A8%88%E6%A8%A1%E5%BC%8F-11-%E8%BF%AA%E7%B1%B3%E7%89%B9%E6%B3%95%E5%89%87-demeter-law-931fefc4abda
// 以客戶端來說不需要知道有哪些子系統，透過外觀模式就可訪問各個子系統。 shanni

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

// API is facade interface of facade package
type API interface {
	Test() string
}

// apiImpl facade implement
type apiImpl struct {
	a AModuleAPI // A子系統 shanni
	b BModuleAPI // B子系統 shanni
}

func (a *apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

// A子系統 shanni
// NewAModuleAPI return new AModuleAPI
func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

// AModuleAPI ...
type AModuleAPI interface {
	TestA() string
}

type aModuleImpl struct{}

func (*aModuleImpl) TestA() string {
	return "A module running"
}

// B子系統 shanni
// NewBModuleAPI return new BModuleAPI
func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

// BModuleAPI ...
type BModuleAPI interface {
	TestB() string
}

type bModuleImpl struct{}

func (*bModuleImpl) TestB() string {
	return "B module running"
}
