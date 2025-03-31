package channel

import "fmt"

/*
無緩衝通道的基本操作

建立一個無緩衝的 channel，名稱為 ch。
啟動一個 goroutine，該 goroutine 會向 ch 發送數字 100。
在 main 函數中接收該數字並打印出來。
*/
func TestOne() {
	var intChan = make(chan int)
	go func() {
		intChan <- 100
	}()
	x := <-intChan
	fmt.Println(x)
}

/*
緩衝通道的行為

建立一個緩衝大小為 2 的 channel，名稱為 ch。
在 main 函數中，先後向 ch 發送 10 和 20，然後接收並打印這兩個數字。
試著向 ch 發送第三個數字 30，看看會發生什麼。
*/

func TestTwo() {
	// 符合先進先出
	var intChan = make(chan int, 2)
	intChan <- 10
	intChan <- 20
	//intChan <- 30 // 如果無人讀取，這邊會deadlock；fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-intChan) // 10
	fmt.Println(<-intChan) //20
}

/*
通道的關閉與遍歷

建立一個 channel，名稱為 ch，它傳遞 int。
啟動一個 goroutine，該 goroutine 會依序向 ch 發送 1、2、3，然後關閉 ch。
在 main 函數中，使用 for range 遍歷 ch 並打印所有收到的數字。
*/
func TestThree() {
	var ch = make(chan int)
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println("TestThree: ", v)
	}
}
