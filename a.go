package main

import (
	"fmt"
)

// 输出运行结果
func main() {
	DeferCall()
}
func DeferCall() {
	defer func() { fmt.Println("A") }()
	defer func() { fmt.Println("B") }()
	defer func() { fmt.Println("C") }()

	panic("error")
}

// 输出运行结果
type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA()

	t.ShowB()
}

// 输出运行结果
func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

// 输出运行结果
type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

func main() {

	if live() == nil {
		fmt.Println("AAAAAAA", live())
	} else {
		fmt.Println("BBBBBBB", live())
	}

	var s *Student
	fmt.Println(s == nil)
}
