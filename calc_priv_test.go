package fizzbuzz

import "testing"

func TestReturnfizzbuzz3(t *testing.T) {
	actual := ""
	actual = returnfizzbuzz(3)
	expectation := "Fizz"
	if actual != expectation {
		t.Fatalf("Expected %s got %s ", expectation, actual)
	}
}

func TestReturnfizzbuzz4(t *testing.T) {
	actual := ""
	actual = returnfizzbuzz(4)
	expectation := ""
	if actual != expectation {
		t.Fatalf("Expected %s got %s ", expectation, actual)
	}
}

func TestReturnfizzbuzz15(t *testing.T) {
	actual := ""
	actual = returnfizzbuzz(15)
	expectation := "FizzBuzz"
	if actual != expectation {
		t.Fatalf("Expected %s got %s ", expectation, actual)
	}
}

func TestReturnfizzbuzz5(t *testing.T) {
	actual := ""
	actual = returnfizzbuzz(5)
	expectation := "Buzz"
	if actual != expectation {
		t.Fatalf("Expected %s got %s ", expectation, actual)
	}
}
