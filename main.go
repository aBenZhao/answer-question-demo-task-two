package main

import (
	"answer-question-demo-task-two/pointer"
	"fmt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	// 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
	num := 200
	fmt.Println(pointer.QuestionOne(&num))

	// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
	numSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(pointer.QuestionTwo(&numSlice))
}
