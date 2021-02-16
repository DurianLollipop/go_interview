// day005
package day005

import (
	"fmt"
)

// 参考答案及解析: 编译不通过invalid operation: sm1 == sm2

// 这道题目考查的是结构体的比较，有几个需要注意的地方:

// 1.结构体只能比较是否相等，但不能比较大小
// 2.相同类型的结构体才能进行比较，结构体是否相同不但与属性类型有关，还与属性顺序有关，sn3与sn1就是不同的结构体
// 3.如果struct的所有成员都可以比较，则该struct就可以通过==或！=进行比较是否相等，如果每一项都相等，则两个结构体才相等，否则不相等；
// 什么是都可以比较呢， 常见的类型有bool,数值型,字符,指针,数组，像切片,map,函数等使用reflect.DeepEqual(x, y interface{})进行比较
func compile() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	//sn3 := struct {
	//	name string
	//	age int
	//}{age: 11, name: "qq"}
	//
	//if sn1 == sn3 {
	//	fmt.Println("sn1 == sn2")
	//}

	//sm1 := struct {
	//	age int
	//	m map[string]string
	//}{age: 11, m: map[string]string{"a":"1"}}
	//sm2 := struct {
	//	age int
	//	m map[string]string
	//}{age: 11, m: map[string]string{"a":"1"}}
	//// 使用reflect.DeepEqual(sm1 ,sm2)进行比较
	//if sm1 == sm2 {
	//	fmt.Println("sm1 ==sm2")
	//}
}
