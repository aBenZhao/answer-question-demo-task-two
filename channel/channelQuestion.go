package channel

import (
	"fmt"
	"sync"
	"time"
)

// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。
func UnbufferedChannelQuestion() {
	var unbufferedChan chan int = make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			unbufferedChan <- i

		}
		close(unbufferedChan)
	}()

	go func() {
		for receiveValue := range unbufferedChan {
			fmt.Println(receiveValue)
		}
	}()

	time.Sleep(100 * time.Millisecond)

}

//题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
//考察点 ：通道的缓冲机制。

func BufferedChannelQuestion() {
	var unbufferedChan chan int = make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 1; i <= 100; i++ {
			unbufferedChan <- i

		}
		defer wg.Done()
		close(unbufferedChan)
	}()

	go func() {
		for receiveValue := range unbufferedChan {
			fmt.Println(receiveValue)
		}
		defer wg.Done()
	}()

	wg.Wait()

}
