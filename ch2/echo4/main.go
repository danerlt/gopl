package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "忽略正常输出时结尾的换行符")
var sep = flag.String("s", " ", "分隔符")

func main() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
