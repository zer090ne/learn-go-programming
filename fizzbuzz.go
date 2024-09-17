// FizzBuzz adalah program yang mencetak angka dari 1 hingga 100,
// tetapi jika angka tersebut dapat dibagi dengan 3, maka cetak "Fizz",
// jika dapat dibagi dengan 5, maka cetak "Buzz", dan jika dapat dibagi
// dengan keduanya, maka cetak "FizzBuzz".

package main

import "fmt"

func main() {
    // Loop dari 1 hingga 100
    for i := 1; i <= 100; i++ {
        // Cek jika angka dapat dibagi dengan 3 dan 5
        if i % 3 == 0 && i % 5 == 0 {
            // Jika dapat dibagi dengan keduanya, cetak "FizzBuzz"
            fmt.Println("FizzBuzz")
        } else if i % 3 == 0 {
            // Jika dapat dibagi dengan 3, cetak "Fizz"
            fmt.Println("Fizz")
        } else if i % 5 == 0 {
            // Jika dapat dibagi dengan 5, cetak "Buzz"
            fmt.Println("Buzz")
        } else {
            // Jika tidak dapat dibagi dengan keduanya, cetak angka
            fmt.Println(i)
        }
    }
}
