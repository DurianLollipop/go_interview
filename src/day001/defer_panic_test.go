// 第一天
// 下面这段代码输出的内容

// 输出结果：
// 打印后
// 打印中
// 打印前
// panic: 触发异常

// 参考解析: defer的执行顺序是后进先出.当出现panic语句的时候，会先按照defer的后进先出的顺序执行，最后才会执行panic
package day001

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}
