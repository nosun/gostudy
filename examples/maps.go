package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["a"] = 10
	m["b"] = 100

	fmt.Println("map:", m)

	v1 := m["a"]

	fmt.Println("v1:",v1)

	fmt.Println("len:",len(m))

	delete(m,"a")
	fmt.Println("map:",m)

	v,prs := m["b"]

	fmt.Println("prs:",v,prs)

	n :=map[string]int{"a":1,"b":2}

	fmt.Println(n)

}