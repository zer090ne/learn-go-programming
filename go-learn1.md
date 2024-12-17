Soal Tantangan Sederhana:
Buat program yang:
1. Menerima data beberapa mahasiswa (nama, umur, dan nilai).
2. Data ini disimpan dalam sebuah slice of struct.
3. Program harus memiliki function untuk:
   * Menampilkan semua data mahasiswa.
   * Menghitung rata-rata nilai mahasiswa.
   * Mencari mahasiswa dengan nilai tertinggi.
4. Gunakan control flow (for, if, atau switch) sesuai kebutuhan.
5. Coba gunakan pointer untuk memperbarui nilai tertentu jika diperlukan.

Saya akan membantu membuat program Go sesuai dengan spesifikasi yang diminta.



```go
package main

import "fmt"

// Struct untuk menyimpan data mahasiswa
type Mahasiswa struct {
    Nama  string
    Umur  int
    Nilai float64
}

// Function untuk menampilkan semua data mahasiswa
func tampilkanData(mahasiswas []Mahasiswa) {
    fmt.Println("\nData Semua Mahasiswa:")
    for i, m := range mahasiswas {
        fmt.Printf("%d. Nama: %s, Umur: %d, Nilai: %.2f\n", i+1, m.Nama, m.Umur, m.Nilai)
    }
}

// Function untuk menghitung rata-rata nilai
func hitungRataRata(mahasiswas []Mahasiswa) float64 {
    total := 0.0
    for _, m := range mahasiswas {
        total += m.Nilai
    }
    return total / float64(len(mahasiswas))
}

// Function untuk mencari mahasiswa dengan nilai tertinggi
func cariNilaiTertinggi(mahasiswas []Mahasiswa) Mahasiswa {
    tertinggi := mahasiswas[0]
    for _, m := range mahasiswas {
        if m.Nilai > tertinggi.Nilai {
            tertinggi = m
        }
    }
    return tertinggi
}

// Function untuk mengupdate nilai mahasiswa menggunakan pointer
func updateNilai(m *Mahasiswa, nilaiBaru float64) {
    m.Nilai = nilaiBaru
}

func main() {
    // Membuat slice of struct untuk menyimpan data mahasiswa
    mahasiswas := []Mahasiswa{
        {"Budi", 20, 85.5},
        {"Ana", 19, 90.0},
        {"Citra", 21, 87.5},
        {"Deni", 20, 82.5},
    }

    // Menampilkan data awal
    tampilkanData(mahasiswas)

    // Menghitung dan menampilkan rata-rata
    rataRata := hitungRataRata(mahasiswas)
    fmt.Printf("\nRata-rata nilai: %.2f\n", rataRata)

    // Mencari dan menampilkan mahasiswa dengan nilai tertinggi
    tertinggi := cariNilaiTertinggi(mahasiswas)
    fmt.Printf("\nMahasiswa dengan nilai tertinggi:\nNama: %s\nNilai: %.2f\n", 
        tertinggi.Nama, tertinggi.Nilai)

    // Contoh penggunaan pointer untuk mengupdate nilai
    fmt.Printf("\nMengupdate nilai %s\n", mahasiswas[0].Nama)
    updateNilai(&mahasiswas[0], 88.0)
    tampilkanData(mahasiswas)
}
```

Program ini memenuhi semua kriteria yang diminta:

1. Menggunakan struct `Mahasiswa` untuk menyimpan data nama, umur, dan nilai
2. Menyimpan data dalam slice of struct (`[]Mahasiswa`)
3. Memiliki fungsi-fungsi yang diminta:
   - `tampilkanData()` untuk menampilkan semua data
   - `hitungRataRata()` untuk menghitung rata-rata nilai
   - `cariNilaiTertinggi()` untuk mencari mahasiswa dengan nilai tertinggi
4. Menggunakan control flow:
   - `for` untuk iterasi
   - `if` untuk perbandingan nilai
5. Menggunakan pointer dalam fungsi `updateNilai()` untuk memperbarui nilai mahasiswa

Ketika dijalankan, program akan:
1. Menampilkan data semua mahasiswa
2. Menghitung dan menampilkan rata-rata nilai
3. Menampilkan data mahasiswa dengan nilai tertinggi
4. Mendemonstrasikan penggunaan pointer dengan mengupdate nilai salah satu mahasiswa
5. Menampilkan kembali data setelah update

Anda bisa mencoba program ini dengan menyalin kode tersebut ke editor Go dan menjalankannya.
