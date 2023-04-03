# 基本语法

```go
package main

import "fmt"

func main() {
   fmt.Println("Hello 世界!")
}
```

> package关键字代表当前go文件属于哪一个包，启动文件通常是main包，启动函数是main函数
>
> import是导入关键字，后面跟着被导入的包名
>
> func是函数声明关键字，用于声明一个函数
>
> fmt.Println("Hello 世界!")是一个语句，调用fmt包下的Println函数进行标准输出

# 包

> 在Go中，程序是通过将包链接在一起来构建的，也可以理解为最基本的调用单位是包，而不是go文件。包其实就是一个文件夹，包内共享所有源文件的变量，常量，函数以及其他类型。包的命名风格建议都是小写字母，并且要尽量简短。

## 导入

```go
//定义一个example的包
package example

import "fmt"

func SayHello() {
   fmt.Println("Hello")
}

//在main函数中调用
package main

import "example"

func main() {
   example.SayHello()
}

//可以给包起别名
package main

import e "example"

func main() {
   e.SayHello()
}

//导入多个包用括号括起来
package main

import (
   "fmt"
   "math"
)

func main() {
   fmt.Println(math.MaxInt64)
}

//可以只导入不调用，通常这么做是为了调用该包下的init函数
package main

import (
   "fmt"
    _ "math" // 下划线表示匿名导入
)

func main() {
   fmt.Println(1)
}
```

## 导出

如果想要对外暴露一个函数或者一个变量，只需要首字母大写

首字母小写表示只能在包内使用

本规则适用于整个golang

# 注释

```go
// 这是main包
package main

// 导入了fmt包
import "fmt"

/*
*
这是启动函数main函数
*/
func main() {
	// 这是一个语句
	fmt.Println("Hello 世界!")
}

```

# 标识符

用于命名，如包命名，函数命名，变量命名等，规则如下：

- 只能由字母，数字，下划线组成

- 只能以字母和下划线开头

- 严格区分大小写

- 不能与任何已存在的标识符重复，即包内唯一的存在

- 不能与Go任何内置的关键字冲突

  - 内置关键字包括：

    ```go
    break        default      func         interface    select
    case         defer        go           map          struct
    chan         else         goto         package      switch
    const        fallthrough  if           range        type
    continue     for          import       return       var
    ```

# 运算符

```go
Precedence    Operator
    5             *  /  %  <<  >>  &  &^
    4             +  -  |  ^
    3             ==  !=  <  <=  >  >=
    2             &&
    1             ||
```

## 赋值运算符

```go
a += 1
a /= 2
a &^= 2
```

## 自增自减

go只有一种写法

```go
a++
a--
--a //是错误的
a = b++ //也是错的，因为go的自增自减不具有返回值
```

# 字面量

> 字面量，按照计算机科学的术语来讲是用于表达源代码中一个固定值的符号，也叫字面值。两个叫法都是一个意思，写了什么东西，值就是什么，值就是“字面意义上“的值

## 整型字面量

为了便于阅读，允许使用下划线`_`来进行数字划分，但是仅允许在**前缀符号之后**和**数字之间**使用。

```go
24 // 24
024 // 24
2_4 // 24
0_2_4 // 24
10_000 // 10k
100_000 // 100k
0O24 // 20
0b00 // 0
0x00 // 0
0x0_0 // 0
```

## 浮点数字面量

通过不同的前缀可以表达不同进制的浮点数

```go
0.
72.40
072.40       // == 72.40
2.71828
1.e+0
6.67428e-11
1E6
.25
.12345E+5
1_5.         // == 15.0
0.15e+0_2    // == 15.0

0x1p-2       // == 0.25
0x2.p10      // == 2048.0
0x1.Fp+0     // == 1.9375
0X.8p-0      // == 0.5
0X_1FFFP-16  // == 0.1249847412109375
0x15e-2      // == 0x15e - 2 (integer subtraction)
```

## 复数字面量

```go
0i
0123i         // == 123i
0o123i        // == 0o123 * 1i == 83i
0xabci        // == 0xabc * 1i == 2748i
0.i
2.71828i
1.e+0i
6.67428e-11i
1E6i
.25i
.12345E+5i
0x1p-2i       // == 0x1p-2 * 1i == 0.25i
```

## 字符字面量

字符字面量必须使用`''`括起来，go中的字符完全兼容utf8

```go
'a'
'ä'
'你'
'\t'
'\000'
'\007'
'\377'
'\x07'
'\xff'
'\u12e4'
'\U00101234'
```

## 字符串字面量

字符串字面量必须使用双引号`""`括起来或者反引号（反引号字符串不允许转义）

```go
`abc`                // "abc"
`\n
\n`                  // "\\n\n\\n"
"\n"
"\""                 // `"`
"Hello, world!\n"
"今天天气不错"
"日本語"
"\u65e5本\U00008a9e"
"\xff\u00FF"
```

## 转移字符

```go
\a   U+0007 响铃符号（建议调高音量）
\b   U+0008 回退符号
\f   U+000C 换页符号
\n   U+000A 换行符号
\r   U+000D 回车符号
\t   U+0009 横向制表符号
\v   U+000B 纵向制表符号
\\   U+005C 反斜杠转义
\'   U+0027 单引号转义 (该转义仅在字符内有效)
\"   U+0022 双引号转义 (该转义仅在字符串内有效)
```



