# Kontrak API

Base URL: `/api/v1`

## Catatan Umum

- Semua response mengikuti satu format envelope:
  ```json
  {
    "message": "string",
    "data": {}
  }
  ```
  `data` berupa **objek** untuk response satu resource, dan berupa **array** untuk response list/koleksi.
- Path parameter (`{id}`, `{name}`) digunakan untuk menunjuk resource yang spesifik.
- Query parameter hanya digunakan untuk filtering opsional pada endpoint koleksi (`GET`).

---

## 1. Autentikasi

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/auth/register` | Mendaftarkan akun baru |
| POST | `/auth/login` | Login dan mendapatkan access token |

### Input Parameter — Register

| Field | Tipe | Deskripsi |
|-------|------|-----------|
| name | string | Nama lengkap user |
| email | string | Alamat email, digunakan sebagai identitas login |
| password | string | Kata sandi akun |

### Input Parameter — Login

| Field | Tipe | Deskripsi |
|-------|------|-----------|
| email | string | Alamat email yang sudah terdaftar |
| password | string | Kata sandi akun |

### Output — Register

```json
{
  "message": "Registrasi berhasil",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "created_at": "2026-09-10T10:00:00Z"
  }
}
```

### Output — Login

```json
{
  "message": "Login berhasil",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "token_type": "Bearer",
    "expires_in": 3600,
    "user": {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com"
    }
  }
}
```

> Catatan: menggunakan standar umum berbasis JWT (access token + masa berlaku). Sesuaikan jika Anda menambahkan refresh token atau strategi auth yang berbeda.

---

## 2. User

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/users` | Menambah user baru |
| PATCH | `/users/{id}` | Mengupdate data user berdasarkan ID |
| DELETE | `/users/{id}` | Menghapus user berdasarkan ID |

### Input Parameter (Create / Update)

| Field | Tipe | Deskripsi |
|-------|------|-----------|
| name | string | Nama lengkap user |
| email | string | Alamat email user |
| password | string | Kata sandi akun user |

### Output

```json
{
  "message": "User berhasil dibuat",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "created_at": "2026-09-10T10:00:00Z"
  }
}
```

---

## 3. Categories (Kategori)

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/categories` | Membuat kategori baru |
| GET | `/categories` | Mengambil semua data kategori |
| GET | `/categories/{name}` | Mengambil kategori berdasarkan nama |
| PATCH | `/categories/{id}` | Mengupdate kategori berdasarkan ID |
| DELETE | `/categories/{id}` | Menghapus kategori berdasarkan ID |

### Input Parameter (Create / Update)

| Field | Tipe | Deskripsi |
|-------|------|-----------|
| name | string | Nama kategori |

### Output — satu kategori

```json
{
  "message": "Kategori berhasil diambil",
  "data": {
    "id": 1,
    "user_id": 10,
    "name": "Food"
  }
}
```

### Output — daftar kategori

```json
{
  "message": "Daftar kategori berhasil diambil",
  "data": [
    { "id": 1, "user_id": 10, "name": "Food" },
    { "id": 2, "user_id": 10, "name": "Transport" }
  ]
}
```

---

## 4. Expenses (Pengeluaran)

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/expenses` | Membuat data pengeluaran baru |
| GET | `/expenses` | Mengambil semua data pengeluaran (mendukung query filter opsional di bawah) |
| PATCH | `/expenses/{id}` | Mengupdate pengeluaran berdasarkan ID |
| DELETE | `/expenses/{id}` | Menghapus pengeluaran berdasarkan ID |

### Query Parameter (opsional, `GET /expenses`)

| Parameter | Tipe | Deskripsi |
|-----------|------|-----------|
| start_date | string (`YYYY-MM-DD`) | Awal rentang tanggal filter |
| end_date | string (`YYYY-MM-DD`) | Akhir rentang tanggal filter |
| type | string | Filter berdasarkan tipe transaksi, mis. `expense` / `income` |
| category_id | integer | Filter berdasarkan ID kategori |

### Input Parameter (Create / Update)

| Field | Tipe | Deskripsi |
|-------|------|-----------|
| category_id | uint64 | ID kategori terkait |
| amount | uint64 | Jumlah nominal transaksi |
| description | *string | Deskripsi tambahan (boleh kosong/null) |
| expense_date | time.Time | Tanggal transaksi terjadi |
| type | model.ExpenseType | Tipe transaksi (mis. expense/income) |

### Output — satu data pengeluaran

```json
{
  "message": "Pengeluaran berhasil dibuat",
  "data": {
    "id": 1,
    "user_id": 10,
    "category_id": 2,
    "amount": 50000,
    "description": "Makan siang bersama klien",
    "expense_date": "2026-09-10T00:00:00Z",
    "type": "expense",
    "updated_at": "2026-09-10T10:00:00Z"
  }
}
```

### Output — daftar data pengeluaran

```json
{
  "message": "Daftar pengeluaran berhasil diambil",
  "data": [
    {
      "id": 1,
      "user_id": 10,
      "category_id": 2,
      "amount": 50000,
      "description": "Makan siang bersama klien",
      "expense_date": "2026-09-10T00:00:00Z",
      "type": "expense",
      "updated_at": "2026-09-10T10:00:00Z"
    }
  ]
}
```

---

## 5. Budgets (Anggaran)

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/budgets` | Membuat data anggaran baru |
| GET | `/budgets` | Mengambil semua data anggaran (mendukung query filter opsional di bawah) |
| PATCH | `/budgets/{id}` | Mengupdate anggaran berdasarkan ID |
| DELETE | `/budgets/{id}` | Menghapus anggaran berdasarkan ID |

### Query Parameter (opsional, `GET /budgets`)

| Parameter | Tipe | Deskripsi |
|-----------|------|-----------|
| month | integer | Filter berdasarkan bulan |
| year | integer | Filter berdasarkan tahun |

### Input Parameter (Create / Update)

| Field | Tipe | Deskripsi |
|-------|------|-----------|
| total_income | uint64 | Total pemasukan untuk periode anggaran |
| needs_percentage | uint8 | Persentase alokasi kebutuhan (needs) |
| wants_percentage | uint8 | Persentase alokasi keinginan (wants) |
| savings_percentage | uint8 | Persentase alokasi tabungan (savings) |
| year | uint16 | Tahun anggaran |
| month | uint8 | Bulan anggaran |

### Output — satu data anggaran

```json
{
  "message": "Anggaran berhasil dibuat",
  "data": {
    "id": 1,
    "user_id": 10,
    "total_income": 5000000,
    "needs_percentage": 50,
    "wants_percentage": 30,
    "savings_percentage": 20,
    "year": 2026,
    "month": 9,
    "updated_at": "2026-09-10T10:00:00Z"
  }
}
```

### Output — daftar data anggaran

```json
{
  "message": "Daftar anggaran berhasil diambil",
  "data": [
    {
      "id": 1,
      "user_id": 10,
      "total_income": 5000000,
      "needs_percentage": 50,
      "wants_percentage": 30,
      "savings_percentage": 20,
      "year": 2026,
      "month": 9,
      "updated_at": "2026-09-10T10:00:00Z"
    }
  ]
}
```

---

## 6. Dashboard

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/dashboard` | Mengambil keseluruhan data dashboard |

### Input Parameter

_Masih dalam tahap diskusi._

### Output

_Masih dalam tahap diskusi._