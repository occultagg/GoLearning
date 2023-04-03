package main

import "fmt"

//起重机接口
type Crane interface {
	JackUp() string
	Hoist() string
}

//起重机A
type CraneA struct {
	work int
}

func (c CraneA) Work() {
	fmt.Println("使用技术A")
}

func (c CraneA) JackUp() string {
	c.Work()
	return "jackup"
}

func (c CraneA) Hoist() string {
	c.Work()
	return "hoist"
}

//起重机B
type CraneB struct {
	boot string
}

func (c CraneB) Work() {
	fmt.Println("使用技术B")
}

func (c CraneB) JackUp() string {
	c.Work()
	return "jackup"
}

func (c CraneB) Hoist() string {
	c.Work()
	return "hoist"
}

//建筑公司
type ConstructionCompany struct {
	//建筑公司结构体里有一个Crane字段，类型是Crane接口，接口的本质是结构体，所以也是一种类型
	Crane Crane //只根据Crane类型来存放起重机
}

//建筑公司做的事情不会改变
func (c *ConstructionCompany) Build() {
	fmt.Println(c.Crane.JackUp())
	fmt.Println(c.Crane.Hoist())
	fmt.Println("建筑完成")
}

func main() {
	//使用起重机A
	company := ConstructionCompany{CraneA{}}
	company.Build()
	fmt.Println()
	//更换起重机B
	company.Crane = CraneB{}
	company.Build()
}
