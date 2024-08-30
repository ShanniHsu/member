package singleton

import "sync"

// 單例模式
// 使用懶惰模式的單例模式，使用雙重檢查加鎖保證線程安全

// singleton 是單例模式接口，導出的
// 通過接口可以避免 GetInstance返回一個包私有類型的指針
type Singleton interface {
	foo()
}

// singleton 是單例模式類，包私有的
type singleton struct{}

func (s singleton) foo() {}

var (
	instance *singleton
	once     sync.Once
)

// GetInstance 用於獲取單例模式對象
func GetInstance() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
