//Package fizzbuzz genrates fizzbuzz game sequence
package fizzbuzz

import "fmt"

//returnfizzbuzz returns string representation of fizz buzz sequence
func returnfizzbuzz(number int) (result string) {

	if number%3 == 0 {
		result += "Fizz"
	}
	if number%5 == 0 {
		result += "Buzz"
	}
	return
}

//GoUptoANumber prints fizz buzz sequence
func GoUptoANumber(number int) (result string) {

	for i := 1; i <= number; i++ {
		s := returnfizzbuzz(i)

		if s == "" {
			result += fmt.Sprintf("%d,", i)
		} else {

			result += s + ","
		}

	}

	return
}
