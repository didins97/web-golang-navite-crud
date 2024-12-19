# Web Golang Native CRUD

Proyek ini adalah aplikasi web sederhana yang mengimplementasikan operasi CRUD (Create, Read, Update, Delete) menggunakan bahasa pemrograman Go (Golang) tanpa menggunakan framework tambahan.

## Fitur
- Create: Menambahkan data baru ke dalam sistem.
- Read: Menampilkan data yang tersimpan.
- Update: Memperbarui data yang ada.
- Delete: Menghapus data dari sistem.

## Prasyarat
Sebelum menjalankan aplikasi ini, pastikan Anda telah menginstal:
- Go versi 1.16 atau lebih baru.
- MySQL sebagai basis data.

## Instalasi
- Clone repositori ini:
```bash
git clone https://github.com/didins97/web-golang-navite-crud.git
cd web-golang-navite-crud
```
## Inisialisasi modul Go:

```bash
go mod tidy
```

## Konfigurasi basis data:
Buat basis data MySQL baru.
Perbarui informasi koneksi basis data di file config/config.go sesuai dengan kredensial Anda.
Jalankan migrasi untuk membuat tabel yang diperlukan:
``` bash
go run migrate/migrate.go
```

## Menjalankan Aplikasi
Setelah menyelesaikan langkah-langkah di atas, Anda dapat menjalankan aplikasi dengan perintah:
```bash
go run main.go
```
Aplikasi akan berjalan pada http://localhost:8080.

### Struktur Proyek
Berikut adalah struktur direktori utama dalam proyek ini:

```graphql
web-golang-navite-crud/
├── config/             # Konfigurasi aplikasi dan basis data
├── controllers/        # Logika pengendali untuk setiap entitas
├── models/             # Definisi model dan interaksi dengan basis data
├── views/              # Template HTML untuk rendering halaman
├── migrate/            # Skrip migrasi basis data
├── main.go             # Entry point aplikasi
└── README.md           # Dokumentasi proyek
```
