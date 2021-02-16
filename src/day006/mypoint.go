// day006
package day006

import "fmt"

// 1.通过指针变量 p 访问其成员变量 name，有哪几种方式？
type person struct {
	age  int
	name string
}

func usePoint() {
	p := &person{13, "张三"}
	fmt.Println(p.name) //张三
	//fmt.Println(&p.name) //0xc00000c0c8
	fmt.Println((*p).name) //张三
	//fmt.Println(p->name) //无该种写法
}

// 2.下面这段代码能否通过编译？如果通过，输出什么？

// 参考答案与解析: 编译不通过， cannot use i (type int) as type MyInt1 in assignment.
// 这道题考的是类型别名与类型定义的区别。

// type MyInt1 int 是基于类型int创建了新类型MyInt1, type MyInt2 = int 是创建了int的类型别名, 注意类型别名的定义时=。
// 所以var i1 MyInt1 = i 相当于将int类型的变量赋值给MyInt1类型的变量,Go是强语言类型, 编译不通过;
// 而MyInt2只是int的别名,本质上还是int,可以赋值。
// 可以使用强制类型转换进行赋值 var i1 MyInt1 = MyInt1(i)

type MyInt1 int
type MyInt2 = int

func typeAlisa() {
	var i int = 0
	//var i1 MyInt1 = i  类型转换错误
	var i1 MyInt1 = MyInt1(i)
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
}
