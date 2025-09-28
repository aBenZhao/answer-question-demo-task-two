package sync

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。
func SyncMutexQuestion() {
	var wg sync.WaitGroup
	mutex := sync.Mutex{}
	num := 0
	wg.Add(1)
	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()
			num++
			mutex.Unlock()
		}
		defer wg.Done()
	}()
	wg.Wait()
	fmt.Println(num)
}

// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。
func UnSyncMutexQuestion() {
	var wg sync.WaitGroup
	var num int64 = 0
	wg.Add(1)
	go func() {
		for i := 0; i < 1000; i++ {
			atomic.AddInt64(&num, 1)
		}
		defer wg.Done()
	}()
	wg.Wait()
	atoNum := atomic.LoadInt64(&num)
	fmt.Println(atoNum)
}
