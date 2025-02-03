package main

import "fmt"

func add(a, b int) {
	c := a + b
	fmt.Println(c)
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

type address struct {
	street  string
	zip     int
	country string
}

type info struct {
	name string
	age  int
	address
}

func main() {
	person1 := info{"John", 25, address{"123 Main St", 12345, "USA"}}
	fmt.Println(person1)
	fmt.Println(sum(1, 2, 3, 4, 5))
	add(1, 2)

}
