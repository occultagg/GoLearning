# 结构体

> Go抛弃了类与继承，同时也抛弃了构造方法，刻意弱化了面向对象的功能，Go并非是一个OOP的语言，但是Go依旧有着OOP的影子，通过结构体和方法也可以模拟出一个类。结构体可以存储一组不同类型的数据，是一种复合类型

## 声明

```go
type Programmer struct {
	Name     string
	Age      int
	Job      string
	Language []string
}
```

- 结构体本身以及其内部的字段都遵守大小写命名的暴露方式
- 类型相同的字段可以合并声明

```go
type Rectangle struct {
	height, width, area int
	color string
}
```

## 创建

golang不存在构造方法，大多数情况下采用如下方式创建

```go
programmer := Programmer{
    Name: "jack",
    Age: "19",
    Job: "coder",
    Language: []string{"Go", "Python"},
}
```

- 可以省略字段名创建，但必须按顺序初始化所有字段
- 也可以专门写一个函数初始化，这种函数被称为`工厂方法`，这也是go为什么没有构造方法的原因

```go
package main

import "fmt"

//声明结构体
type Programmer struct {
	Name     string
	Age      int
	Job      string
	Language []string
}

//工厂方法
func NewProgrammer() Programmer {
	return Programmer{
		"Jack",
		19,
		"coder",
		[]string{"Go", "Python"},
	}
}

func main() {
	p := NewProgrammer()
	fmt.Println(p)
}
```

## 组合

结构体之间的关系是通过组合来表示的，可以显式也可以匿名，后者使用起来更类似于继承，但本质上没有任何变化

### 显式组合

```go
package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Student struct {
	p      Person
	school string
}

type Employee struct {
	p   Person
	job string
}

//使用时需要显式指定字段p
func main() {
	student := Student{
		p:      Person{name: "jack", age: 18},
		school: "foda",
	}
	fmt.Println(student.p.name)
}

```

### 匿名组合

```go
package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Student struct {
	Person
	school string
}

type Employee struct {
	Person
	job string
}

func main() {
	student := Student{
        //匿名字段的名称默认为类型名
		Person: Person{name: "jack", age: 18},
		school: "foda",
	}
    //调用者可以直接访问该类型的字段和方法
	fmt.Println(student.name)
}
```

匿名组合除了更方便外和显式组合无任何区别

## 结构体指针

不需要解引用就能直接访问结构体指针的内容

```go
package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	p := &Person{
		name: "jack",
		age:  18,
	}
	fmt.Println(p.age, p.name)
}
```

实际上还是需要解引用，编译时go自动会转换成`(*p).name, (*p).age`

## 标签

> 结构体标签是一种元编程的形式，结合反射可以做出很多奇妙的功能

```go
`key1:"val1" key2:"val2"`
```

- 标签是一种键值对的形式，使用空格分隔
- 容错性很低，如果没有按照正确格式书写结构体，会导致无法正常读取，但是编译时没有任何报错
- 结构体标签最广泛的应用就是在各种序列化格式中的别名定义，标签的使用需要结合反射才能完整发挥出其功能

```go
type Programmer struct {
	Name     string   `json:"name"`
	Age      int      `yaml:"age"`
	Job      string   `toml:"job"`
	Language []string `properties:"language"`
}
```
# 函数

> 在go中，函数是一等公民，函数是go最基础的组成部分，也是核心

## 声明

```go
func funcName(params) returnType {
    body
}
//直接声明
func DoSomething() {
    
}

//字面量声明
var DoSomething func()

//作为一个类型
type DoSomething func()
```

## 参数

- 可以有名称也可以没有名称
- 形参：函数中定义的变量，只能在函数体内使用
- 实参：调用者调用函数时传的值

```go
//声明一个函数字面量时可以省略名称但在赋值时依旧需要名称
var sum func(int, int) int

sum = func(a int, b int) int {
    return a + b
}

//类型相同且相邻的参数，可以只声明一次类型
func sum(a, b int) int {
    return a + b
}

//       形参
func sum(a int, b int) int {
	return a + b
}

func main() {
    //            实参
	result := sum(1, 2)
	fmt.Println(result)
}
```

##  返回值

- 如果函数有多个返回值，returnType要用()括起来

```go
package main

import (
	"errors"
	"fmt"
)

func sum(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或b不能是负数")
	}
	return a + b, nil
}

func main() {
	result, err := sum(1, -2)
	fmt.Println(result, err)
}
```

- 返回值有多个时，也需要多个变量接收
- 不需要的返回值可以用占位符接收“_”

### 返回值命名

- 名字在前，类型在后

- 效果一致，不常用

```go
func sum(a, b int) (sum int, err error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或b不能是负数")
	}
	sum = a + b
	err = nil
	return
}

func main() {
	result, err := sum(1, -2)
	fmt.Println(result, err)
}

```

## 可变参数

- 参数的数量是可变的，可以为0，也可以一个或多个
- 可变参数的类型其实就是切片，for range 返回第一个值是index，这里不需要用占位符把它去掉
- 定义可变参数，就是在参数类型前加”...“

```go
func sum(a ...int) int {
	sum := 0
	for _, i := range a {
		sum += i
	}
	return sum
}

func main() {
	result := sum(1, -2, 3)
	fmt.Println(result)
}
```

- 既有普通参数，又有可变参数，可变参数一定要放在参数列表的最末尾

``` go
func sum1(tip string, params ...int) {
} 
```

## 包级函数

- 函数名称首字母小写代表私有函数，只有在同一个包中才可以被调用
- 函数名称首字母大写代表公有函数，在不同包中也可以被调用
- 任何一个函数都会从属于一个包

## 匿名函数

只能在函数内部存在

```go
func main() {
    func(a, b int) int {
        return a + b
    }(1, 2)
}

//或者当函数参数是一个函数类型，此时名称不再重要，可以直接传递一个匿名函数
func main() {
    DoSum(1, 2, func(a, b int) int {
        return a + b
    })
}

func DoSum(a, b int, f func(int, int) int) int {
    return f(a, b)
}
```

## 闭包（Closure）

函数内部的匿名函数使用定义它的作用域之外的变量

主要应用场景主要有：

1. 实现高阶函数，即将一个函数作为另一个函数的参数或返回值。
2. 实现计数器，每次调用闭包时将计数器的值加1。
3. 实现工厂函数，每次调用都返回一个新的闭包，用于创建不同的对象。
4. 实现延迟执行，在需要时才执行闭包中的代码。
5. 实现私有变量，通过闭包隐藏内部状态，只暴露必要的接口。
6. 实现单例模式，通过闭包限制变量的作用域，保证全局只有一个实例。

```go
func main() {
    //这里Sum函数已经执行完毕
    sum := Sum(5)
    //这里是对其返回的匿名函数传参
    fmt.Println(sum(1, 2))
    fmt.Println(sum(1, 2))
    fmt.Println(sum(1, 2))
}

func Sum(sum int) func(int, int) int {
    return func(a, b int) int {
       //Sum函数内部的匿名函数可以使用外部Sum函数的sum变量
       sum += a + b
       return sum
   }
}
```

> 匿名函数引用了参数`sum`，即便`Sum`函数已经执行完毕，虽然已经超出了它的生命周期，但是对其返回的函数传入参数，依旧可以成功的修改其值，这一个过程就是闭包。事实上参数`sum`已经逃逸到了堆上，只要其返回值函数的生命周期没有结束，就不会被回收掉。

## 延迟调用

- `defer`关键字描述一个匿名函数会在函数返回之前执行，多个defer会放在一个堆栈中，遵循FILO
- 通常用于释放文件资源关闭连接等操作，或者用于捕获panic

```go
package main

import "fmt"

func Sum(a, b int) int {
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(2)
	}()
	return a + b
}

func main() {
	s := Sum(2, 2)
	fmt.Println(s)
}

//输出：
//2
//1
//4
```

# 方法

>  方法与函数是两个概念，但又非常相似，不同点在于方法必须有一个接收者，这个接收者是一个自定义类型，这样方法就与这个类型绑定在一起，称为这个类型的方法.

```go
type IntSlice []int

func (i IntSlice) Get(index int) int {
    return i[index]
}

func (i InSlice) Set(index, val int) {
    i[index] = val
}

func (i IntSlice) Len() int {
    return len(i)
}
```

1. 先声明一个类型IntSlice，其底层类型是[]int
2. 再声明三个方法Get,Set,Len
3. 方法的长相与函数无太大区别，只是多了个接收者和接收者的类型`(i IntSlice)`
4. `i`就是接收者，`IntSlice`就是接收者的类型
5. 接收者`i`就类似其他语言的`this`或`self`，只不过golang需要显式指明

```go
//方法的调用类似调用一个类的方法，先声明，再初始化，再调用
func main() {
    //先声明
	var intSlice IntSlice
    //初始化，类似其他语言类的实例化
	intSlice = []int{1, 2, 3, 4, 5}
	fmt.Println(intSlice.Get(0))
	intSlice.Set(0, 2)
	fmt.Println(intSlice)
	fmt.Println(intSlice.Len())
}
```

## 值类型接收者于指针类型接收者

```go
type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Email is %s \n", u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	peter := user{"peter", "peter@oldexample.com"}
	peter.notify()
	peter.changeEmail("peter@example.com")
	peter.notify()
}
```

当接收者是一个值类型时`(u user)`，此时的接收者是一个副本，修改副本的值不影响接收者本身

当接收者是一个指针类型`(u *user)`，此时的接收者就是引用的，可以修改

另外，golang在编译时会自动隐式转换，调用时将其转换为正确的接收者类型

```go
//peter.changeEmail("peter@example.com")
(&peter).changeEmail("peter@example.com")

//u.email = email
(*u).email = email
```

## 方法表达式

```go
func (age Age) String() {
	fmt.Println("the age is", age)
}

func main() {
	age := Age(25)
    //方法赋值给变量，方法表达式
	sm := Age.String
    //通过变量调用方法要传一个接收者age
	sm(age)
}
```

- 方法string是没有参数的，在通过方法表达式赋值给变量sm后，在调用的时候，必须传一个接收者，这样sm才知道怎么调用
- 不管方法是否有参数，通过方法表达式调用时，第一个参数必须是接收者，然后才是方法自身的参数

## 小结

> 在调用方法时传递的接收者本质上都是副本，只不过一个是这个值的副本，一个是指向这个值的指针的副本。指针具有指向原有值的特性，所以修改了指针指向的值，也就修改了原有的值。我们可以简单理解为值接收者使用的是值的副本来调用方法
