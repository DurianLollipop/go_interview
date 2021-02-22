// day013

package day013

import "fmt"

// 1.定义一个包内全局字符串变量，下面语法正确的是（）

// 参考答案及解析: A, D。 B只支持局部变量声明。 C是赋值, strC必须在之前已经声明。
// A:
var strA string

// B:
//strB := ""
// C:
//strC = ""
// D:
var strD = ""

// 2.下面这段代码输出什么?
// 参考答案及解析: 5。 hello(i int)函数的参数在执行defer函数时会保存一份副本，在实际调用hello()函数时用，所以是5。
func myDefer() {
	i := 5
	defer hello(i)
	i = i + 10
}

func hello(i int) {
	fmt.Println(i)
}

// 3.下面这段代码输出什么？
// 参考答案及解析： ShowA ShowB。
// 结构体嵌套，Teacher没有自己的ShowA方法，所以调用内部类型People的ShowA方法。
type People struct{}

func (p *People) ShowA() {
	fmt.Println("ShowA")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("ShowB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher ShowB")
}

func structNested() {
	teacher := Teacher{}
	teacher.ShowA()
}
