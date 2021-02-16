// day004
package day004

// 1.下面这段代码能否通过编译，不能的话原因是什么；如果能，输出什么。

// 参考答案及解析: 不能通过编译,new([]int)之后的list是一个 *[]int 类型的指针，不能对指针进行append操作。
// 可以使用make()初始化之后再用。 同样的, map和channel建议使用make()或字面量的方式初始化，不要使用new().
//func newSlice() {
//	list := new([]int)
//	list = append(list, 1)
//	fmt.Println(list)
//}

// 2.下面这段代码能否通过编译，如果可以，输出什么？

// 参考答案及解析: 不能通过编译。append()的第二个参数不能直接使用slice,需要使用...操作符，讲一个切片追加到另一个切片上:
// append(s1, s2...)。或者直接跟上元素,形如append(s1, 1, 2, 3)。
//func appendSlice() {
//	s1 := []int{1, 2, 3}
//	s2 := []int{4, 5}
//	s1 = append(s1, s2)
//	fmt.Println(s1)
//}

// 3.下面这段代码能否通过编译，如果可以，输出什么？

// 参考答案及解析: 不能通过编译。这道题的主要知识点是变量声明的简短模式，形如： x := 100。但这种声明有限制。
// 1.必须使用显示初始化;
// 2.不能提供数据类型，编译器会自动推导;
// 3.只能在函数内部使用;

//var (
//	size := 1024
//	max_size = size * 2
//)
//
//func declareVariable()  {
//	fmt.Println(size, max_size)
//}
