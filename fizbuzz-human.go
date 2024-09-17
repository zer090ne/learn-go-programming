package main

import (
    "encoding/json" // Digunakan untuk decoding data JSON
    "fmt"          // Digunakan untuk output teks ke terminal
)

// Definisi struct Orang untuk mewakili data orang
type Orang struct {
    Nama        string `json:"nama"`           // Field Nama di-mapping ke properti "nama" di JSON
    Umur        int    `json:"umur"`           // Field Umur di-mapping ke properti "umur" di JSON
    JenisKelamin string `json:"jenis_kelamin"` // Field JenisKelamin di-mapping ke properti "jenis_kelamin" di JSON
}

func main() {
    // Data JSON dalam bentuk string
    data := `[
        {
            "nama": "John Doe",
            "umur": 25,
            "jenis_kelamin": "Laki-laki"
        },
        {
            "nama": "Jane Doe",
            "umur": 30,
            "jenis_kelamin": "Perempuan"
        },
        {
            "nama": "Bob Smith",
            "umur": 35,
            "jenis_kelamin": "Laki-laki"
        }
    ]`

    // Deklarasi slice orang untuk menampung data dari JSON
    var orang []Orang
    // Mengubah data JSON menjadi slice orang
    err := json.Unmarshal([]byte(data), &orang)
    if err != nil {
        // Menampilkan error jika terjadi kesalahan dalam decoding JSON
        fmt.Println(err)
        return
    }

    // Looping melalui setiap data orang
    for _, o := range orang {
        // Jika umur orang habis dibagi 3 dan 5, tampilkan FizzBuzz
        if o.Umur%3 == 0 && o.Umur%5 == 0 {
            fmt.Println(o.Nama, o.Umur, "FizzBuzz")
        // Jika umur orang habis dibagi 3, tampilkan Fizz
        } else if o.Umur%3 == 0 {
            fmt.Println(o.Nama, o.Umur, "Fizz")
        // Jika umur orang habis dibagi 5, tampilkan Buzz
        } else if o.Umur%5 == 0 {
            fmt.Println(o.Nama, o.Umur, "Buzz")
        }
    }
}
