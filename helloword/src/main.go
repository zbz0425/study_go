//当前程序包名 入口包 main
package main

//导入其他包
import (
	"fmt"
)

// 全局变量的声明与赋值
const PI = 3.14

// 一般类型声明
var name = "zbz"

//声明一个变量
type newType int

//结构声明
type gotest struct{}

// 接口声明
type golang interface{}

// main函数程序入口
func main() {
	fmt.Println("hello " + name)
}
