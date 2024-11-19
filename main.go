package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	//redis.ExampleClient()
	//config.Init()
	//storage.Init()
	//migrate.Init()
	//go router.Init()
	_, err := fmt.Println("Hello, ithome")
	if err == nil {
		gorace()
	}
}

func gorace() {
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		m["1"] = "a" // First conflicting access.
		c <- true
		fmt.Println("傳送了")
	}()
	fmt.Println("c8 c8 c8")
	m["2"] = "b" // Second conflicting access.
	fmt.Println("等接true", "m:", m)
	<-c
	fmt.Printf("接到了m: %+v\n", m)
	for k, v := range m {
		fmt.Println(k, v)
	}

	// 使用 http://localhost:6060/debug/pprof/ 開啟
	// 如果要用graphviz查看可下指令 go tool pprof http://localhost:6060/debug/pprof/heap
	// 會獲得像是這樣的資料 Saved profile in /Users/shannihsu/pprof/pprof.alloc_objects.alloc_space.inuse_objects.inuse_space.003.pb.gz
	// 再下指令 go tool pprof -http=:8080 /Users/shannihsu/pprof/pprof.alloc_objects.alloc_space.inuse_objects.inuse_space.002.pb.gz
	// 會獲得Serving web UI on http://localhost:8080的回應
	http.ListenAndServe("localhost:6060", nil)
}
