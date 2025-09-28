package goroutine

import (
	"fmt"
	"sync"
	"time"
)

// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。
func GoroutineQuestionOne() {
	go oddNum()
	go evenNum()
	time.Sleep(10 * time.Second)

}

func oddNum() {
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			time.Sleep(1 * time.Second)
			fmt.Println("奇数：", i)
		}
	}
}

func evenNum() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			time.Sleep(1 * time.Second)
			fmt.Println("偶数：", i)
		}
	}
}

type Task struct {
	Id      string
	TaskFun func()
}

// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。
func GoroutineQuestionTwo(tasks []Task) map[string]time.Duration {
	result := make(map[string]time.Duration)
	var mut sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(tasks))

	for _, task := range tasks {
		currentTask := task
		go func() {
			start := time.Now()
			currentTask.TaskFun()
			duration := time.Since(start)

			mut.Lock()
			result[currentTask.Id] = duration
			mut.Unlock()

			defer wg.Done()
		}()
	}

	wg.Wait()
	return result
}
