// 牌七(撲克牌遊戲)
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	var player int = 4                // 4位玩家
	var startNum int = 7              // 初始牌號
	var minNum int = 1                // 撲克牌最小號碼
	var maxNum int = 13               // 撲克牌最大號碼
	ch := make([]chan string, player) //玩家通道

	wg.Add(player)

	for i := 0; i < player; i++ {
		// 隨機初始牌
		// rand.Seed(time.Now().UnixNano())
		// startNum = rand.Intn(maxNum-minNum) + minNum
		ch[i] = make(chan string) // 創建每个玩家的通道
		go gameStart(maxNum, minNum, startNum, "玩家"+strconv.Itoa(i+1), ch[i])
	}

	i := 0

	for {
		select {
		case msg := <-ch[0]:
			fmt.Println("=======>", msg)
			i++
		case msg := <-ch[1]:
			fmt.Println("=======>", msg)
			i++
		case msg := <-ch[2]:
			fmt.Println("=======>", msg)
			i++
		case msg := <-ch[3]:
			fmt.Println("=======>", msg)
			i++
		case <-time.After(5 * time.Second): // 超時
			fmt.Println("Timed out", i)
		}

		if i == 4 { // 4位玩家
			break
		}
	}
	wg.Wait()

}

func gameStart(max, min, startNum int, player string, ch chan string) {
	// allList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	var allList = make([]int, max+min+1)

	for i := 0; i < max+min+1; i++ {
		allList[i] = i + 1
	}

	pos := valuePosition(startNum, allList...)
	allList = append(allList[:pos], allList[pos+1:]...)
	nStart := startNum - 1
	nEnd := startNum + 1
	imax := len(allList)

	// 計算排完
	for len(allList) > 0 {
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(imax)
		rNum := allList[index]

		if rNum == nStart {
			allList = append(allList[:index], allList[index+1:]...)
			imax--
			fmt.Println(player, "Get", nStart)
			nStart = nStart - 1
		} else if rNum == nEnd {
			allList = append(allList[:index], allList[index+1:]...)
			imax--
			fmt.Println(player, "Get", nEnd)
			nEnd = nEnd + 1
		}

		// 先到 1 或 13 贏
		if nStart+1 == min || nEnd-1 == max {
			// fmt.Println("========", player, nStart, nEnd, "========")
			break
		}
	}

	ch <- "Player " + player + " finished"
	wg.Done()
}

func valuePosition(targetValue int, arr ...int) int {
	// var position int
	for i, value := range arr {
		if value == targetValue {
			// position = i
			// return position
			return i
			// break // 找到后可以提前退出循环
		}
	}
	return -1
}
