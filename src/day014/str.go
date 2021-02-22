// day014
package day014

import "fmt"

// 1.下面代码输出什么？
// 参考答案及解析: compilation error, GO语言中的字符串是只读的。
func Str() {
	str := "hello"
	//str[0] = 'x'
	fmt.Println(str)
}

// 2.下面代码输出什么?
// 参考答案及解析: 2。 指针，incr()函数里的p是 *int类型的指针，指向的是调用函数的变量p的地址。
// incr()是将该地址的值执行一个自增操作。
func incr(p *int) {
	*p++
}

func PointAdd() {
	p := 1
	incr(&p)
	fmt.Println(p)
}

// 3.对 add() 函数调用正确的是（）
// 	A: add(1, 2)
//	B: add(1, 2, 3)
//  C: add([]int{1, 2})
//	D: add([]int{1, 2, 7}...)

// 参考答案及解析: A, B, D
func add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

func invokeAdd() {
	add(1, 2)
	add(1, 2, 3)
	add([]int{1, 2, 7}...)
}
