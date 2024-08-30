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

func TestParallelSingleton(t *testing.T) {
	start := make(chan struct{})
	wg := sync.WaitGroup{}
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
	// 關閉channel，所有攜程同時開始進行，實現并行 (parallel)
	close(start)
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}
