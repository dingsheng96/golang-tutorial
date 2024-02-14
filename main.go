package main

import "fmt"

var someName string = "hello"

func main() {
	// strings
	var nameOne string = "mario" // 1st method of declaring string variables
	var nameTwo = "luigi"        // 2nd method of declaring string variables
	var nameThree string         // default value is empty string
	nameFour := "yoshi"          // shorthand
	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	// integers
	var ageOne int = 20 // 1st method of declaring integer variables
	var ageTwo = 30     // 2nd method of declaring integer variables
	var ageThree int    // default value is 0
	ageFour := 40       // shorthand

	fmt.Println(ageOne, ageTwo, ageThree, ageFour)

	// bits & memory
	var numOne int8 = 25     // 8 bits will be used (range: -128 to 127)
	var numTwo uint = 25     // only unsigned integer can be use
	var numThree uint8 = 225 // (range: 0 to 255)

	fmt.Println(numOne, numTwo, numThree)

	// floats
	var scoreOne float32 = -1.5
	var scoreTwo float32 = 25.98
	var scoreThree float64 = 8881231289867.7
	scoreFour := 1.5

	fmt.Println(scoreOne, scoreTwo, scoreThree, scoreFour)

	// For more variables type, please refer to https://go.dev/ref/spec#Numeric_types

}
