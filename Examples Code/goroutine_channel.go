// 併發-通道(channel)應用
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 建立跑者&動作
func runner2(roadNo int, runnerNo int, ch chan int) {
	if runnerNo != 4 { // 判斷跑者是否為最後一位
		nextCh := make(chan int) // 建立下一位跑者的通道

		go runner2(roadNo, runnerNo+1, nextCh) // 同跑道下位跑者
		<-nextCh                               // 等待下位跑者傳入通道後呼叫
	}

	counT := 1000000

	for counT > 0 {
		counT -= rand.Intn(10)
	}

	fmt.Println("第" + strconv.Itoa(roadNo) + "跑道 - 第" + strconv.Itoa(runnerNo) + "號")
	ch <- runnerNo // ?號跑者傳入通道

}

// 建立跑道
func road(roadNo int) {
	ch := make(chan int)

	go runner2(roadNo, 1, ch) // 每位跑道先建立1號跑者
	<-ch                      // 等待建立的跑者都傳入通道後，呼叫，往下執行
	fmt.Println("第" + strconv.Itoa(roadNo) + "跑道" + " - finish")
	wg.Done()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	wg.Add(5) // 產生5個goroutine

	for i := 1; i <= 5; i = i + 1 {
		go road(i)
	}
	wg.Wait() // 5個goroutine跑完呼叫完在往下執行
	fmt.Println("Game Over")
}
