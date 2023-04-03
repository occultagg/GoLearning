# 流程控制

## if

- 条件表达式不需要()
- 一定要有大括号
- if后的 “{” 和else前的 “}” 不能单独一行

```go
package main

import (
	"fmt"
)

func main() {
	if i := 6; i > 10 {
		fmt.Println("i>10")
	} else if i > 5 && i <= 10 {
		fmt.Println("5<i<=10")
	} else {
		fmt.Println("i<=5")
	}
}
```

## switch

- case后自带break，一旦符合条件就不再往下执行

```go
func main() {
	switch i := 6; {
	case i > 10:
		fmt.Println("i>10")
	case i > 5 && i <= 10:
		fmt.Println("5<i<=10")
	default:
		fmt.Println("i<=5")
	}
}
```

- fallthourgh关键字表示执行下一个紧跟的case
- switch后可以跟表达式，比较表达式也可以
- 当switch后跟表达式时，case的值要跟表达式的结果类型一致，不然不能编译

```go
func main() {
	switch i := 6; i {
	case 6:
		fallthrough
	case 5:
		fmt.Println("5")
	default:
		fmt.Println("没有匹配")
	}
}

//输出：5
```

## for

- 经典例子  for <表达式>;<条件表达式>;<更新表达式> {}

```go
sum := 0
for i := 1; i <= 100; i++ {
    sum += i
}
```

- for后面的三个表达式都可以省略，达到while true的效果

```go
func main() {
	i := 1
	for {
		i++
		fmt.Println(i)
	}
}
```

- 支持continue和break

```go
func main() {
	sum := 0
	i := 1
	for {
		sum += i
		i++
        //i>100时跳出循环
		if i > 100 {
			break
		}
	}
	fmt.Println("sum is:", sum)
}
```

```go
func main() {
	sum := 0
	for i := 1; i < 100; i++ {
        //i不是偶数时跳过本次循环
		if i%2 != 0 {
			continue
		}
		sum += i
	}
	fmt.Println("sum is:", sum)
}
```

## label

给代码块打上标签，配合goto，break，continue使用

```go
func main() {
    a := 1
    if a == 1 {
        //执行A标签的代码块，也就是fmt.Println("a")
        goto A
    } else {
        fmt.Println("b")
    }
    //给fmt.Println("a")打上A标签
    A:
    fmt.Println("a")
}
```

- 一般很少使用，跳来跳去影响可读性，性能消耗也是一个问题