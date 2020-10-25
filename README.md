## Go 学习笔记

### 一. 安装配置



### 二. 基本语法



### 三. 库函数

#### 1. flag

flag包是用来解析启动程序时传入的参数的 `func Int(name string, value int, usage string) *int`

- 参数1：参数的名称
- 参数2：启动程序是没有指定参数时的缺省值
- 参数3：参数的描述，运行`-help`时显示出来

```go
package main
import (
    "flag"
    "fmt"
)
func main() {
    num := flag.Int("n", 10, "number")
  flag.Parse()	// 必须执行 flag.Parse()，命令参数才能解析出来，不然访问变量只会得到默认值。
    fmt.Println(*num)
}

启动： ./flag_learn -n 123
输出： 123
```

