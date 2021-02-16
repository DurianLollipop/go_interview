// 第二天
// 下面这段代码输出什么，说明原因
package main

import (
	"fmt"
	"testing"
)

// 错误的写法：

// 输出结果
// 0 -> 3
// 1 -> 3
// 2 -> 3
// 3 -> 3

// 参考解析: 这是新手常会犯的错误写法，for range循环的时候会创建每个元素的副本，而不是元素的引用，所以m[key] = &val取的都是变量val的地址，
// 所以最后map中的所有元素的值都是val的地址，最后因为val被赋值为3，所有输出的都是3.
func TestA(t *testing.T) {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		m[key] = &val
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}

// 正确的写法

// 输出结果
// 1 => 1
// 2 => 2
// 3 => 3
// 0 => 0
func TestB(t *testing.T) {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		value := val
		m[key] = &value
	}

	for k, v := range m {
		fmt.Println(k, "=>", *v)
	}
}
