package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			// 多数情况未打印任何东西
			fmt.Println(i)
		}()
	}
}