// 第三天
package day003

import (
	"fmt"
	"testing"
)

// 1.下面两段代码输出什么

// 参考解析
// 这道题考的是使用append向slice添加元素，make初始化slice，slice的size会初始零值

// 输出结果
// [0 0 0 0 0 1 2 3]
func TestAppendA(t *testing.T) {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

// 输出结果
// [1,2,3,4]
func TestAppendB(t *testing.T) {
	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4)
	fmt.Println(s)
}

// 2.下面这段代码的缺陷是什么

// 参考答案: 第二个返回值没有命名
// 参考解析: 在函数有多个返回值时，只要有一个返回值有命名，其他的也必须命名。如果有多个返回值必须加上括号();如果只有一个返回值并且命名也必须加上括号(),
// 这里的第一个返回值有命名sum，第二个没有命名，所以错误。

//func funcMui(x, y int)(sum int, error){
//	return x + y, nil
//}

// 3.new()和make()的区别

// 参考答案:
// new(T)和make(T, args)是GO语言的内建函数，但适用的类型不同。
// new(T)会为T类型的新值分配已置零的内存空间，并返回地址(指针),即类型为 *T的值。换句话说，返回一个指针，该指针指向新分配的，类型为T的零值。
// 适用于值类型，如数组, 结构体等。

//make(T, args)返回初始化之后的T类型的值，这个值并不是T类型的零值，也不是指针 *T,是经过初始化之后的T的引用。 make()只适用于slice, map, 和channel。
