# 库

library： 模块的集合

- 标准库

  由go官方开发维护，go安装目录的src下有源代码，[这里](https://pkg.go.dev/std)是文档

- 第三方库

  [这里](https://pkg.go.dev/)可以搜索到github上开源go项目的依赖名，通过依赖名就可以下载，例如gin的依赖名：github.com/gin-gonic/gin

# Go mod

1.11后才推出的官方依赖管理解决方案，主要靠go.mod文件，用于记录项目中的依赖和版本。

```shell
#开启
go env -w GO111MODULE=on
```

# 下载

## 代理

- [七牛云 - Goproxy.cnopen in new window](https://goproxy.cn/)
- [GOPROXY.IO - 一个全球代理 为 Go 模块而生](https://goproxy.io/zh/)

``` shell
go env -w GO111MODULE=on
#七牛云
go env -w GOPROXY=https://goproxy.cn,direct
```

## go get

``` shell
go get [-t] [-u] [-x] [build flags] [packages]
# 一些例子
go get github.com/gin-gonic/gin@latest
go get golang.org/x/text@master
go get golang.org/x/text@v0.3.2
# 删除一个依赖
go get github.com/gin-gonic/gin@none
```

`-t`：构建测试所需要的模块，这些模块是在命令行指定的`packages`。

`-u`：更新指定的模块，当这些模块有新的镜像或者发行补丁版本。

`-x`：打印过程中执行的命令，通常用于调试。

> go get命令专门用于调整和修改`go.mod`文件中的依赖，如果没有特殊需求，用go get就足够了。

## go install

- 该命令会下载远程的包并编译成二进制文件放在`$GOROOT/bin`或者`$GOBIN`目录下
- 如果指定版本，会自动忽略go.mod文件，如果在模块外部执行该命令就必须指定版本
- 在模块内部使用该命令可以不用指定版本，但只能指定模块内已有的包

```shell
go install [build flags] [packages]
go install golang.org/x/tools/gopls@latest
```

# 使用流程

## 创建项目

项目即一个文件夹，例如创建一个new-project文件夹

## 初始化

```shell
#在项目（文件夹）内执行
go mod init new-project
```

这里的new-project也就是模块名，最好跟项目同名，此时go.mod文件会自动生成

```shell
#go.mod初始化后内容
module new-project // 模块名
go 1.19 // 使用Go的版本
```

## 下载依赖

```go
go get github.com/gin-gonic/gin@latest
```

此时会生成go.sum文件以及更新go.mod

## 导入运行

```go
package main

package main

import "github.com/gin-gonic/gin" // 导入包

func main() {
    // 运行
	gin.Default().Run(":8080")
}
```

```shell
#在项目内运行
go run new-project
```

# go.mod

一个模块的定义通常由go.mod文件中的module指定，每一行都包括一个指令，但是一般不建议手动编辑

```
module learn

go 1.19

require github.com/gin-gonic/gin v1.8.2

require (
   github.com/gin-contrib/sse v0.1.0 // indirect
   github.com/go-playground/locales v0.14.0 // indirect
   ...
)
```

- module：一个module指令定义了一个主模块的路径，每个go.mod文件都必须有且只有一个module指令
- require：声明了一个模块依赖项所需要的最小版本。有时会自动加上`// indirect` 注释表示间接依赖，比如这里，主模块依赖gin，gin依赖其他包，则其他包对于主模块而言就是间接依赖
- exclude：可以防止go命令加载该版本的依赖，只在主模块中应用，在子模块会被忽略
- replace：将会替换掉指定版本的依赖，仅`=>`左边的版本被替换

```
replace golang.org/x/net v1.2.3 => example.com/fork/net v1.4.5

replace (
    golang.org/x/net v1.2.3 => example.com/fork/net v1.4.5
    golang.org/x/net => example.com/fork/net v1.4.5
    golang.org/x/net v1.2.3 => ./fork/net
    golang.org/x/net => ./fork/net
)
```

- retract： 表示撤回，意思是不应该依赖所指定的版本，比如一个新的版本发布后发现重大问题，这时候就可以使用

# go.sum

> `go.sum`文件的存在是为了解决一致性构建的问题。依赖下载到本地后，会缓存到本地，以方便下次构建，下载的依赖和缓存的依赖都会有被纂改的可能，`go.sum`文件会记录每一个依赖的哈希值，如果`go.sum`文件中的哈希值与本地依赖的哈希值不同，则会拒绝构建，并且一般在下载依赖完成后还会请求环境变量`GOSUMDB`所指定的服务器查询一个可信的公共哈希值，如果哈希值不同的话也不会继续执行。可以前往[Go Modules - go.sum fileopen in new window](https://go.dev/ref/mod#go-sum-files)以查看更多细节。总的来说，该文件存在的意义就是确保下载的依赖是安全的与远程仓库是内容一致且未被修改的。

这个文件就更不应该手动改了。

# go常用命令

| 命令                 | 说明                       |
| -------------------- | -------------------------- |
| `go mod download`    | 下载当前项目的依赖包       |
| `go mod edit`        | 编辑go.mod文件             |
| `go mod graph`       | 输出模块依赖图             |
| `go mod init`        | 在当前目录初始化go mod     |
| `go mod tidy`        | 清理项目模块               |
| `go mod verify`      | 验证项目的依赖合法性       |
| `go mod why`         | 解释项目哪些地方用到了依赖 |
| `go clean -modcache` | 用于删除项目模块依赖缓存   |
| `go list -m`         | 列出模块                   |

# 工作区

workspace是1.18后引入的关于多模块管理的一个新解决方案。在就旧版本，如果想要在本地依赖其他模块但又没有上传到远程仓库，一般都需要使用replace指令，例如：

```shell
#有如下文件结构
learn
	-main
	|--	main.go
	|--	go.mod
	-tool
	|--	util.go
	|--	go.mod
#假如main.go想要导入tool模块下的一个函数，则需要将main.go的go.mod文件修改如下
module main

go 1.19

require (
   tool v0.0.0
)
replace (
   tool => "../utils" // 使用replace指令指向本地模块
)
#这样main.go才能直接导入
package main

import (
   "fmt"
   "tool"
)

func main() {
   fmt.Println(tool.StringMsg())
}
```

也可以将tool模块上传到远程仓库让后发布tag，然后main模块使用go get  -u更新。

workspace就是为了解决这样的问题：

```shell
#直接在项目目录下使用
go work init main tool
#此时会自动生成一个go.work文件
go 1.19

use (
   ./main
   ./tool
)
```

这样子main模块就能直接使用tool模块了，无需replace也无需上传。

```shell
#禁用工作区模式
go run -workfile=off main.go
```





