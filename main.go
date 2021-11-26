package main

import "fmt"

func plus(a ...int) int {
    var total int
    for _, item := range a{
        total += item
    }
    return total
}
func main() {
	result := plus(2, 3, 4, 5)
	fmt.Println(result)
}
