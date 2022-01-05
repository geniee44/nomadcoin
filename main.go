package main

import (
	"fmt"
)

func plus(a ...int) int {
	var total int
	for _, item := range a {
		total += item
	}
	return total
}
func main() {
	result := plus(2, 3, 4, 5)
	fmt.Println(result)

	//array
	foods := [3]string{"potato", "pizza", "pasta"}

	for _, food := range foods {
		fmt.Println(food)
	}

	//위와 동일하지만 loop 방식으로
	for i := 0; i < len(foods); i++ {
		fmt.Println(foods[i])
	}
}
