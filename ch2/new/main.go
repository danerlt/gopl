package main

import "fmt"

func main() {
	p := new(int)
	q := new(int)
	fmt.Printf("p的地址:%v， p的值%v， q的指针%v， q的值%v， p==q:%v, *p==*q:%v \n", p, q, *p, *q, p == q, *p == *q)
}
