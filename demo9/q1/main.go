package main

import "fmt"

func main() {
	aMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	k := "two"
	v, ok := aMap[k]
	if ok {
		fmt.Printf("The element of key %q: %d\n", k, v)
	} else {
		fmt.Println("Not Found")
	}

	// demo
	// panic: runtime error: hash of unhashable type []int
	// goroutine 1 [running]:
	// main.main()

	//var badMap2 = map[interface{}]int{
	//	"1":      1,
	//	[]int{2}: 2, // panic
	//	3:        3,
	//}
	//fmt.Println(badMap2)
}
