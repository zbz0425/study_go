package main

import (
	"fmt"
)

// 声明类型
type (
	isLegal  bool
	byte     int8
	rune     int32
	byteSize int64
    文本 string
)

var name = "zbz"

func main() {
    var byteSize = 66
    a,_,c,d,e,f := 2,1,2,3,4,1
    fmt.Println(byteSize)
    fmt.Println(a)
    fmt.Println(c)
    fmt.Println(d)
    fmt.Println(e)
    fmt.Println(f)
    var q float32 = 100.11
    fmt.Println(q)
    b:=int(q)
    fmt.Println(b)
}
