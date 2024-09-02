package singleton

import "sync"

// 單例模式
// 使用懶惰模式的單例模式，使用雙重檢查加鎖保證線程安全

// singleton 是單例模式接口，導出的
// 通過接口可以避免 GetInstance返回一個包私有類型的指針

// https://huashen87.medium.com/%E5%96%AE%E4%BE%8B%E6%A8%A1%E5%BC%8F-singleton-pattern-5e09dc6d11e7
// https://geektutu.com/post/hpg-sync-once.html
// sync.Once是Go標準庫提供的使函數只執行一次的實現，常應用於單例模式，例如初始化配置、保持數據庫連接等。 shanni
// sync.Once作用與Init函數類似，但有區別，請參考連結二 shanni
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
