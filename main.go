package main

import "fmt"

func main() {
	// arrays (sizes cannot be change)

	// var ages [3]int = [3]int{20, 25, 30} // standard format
	// var ages = [3]int{20, 25, 30} // shorthand
	ages := [3]int{20, 25, 30} // shorthand
	names := [4]string{"yoshi", "mario", "peach", "bowser"}

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	names[1] = "luyigi"
	fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	scores := []int{100, 50, 60}
	fmt.Println(scores, len(scores))

	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:3]  // start from position 1 and exclude index 3
	rangeTwo := names[2:]   // start from position 2
	rangeThree := names[:3] // start from position 0 and exclude 3
	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)
}
