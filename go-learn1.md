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

Saya akan meningkatkan program dengan menambahkan penanganan slice kosong, membuat fungsi update yang lebih umum, dan error handling yang lebih baik.

```go
package main

import (
    "errors"
    "fmt"
)

type Mahasiswa struct {
    Nama  string
    Umur  int
    Nilai float64
}

// Menambahkan error handling untuk slice kosong
func tampilkanData(mahasiswas []Mahasiswa) error {
    if len(mahasiswas) == 0 {
        return errors.New("data mahasiswa kosong")
    }

    fmt.Println("\nData Semua Mahasiswa:")
    for i, m := range mahasiswas {
        fmt.Printf("%d. Nama: %s, Umur: %d, Nilai: %.2f\n", i+1, m.Nama, m.Umur, m.Nilai)
    }
    return nil
}

func hitungRataRata(mahasiswas []Mahasiswa) (float64, error) {
    if len(mahasiswas) == 0 {
        return 0, errors.New("tidak dapat menghitung rata-rata: data mahasiswa kosong")
    }

    total := 0.0
    for _, m := range mahasiswas {
        total += m.Nilai
    }
    return total / float64(len(mahasiswas)), nil
}

func cariNilaiTertinggi(mahasiswas []Mahasiswa) (Mahasiswa, error) {
    if len(mahasiswas) == 0 {
        return Mahasiswa{}, errors.New("tidak dapat mencari nilai tertinggi: data mahasiswa kosong")
    }

    tertinggi := mahasiswas[0]
    for _, m := range mahasiswas {
        if m.Nilai > tertinggi.Nilai {
            tertinggi = m
        }
    }
    return tertinggi, nil
}

// Generalized update function yang bisa mengupdate semua field
func updateMahasiswa(m *Mahasiswa, field string, value interface{}) error {
    if m == nil {
        return errors.New("pointer mahasiswa tidak valid")
    }

    switch field {
    case "nama":
        if val, ok := value.(string); ok {
            m.Nama = val
            return nil
        }
        return errors.New("tipe data tidak sesuai untuk nama")

    case "umur":
        if val, ok := value.(int); ok {
            m.Umur = val
            return nil
        }
        return errors.New("tipe data tidak sesuai untuk umur")

    case "nilai":
        switch v := value.(type) {
        case float64:
            m.Nilai = v
            return nil
        case int:
            m.Nilai = float64(v)
            return nil
        default:
            return errors.New("tipe data tidak sesuai untuk nilai")
        }

    default:
        return fmt.Errorf("field tidak valid: %s", field)
    }
}

func main() {
    // Test dengan slice kosong
    var mahasiswaKosong []Mahasiswa
    if err := tampilkanData(mahasiswaKosong); err != nil {
        fmt.Printf("Error: %v\n", err)
    }

    // Data normal
    mahasiswas := []Mahasiswa{
        {"Budi", 20, 85.5},
        {"Ana", 19, 90.0},
        {"Citra", 21, 87.5},
        {"Deni", 20, 82.5},
    }

    // Menampilkan data dengan error handling
    if err := tampilkanData(mahasiswas); err != nil {
        fmt.Printf("Error: %v\n", err)
    }

    // Menghitung rata-rata dengan error handling
    if rataRata, err := hitungRataRata(mahasiswas); err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("\nRata-rata nilai: %.2f\n", rataRata)
    }

    // Mencari nilai tertinggi dengan error handling
    if tertinggi, err := cariNilaiTertinggi(mahasiswas); err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("\nMahasiswa dengan nilai tertinggi:\nNama: %s\nNilai: %.2f\n",
            tertinggi.Nama, tertinggi.Nilai)
    }

    // Demonstrasi penggunaan generalized update function
    fmt.Println("\nMencoba update berbagai field:")
    
    // Update nilai
    if err := updateMahasiswa(&mahasiswas[0], "nilai", 88.5); err != nil {
        fmt.Printf("Error update nilai: %v\n", err)
    }

    // Update nama
    if err := updateMahasiswa(&mahasiswas[0], "nama", "Budi Revisi"); err != nil {
        fmt.Printf("Error update nama: %v\n", err)
    }

    // Update umur
    if err := updateMahasiswa(&mahasiswas[0], "umur", 21); err != nil {
        fmt.Printf("Error update umur: %v\n", err)
    }

    // Contoh error handling untuk field tidak valid
    if err := updateMahasiswa(&mahasiswas[0], "invalid", "value"); err != nil {
        fmt.Printf("Error expected: %v\n", err)
    }

    // Menampilkan data setelah update
    if err := tampilkanData(mahasiswas); err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}
```

Peningkatan yang telah ditambahkan:

1. **Penanganan Slice Kosong**:
   - Semua fungsi sekarang memeriksa apakah slice kosong
   - Mengembalikan error yang sesuai jika slice kosong
   - Main function mendemonstrasikan penanganan slice kosong

2. **Generalized Update Function**:
   - Fungsi `updateMahasiswa` yang bisa mengupdate semua field
   - Menggunakan interface{} untuk menerima berbagai tipe data
   - Type assertion untuk memastikan tipe data yang benar
   - Switch statement untuk menangani berbagai field
   - Error handling untuk tipe data yang tidak sesuai

3. **Error Handling yang Lebih Baik**:
   - Semua fungsi sekarang mengembalikan error
   - Penggunaan package errors untuk membuat custom error
   - Pengecekan error di main function
   - Pesan error yang lebih deskriptif
   - Penanganan error untuk berbagai kasus (slice kosong, tipe data tidak sesuai, field tidak valid)

Program ini sekarang lebih robust dan aman digunakan karena:
- Menangani kasus slice kosong
- Memvalidasi tipe data input
- Memberikan pesan error yang jelas
- Memungkinkan update yang lebih fleksibel untuk semua field
- Menggunakan proper error handling di semua operasi

  Tantangan Tambahan
Jika kamu ingin meningkatkan program ini lebih lanjut, coba lakukan:
1. Input Dinamis: Izinkan pengguna memasukkan data mahasiswa melalui terminal (gunakan fmt.Scan).
2. Fitur Sorting: Urutkan mahasiswa berdasarkan nilai tertinggi ke terendah menggunakan sort.
3. Menu Interaktif: Tambahkan menu untuk memilih operasi: tampilkan data, tambah mahasiswa, update nilai, dll.

Saya akan meningkatkan program dengan menambahkan fitur-fitur tersebut.

```go
package main

import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

type Mahasiswa struct {
    Nama  string
    Umur  int
    Nilai float64
}

// Untuk implementasi sort.Interface
type DaftarMahasiswa []Mahasiswa

func (d DaftarMahasiswa) Len() int           { return len(d) }
func (d DaftarMahasiswa) Less(i, j int) bool { return d[i].Nilai > d[j].Nilai } // Descending
func (d DaftarMahasiswa) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

// Input mahasiswa baru dengan validasi
func inputMahasiswaBaru(scanner *bufio.Scanner) (Mahasiswa, error) {
    var m Mahasiswa
    
    fmt.Print("Masukkan nama: ")
    scanner.Scan()
    m.Nama = strings.TrimSpace(scanner.Text())
    if m.Nama == "" {
        return m, errors.New("nama tidak boleh kosong")
    }

    fmt.Print("Masukkan umur: ")
    scanner.Scan()
    umur, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
    if err != nil {
        return m, errors.New("umur harus berupa angka")
    }
    if umur < 15 || umur > 100 {
        return m, errors.New("umur tidak valid (15-100)")
    }
    m.Umur = umur

    fmt.Print("Masukkan nilai: ")
    scanner.Scan()
    nilai, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
    if err != nil {
        return m, errors.New("nilai harus berupa angka")
    }
    if nilai < 0 || nilai > 100 {
        return m, errors.New("nilai tidak valid (0-100)")
    }
    m.Nilai = nilai

    return m, nil
}

func tampilkanMenu() {
    fmt.Println("\n=== Menu Manajemen Data Mahasiswa ===")
    fmt.Println("1. Tampilkan semua data")
    fmt.Println("2. Tambah mahasiswa baru")
    fmt.Println("3. Update data mahasiswa")
    fmt.Println("4. Urutkan berdasarkan nilai")
    fmt.Println("5. Hitung rata-rata nilai")
    fmt.Println("6. Tampilkan nilai tertinggi")
    fmt.Println("0. Keluar")
    fmt.Print("Pilih menu (0-6): ")
}

func tampilkanData(mahasiswas []Mahasiswa) error {
    if len(mahasiswas) == 0 {
        return errors.New("data mahasiswa kosong")
    }

    fmt.Println("\nData Semua Mahasiswa:")
    for i, m := range mahasiswas {
        fmt.Printf("%d. Nama: %s, Umur: %d, Nilai: %.2f\n", i+1, m.Nama, m.Umur, m.Nilai)
    }
    return nil
}

func updateMahasiswa(m *Mahasiswa, field string, value interface{}) error {
    if m == nil {
        return errors.New("pointer mahasiswa tidak valid")
    }

    switch field {
    case "nama":
        if val, ok := value.(string); ok {
            m.Nama = val
            return nil
        }
        return errors.New("tipe data tidak sesuai untuk nama")
    case "umur":
        if val, ok := value.(int); ok {
            if val < 15 || val > 100 {
                return errors.New("umur tidak valid (15-100)")
            }
            m.Umur = val
            return nil
        }
        return errors.New("tipe data tidak sesuai untuk umur")
    case "nilai":
        switch v := value.(type) {
        case float64:
            if v < 0 || v > 100 {
                return errors.New("nilai tidak valid (0-100)")
            }
            m.Nilai = v
            return nil
        case int:
            if v < 0 || v > 100 {
                return errors.New("nilai tidak valid (0-100)")
            }
            m.Nilai = float64(v)
            return nil
        default:
            return errors.New("tipe data tidak sesuai untuk nilai")
        }
    default:
        return fmt.Errorf("field tidak valid: %s", field)
    }
}

func main() {
    var mahasiswas []Mahasiswa
    scanner := bufio.NewScanner(os.Stdin)
    
    for {
        tampilkanMenu()
        scanner.Scan()
        pilihan := strings.TrimSpace(scanner.Text())

        switch pilihan {
        case "1": // Tampilkan data
            if err := tampilkanData(mahasiswas); err != nil {
                fmt.Printf("Error: %v\n", err)
            }

        case "2": // Tambah mahasiswa
            fmt.Println("\n=== Tambah Mahasiswa Baru ===")
            mhs, err := inputMahasiswaBaru(scanner)
            if err != nil {
                fmt.Printf("Error: %v\n", err)
                continue
            }
            mahasiswas = append(mahasiswas, mhs)
            fmt.Println("Mahasiswa berhasil ditambahkan!")

        case "3": // Update data
            if len(mahasiswas) == 0 {
                fmt.Println("Error: Belum ada data mahasiswa")
                continue
            }

            fmt.Print("\nMasukkan nomor mahasiswa (1-", len(mahasiswas), "): ")
            scanner.Scan()
            idx, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
            if err != nil || idx < 1 || idx > len(mahasiswas) {
                fmt.Println("Error: Nomor mahasiswa tidak valid")
                continue
            }

            fmt.Print("Pilih field yang akan diupdate (nama/umur/nilai): ")
            scanner.Scan()
            field := strings.ToLower(strings.TrimSpace(scanner.Text()))

            fmt.Print("Masukkan nilai baru: ")
            scanner.Scan()
            newValue := strings.TrimSpace(scanner.Text())

            var updateErr error
            switch field {
            case "nama":
                updateErr = updateMahasiswa(&mahasiswas[idx-1], field, newValue)
            case "umur":
                if val, err := strconv.Atoi(newValue); err == nil {
                    updateErr = updateMahasiswa(&mahasiswas[idx-1], field, val)
                } else {
                    updateErr = errors.New("nilai umur harus berupa angka")
                }
            case "nilai":
                if val, err := strconv.ParseFloat(newValue, 64); err == nil {
                    updateErr = updateMahasiswa(&mahasiswas[idx-1], field, val)
                } else {
                    updateErr = errors.New("nilai harus berupa angka")
                }
            default:
                updateErr = errors.New("field tidak valid")
            }

            if updateErr != nil {
                fmt.Printf("Error: %v\n", updateErr)
            } else {
                fmt.Println("Data berhasil diupdate!")
            }

        case "4": // Sort berdasarkan nilai
            if len(mahasiswas) == 0 {
                fmt.Println("Error: Belum ada data mahasiswa")
                continue
            }
            sort.Sort(DaftarMahasiswa(mahasiswas))
            fmt.Println("\nData telah diurutkan berdasarkan nilai (tertinggi ke terendah)")
            tampilkanData(mahasiswas)

        case "5": // Hitung rata-rata
            if len(mahasiswas) == 0 {
                fmt.Println("Error: Belum ada data mahasiswa")
                continue
            }
            total := 0.0
            for _, m := range mahasiswas {
                total += m.Nilai
            }
            fmt.Printf("\nRata-rata nilai: %.2f\n", total/float64(len(mahasiswas)))

        case "6": // Nilai tertinggi
            if len(mahasiswas) == 0 {
                fmt.Println("Error: Belum ada data mahasiswa")
                continue
            }
            tertinggi := mahasiswas[0]
            for _, m := range mahasiswas {
                if m.Nilai > tertinggi.Nilai {
                    tertinggi = m
                }
            }
            fmt.Printf("\nMahasiswa dengan nilai tertinggi:\nNama: %s\nNilai: %.2f\n",
                tertinggi.Nama, tertinggi.Nilai)

        case "0": // Keluar
            fmt.Println("\nTerima kasih telah menggunakan program ini!")
            return

        default:
            fmt.Println("Error: Menu tidak valid")
        }
    }
}
```

Peningkatan yang telah ditambahkan:

1. **Input Dinamis**:
   - Menggunakan `bufio.Scanner` untuk input yang lebih baik
   - Validasi input untuk memastikan data valid
   - Penanganan error untuk input yang tidak sesuai

2. **Fitur Sorting**:
   - Implementasi interface `sort.Interface`
   - Pengurutan berdasarkan nilai (descending)
   - Pembuatan type `DaftarMahasiswa` untuk mendukung sorting

3. **Menu Interaktif**:
   - Menu lengkap dengan berbagai operasi
   - Loop utama untuk interaksi berkelanjutan
   - Penanganan error untuk setiap operasi
   - Validasi input untuk setiap operasi

4. **Fitur Tambahan**:
   - Validasi range nilai (0-100)
   - Validasi range umur (15-100)
   - Pesan error yang lebih informatif
   - Konfirmasi operasi berhasil

5. **Peningkatan UX**:
   - Clear menu structure
   - Feedback yang jelas untuk setiap operasi
   - Konsistensi dalam format output
   - Validasi input yang lebih ketat

Program ini sekarang jauh lebih interaktif dan user-friendly, dengan penanganan error yang lebih baik dan fitur-fitur tambahan yang berguna.
5. Menampilkan kembali data setelah update

Anda bisa mencoba program ini dengan menyalin kode tersebut ke editor Go dan menjalankannya.
