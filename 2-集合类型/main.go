package main

import (
	"fmt"
)

// 定义animal接口，接口需要一个description的方法（功能）
type animal interface {
	description() string
}

// 定义cat结构体
type cat struct {
	Type  string
	Sound string
}

// cat实现接口的description方法
func (c cat) description() string {
	return fmt.Sprintf("Sound: %v", c.Sound)
}

// 定义snake结构体
type snake struct {
	Type      string
	Poisonous bool
}

// snake实现接口的description方法
func (s snake) description() string {
	return fmt.Sprintf("Poisonous: %v", s.Poisonous)
}

func main() {
	//声明接口animal给变量a
	var a animal
	//将接口animal即变量a初始化为snake
	a = snake{Poisonous: true}
	//调用接口
	fmt.Println(a.description())
	a = cat{Sound: "Meow"}
	fmt.Println(a.description())
}

/*
Poisonous: true
Sound: Meow
*/
