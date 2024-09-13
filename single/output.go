package single

import (
	"fmt"
)

func Output() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}
	// 如果不加這個，這個線程不會執行完
	// 可以改成 time.Sleep(10 *time.Second) 等他執行完
	fmt.Scanln()
}

/*
Creating single instance now.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
*/
