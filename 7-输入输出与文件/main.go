package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Rename("test.txt", "TEST.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("重命名成功")
	}
}
