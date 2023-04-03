# 输入输出

## 标准

控制台标准输出输入一般使用os包下：(明显就是运维必备包)

- Stdin：标准输入
- Stdout：标准输出
- Stderr：便准错误

## 输出

```
//调用os.Stdout
os.Stdout.WriteString("Hello World!")
//内置函数println
println("Hello World!")
//最推荐使用fmt.Println
fmt.Println("Hello World!")
```

> `fmt.Println`会用到反射，因此输出的内容通常更容易使人阅读，不过性能很差强人意。

## 格式化

| 0    | 格式化    | 描述                                           | 接收类型           |
| ---- | --------- | ---------------------------------------------- | ------------------ |
| 1    | **`%%`**  | 输出百分号`%`                                  | `任意类型`         |
| 2    | **`%s`**  | 输出`string`/`[] byte`值                       | `string`,`[] byte` |
| 3    | **`%q`**  | 格式化字符串，输出的字符串两端有双引号`""`     | `string`,`[] byte` |
| 4    | **`%d`**  | 输出十进制整型值                               | `整型类型`         |
| 5    | **`%f`**  | 输出浮点数                                     | `浮点类型`         |
| 6    | **`%e`**  | 输出科学计数法形式 ,也可以用于复数             | `浮点类型`         |
| 7    | **`%E`**  | 与**`%e`**相同                                 | `浮点类型`         |
| 8    | **`%g`**  | 根据实际情况判断输出`%f`或者`%e`,会去掉多余的0 | `浮点类型`         |
| 9    | **`%b`**  | 输出整型的二进制表现形式                       | `数字类型`         |
| 10   | **`%#b`** | 输出二进制完整的表现形式                       | `数字类型`         |
| 11   | **`%o`**  | 输出整型的八进制表示                           | `整型`             |
| 12   | **`%#o`** | 输出整型的完整八进制表示                       | `整型`             |
| 13   | **`%x`**  | 输出整型的小写十六进制表示                     | `数字类型`         |
| 14   | **`%#x`** | 输出整型的完整小写十六进制表示                 | `数字类型`         |
| 15   | **`%X`**  | 输出整型的大写十六进制表示                     | `数字类型`         |
| 16   | **`%#X`** | 输出整型的完整大写十六进制表示                 | `数字类型`         |
| 17   | **`%v`**  | 输出值原本的形式，多用于数据结构的输出         | `任意类型`         |
| 18   | **`%+v`** | 输出结构体时将加上字段名                       | `任意类型`         |
| 19   | **`%#v`** | 输出完整Go语法格式的值                         | `任意类型`         |
| 20   | **`%t`**  | 输出布尔值                                     | `布尔类型`         |
| 21   | **`%T`**  | 输出值对应的Go语言类型值                       | `任意类型`         |
| 22   | **`%c`**  | 输出Unicode码对应的字符                        | `int32`            |
| 23   | **`%U`**  | 输出字符对应的Unicode码                        | `rune`,`byte`      |
| 24   | **`%p`**  | 输出指针所指向的地址                           | `指针类型`         |

```go
//一般使用fmt.SprintF或fmt.Printf来格式化输出字符串

type person struct {
	name    string
	age     int
	address string
}

func main() {
	fmt.Printf("%%%s\n", "hello world")
	fmt.Printf("%s\n", "hello world")
	fmt.Printf("%q\n", "hello world")
	fmt.Printf("%v\n", person{"lihua", 22, "beijing"})
	fmt.Printf("%+v\n", person{"lihua", 22, "beijing"})
}
```

## 输入

一般使用fmt包的三个函数

```go
// 扫描从os.Stdin读入的文本，根据空格分隔，换行也被当作空格
func Scan(a ...any) (n int, err error) 

// 与Scan类似，但是遇到换行停止扫描
func Scanln(a ...any) (n int, err error)

// 根据格式化的字符串扫描
func Scanf(format string, a ...any) (n int, err error)
```

```go
//fmt.Scan()
func main() {
    var s, s2 string
    //回车也被当作空格，输入两个值后Scan函数执行完毕
    fmt.Scan(&s, &s2)
    fmt.Println(s, s2)
}

//fmt.Scanln()
func main() {
	var s, s2 string
	fmt.Scanln(&s, &s2)
	fmt.Println(s, s2)
}

//fmt.Scanf()
func main() {
    var s, s2, s3 string
	//根据给定格式化，判断输入是否结束
    scanf, err := fmt.Scanf("%s %s \n %s", &s, &s2, &s3)
    if err != nil {
        fmt.Println(scanf, err)
    }
    fmt.Println(s)
    fmt.Println(s2)
    fmt.Println(s3)
}
//比如本例子中Scanf获取三个参数，输入格式应该如下：
/*
<s1> <s2>
<s3>
*/
```

## 缓冲

对性能有要求可以使用`bufio`包进行读写

```go
//读
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    fmt.Println(scanner.Text())
}


//写
func main() {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString("hello world!\n")
	//刷新缓冲区，同时输出缓冲区内容
	writer.Flush()
	//返回当前缓冲区的字节数
	fmt.Println(writer.Buffered())
}
```

# 文件操作

> go文件操作的基本数据类型支持的式字节切片[]byte，标准库的包包括os，io等，ioutil已经被弃用，不建议使用

## 打开

一般使用os包的两个函数：

- Open()： 返回一个文件指针和一个错误

  ```go
  func Open(name string) (*File, error)
  ```

  ```go
  func main() {
  	file, err := os.Open("README.txt")
  	fmt.Println(file, err)
  }
  ```

  - 文件不存在返回错误： <nil> open README.txt: The system cannot find the file specified.

    ```go
    //判断文件是否存在
    func main() {
    	file, err := os.Open("README.txt")
    	if os.IsNotExist(err) {
    		fmt.Println("文件不存在")
    	} else if err != nil {
    		fmt.Println("文件访问异常")
    	} else {
    		fmt.Println("文件读取成功", file)
    	}
    }
    ```

  > 事实上open函数读取的文件只是只读无法修改

- OpenFile()：事实上Open()是OpenFile()的一个简单封装，后者提供更加细粒度的控制，例如修改文件描述符和文件权限

  ```go
  func OpenFile(name string, flag int, perm FileMode) (*File, error)
  ```

  - 文件描述符

    ```go
    const (
       // 只读，只写，读写 三种必须指定一个
       O_RDONLY int = syscall.O_RDONLY // 以只读的模式打开文件
       O_WRONLY int = syscall.O_WRONLY // 以只写的模式打开文件
       O_RDWR   int = syscall.O_RDWR   // 以读写的模式打开文件
       // 剩余的值用于控制行为
       O_APPEND int = syscall.O_APPEND // 当写入文件时，将数据添加到文件末尾
       O_CREATE int = syscall.O_CREAT  // 如果文件不存在则创建文件
       O_EXCL   int = syscall.O_EXCL   // 与O_CREATE一起使用, 文件必须不存在
       O_SYNC   int = syscall.O_SYNC   // 以同步IO的方式打开文件
       O_TRUNC  int = syscall.O_TRUNC  // 当打开的时候截断可写的文件
    )
    ```

  - 文件权限

    ```go
    const (
       ModeDir        = fs.ModeDir        // d: 目录
       ModeAppend     = fs.ModeAppend     // a: 只能添加
       ModeExclusive  = fs.ModeExclusive  // l: 专用
       ModeTemporary  = fs.ModeTemporary  // T: 临时文件
       ModeSymlink    = fs.ModeSymlink    // L: 符号链接
       ModeDevice     = fs.ModeDevice     // D: 设备文件
       ModeNamedPipe  = fs.ModeNamedPipe  // p: 具名管道 (FIFO)
       ModeSocket     = fs.ModeSocket     // S: Unix 域套接字
       ModeSetuid     = fs.ModeSetuid     // u: setuid
       ModeSetgid     = fs.ModeSetgid     // g: setgid
       ModeCharDevice = fs.ModeCharDevice // c: Unix 字符设备, 前提是设置了 ModeDevice
       ModeSticky     = fs.ModeSticky     // t: 黏滞位
       ModeIrregular  = fs.ModeIrregular  // ?: 非常规文件
    
       // 类型位的掩码. 对于常规文件而言，什么都不会设置.
       ModeType = fs.ModeType
    
       ModePerm = fs.ModePerm // Unix 权限位, 0o777
    )
    ```

  ```go
  func main() {
      //以读写的模式打开文件 如果文件不存在则创建文件 权限：0666
      file, err := os.OpenFile("README.txt", os.O_RDWR|os.O_CREATE, 0666)
      if os.IsNotExist(err) {
          fmt.Println("文件不存在")
      } else if err != nil {
          fmt.Println("文件访问异常")
      } else {
          fmt.Println("文件打开成功", file.Name())
          file.Close()
      }
  }
  ```

## 读取

### 读取文件信息

`os.lstat()`

```go
func main() {
	fileInfo, err := os.Lstat("README.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fmt.Sprintf("%+v", fileInfo))
	}
}
```

### 读取文件内容

先打开，再读取，`os.File`类型提供了以下几个公开的方法

```go
// 将文件读进传入的字节切片
func (f *File) Read(b []byte) (n int, err error) 

// 相较于第一种可以从指定偏移量读取
func (f *File) ReadAt(b []byte, off int64) (n int, err error) 
```

一般使用第一种，但是需要自己处理切片容量的逻辑，比较麻烦，具体看[这里](https://golang.halfiisland.com/%E8%AF%AD%E8%A8%80%E5%85%A5%E9%97%A8/%E8%AF%AD%E6%B3%95%E8%BF%9B%E9%98%B6/100.io.html#%E8%AF%BB%E5%8F%96)

#### 方便的函数

- `os.ReadFile()`：只需要提供文件名

  ```go
  func main() {
  	bytes, err := os.ReadFile("README.txt")
  	if err != nil {
  		fmt.Println(err)
  	} else {
  		fmt.Println(string(bytes))
  	}
  }
  ```

- `io.ReadAll()`: 需要提供一个io.Reader类型的实现（比如`OpenFile()`返回的文件指针）

  ```go
  func main() {
     file, err := os.OpenFile("README.txt", os.O_RDWR|os.O_CREATE, 0666)
     if err != nil {
        fmt.Println("文件访问异常")
     } else {
        fmt.Println("文件打开成功", file.Name())
        bytes, err := io.ReadAll(file)
        if err != nil {
           fmt.Println(err)
        } else {
           fmt.Println(string(bytes))
        }
        file.Close()
     }
  }
  ```

## 写入

```go
//os.O_TRUNC: truncate 清空
//os.O_APPEND: 数据添加到尾部
func main() {
    //必须以O_RDWR或O_WRONLY的模式打开才能写入
	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("文件访问异常")
	} else {
		fmt.Println("文件打开成功", file.Name())
		for i := 0; i < 5; i++ {
			offset, err := file.WriteString("Hello World!\n")
			if err != nil {
				fmt.Println(offset, err)
			}
		}
		fmt.Println(file.Close())
	}
}
```

### 方便的函数

- os.WriteFile

```go
//打开模式默认跟上面的例子一样，无需手动打开关闭
func main() {
	err := os.WriteFile("test.txt", []byte("hello world!\n"), 0666)
	if err != nil {
		fmt.Println(err)
	}
}
```

- io.WriteString

```go
func main() {
	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("文件访问异常")
	} else {
		fmt.Println("文件打开成功", file.Name())
		for i := 0; i < 5; i++ {
			offset, err := io.WriteString(file, "hello world!\n")
			if err != nil {
				fmt.Println(offset, err)
			}
		}
		fmt.Println(file.Close())
	}
}
```

- os.Create

```go
//用于创建文件,如果其父目录不存在，将创建失败并会返回错误
func Create(name string) (*File, error) {
   return OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
}
```

## 复制

复制文件的时候需要同时打开两个文件

一种方法是将原文件数据读出来，然后写入目标文件中，使用`os.ReadFile()`和`os.WriteFile()`

```go
func main() {
    // 从原文件中读取数据
	data, err := os.ReadFile("README.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
    // 写入目标文件
	err = os.WriteFile("README(1).txt", data, 0666)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("复制成功")
	}
}
```

另一种方法是使用`os.File`的`ReadFrom`,打开文件时，一个只读一个只写

```go
func main() {
    // 以只读的方式打开原文件
	origin, err := os.OpenFile("README.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
    defer origin.Close()
    // 以只写的方式打开副本文件
	target, err := os.OpenFile("README(1).txt", os.O_RDONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
    defer target.Close()
    // 从原文件中读取数据，然后写入副本文件
	offset, err := target.ReadFrom(origin)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("文件复制成功", offset)
}
```

### 方便的函数

- io.Copy()

  除了复制文件，也能复制文件夹

```go
func main() {
	//以只读方式打开源文件
	origin, err := os.OpenFile("test.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer origin.Close()
	//以只写方式打开目标文件
	dst, err := os.OpenFile("test1.txt", os.O_RDONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dst.Close()
	//复制
	written, err := io.Copy(dst, origin)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(written)
	}
}
```

## 重命名

- os.Rename()

```go
func main() {
	err := os.Rename("test.txt", "TEST.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("重命名成功")
	}
}
```

## 删除

- os.Remove()

  删除单个文件或者空目录

  ```go
  func main() {
  	// 删除当前目录下文件
  	err := os.Remove("README.txt")
  	if err != nil {
  		fmt.Println(err)
  	} else {
  		fmt.Println("删除成功")
  	}
  }
  ```

- os.RemoveAll()

  删除指定目录的所有文件和目录包括的子目录和子文件

  ```go
  func main() {
  	// 删除当前目录下所有的文件与子目录
  	err := os.RemoveAll(".")
  	if err != nil {
  		fmt.Println(err)
  	}else {
  		fmt.Println("删除成功")
  	}
  }
  ```

# 文件夹

## 打开文件夹

- `os.ReadDir()`

  ``` go
  func main() {
     // 当前目录
     dir, err := os.ReadDir(".")
     if err != nil {
        fmt.Println(err)
     } else {
        for _, entry := range dir {
           fmt.Println(entry.Name())
        }
     }
  }
  ```

- `*os.File.ReaDir`

  ```go
  // n < 0时，则读取文件夹下所有的内容
  func (f *File) ReadDir(n int) ([]DirEntry, error)
  ```

  ```go
  func main() {
     // 当前目录
     dir, err := os.Open(".")
     if err != nil {
        fmt.Println(err)
     }
     defer dir.Close()
     dirs, err := dir.ReadDir(-1)
     if err != nil {
        fmt.Println(err)
     } else {
        for _, entry := range dirs {
           fmt.Println(entry.Name())
        }
     }
  }
  ```

## 创建文件夹

os包下的两个函数

```go
// 用指定的权限创建指定名称的目录
func Mkdir(name string, perm FileMode) error 

// 相较于前者该函数会创建一切必要的父目录
func MkdirAll(path string, perm FileMode) error
```



