// day 011
package day011

import "fmt"

// 1.关于 cap() 函数的适用类型，下面说法正确的是()
// 参考答案及解析： array, slice, channel, cap()函数不适用于map
func capUse() {
	arr := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("arr %T, %d\n", arr, cap(arr))
	sli := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("sli %T, %d\n", sli, cap(sli))
	//m := map[string]int{"A":1, "B":2,"C":3}
	//fmt.Printf("map %T, %d\n", m, cap(m))
	cha := make(chan int, 6)
	fmt.Printf("cha %T, %d\n", cha, cap(cha))
}

// 2.下面这段代码输出什么？
// 参考答案及解析：A。 当且仅当接口的动态值和动态类型都为nil时，接口类型才为nil
func interfaceNil() {
	var i interface{}
	if i == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println("not nil")
}

// 3.下面这段代码输出什么？
// 参考答案及解析： 0， 输出map不存在的键值对时，不会报错，相当于没有任何操作；
// 获取不存在的键值对时，返回值类型对应的零值，所以返回0.
func mapDelete() {
	s := make(map[string]int)
	delete(s, "h")
	fmt.Println(s["h"])
	val, ok := s["h"]
	fmt.Println(val, ok)
}
