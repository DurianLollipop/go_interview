// day012
package day012

import "fmt"

// 1.下面属于关键字的是？
// A: func
// B: struct
// C: class
// D: defer
// 参考答案及解析： A,B,D.
// Go语言有25个关键字：
// break    default     func   interface select
// case     defer       go     map       struct
// chan     else        goto   package   switch
// const    fallthrough if     range     type
// continue for         import return    var

// 2.下面这段代码输出什么？
// 格式化字符串里的+号代表输出数值的符号
func formatStr() {
	i := -5
	j := +5
	fmt.Printf("%+d %+d\n", i, j) //-5 +5
	fmt.Printf("%d %d\n", i, j)   //-5 5
}

// 3.下面这段代码输出什么？
// 参考答案及解析： teacher showB.
// 知识点： 结构体嵌套。 在结构体嵌套中，People称为内部类型， Teacher称为外部类型；
// 通过嵌套， 内部类型的属性，方法，可以为外部类型所有，就好像是外部类型自己的一样。
// 此外，外部类型还可以定义自己的属性和方法，甚至可以和内部定义相同的方法， 这样
// 内部类型的方法就会被"屏蔽".
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

func Nested() {
	t := Teacher{}
	t.ShowB()
}
