package main

import (
	"encoding/json"
	"fmt"
)

type Hewan struct {
	Nama		string `json:"nama"`
	Kaki		int 	`json:"kaki"`
	JenisKelamin 	string `json:"jenis_kelamin"`
}

func main() {
	data := `[
		{
			"nama": "Doggie",
			"kaki": 40,
			"jenis_kelamin": "Jantan"
		},
		{
			"nama": "Kucing",
			"kaki": 40,
			"jenis_kelamin": "Betina"
		},
		{
			"nama": "Ayam",
			"kaki":  25,
			"jenis_kelamin": "Jantan"
		}			
	]`

	var hewan []Hewan
	err := json.Unmarshal([]byte(data), &hewan)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	for  _, h := range hewan {
		if h.Kaki%3 == 0 &&  h.Kaki%5  == 0 {
			fmt.Println(h.Nama, h.Kaki, "FizzBuzz")
			} else if h.Kaki%3 == 0 {
				fmt.Println(h.Nama, h.Kaki, "Fizz")
				} else if h.Kaki%5 == 0 {
					fmt.Println(h.Nama, h.Kaki, "Buzz")
				}
				}

}