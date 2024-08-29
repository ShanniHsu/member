package simpleFactory

import "fmt"

// 簡單工廠模式
// go語言沒有構造函數一說，一般會定義NewXXX函數來初始化相關類。
// NewXXX函數返回接口時就是簡單工廠模式，也就是說Golang的一般推薦做法就是簡單工廠
// 在這個simplefactory包中只有API接口和NewAPI函數為包外可見，封裝了實現細節。

// API is interface
// 這個對簡單工廠模式來說是產品(Product) shanni
type API interface {
	// 告訴別人你要用哪一種打招呼方式
	Say(name string) string
}

// NewAPI return API instance by type
// 這個對簡單工廠模式來說是工廠(Factory) shanni
func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

/*有Hi這個方法 Class1*/
// hiAPI is one of API implement
// 這個對簡單工廠模式來說是產品(Product) shanni
type hiAPI struct{}

// Say hi to name
func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

/*有Hello這個方法 Class2*/
// helloAPI is another API implement
// 這個對簡單工廠模式來說是產品(Product) shanni
type helloAPI struct{}

// Say hello to name
func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
