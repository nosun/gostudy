package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s) // constant

	const n = 100

	const d = 3e20 /n

	fmt.Println(d) // 3e18

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}