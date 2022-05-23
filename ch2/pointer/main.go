package main

import "fmt"

func main() {
	v := 1 //定义一个变量赋值为1
	fmt.Printf("调用前变量v的地址：%v, v的值：%v\n", &v, v)
	fmt.Println("调用第一次")
	v2 := incr(&v)
	fmt.Printf("v的地址：%v, v的值：%v\n", &v, v)
	fmt.Printf("v2的地址：%v, v2的值：%v\n", &v2, v2)
	fmt.Println("调用第二次")
	v3 := incr(&v)
	fmt.Printf("v的地址：%v, v的值：%v\n", &v, v)
	fmt.Printf("v3的地址：%v, v3的值：%v\n", &v3, v3)
	fmt.Println(&v == &v2)
	fmt.Println(&v == &v3)
}

func incr(p *int) int {
	fmt.Printf("***********************修改前指针p的地址：%v, 指针p的值: %v\n", p, *p)
	*p++ // 将指针p指向的值加1
	fmt.Printf("***********************修改后指针p的地址：%v, 指针p的值: %v\n", p, *p)
	return *p
}
