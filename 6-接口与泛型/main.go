package main

import "fmt"

type SignedInt interface {
	int8 | int16 | int | int32 | int64
}

type Integer interface {
	int8 | int16 | int | int32 | int64 | uint8 | uint16 | uint | uint32 | uint64
}

// Number接口就是包含的SignedInt和Integer两个类型集的交集
// 所以该接口允许的类型集式：int8 | int16 | int | int32 | int64
type Number interface {
	SignedInt
	Integer
}

func Do[T Number](n T) T {
	return n
}

func main() {
	r := Do[int](2)
	fmt.Println(r)
}
