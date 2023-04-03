# 面向接口编程

go不是绝对的面向对象的语言，结构体，方法，接口，是golang对面向对象的支持。

## 基本理解

- 结构体：不同字段的类型集合，它一种复合类型，将数据分组在一起。
- 方法：是一个特殊类型带有返回值的函数，返回值可以是值也可以是指针。
- 接口：
  - 1.17-：一系列方法的集合。（基本接口）
  - 1.18+：一组类型的集合。为泛型服务。（通用接口）

# 基本接口

方法集，一组方法的集合

## 声明

```go
type interface_name interface {
   method_name1([args ...arg_type]) [return_type]
   method_name2([args ...arg_type]) [return_type]
   method_name3([args ...arg_type]) [return_type]
   ...
   method_namen([args ...arg_type]) [return_type]
}
```

在接口里，函数的参数名变得不再重要，所以可以省略

```go
//定义一个Person接口
type Person interface {
    //接口定义两个方法
	Say(string) string
	Walk(int)
}
```

可以理解为定义了Person这么一个接口（合约），在接口中约定要有Say和Walk两个方法（功能），至于你怎么实现（方法实现）这两个方法（功能），接口（合约）不管，调用时直接通过接口调用就行。这样一样，当你需要升级更新你的方法实现时，比如优化性能，你只需要修改你的方法，作为调用方是无感的，也无需改变调用方式，这就是面向接口编程。

## 实现

### 一个容易理解的例子

```go
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
```

这里定义了animal这个接口，这个接口需要实现description这个功能，cat结构体实现它是通过sound，snake结构体实现它通过打印它是否有毒，接口不管你具体怎么实现，只要你有描述就行。

### 又一个例子

> 一个建筑公司想一种特殊规格的起重机，于是给出了起重机的特殊规范和图纸，并指明了起重机应该有起重和吊货的功能，**建筑公司并不负责造起重机，只是给出了一个规范，这就叫接口**，于是公司A接下了订单，根据自家公司的独门技术造出了绝世起重机并交给了建筑公司，建筑公司不在乎是用什么技术实现的，也不在乎什么绝世起重机，只要能够起重和吊货就行，仅仅只是当作一台普通起重机来用，**根据规范提供具体的功能，这就叫实现**，。只**根据接口的规范来使用功能，屏蔽其内部实现，这就叫面向接口编程**。过了一段时间，绝世起重机出故障了，公司A也跑路了，于是公司B依据规范造了一台更厉害的巨无霸起重机，由于同样具有起重和吊货的功能，可以与绝世起重机无缝衔接，并不影响建筑进度，建筑得以顺利完成，**内部实现改变而功能不变，不影响之前的使用，可以随意替换，这就是面向接口编程的好处。**

用go来描述上面的情况：

```go
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
	//更换起重机B，将里面的Crane字段更换为结构体CraneB
	company.Crane = CraneB{}
	company.Build()
}

/*
使用技术A
jackup
使用技术A
hoist
建筑完成

使用技术B
jackup
使用技术B
hoist
建筑完成
*/
```

> 只要是实现了一个接口的全部方法，那就是实现了该接口。有了实现之后，就可以初始化接口了，建筑公司结构体内部声明了一个`Crane`类型的成员变量，可以保存所有实现了`Crane`接口的值，由于是`Crane` 类型的变量，所以能够访问到的方法只有`JackUp` 和`Hoist`，内部的其他方法例如`Work`和`Boot`都无法访问。

# 空接口

```go
type Any interface{

}
```

TBD

# 通用接口

通用接口就是为了泛型服务的，只要掌握了泛型，就掌握了通用接口

# 泛型

比如有一个sum函数计算两数之和，但是定义时定义了整型。

```go
func Sum(a, b int) int {
    return a + b
}
```

那如果要计算两个浮点型的和呢？重新写一个新函数？那如果要写一个计算器程序，难道要有多少种类型就写多少个函数？然而里面的逻辑是一样的，只是两数相加而已。显然不科学。

泛型是为了解决执行逻辑与类型无关的问题。

## 写法

```go
func Sum[T int | float64](a, b T) T {
	return a + b
}

func main() {
    //可以显式指明类型实参
	s1 := Sum[int](2, 3)
    //也可以隐式，让编译器自己判断
	s2 := Sum(3.1415, 1.1144)
	fmt.Println("s1:", s1, "s2:", s2)
}
```

`T`：类型形参，形参具体什么类型取决于传进来什么类型。

`int | float64`：类型约束，规定了允许的类型形参的范围。

`Sum[int](2, 3)`：int就是了类型实参

## 泛型结构

泛型可以用于多个地方，例如

```go
//切片
type GenericSlice[T int | int32 | int64] []T
//这里使用时不能省略掉类型实参
s1 := GenericSlice[int]{1, 2, 3}

//泛型映射表，使用comparable接口确定键式可比较的，值使用泛型约束
type GenericMap[K comparable, V int | string | byte] map[K][V]
//使用
gmap1 := GenricMap[int, string]{1: "hello world"}
gmap2 := make(GenricMap[string, byte], 0)

//泛型结构体
type Company[T int | string] struct {
    Name string
    Id T
}

//使用
c := Company[int]{
    Name: "Jack",
    Id: 1
}
```

更多：[参考这里](https://golang.halfiisland.com/%E8%AF%AD%E8%A8%80%E5%85%A5%E9%97%A8/%E8%AF%AD%E6%B3%95%E8%BF%9B%E9%98%B6/90.generic.html#%E6%B3%9B%E5%9E%8B%E7%BB%93%E6%9E%84)

# 类型集

1.18后，接口的定义变为类型集（type set），含有类型集的接口又称为Genral interface即通用接口。

类型集主要用于约束类型，不能用于类型声明。

## 并集

```go
//接口SigneInt式一个类型集，包含了全部有符号整数的并集
type SignedInt interface {
    int8 | int16 | int | int32 | int64
}
```

## 交集

如果一个接口包含多个非空类型集，那么该接口就是这些类型集的交集

```go
package main

import "fmt"

type SignedInt interface {
	int8 | int16 | int | int32 | int64
}

type Integer interface {
	int8 | int16 | int | int32 | int64 | uint8 | uint16 | uint | uint32 | uint64
}

// Number接口就是包含的SignedInt和Integer两个类型集的交集
// 所以该接口允许的类型集式：int8 | int16 | int | int32 | int64
type Number interface {
	SignedInt
	Integer
}

func Do[T Number](n T) T {
	return n
}

func main() {
	r := Do[int](2)
	fmt.Println(r)
}
```

## 空集

就是没有交集

```go
type SignedInt interface {
	int8 | int16 | int | int32 | int64
}

type Integer interface {
	uint8 | uint16 | uint | uint32 | uint64
}

//这个就是空集
type Integer interface {
	SignedInt
	Integer
}

//空集不管传什么类型都无法编译
```

更多： [参考这里](https://golang.halfiisland.com/%E8%AF%AD%E8%A8%80%E5%85%A5%E9%97%A8/%E8%AF%AD%E6%B3%95%E8%BF%9B%E9%98%B6/90.generic.html#%E7%B1%BB%E5%9E%8B%E9%9B%86)