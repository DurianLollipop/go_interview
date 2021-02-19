// day007
package day007

import (
	"fmt"
	"strings"
)

// 1.关于字符串连接，下面语法正确的是?
// 参考答案及解析: B,D。 知识点：字符串连接还可以使用strings.Join(), buffer.WriteString等。
func stringConnect() {
	// A: 错误
	//strA :='abc' + '123'
	// B: 正确
	strB := "abc" + "123"
	// C: 错误
	//strC := '123' + "abc"
	// D: 正确
	strD := fmt.Sprintf("abc%d", 123)
	fmt.Println(strB, strD)

	strE := strings.Join([]string{"def", "456"}, ",")
	fmt.Println(strE)
}

// 2.下面这段代码能否编译通过？如果可以，输出什么？
// 参考答案及解析: 能编译通过，输出L： -1 2 zz zz 5
// 安行计数，不看从哪行开始
const (
	x = -1   // iota= 0
	_        // iota= 1
	y = iota // iota= 2
	z = "zz" // iota= 3
	k        // iota= 4
	p = iota // iota= 5
)

func testIota() {
	fmt.Println(x, y, z, k, p)
}

// 3.下面赋值正确的是()
// 参考答案及解析：B,D。 nil值只能赋给指针, chan, func, interface, map, slice类型的变量.选项D的error类型，它是一种内置接口类型。
func assignment() {
	// A:
	//var a = nil
	// B:
	var b interface{} = nil
	// C:
	//var c string = nil
	// D:
	var d error = nil
	fmt.Println(b, d)
}
