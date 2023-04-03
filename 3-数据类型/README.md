# 变量

## 声明

``` go
// var <变量名> <类型名>
var i int
```

- 变量必须被使用，否则不能编译，避免浪费内存

``` go
// 多变量声明
var (
	j int
	k int
)
```

## 变量的简短声明

编译器自动判断类型

``` go
// <变量名> := <表达式>
i := 10
j, k := "peter", 22
```

- 使用短声明时候不能重复对同一个变量重复声明，除非同时声明新变量

```go
//报错
a := 1
a := 2
fmt.Println(a)

//成功
a := 1
a, b := 2, 3
fmt.Println(a, b)
```



## 赋值

``` go
i = 20
j, k = "peter", 22
```

> note:
>
> := 可以声明变量也可以赋值
>
> = 只能赋值

## 变量交换

在go中，如果要交换两个变量的值，不需要用到指针

```go
num1, num2 := 25, 36
nam1, num2 = num2, num1
//
num1, num2, num3 := 25, 36, 49
nam1, num2, num3  = num3, num2, num1
```

## 变量的使用

- 当变量的作用域为函数内部时，如果该变量没有被调用，会无法通过编译
- 当变量的作用域为整个包时，就允许该变量可以未被调用

# 基本数据类型

## 整形

- 有符号：int/int8/int16/int32/int64
  - 负数，零，正数
- 无符号: uint/uint8/uint16/uint32/uint64
  - 零，正数
- uintptr：一个足够大的整数类型，足以容纳任何位数的整数指针（特殊用途）
- 后面的数字表示整型的大小，单位是bit，不写则表示没有具体大小，该整型的大小与cpu有关
- 能明确最好明确表明大小，增加可移植性

## 复数类型

- complex128：64位实数和虚数
- complex64：32位实数和虚数

## 浮点型

代表现实中的小数

- float32
- float64： 更常用，精度更高，误差更小

## 布尔型

bool

- 一元操作符： “！” 代表  “非”
- 二元操作符
  - “&&” => 与
  - “||” => 或

## 字符类型

go对于Unicode编码百分百兼容和支持

- byte： 等价于uint8可以表达ANSCII字符
- rune：等价于int32可以表达Unicode字符
- string：字符串即字节序列，可以 转换为[]byte类型即字节切片
  - 可以通过操作符“+”连起来
  - 也可以通过“+=”运算符号操作

## 派生类型

- 数组： 如 [5]int 表示长度为5的整型数组
- 切片：[]float64 表示64位浮点数切片（python的list）
- 映射表：map[string]int 键为字符串类型，值为整型的映射表（python的dict）
- 结构体： type Gopher struct{}，名为Gopher的结构体
- 指针： *int 整型指针
- 接口：type Gopher interface{} 名为Gopher的接口
- 通道：chan int 整型通道

## 零值

就是变量的默认值，单声明变量时没有赋值，golang就会将其初始化成零值

- 数字类型的零值是0
- 布尔的零值是false
- 字符串的零值是空字符串“”
- 数组的零值是固定长度的对应类型的零值集合
- 结构体的零值是内部字段都是零值的结构体
- 切片，映射表，函数，接口，通道，指针的零值都是nil

## nil

仅仅只是一个变量，仅仅只是一些类型的零值但不属于任何类型所以nil == nil是错误的

## 指针

`&`：取地址符，对一个变量进行取址，返回对应类型的指针

`*`：解引用符

1. 访问指针所指向的元素，就是解引用

   ```go
   func main() {
   	num := 2
       //通过取地址符获取num变量的指针，初始化给p，p就成了一个指针
   	p := &num
       //通过对p解引用就可以访问该指针指向的元素
   	rawNum := *p
   	fmt.Println(rawNum)
   }
   ```

2. 声明一个指针

   ```go
   func main() {
       var numPtr *int
       fmt.Println(numPtr)
   }
   //输出：nil
   ```

   `*int`表示numPtr变量是一个int类型的指针，但是光声明没有初始化则只是一个空指针，没有内存空间，不能正常使用，所以放回nil，要么使用取地址符将其他变量的地址赋给该指针，要么使用内置函数`new()`手动分配

   ```go
   func main() {
       numPtr := new(int)
       fmt.Println(numPtr)
   }
   //输出： 0xc000018098
   //此时numPtr就是一个有确切内存地址的指针
   ```

   new函数只有一个参数那就是类型，并返回一个对应类型的指针，该指针指向对应类型的零值，并为该指针分配内存

   ```go
   func main() {
      fmt.Println(*new(string)) //""
      fmt.Println(*new(int)) //0
      fmt.Println(*new([5]int)) //[0 0 0 0 0]
      fmt.Println(*new([]float64)) //[]
   }
   ```

   golang中的指针禁止运算，标准库中的unsafe提供一些指针运算的操作

## new和make

```go
func new(Type) *Type
```

- 返回值是类型指针
- 接收参数是类型
- 专用于给指针分配内存空间

```go
func make(t Type, size ...IntegerType) Type
```

- 返回值是值，不是指针
- 接收的第一个参数是类型，不定长参数根据传入类型的不同而不同
- 专用于给切片，映射表，通道分配内存。

一些例子：

```go
new(int) // int指针
new(string) // string指针
new([]int) // 整型切片指针
make([]int, 10, 100) // 长度为10，容量100的整型切片 
make(map[string]int, 10) // 容量为10的映射表
make(chan int, 10) // 缓冲区大小为10的通道
```

## 常量

不变的量，防止程序运行时被恶意修改

只能来源于:

- 字面量
- 其他常量标识符
- 常量表达式
- 结果是常量的类型转换
- iota
- 等等
- 除此以外常量不能是其他类型，例如切片这类派生类型

```go
//常量的声明需要用到const关键字，常量在声明时就必须初始化一个值，并且常量的类型可以省略
const name string = "Jack" // 字面量
const msg = "hello world" // 字面量
const num = 1 // 字面量
const numExpression = (1+2+3) / 2 % 100 + num // 常量表达式
```

### iota

内置的常量标识符,用来初始化规则相似的常量,初始值是0.

``` go
//批量声明常量可以用()括起来以提升可读性，可以存在多个()达到分组的效果
const (
	one = 1
    two = 2
    three = 3
    four = 4
)
// 使用iota
const (
    zero = iosta //0
	one //1
    two //2
    three //3
    four//4
)

//也可以这样写
const (
   Num = iota*2 // 0
   Num1 // 2
   Num2 // 4
   Num3 // 6
   Num4 // 8
)
```

- iota的值本质上是iota所在行相对于当前const分组的第一行的差值，不同的const分组则不会相互影响。

## 注意

- 常量值不能修改，否则会编译失败
- 数字常量在golang中可以为任意精度，但是具体还是有一定限制：
  - 表示整型常量时至少有256 bits、
  - 表示浮点数和复数时，尾数部分至少256 bits，二进制指数部分至少16 bits
  - 当无法表示相应的精度时，浮点数会进行相应的舍入。
  - 否则编译器会报错

# 字符串

字符串的本质是一个不可变的只读的字节数组，也是一片连续的内存空间

## 普通字符串

用`""`表示，支持转义，不支持多行书写

``` go
"这是一个普通字符串\n"
"abcdefghijlmn\nopqrst\t\\uvwxyz"
```

## 原生字符串

用反引号表示，不支持转义，支持多行书写，原生字符串里面的所有字符都会原封不动的输出，包括换行和缩进

```go
`这是一个原生字符串，换行
	tab缩进，\t制表符但是无效,换行
	"这是一个普通字符串"
	
	结束
`
```

## 访问

由于字符串本质是字节数组，所以字符串的访问形式跟数组切片完全一致，但输出的是字节而不是字符

```go
func main() {
    str := "this is a string"
    fmt.Println(str[0]) //116
}
```

与数组同样也可以切割

```go
func main() {
   str := "this is a string"
   //使用strings函数将字节转为字符
   fmt.Println(string(str[0:4]))
}
```

字符串无法修改但是可以覆盖

```go
func main() {
   str := "this is a string"
   str[0] = 'a' // 无法通过编译
   fmt.Println(str)
}

func main() {
   str := "this is a string"
   str = "that is b string"
   fmt.Println(str)
}
```

## 转换

字符串和字节切片可以相互转换

```go
func main() {
	str := "this is a string"
    //显式类型转换为字节切片
	bytes := []byte(str)
	fmt.Println(bytes)
    //显示类型转换为字符串
	fmt.Println(string(bytes))
}
```

字符串的内容是只读不可变的，但是字节切片可以修改

```go
func main() {
	str := "this is a string"
	fmt.Println(&str)
	bytes := []byte(str)
    // 修改字节切片
	bytes = append(bytes, 96, 97, 98, 99)
    // 赋值给原字符串
	str = string(bytes)
	fmt.Println(str)
}
```

## 字符串的长度

Unicode下，一个汉字三个字节，UTF-8就是互联网上使用最广的一种unicode实现方式，使用len()获取字符串长度，实际上是字节数组的长度

```go
//UTF-8下，一个汉字三个字节
func main() {
	s := "Peter Pan 小飞侠"
	bs := []byte(s)
	fmt.Println(bs)
    fmt.Println(s[0], s[3], s[5])
}

//使用utf8.RuneCountInString函数一个汉字算一个
func main() {
	s := "Peter Pan 小飞侠"
	fmt.Println(utf8.RuneCountInString(s))
}

//使用for range循环时会自动地隐式解码为Unicode
func main() {
	s := "Peter Pan 小飞侠"
	for i, r := range s {
		fmt.Println(i, r)
	}
}
```

## 拷贝

字符串拷贝其实是字节切片拷贝，使用内置函数copy()或者strings.clone()

```go
func main() {
   var dst, src string
   src = "this is a string"
   desBytes := make([]byte, len(src))
   copy(desBytes, src)
   dst = string(desBytes)
   fmt.Println(src, dst)
}

func main() {
   var dst, src string
   src = "this is a string"
   dst = strings.Clone(src)
   fmt.Println(src, dst)
}
```

## 拼接

与python一样可以使用`+`拼接

也可以转换后再添加元素

```go
//使用+
func main() {
   str := "this is a string"
   str = str + " that is a int"
   fmt.Println(str)
}
//
func main() {
	str := "this is a string"
	str2 := " + this is a new string"
	bytes := []byte(str)
	bytes = append(bytes, str2...)
	str = string(bytes)
	fmt.Println(str)
}
```

以上两种方式性能都很差，一般情况可以使用，要是对性能有更高要求，使用`strings.Builder`

```go
func main() {
	builder := strings.Builder{}
	builder.WriteString("this is a string ")
	builder.WriteString("this is another string")
	fmt.Println(builder.String())
}
```



## 字符串的使用

强类型语言： 不同类型的变量无法相互使用和计算

## strconv

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 10
    //int转string
	i2s := strconv.Itoa(i)
    //string装int
	s2i, err := strconv.Atoi(i2s)
	fmt.Println(i2s, s2i, err)
}

//输出：10 10 <nil>

//字符串（数字）转浮点型
strconv.ParseFloat()
//字符串（0/1）转布尔
strconv.ParseBool()
//布尔转字符串
strconv.FormatBool()
//浮点型转字符串
strconv.FormatFloat()
```

## 数字间转换

```go
//直接显式转换
i2f := float64(i)
f2i := int(64)
```

## strings

golang用于处理字符串的标准包

``` go
package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "Hello World!"
    //判断前缀是否为H
	fmt.Println(strings.HasPrefix(s1, "H"))
    //在s1中查找“o”，返回第一个匹配的index
	fmt.Println(strings.Index(s1, "o"))
    //全部转成大写
	fmt.Println(strings.ToUpper(s1))
}
```

