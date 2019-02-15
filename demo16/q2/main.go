package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
    // 睡眠500毫秒
	time.Sleep(time.Millisecond * 500)
}
