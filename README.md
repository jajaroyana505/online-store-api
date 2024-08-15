# API Documentation

## Authentication Endpoint

| Method | Endpoint         | Description    |
| ------ | ---------------- | -------------- |
| POST   | `/login`         | Login user     |
| POST   | `/register`      | Register user  |
| POST   | `/{id}/password` | Ganti password |
| GET    | `/logout`        | Logout         |

### Contoh Request:

**Login:**

```http
POST /login
Content-Type: application/json
{
  "username": "usernameexample",
  "password": "password123"
}
```

**Ganti Password:**

```http
POST /id/login
Content-Type: application/json
{
  "current_password": "currentpass",
  "new_password": "newPass"
}
```

## Product Endpoint (Protected by JWT)

| Method | Endpoint        | Description                             |
| ------ | --------------- | --------------------------------------- |
| POST   | `/product`      | Tambah produk                           |
| GET    | `/product`      | Ambil semua data produk                 |
| GET    | `/product/{id}` | Ambil data produk berdasarkan ID produk |
| PUT    | `/product/{id}` | Update data produk                      |
| DELETE | `/product/{id}` | Hapus data produk                       |

### Contoh Request:

**Tambah produk:**

```http
POST /api/product
Content-Type: application/json
{
  "nama_product": "piring",
  "stok": 100,
  "harga": 25000
}
```

## Order Endpoint (Protected by JWT)

| Method | Endpoint             | Description                           |
| ------ | -------------------- | ------------------------------------- |
| POST   | `/order`             | Tambah order                          |
| GET    | `/order`             | Ambil semua data order                |
| GET    | `/order/{id}`        | Ambil data order berdasarkan ID Order |
| PUT    | `/order/{id}/status` | Ubah status order                     |
| DELETE | `/order/{id}`        | Hapus data order                      |

### Contoh Request:

**Tambah order:**

```http
POST /api/order
Content-Type: application/json
{
  "id_product": 100,
  "jumlah": 100
}
```

**Update status:**

```http
PUT /api/order/id/status
Content-Type: application/json
{
  "status": "diproses"
}
```
