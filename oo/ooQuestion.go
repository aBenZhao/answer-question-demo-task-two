package oo

//题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
//考察点 ：接口的定义与实现、面向对象编程风格。

type Shape interface {
	Area(a float32, b float32) float32
	Perimeter(a float32, b float32) float32
}
type Rectangle struct {
	Long float32
	Wind float32
}

func (rectangle Rectangle) Area(l float32, w float32) float32 {
	return l * w
}
func (rectangle Rectangle) Perimeter(l float32, w float32) float32 {
	return 2 * (l + w)
}

type Circle struct {
	Radius float32
}

func (circle Circle) Area(p float32, r float32) float32 {
	return p * r * r
}

func (circle Circle) Perimeter(p float32, r float32) float32 {
	return 2 * p * r
}
