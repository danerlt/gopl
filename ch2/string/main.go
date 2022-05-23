package main

import "fmt"

func main() {
	a := "1111"
	b := "1111"
	fmt.Printf("a: %v, &a: %v \n", a, &a)
	fmt.Printf("b: %v, &b: %v \n", b, &b)
	fmt.Printf("a == b: %v, addr a == addr b: %v \n", a == b, &a == &b)
	c := 1
	d := 1
	fmt.Printf("c: %v, &c: %v \n", c, &c)
	fmt.Printf("d: %v, &d: %v \n", d, &d)
	fmt.Printf("c == d: %v, addr c == addr d: %v \n", c == d, &c == &d)
}

/**
执行结果 go里面没有Python内存共享的机制
a: 1111, &a: 0xc00004c250
b: 1111, &b: 0xc00004c260
a == b: true, addr a == addr b: false
c: 1, &c: 0xc000010140
d: 1, &d: 0xc000010148
c == d: true, addr c == addr d: false
*/
