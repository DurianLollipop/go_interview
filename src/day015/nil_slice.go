// day015
package day015

import "fmt"

// 1.下面代码下划线处可以填入哪个选项？
// A: s1
// B: s2
// C: s1, s2都可以
// 参考答案及解析，nil切片和空切片。 nil切片和nil相等，一般用来表示一个不存在的切片；
// 空切片和nil不相等，表示一个空的集合。
func nilSlice() {
	var s1 []int
	var s2 = []int{}
	fmt.Printf("%#v\n", s1)
	if s1 == nil {
		fmt.Println("s1 yes nil")
	} else {
		fmt.Println("s1 no nil")
	}

	fmt.Printf("%#v\n", s2)
	if s2 == nil {
		fmt.Println("s2 yes nil")
	} else {
		fmt.Println("s2 no nil")
	}
}

// 2.下面这段代码输出是什么？
// 参考答案及解析： A, UTF-8编码中，十进制数字65对应的符号是A。
func iToString() {
	i := 65
	fmt.Println(string(rune(i))) // A
}

// 3.下面这段代码输出什么？
// 参考答案及解析: 13 23。
// 知识点: 接口。一种类型实现多个接口，结构体Work分别实现了接口A, B, 所以接口变量a, b调用各自的方法ShowA()
// 和ShowB().
type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w *Work) ShowA() int {
	return w.i + 10
}

func (w *Work) ShowB() int {
	return w.i + 20
}

func interfaceUse() {
	c := &Work{3}
	var a A = c
	var b B = c
	fmt.Println(a.ShowA())
	fmt.Println(b.ShowB())
}
