# array（数组）

- 存放固定长度，相同类型的数据
- 数组中的元素在内存中是连续的
- 数据类型没有限定，自定义类型也行
- 数组有下标（index），从0开始
- 如果事先就知道了要存放数据的长度，且后续使用中不会有扩容的需求，就可以考虑使用数组
- go中的数组是值类型，也就是说数组是一个单独的类型，并不是指向头部元素的指针

```go
//初始化,长度必须是一个常量
var nums [5]int{1, 2, 3}
nums := [5]int{1, 2, 3}
//通过new函数获得一个指针
nums new([5]int)
//通过len()获得数组元素的数量cap()获取容量
len(nums)
cap(nums)
//
array := [5]string{"a", "b", "c", "d", "e"}
//长度可以省略，但要写...，还不如写了算了
array := [...]string{"a", "b", "c", "d", "e"}
//根据下标初始化值，没有初始化的默认为数组类型的零值
array := [5]string{1: "b", 3: "d"}
//使用for循环打印array中所有的元素,不推荐
func main() {
	array := [...]string{"a", "b", "c", "d", "e"}
	for i := 0; i < 5; i++ {
		fmt.Printf("数组索引:%d,对应值:%s\n", i, array[i])
	}
}
```

## 数组循环

推荐使用for...range...

```go
func main() {
	array := [...]string{"a", "b", "c", "d", "e"}
	for i, v := range array {
		fmt.Printf("数组索引:%d, 对应的值:%s", i, v)
	}
}
```

- range返回两个结果，一个是索引，一个是值
- 上面的例子分别将range的返回值赋给i和v

# Slice（切片）

- 左包含右不包含
- 切片基于数组生成，单独索引，修改切片的值等于修改数组的值
- [start : end]，start和end都可以省略，start默认0，end默认为数组长度，全省略就是整个数组（跟python一样）
- 切片在go中应用更广泛，用于存放不知道长度的数据，且后续使用过程中可能会频繁地插入和删除元素

## 切片的声明

```go
var slice1 []int //默认值为nil，不会为其分配内存
//make([]<type>, <长度>, <容量>)
slice1 := make([]string, 4, 8) // 使用make()初始化建议预分配同一个足够的容量，有效减少后续扩容的内存消耗
slice1 := new([]int) //返回指针
```

- 容量不能小于长度
- 长度就是切片内元素的个数
- 容量： 在上面的例子中golang在内存上划分了一块容量为8的内容空间，但是只有4个内存空间有元素（长度4），其他内存空间处于空闲状态，通过append函数往切片追加容量时，会追加到空闲的内存上

字面量声明和初始化

```go
//字面量初始化时长度==容量
func main() {
	slice2 := []string{"a", "b", "c", "d", "e"}
	fmt.Println(len(slice2), cap(slice2))
}
```

## append

### 追加元素

- 只能追加切片

```go
nums := make([]int, 0, 0)
nums = append(nums, 1, 2, 3, 4, 5, 6, 7)
fmt.Println(len(nums), cap(nums)) // 7 8 可以看到长度与容量并不一致。
```

### 插入元素

从头部插入

```go
nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
nums = append([]int{-1, 0}, nums...)
fmt.Println(nums) // [-1 0 1 2 3 4 5 6 7 8 9 10]
```

从中间下标i插入

```go
i := 3
nums = append(nums[:i+1], append([]int{999, 999}, nums[i+1:]...)...)
fmt.Println(nums) // [1 2 3 4 999 999 5 6 7 8 9 10]
```

从尾部插入（就是上面的追加）

## 删除元素

```go
nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
```

从头部删除n个元素

```go
n := 3
nums = nums[n:]
fmt.Println(nums) //[4 5 6 7 8 9 10]
```

从尾部删除n个元素

```go
n := 3
nums = nums[:len(nums)-n]
fmt.Println(nums) // [1 2 3 4 5 6 7]
```

从中间下标i位置开始删除n个元素

```go
i, n := 2, 3
nums = append(nums[:i], nums[i+n:]...)
fmt.Println(nums)//[1 2 6 7 8 9 10]
```

删除所有元素

```go
nums = nums[:0]
fmt.Println(nums) // []
```

## 切片循环

跟数组一样推荐使用for...range...

## 切片拷贝

必须确保目标切片有足够的长度

```go
func main() {
    //保证目标切片有足够长度
	dest := make([]int, 10)
	src := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(src, dest)
	fmt.Println(copy(dest, src))
	fmt.Println(src, dest)
}
```

## 多维切片

```go
var nums [5][5]int
for _, num := range nums {
   fmt.Println(num)
}
fmt.Println()
slices := make([][]int, 5)
for _, slice := range slices {
   fmt.Println(slice)
}
//输出
[0 0 0 0 0]
[0 0 0 0 0]
[0 0 0 0 0]
[0 0 0 0 0]
[0 0 0 0 0]

[]
[]
[]
[]
[]
//因为数组的长度早已固定，但是切片的长度是不固定的
```

应该改成这样初始化二维切片

```go
func main() {
	var nums [5][5]int
	for _, num := range nums {
		fmt.Println(num)
	}
	fmt.Println()
	slices := make([][]int, 5)
	for i := 0; i < len(slices); i++ {
		slices[i] = make([]int, 5)
	}
	for _, slice := range slices {
		fmt.Println(slice)
	}
}
```

## 扩展表达式

- 数组和切片都能使用表达式切割，只有切片能使用扩展表达式

  ```go
  //low <= high <= max <= cap
  slice[low:high:max]
  //切割出来的切片容量为max-low
  //省略max则表示最大容量
  s1 := []int(1, 2, 3, 4, 5, 6, 7, 8, 9) // cap=9
  s2 := s1[3:4] // cap=cap(s1)-3=9-3=6
  ```

  此时会有一个问题：

  由于s1和s2共享同一个数组，所以在对s2进行读写时，有可能会影响s1的数据

  ```go
  s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} // cap = 9
  s2 := s1[3:4]                          // cap = 9 - 3 = 6
  // 添加新元素，由于容量为6.所以没有扩容，直接修改底层数组
  s2 = append(s2, 1)
  fmt.Println(s1) //[1 2 3 4 1 6 7 8 9]
  fmt.Println(s2) //[4 1]
  ```

  此时就需要用到扩展表达式

  ```go
  func main() {
     s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} // cap = 9
     s2 := s1[3:4:4]                        // cap = 4 - 3 = 1
     // 容量不足，分配新的底层数组
     s2 = append(s2, 1)
     fmt.Println(s2)
     fmt.Println(s1)
  }
  ```

## 总结

> golang开发中优先选择切片作为函数的参数，因为它高效，内存占用小，通常情况下推荐使用make()来创建空切片；切片与数组的最大区别在于，切片的容量会自动扩张，而数组不会

# map（映射）

- go的map实现基于hash桶，所以是无序的

- 无序的K-V对，map[k]v，类似python的字典
- key的类型必须相同，value的类型可以不同
- key的类型必须支持“==”比较运算符，这样才能判断它是否存在并保证key的唯一性

## 创建与初始化

```go
//以这种方式创建的空map无法读写
var mp map[string]int
//字面量创建
map[keyType]valueType{}
//make创建,接收两个参数，分别是类型和初始容量(可以省略)
nameAgeMap := make(map[string]int, 8)
//添加键值对
nameAgeMap["Peter"] = 18
//字面量创建+初始化
nameAgeMap := map[string]int{"peter": 18}
//空map使用空的大括号，大括号不能省略
```

## 添加获取与删除

与array类似都是使用[]操作符

```go
//添加，key已存在则会更新value
nameAgeMap["adm"] = 20

//获取，当key不存在时，会返回value类型的零值，而不会报错
age := nameAgeMap["peter"]

//map的[]操作符会返回两个值
//一个是对应的value
//一个标记该key是否存在，布尔型
age, ok := nameAgeMap["aaa"]
if ok {
    fmt.Println(nameAgeMap, age)
}

//删除使用内置delete(<mapName>, <key>)
delete(nameAgeMap, "peter")
```

## 特殊情况

```go
//键为math.NaN()时，相同的键值不会被覆盖反而可以存在多个，也无法判断其是否存在，所以也无法正常取值
func main() {
	mp := make(map[float64]string, 10)
	mp[math.NaN()] = "a"
	mp[math.NaN()] = "b"
	mp[math.NaN()] = "c"
	_, exist := mp[math.NaN()]
	fmt.Println(exist) //false
	fmt.Println(mp) //map[NaN:c NaN:a NaN:b]
}
```

> 因为NaN是IEE754标准所定义的，其实现是由底层的汇编指令`UCOMISD`完成，这是一个无序比较双精度浮点数的指令，该指令会考虑到NaN的情况，因此结果就是任何数字都不等于NaN，NaN也不等于自身，这也造成了每次哈希值都不相同。关于这一点社区也曾激烈讨论过，但是官方认为没有必要去修改，所以应当尽量避免使用NaN作为map的键。但是NaN可以遍历访问。

## 遍历map

一样使用for range，返回两个值

- 第一个是map的key
- 第二个是map的value

```go
func main() {
	nameAgeMap := map[string]int{"peter": 18}
	nameAgeMap["adm"] = 20
	nameAgeMap["eve"] = 21

	for k, v := range nameAgeMap {
		fmt.Println("Key is", k, "value is", v)
	}
}
```

- map是无序的所以每次遍历出来的键值对顺序可能不一致
- 如果想按顺序遍历，可以根据排序好的key获取对应的value
- for range map也可以只返回一个值，默认是map的key

## map的长度

```go
len(nameAgeMap)
```

## set

set是一种无序的，不包含重复元素的的集合，但是golang没有提供类似的数据结构实现，而map的key正是无序且不能重复，所以可以用map来代替set。

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
    //以空结构体为value
	set := make(map[int]struct{}, 10)
	for i := 0; i < 10; i++ {
        //空结构体什么都不是，也不占内存空间
		set[rand.Intn(100)] = struct{}{}
	}
	fmt.Println(set)
}

```









