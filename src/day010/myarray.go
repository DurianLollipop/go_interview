// day010
package day010

import "fmt"

// 1.下面这段代码输出什么？
// A: 13.1
// B: 13
// C:compilation error

// 参考答案及解析: C, A的类型是int, B的类型是float, 两个不同类型的数值不能相加
func intAddFloat() {
	//a := 5
	//b := 8.1
	//fmt.Println(a + b)
}

// 2.下面这段代码输出什么？
// b := a[low:high:max], 新切片cap(b) = max-low, 0 <= low <= high <= max <= cap(a), 左包含右不包含
// 参考答案及解析: 操作符[low:high:max].基于数组(切片)可以使用操作符[i:j]，创建新的切片，
// 从索引i(包含),到索引j(不包含), 截取已有数组(切片)的任意部分,返回新的切片,i,j都是可选的， i省略，默认为0，
// j省略，默认原数组(切片)的长度.
func sliceArr() {
	a := [5]int{1, 2, 3, 4, 5}
	j := a[3:4]
	fmt.Printf("j j[0]:%d len:%d, cap:%d\n", j[0], len(j), cap(j))
	k := a[3:4:4]
	fmt.Printf("k k[0]:%d len:%d, cap:%d\n", k[0], len(k), cap(k))
}

// 3.下面这段代码输出什么？
// 参考答案及解析: compilation error
// Go的数组是值类型，可比较。另外一方面，数组的长度也是数组类型的一部分，所以a和b是不同的类型，是不能进行比较的，编译报错。
func compareArr() {
	//a := [2]int{5, 6}
	//b := [3]int{5,6}
	//if a == b {
	//	fmt.Println("equals")
	//}else {
	//	fmt.Println("not equals")
	//}
}
