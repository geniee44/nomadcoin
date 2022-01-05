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

	// for _, food := range foods {
	// 	fmt.Println(food)
	// }

	//위와 동일하지만 loop 방식으로
	for i := 0; i < len(foods); i++ {
		fmt.Println(foods[i])
	}

	fmt.Println("\nslices ver.")
	//slice <- array는 길이를 미리 명시해둬야 하기 때문에 블록체인 같이 길이를 예측할 수 없는 경우 slice 사용
	dishes := []string{"potato", "pizza", "pasta"}

	for i := 0; i < len(dishes); i++ {
		fmt.Println(dishes[i])
	}
}
