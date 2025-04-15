package channel

import (
	"fmt"
	"sync"
)

var ch = make(chan int)

func Channel() {

	go handle()

	// 讀取channel方式，要用for range
	for v := range ch {
		fmt.Println(v)
	}
}

func handle() {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

var cc = make(chan int)

func Ch() {
	go chHandle()
	// 讀取channel方式，要用for、v, ok := <-cc
	for {
		v, ok := <-cc
		if !ok {
			return
		}
		fmt.Println(v)
	}
}

func chHandle() {
	for i := 0; i < 10; i++ {
		cc <- i
	}
	close(cc)
}

// https://blog.wu-boy.com/2022/05/read-data-from-channel-in-go/
// 練習使用兩個goroutine取值跑資料
func Foobar() {
	str := []byte("foobar")
	chx := make(chan byte, len(str))
	next := make(chan struct{})
	wg := &sync.WaitGroup{}
	wg.Add(2)

	for i := 0; i < len(str); i++ {
		chx <- str[i]
	}
	close(chx)

	go func() {
		defer wg.Done()

		for {
			stop, ok := <-next
			if !ok {
				return
			}
			v, ok := <-chx
			if !ok {
				close(next)
				return
			}
			fmt.Println("goroutine01: ", string(v))
			next <- stop
		}

	}()

	go func() {

		defer wg.Done()

		for {
			stop, ok := <-next
			if !ok {
				return
			}
			v, ok := <-chx
			if !ok {
				close(next)
				return
			}
			fmt.Println("goroutine02: ", string(v))
			next <- stop
		}

	}()
	next <- struct{}{}
	wg.Wait()
}

type Message struct {
	User    string `json:"user"`
	Message string `json:"message"`
}

var broadcast = make(chan Message)

func MessageSend() {
	msg := Message{
		User:    "Shanni",
		Message: "Hello",
	}

	go func() {

		broadcast <- msg

		close(broadcast)
	}()

	for {
		v, ok := <-broadcast
		if !ok {
			return
		}
		fmt.Println("v:", v)
	}
}

// 使用兩個goroutine接值
func TwoGo() {
	testStr := []byte("foobar")
	receiveChan := make(chan byte, len(testStr))
	stopChan := make(chan struct{})
	wg := &sync.WaitGroup{}
	wg.Add(2)

	for i := 0; i < len(testStr); i++ {
		receiveChan <- testStr[i]
	}
	close(receiveChan)

	go func() {
		defer wg.Done()

		for {
			stop1, ok := <-stopChan
			if !ok {
				return
			}
			v, ok := <-receiveChan
			if !ok {
				return
			}
			fmt.Println("goroutine01:", string(v))
			stopChan <- stop1
		}
	}()

	go func() {
		defer wg.Done()

		for {
			stop2, ok := <-stopChan
			if !ok {
				return
			}
			v, ok := <-receiveChan
			if !ok {
				return
			}
			fmt.Println("goroutine02:", string(v))
			stopChan <- stop2
		}
	}()
	stopChan <- struct{}{}
	wg.Wait()
}

func WorkPool() {
	fmt.Println("進入WorkPool")
	// 題目: 將一組整數作為「任務」，每個任務的工作是: 把這個整數平方
	jobs := []int{2, 3, 4, 5, 6}
	numWorkers := 3
	jobsCh := make(chan int, len(jobs))
	resultCh := make(chan int, len(jobs))

	// 開固定數量的goroutine (worker goroutine)
	wg := &sync.WaitGroup{}
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		// 這邊傳的id對於閉包來說不是數值，而是記憶體位置
		// 對於閉包來說，如果在迴圈內，會捕捉變數i的記憶體位置，而i在迴圈內會一直變動，當goroutine真的執行的時候，i已經是變成最後的值(也就是最後一次記憶體的位置)
		go func(id int) {
			defer wg.Done() // 用來通知主程式: 這個worker處理完了
			for resultJob := range jobsCh {
				resultCh <- resultJob * resultJob
			}
		}(i)
	}

	// 主goroutine丟任務
	for _, job := range jobs {
		jobsCh <- job
	}
	close(jobsCh) // 所有任務送出後關閉，這邊通知所有worker:「任務都送完了」
	// 這邊 worker goroutine的 for resultJob := range jobsCh就會自動結束

	// 使用額外的goroutine監控所有worker是否完成(透過WaitGroup)
	go func() {
		wg.Wait() // 等上面的goroutine結束
		close(resultCh)
	}()

	// 這邊for range會一直等待channel被關閉，如果沒被關閉會被卡死
	for result := range resultCh {
		fmt.Println("result:", result)
	}
}

/*
補充: 什麼時機使用close channel
不是用完一定要關，而是有些情況要關閉，不然程式會卡死、漏資料，或goroutine leak
＊＊＊你是寫入端(Sender)，而且
1. 沒有人再寫入該channel了
2. 有人會用 range channel或v, ok := <-channel來讀取channel

如果這時候不關channel:
range channel永遠不會結束 -> 卡住
v, ok := <-channel的ok永遠是true -> 無法判斷結束

使用close channel
我寫入，我負責關；
只有我寫，才可以關；
要 range 的話，一定得關。
*/
