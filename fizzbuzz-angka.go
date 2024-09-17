package main

import "fmt"

func main() {
	fmt.Println("Angka FizzBuzz:")
	for  i := 1; i <= 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("Angka Fizz:")
	for  i := 1; i <= 100; i++ {
		if i % 3 == 0 && i  % 5 != 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("Angka Buzz:")
	for  i := 1; i <= 100; i++ {
		if i % 5 == 0 && i  % 3 != 0 {
			fmt.Println(i)
		}
	}
}