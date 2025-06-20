# Marketplace Backend API

Backend API untuk aplikasi marketplace, dibangun dengan Go, Gin, GORM, dan PostgreSQL.

## Fitur

- **Autentikasi JWT** (login, register, role-based)
- **Manajemen Produk Merchant** (CRUD produk, upload gambar)
- **Sistem Order Customer** (buat order, lihat order, diskon & free shipping otomatis)
- **Manajemen Kategori Produk** (CRUD kategori)
- **Manajemen User (Admin)**
- **Pencarian & filter produk**
- **Soft delete** (`deleted_at`) di semua tabel utama

## Fitur Tambahan

Selain requirement minimal dari soal test, project ini juga sudah mendukung fitur-fitur berikut:

1. **Update & Delete Produk (Merchant)** — CRUD penuh untuk produk.
2. **Manajemen Kategori Produk** — Admin/merchant dapat CRUD kategori.
3. **Upload Gambar Produk (belum di test)** — Merchant dapat upload gambar produk.
4. **Manajemen User (Admin)** — Admin dapat melihat daftar dan detail user.
5. **Pencarian & Filter Produk** — Customer dapat mencari/filter produk.
6. **Soft Delete** — Semua tabel utama mendukung soft delete (`deleted_at`).
7. **Pagination & Meta Data** — Endpoint list mendukung pagination & meta.
8. **Validasi Lengkap** — Validasi input & JWT di semua endpoint.
9. **Response API Konsisten** — Format JSON konsisten (`success`, `message`, `data`, `meta`).
10. **Hash Password** — Password di-hash dengan bcrypt.
11. **Seed Data Otomatis** — Migrasi awal menambah user admin/customer.
12. **Testing dengan Postman** — Disediakan koleksi Postman.
13. **Konfigurasi Environment & Dokumentasi Lengkap** — Mendukung `.env` & README lengkap.
14. **Support Docker (opsional)** — Dockerfile & docker-compose tersedia.
15. **Manajemen Order Detail** — Merchant dapat melihat customer yang membeli produknya.

## Kekurangan / Catatan

- Beberapa response API masih inkonsisten antara bahasa Indonesia dan English.
- Fitur upload gambar produk **belum sepenuhnya dites** di berbagai environment.
- Belum ada fitur approval admin untuk produk/kategori baru.
- Belum ada fitur notifikasi/email.
- Belum ada unit test/coverage yang menyeluruh.
- Belum ada rate limiting atau proteksi brute force pada endpoint login.
- Dokumentasi Postman bisa lebih detail (misal: contoh response error untuk semua endpoint).
- Belum ada ERD/diagram relasi database di README.
- Belum ada fitur laporan penjualan atau dashboard admin.
- Beberapa test yang belum menyeluruh menyebabkan kemungkinan bug muncul.

## ERD
![ERD](https://raw.githubusercontent.com/sm888sm/backend-marketplace/refs/heads/main/erd.png)

## Struktur Project

```
backend-marketplace/
├── cmd/main/main.go
├── internal/
│   ├── config/
│   ├── controllers/
│   ├── middleware/
│   ├── models/
│   ├── repositories/
│   ├── routes/
│   └── services/
├── migrations/
├── pkg/
│   ├── auth/
│   └── utils/
├── Dockerfile
├── docker-compose.yml
├── Makefile
├── go.mod
├── README.md
└── .env (optional)
```

## Instalasi & Setup

1. **Clone repository**
   ```bash
   git clone https://github.com/sm888sm/backend-marketplace.git
   cd backend-marketplace
   ```

2. **Atur environment variable**
   - Bisa menggunakan file `.env` atau environment variable langsung.
   - Contoh `.env`:
     ```
     DB_HOST=localhost
     DB_PORT=5432
     DB_USER=myuser
     DB_PASSWORD=mypassword
     DB_NAME=marketplace
     SERVER_PORT=8080
     JWT_SECRET=pelindo888
     ```

3. **Install dependencies**
   ```bash
   go mod download
   ```

4. **Setup database**
   - Pastikan PostgreSQL sudah berjalan.
   - Buat database:
     ```bash
     make createdb
     ```
   - Jalankan migrasi:
     ```bash
     make migrate
     ```
   - Untuk reset database:
     ```bash
     make resetdb
     ```

5. **Jalankan aplikasi**
   ```bash
   make run
   ```
   Aplikasi akan berjalan di port sesuai `SERVER_PORT` (default: 8080).

## Perintah Makefile

- `make run` — Jalankan aplikasi
- `make migrate` — Jalankan migrasi database
- `make createdb` — Buat database
- `make dropdb` — Hapus database
- `make resetdb` — Drop lalu create database
- `make test` — Jalankan unit test

## Format Response API

- **Sukses:**
  ```json
  {
    "success": true,
    "message": "Pesan sukses",
    "data": { ... },
    "meta": { ... } // jika ada paginasi
  }
  ```
- **Error:**
  ```json
  {
    "success": false,
    "error": "Pesan error"
  }
  ```

## Contoh Endpoint

- **Register**
  - `POST /api/auth/register`
  - Body:
    ```json
    {
      "username": "merchant1",
      "password": "password",
      "email": "merchant1@example.com",
      "role": "merchant"
    }
    ```
- **Login**
  - `POST /api/auth/login`
  - Response:
    ```json
    {
      "success": true,
      "message": "Login berhasil",
      "data": {
        "token": "<jwt_token>",
        "user": {
          "id": 1,
          "username": "merchant1",
          "email": "merchant1@example.com",
          "role": "merchant",
          "created_at": "..."
        }
      }
    }
    ```

- **CRUD Produk (Merchant)**
  - Semua endpoint `/api/merchant/products/...` membutuhkan JWT dan role merchant.

- **CRUD Kategori**
  - Semua endpoint POST/PUT/DELETE `/api/categories` membutuhkan JWT dan role admin/merchant.

- **Order**
  - Semua endpoint `/api/orders` membutuhkan JWT dan role customer.

## Testing dengan Postman

- Gunakan file `Marketplace API Pelindo.postman_collection.json` untuk import koleksi endpoint ke Postman.
- Pastikan mengisi header:
  ```
  Authorization: Bearer <jwt_token>
  ```

## Catatan

- Password user di-hash dengan bcrypt.
- Semua tabel utama mendukung soft delete (`deleted_at`).
- Untuk development, seed user admin/customer sudah tersedia di migrasi awal.
