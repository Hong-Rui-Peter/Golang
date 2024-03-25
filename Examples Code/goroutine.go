// 併發
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var no int

func runner(people string) {
	time.Sleep(1 * time.Second) // 使系統平均分配執行緒
	counT := 1000000

	for counT > 0 {
		counT -= rand.Intn(10)
	}

	no += 1 // 名次
	fmt.Printf("第%d名:%s\n", no, people)
	wg.Done() // 跑完呼叫
}

func main() {
	rand.Seed(time.Now().UnixNano()) // 先產生隨機值
	peoples := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	wg.Add(10) // 產生10個goroutine

	for _, people := range peoples {
		go runner(people)
	}

	wg.Wait() // 10個goroutine跑完呼叫完在往下執行
	fmt.Println("Game Over")
}
