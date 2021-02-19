// day 008

// 1.关于init()函数，下列说法正确的是()
// A:一个包中，可以包含多个init()函数
// B:程序编译时，先执行依赖包的的init()函数，再执行main包内的init()函数
// C:main包中，不能有init()函数
// D:init()函数能被其他调用

// 参考答案及解析: A,B.
// 关于init()函数有几个需要注意的地方：
// 1.init函数适用于程序执行前做包初始化的函数，比如初始化包里的变量等；
// 2.一个包里可以出现多个init函数，一个源文件里也可以包含多个init()函数；
// 3.同一个包中init()函数的执行顺序没有明确定义，一般按照定义顺序执行；
// 4.init()函数在代码中不能被显示调用,不能被引用(赋值给函数变量)，否则出现编译错误；
// 5.一个包被引用多次，如A import B, C import B, A import C, B被引用多次，但B包只会初始化一次；
// 6.引入包，不可出现死循环。及A import B, B import A,这种情况下编译失败；
package main

import (
	"errors"
	"fmt"
)

// 涉及引用时,先加载的先执行, 根据文件名排序进行文件加载。
// 同一文件中,先定义的先执行
func init() {
	fmt.Println("execute main.init()")
}

func main() {
	fmt.Println("execute main")
}

// 2.下面这段代码输出什么及原因？
// A: nil
// B: not nil
// C: compilation error

// 答案及解析: B。 题目里是将hello()赋值给变量h, 而不是函数的返回结果，所以输出 not nil;
func hello() []string {
	return nil
}

func functionVariable() {
	f := hello
	if f == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}

// 3.下面这段代码能否编译通过? 如果可以会输出什么？

// 参考答案及解析: 编译失败。
// 类型选择: 类型选择的语法形如: i.(type), 其中i是接口, type是固定关键字, 需要注意的是，只有接口类型才可以使用类型选择。
func GetValue() int {
	return 1
}

// 错误示范：
func typeSelectError() {
	//i := GetValue()
	//switch i.(type) {
	//case int:
	//	println("int")
	//case string:
	//	println("string")
	//case interface{}:
	//	println("interface")
	//default:
	//	println("unknown")
	//}
}

// 正确示范：
func typeSelectSuccess(i interface{}) {
	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}
}

func TestTypeSelectSuccess() {
	typeSelectSuccess("A")
	typeSelectSuccess(18)
	typeSelectSuccess(errors.New("interface"))
}
