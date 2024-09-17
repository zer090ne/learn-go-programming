Berikut adalah penjelasan dari setiap bagian kode dan komentar untuk code Go tersebut:

```go
package main
```
- **package main**: Menyatakan bahwa ini adalah program utama (executable) di Go. Package `main` diperlukan untuk membuat aplikasi yang dapat dieksekusi.

```go
import (
    "encoding/json"
    "fmt"
)
```
- **import**: Digunakan untuk mengimpor package yang dibutuhkan dalam kode. 
  - `"encoding/json"`: Package ini digunakan untuk encoding dan decoding data JSON.
  - `"fmt"`: Package ini digunakan untuk format dan output teks, seperti menampilkan teks di terminal.

```go
type Orang struct {
    Nama        string `json:"nama"`
    Umur        int    `json:"umur"`
    JenisKelamin string `json:"jenis_kelamin"`
}
```
- **type Orang struct**: Mendefinisikan tipe data `Orang` sebagai struct (struktur data) yang memiliki tiga field: `Nama` (string), `Umur` (integer), dan `JenisKelamin` (string).
  - **`json:"nama"`**: Tag JSON ini menginstruksikan Go untuk memetakan field struct `Nama` ke properti JSON `"nama"`.
  - **`json:"umur"`**: Tag JSON untuk memetakan field `Umur` ke properti `"umur"`.
  - **`json:"jenis_kelamin"`**: Tag JSON untuk memetakan field `JenisKelamin` ke properti `"jenis_kelamin"`.

```go
func main() {
```
- **func main()**: Fungsi utama dari program Go. Fungsi ini akan dipanggil ketika program dieksekusi.

```go
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
```
- **data**: Sebuah string berisi array JSON yang merepresentasikan data beberapa orang dengan properti `nama`, `umur`, dan `jenis_kelamin`.

```go
    var orang []Orang
```
- **var orang []Orang**: Mendeklarasikan variabel `orang` sebagai slice (array dinamis) dari struct `Orang`.

```go
    err := json.Unmarshal([]byte(data), &orang)
```
- **json.Unmarshal**: Fungsi ini mengubah (decode) data JSON yang berbentuk string menjadi tipe data Go. Di sini, JSON dari variabel `data` di-decode ke slice `orang`.
  - **[]byte(data)**: Mengkonversi string `data` ke bentuk byte array, karena fungsi `Unmarshal` memerlukan input dalam bentuk byte.
  - **&orang**: Mengirimkan referensi dari slice `orang` sehingga data yang di-decode akan dimasukkan ke dalam slice ini.

```go
    if err != nil {
        fmt.Println(err)
        return
    }
```
- **if err != nil**: Mengecek apakah terjadi error selama proses decoding JSON. Jika ada error, program mencetak error dan menghentikan eksekusi dengan `return`.

```go
    for _, o := range orang {
```
- **for _, o := range orang**: Looping melalui setiap elemen di slice `orang`. Setiap iterasi, variabel `o` merepresentasikan satu elemen `Orang`.

```go
        if o.Umur%3 == 0 && o.Umur%5 == 0 {
            fmt.Println(o.Nama, "FizzBuzz")
        } else if o.Umur%3 == 0 {
            fmt.Println(o.Nama, "Fizz")
        } else if o.Umur%5 == 0 {
            fmt.Println(o.Nama, "Buzz")
        }
```
- **if o.Umur%3 == 0 && o.Umur%5 == 0**: Mengecek apakah umur orang habis dibagi 3 dan 5. Jika iya, tampilkan "FizzBuzz".
- **else if o.Umur%3 == 0**: Jika hanya habis dibagi 3, tampilkan "Fizz".
- **else if o.Umur%5 == 0**: Jika hanya habis dibagi 5, tampilkan "Buzz".
- **fmt.Println(o.Nama, "FizzBuzz")**: Menampilkan nama orang beserta teks "Fizz", "Buzz", atau "FizzBuzz" sesuai dengan kondisinya.

### Komentar Code:

```go
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
            fmt.Println(o.Nama, "FizzBuzz")
        // Jika umur orang habis dibagi 3, tampilkan Fizz
        } else if o.Umur%3 == 0 {
            fmt.Println(o.Nama, "Fizz")
        // Jika umur orang habis dibagi 5, tampilkan Buzz
        } else if o.Umur%5 == 0 {
            fmt.Println(o.Nama, "Buzz")
        }
    }
}
```

Penjelasan ini membantu memahami setiap bagian dari kode Go di atas, termasuk bagaimana data JSON di-decode dan diproses menggunakan struct serta bagaimana kondisional "FizzBuzz" diterapkan.


###  Code Lengkap

```

package main

import (
    "encoding/json"
    "fmt"
)

type Orang struct {
    Nama        string `json:"nama"`
    Umur        int    `json:"umur"`
    JenisKelamin string `json:"jenis_kelamin"`
}

func main() {
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

    var orang []Orang
    err := json.Unmarshal([]byte(data), &orang)
    if err != nil {
        fmt.Println(err)
        return
    }

    for _, o := range orang {
        if o.Umur%3 == 0 && o.Umur%5 == 0 {
            fmt.Println(o.Nama, "FizzBuzz")
        } else if o.Umur%3 == 0 {
            fmt.Println(o.Nama, "Fizz")
        } else if o.Umur%5 == 0 {
            fmt.Println(o.Nama, "Buzz")
        }
    }
}

```
