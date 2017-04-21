package main

import "fmt"

func zeroval(ival int){
	ival = 0
}

func zeroptr(iptr *int)  { // type pointer
	*iptr = 0
}

func main()  {

	i := 1

	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i) // var pointer
	fmt.Println("zeroptr:",i)
	fmt.Println("pointer:", &i) // print var pointer
}