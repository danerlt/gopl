package main

import "fmt"

var a = "hello"

func main() {
	fmt.Printf("global a is %v, address: %v \n", a, &a)
	a := "hello"
	fmt.Printf("main a is %v, address: %v \n", a, &a)
	for i := 1; i < 2; i++ {
		a := "hello"
		fmt.Printf("for a is %v, address: %v \n", a, &a)
		if i%2 != 0 {
			a := "hello"
			fmt.Printf("if a is %v, address: %v \n", a, &a)
		}

	}

}

/**
执行结果  说明go里面不同作用域的变量,地址不一样,重新生成了对象
global a is hello, address: 0x5a60c0
main a is hello, address: 0xc000088230
for a is hello, address: 0xc000088250
if a is hello, address: 0xc000088270
*/
