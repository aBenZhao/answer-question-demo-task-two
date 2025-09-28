package main

import (
	"answer-question-demo-task-two/oo"
	"fmt"
	"math"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	// 指针题目一
	//num := 200
	//fmt.Println(pointer.QuestionOne(&num))

	// 指针题目二
	//numSlice := []int{1, 2, 3, 4, 5}
	//fmt.Println(pointer.QuestionTwo(&numSlice))

	// goroutine题目一
	//goroutine.GoroutineQuestionOne()

	// goroutine题目二
	//taskSlice := []goroutine.Task{
	//	{"task1", func() { time.Sleep(1000 * time.Millisecond) }},
	//	{"task2", func() { time.Sleep(2000 * time.Millisecond) }},
	//	{"task3", func() { time.Sleep(3000 * time.Millisecond) }},
	//	{"task4", func() { time.Sleep(4000 * time.Millisecond) }},
	//	{"task5", func() { time.Sleep(3000 * time.Millisecond) }},
	//}
	//fmt.Println(goroutine.GoroutineQuestionTwo(taskSlice))

	// OO题目一
	circle := oo.Circle{3}
	fmt.Println("圆的面积：", circle.Area(math.Pi, circle.Radius))
	fmt.Println("圆的周长：", circle.Perimeter(math.Pi, circle.Radius))
	rectangle := oo.Rectangle{3, 4}
	fmt.Println("矩形的面积：", rectangle.Area(rectangle.Long, rectangle.Wind))
	fmt.Println("矩形的周长：", rectangle.Perimeter(rectangle.Long, rectangle.Wind))
}
