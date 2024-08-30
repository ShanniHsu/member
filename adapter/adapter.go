package adapter

// 適配器模式用於轉換一種接口適配另一個接口。
// 實際使用中Adaptee一般為接口，並且使用工廠函數生成實例。
// 在Adapter中匿名組合Adaptee接口，所以Adapter類也擁有SpecificRequest實例方法，
// 又因為Go語言中非入侵式接口特徵，其實Adapter也適配Adaptee接口。

// https://medium.com/bucketing/structural-patterns-adapter-pattern-d7889417cff

// Target 是適配的目標接口
// 要轉換的目標物件 shanni
type Target interface {
	Request() string
}

// Adaptee 是被適配的目標接口
// 要被轉換的物件 shanni
type Adaptee interface {
	SpecificRequest() string
}

// NewAdaptee 是被適配接口的工廠函數
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

// AdapteeImpl 是被適配的目標類
type adapteeImpl struct{}

// SpecificRequest 是目標類的一個方法
func (*adapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

// NewAdapter 是轉換Adapter的工廠函數
func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}

// Adapter 是轉換Adaptee為Target接口的適配器
type adapter struct {
	Adaptee
}

// Request 實現Target接口
func (a *adapter) Request() string {
	return a.SpecificRequest()
}
