package main

import "fmt"

func main() {
	fmt.Println("hello world")
	//array
	a := []int{1, 4}
	a1 := a[0]
	a2 := a[1]
	a[0] = 6
	a[1] = 7 //way of adding a valu to an slice not in array

	a = append(a, 8)  //way of adding a valu to an slice not in array
	a = append(a, 9)  //way of adding a valu to an slice not in array
	a = append(a, 10) //way of adding a valu to an slice not in array
	a = append(a, 11) //way of adding a valu to an slice not in array

	fmt.Printf("type %T value %v\n", a, a)
	fmt.Printf("type %T value %v\n", a1, a1)
	fmt.Printf("type %T value %v\n", a2, a2)
	//fmt.Printf("ba--type %T value %v\n", ba, ba)

	//loop through array
	for index, value := range a {
		fmt.Println(index, value)

	}
}
