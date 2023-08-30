package main

import (
	"fmt"
	
)

func test1(list *[]int){
	list = &[]int{}
}

func main() {
	list := []int{1,2,3,4,5}
	test1(&list)
	fmt.Print(list)
}
