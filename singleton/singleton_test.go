package singleton

import (
	"sync"
	"testing"
)

const parCount = 100

func TestSingleton(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}

// https://blog.kennycoder.io/2020/12/18/Golang%E6%95%99%E5%AD%B8%E7%B3%BB%E5%88%97-%E4%BD%95%E8%AC%82WaitGroup-%E7%AD%89%E5%BE%85Goroutine%E7%9A%84%E5%A5%BD%E5%B9%AB%E6%89%8B/
// https://blog.csdn.net/lixora/article/details/129734624
// goRoutine相關資料 shanni
func TestParallelSingleton(t *testing.T) {
	start := make(chan struct{})
	wg := sync.WaitGroup{}
	// 代表等待Goroutine的數量 shanni
	wg.Add(parCount)
	instances := [parCount]Singleton{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			//協程阻塞，等待channel被關閉才能繼續運行
			<-start
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	// 關閉channel，所有協程同時開始運行，實現并行 (parallel)
	close(start)
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}
