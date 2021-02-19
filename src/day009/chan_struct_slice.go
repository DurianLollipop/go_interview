// day009
package day009

import "fmt"

// 1.关于channel, 下面语法正确的是()

// 参考答案及解析: A,B,C。 A, B都是声明channel; C读取channel; D写channel是必须带上值的，例如ch <- "a", D错误;
func chanUse() {
	// A:
	var chA chan int
	// B:
	chB := make(chan int)
	// C:
	<-chB
	// D:
	//chB <-
	fmt.Println(chA, chB)
}

// 2.下面这段代码输出什么?
// A: 0
// B:1
// C:Compilation error

// 参考答案及解析: 打印一个map中不存在的值时，返回元素类型的零值。
type person struct {
	name string
}

func mapUse() {
	var m map[person]int //此时m为nil
	p := person{"mike"}
	fmt.Println(m[p])
}

// 3.下面这段代码输出什么?
// A:18
// B:5
// C:Compilation error

// 参考答案及解析：18
func hello(num ...int) {
	fmt.Printf("%p\n", &num)
	num[0] = 18
}

func sliceUse() {
	i := []int{5, 6, 7}
	fmt.Printf("%p\n", &i)
	hello(i...)
	fmt.Println(i[0])

}
