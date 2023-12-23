package main

import (
	_ "fmt" //副作用导入的意思是，引入包时，只执行包内的init函数，不执行其他函数
	"testing"
)

//global

func TestName(t *testing.T) {
	name := "hello"
	//fmt.Println(name)
	t.Log(name)
}
