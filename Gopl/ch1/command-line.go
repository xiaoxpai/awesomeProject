// Ech1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

// 连接字符串
// sep + os.Args[i]

// 表示连接字符串sep 和os.Args[i]，并将结果再次赋值给s
// s += sep + os.Args[i]

// 将s的旧值跟sep与os.Args[i]连接后赋值回s，
// s = s + sep + os.Args[i]

// 符号:=是短变量声明（short variable declaration）

// 变量声明的等价写法形式，但是有规约
// s := "" //只能在函数内部，不能用与包变量  --常用
// var s string //默认初始化为零值，""  --常用
// var s = ""
// var s string = ""

//用哪种不用哪种，为什么呢？
//第一种形式，是一条短变量声明，最简洁，但只能用在函数内部，而不能用于包变量。
//第二种形式依赖于字符串的默认初始化零值机制，被初始化为""。
//第三种形式用得很少，除非同时声明多个变量。
//第四种形式显式地标明变量的类型，当变量类型与初值类型相同时，类型冗余，但如果两者类型不同，变量类型就必须了。
//实践中一般使用前两种形式中的某个，初始值重要的话就显式地指定变量的类型，否则使用隐式初始化。

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
