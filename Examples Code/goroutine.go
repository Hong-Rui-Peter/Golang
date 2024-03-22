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
	wg.Done()
}

func main() {
	rand.Seed(time.Now().UnixNano()) // 先產生隨機值
	peoples := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	wg.Add(10)

	for _, people := range peoples {
		go runner(people)
	}

	wg.Wait()
	fmt.Println("Game Over")
}
