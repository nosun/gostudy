package main

import "fmt"

type person struct {
	name string
	age int
}

func main(){

	fmt.Println(person{"nosun",100})

	fmt.Println(person{name:"nosun",age:100})

	fmt.Println(person{name:"nosun1"})

	fmt.Println(&person{name:"nosun2",age:200})

	s := person{name:"nosun3", age: 99}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 30
	fmt.Println(s.age)
}