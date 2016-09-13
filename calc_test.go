package fizzbuzz_test

import (
	"fmt"
	"testing"

	"github.com/eleanorrigby/fizzbuzz"
)

func TestFizzbuzz_1(t *testing.T) {
	actual := ""
	actual = fizzbuzz.GoUptoANumber(1)
	expectation := "1,"
	if actual != expectation {
		t.Fatalf("Expected %s got %s ", expectation, actual)
	}
}

func TestFizzbuzz_4(t *testing.T) {
	actual := ""
	actual = fizzbuzz.GoUptoANumber(4)
	expectation := "1,2,Fizz,4,"
	if actual != expectation {
		t.Fatalf("Expected %s got %s ", expectation, actual)
	}
}

func TestFizzbuzz_15(t *testing.T) {
	actual := ""
	actual = fizzbuzz.GoUptoANumber(15)
	expectation := "1,2,Fizz,4,Buzz,Fizz,7,8,Fizz,Buzz,11,Fizz,13,14,FizzBuzz,"
	if actual != expectation {
		t.Fatalf("Expected %s got %s ", expectation, actual)
	}
}

func BenchmarkFizzbuzz(b *testing.B) {

	for i := 0; i < b.N; i++ {
		fizzbuzz.GoUptoANumber(15)
	}

}

func ExampleGoUptoANumber() {
	fmt.Println(fizzbuzz.GoUptoANumber(5))
	// Output:
	// 1,2,Fizz,4,Buzz,
}
