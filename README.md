# Sprig

Aplikasi pengelolaan keuangan yang sempat tertunda pengerjaannya, kini dibangun ulang menggunakan bahasa pemrograman Golang.

## Daftar Isi

- [Features](#features)
- [Struktur Project](#struktur-project)
- [Tools yang Digunakan](#tools-yang-digunakan)
- [Instalasi](#instalasi)
  - [1. Clone Repository](#1-clone-repository-ini)
  - [2. Konfigurasi .env](#2-konfigurasi-env)
  - [3. Install Image & Jalankan Container Docker](#3-install-image--jalankan-container-docker-untuk-project)
  - [4. Migrasi Database](#4-migrasi-database)
  - [5. Install Package yang Dibutuhkan](#5-install-package-yang-dibutuhkan-golang)
  - [Go Unit Test](#tambahan--go-unit-test)
  - [Port & Route list](#port-list)

## Features

- **Income Management** — Mencatat dan mengelola pemasukan.
- **Expense Management** — Mencatat pengeluaran berdasarkan transaksi.
- **Category Management** — Mengelompokkan transaksi berdasarkan kategori.
- **Financial Summary** — Melihat ringkasan kondisi keuangan.
- **Transaction Search & Filtering** — Mencari dan memfilter transaksi berdasarkan kriteria tertentu.
- **Transaction History** — Melihat riwayat pemasukan dan pengeluaran.
- **Automatic Balance Calculation** — Menghitung saldo secara otomatis berdasarkan transaksi.
- **Budget Management** — Menentukan dan memantau batas pengeluaran.
- **Financial Statistics** — Menampilkan statistik pemasukan dan pengeluaran.
- **User Authentication** — Mengamankan akses data setiap pengguna.
- **User Management** — Mengelola informasi akun pengguna.

## Struktur Project

```
root-folder
├── cmd
│   └── web
│       ├── docs           // dokumentasi (swagger, dsb)
│       └── main.go        // entry point (tempat semua komponen dirakit)
├── docker
│   ├── golang
│   │   └── Dockerfile
│   └── mysql
│       └── my.cnf
├── docs
│   └── api
│       └── contract.md    // list contract API yang sudah disepakati
├── internal
│   ├── bootstrap           // inisialisasi/wiring aplikasi saat startup
│   ├── common               // helper/utility yang dipakai lintas fitur
│   ├── config                // konfigurasi dasar yang diambil dari .env
│   ├── features                // masing-masing fitur aplikasi (CRUD)
│   │   ├── budgets
│   │   │   ├── dto.go
│   │   │   ├── handler.go
│   │   │   ├── repository.go
│   │   │   ├── repository_mysql.go
│   │   │   ├── route.go
│   │   │   └── service.go
│   │   ├── categories
│   │   ├── expenses
│   │   └── users
│   ├── model                  // model database
│   └── router                   // definisi routing
├── migrations
├── go.mod
├── go.sum
└── ...
```

## Tools yang Digunakan

| Backend Tools |
|---------------|
| [`Golang`](https://go.dev/) |
| [`MySQL`](https://github.com/go-sql-driver/mysql) |
| [`Docker`](https://docs.docker.com/engine/install/) |
| [`golang-migrate/migrate`](https://github.com/golang-migrate/migrate) |
| [`swaggo/swag`](https://github.com/swaggo/swag) |
| [`stretchr/testify`](https://github.com/stretchr/testify) |

## Instalasi

Pastikan di komputer kamu sudah terinstall WSL beserta distro Linux pilihanmu. Jika belum, silakan install terlebih dahulu dengan mengikuti dokumentasi resmi Microsoft: [WSL Install](https://learn.microsoft.com/en-us/windows/wsl/install).

##### 1. Clone Repository ini
```bash
git clone https://github.com/Sprigate/Sprig-API.git
```

##### 2. Konfigurasi .env
```bash
# Rename file .env.example
mv .env.example .env
```

Konfigurasi sesuai dengan settingan production
```bash
APP_PORT=8080

DB_HOST=mysql
DB_PORT=3306
DB_USERNAME=root
DB_PASSWORD=sangat-rahasia(default: root)
DB_NAME=sprig_db
```

##### 3. Install Image & Jalankan Container Docker untuk Project
```bash
sudo docker compose up -d --build

# Menghentikan container yang sedang aktif
sudo docker compose stop

# Ingin sekaligus menghapus container?
sudo docker compose down
```

##### 4. Migrasi Database
```bash
# Install dependensi yang dibutuhkan (jika belum terinstall)
sudo apt install make

# Jalankan migrasi database
sudo make migrate-up
```
> **NOTE:** Untuk detail penggunaan command migrasi database lainnya, silakan baca daftar lengkapnya di [Migration Command List](/migrations/README.md).

##### 5. Install Package yang Dibutuhkan Golang
```bash
go mod download

# Ingin package-nya disimpan secara lokal di dalam project?
go mod vendor
```

##### Tambahan : Go Unit Test & swagger
```bash
sudo make test path=./file/destination

# Contoh
sudo make test path=./internal/config

# Swagger 
make swag-init
```

##### Port list

| Port & route | Tujuan |
|--------------|--------|
| `localhost:8080` | app |
| `localhost:8080/swagger/` | Dokumentasi swagger |
| `localhost:8081` | adminer (database manager) |
