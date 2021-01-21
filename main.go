package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Maps in Go can be considered the equivalent of dict in python")
	mymap := make(map[int]string)

	mymap[3] = "three"
	mymap[4] = "four"
	mymap[5] = "five"
	fmt.Println("All element in mymap!!!")
	for key, value := range mymap {
		fmt.Println(key, value)
	}
	fmt.Println("_______")

	// var arrKey []int
	// for key, _ := range mymap {
	// 	arrKey = append(arrKey, key)
	// }
	// fmt.Println("arrKey", arrKey)

	// delete one key and value
	delete(mymap, 4)
	for key, value := range mymap {
		fmt.Println(key, value)
	}
	fmt.Println("_______")

}
