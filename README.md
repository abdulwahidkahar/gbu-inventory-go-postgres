# GBU Inventory API - Go PostgreSQL

![Go](https://img.shields.io/badge/Go-1.25-blue)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-17-336791)
![Docker](https://img.shields.io/badge/Docker-Compose-2496ED)

API CRUD sederhana untuk manajemen data `inventory items` PT. Gemilang Bisnis Universal menggunakan `Go` dan `PostgreSQL`.

## 🚀 Quick Start dengan Docker

### Prasyarat

- Docker & Docker Compose terinstall

### Menjalankan Aplikasi

```bash
# Clone dan masuk ke directory
git clone <repo-url>
cd gbu-inventory-go-postgres

# Start services (PostgreSQL + Go API)
docker compose up -d

# Verifikasi services berjalan
docker compose ps

# Test API
curl http://localhost:8081/health
```

API siap diakses di **http://localhost:8081**
PostgreSQL siap diakses di **localhost:5432**

Untuk detail lebih lanjut, lihat [DOCKER.md](DOCKER.md)

## 📋 Tech Stack

- **Backend**: Go 1.25
- **Database**: PostgreSQL 17
- **Driver**: pgx/v5
- **HTTP Server**: net/http (stdlib)

## 🏗️ Project Structure

```
.
├── cmd/api/              # Entry point aplikasi
├── internal/
│   ├── config/          # Configuration & env management
│   ├── db/              # Database schema & migrations
│   ├── httpapi/         # HTTP handlers & routing
│   └── inventory/       # Business logic
├── Dockerfile           # Docker image definition
├── docker-compose.yml   # Multi-container setup
└── go.mod              # Go dependencies
```

## 📡 API Endpoints

| Method   | Endpoint                    | Deskripsi                  |
| -------- | --------------------------- | -------------------------- |
| `GET`    | `/health`                   | Health check API           |
| `GET`    | `/api/inventory-items`      | Get semua inventory items  |
| `GET`    | `/api/inventory-items/{id}` | Get inventory item by ID   |
| `POST`   | `/api/inventory-items`      | Create inventory item baru |
| `PUT`    | `/api/inventory-items/{id}` | Update inventory item      |
| `DELETE` | `/api/inventory-items/{id}` | Delete inventory item      |

## 📝 Request/Response Example

### Create Item (POST)

**Request:**

```bash
curl -X POST http://localhost:8081/api/inventory-items \
  -H "Content-Type: application/json" \
  -d '{
    "sku": "GBU-001",
    "name": "Printer Thermal",
    "category": "Office Equipment",
    "supplier_name": "PT Sumber Jaya Abadi",
    "purchase_price": 850000,
    "selling_price": 975000,
    "stock": 20,
    "reorder_level": 5,
    "is_active": true
  }'
```

**Response (201):**

```json
{
  "success": true,
  "data": {
    "id": 1,
    "sku": "GBU-001",
    "name": "Printer Thermal",
    "category": "Office Equipment",
    "supplier_name": "PT Sumber Jaya Abadi",
    "purchase_price": 850000,
    "selling_price": 975000,
    "stock": 20,
    "reorder_level": 5,
    "is_active": true,
    "created_at": "2026-04-15T10:30:00Z",
    "updated_at": "2026-04-15T10:30:00Z"
  }
}
```

## � Screenshots - Postman Testing

### List All Items

![List Items](list.png)

### Get Item by ID

![Get by ID](getbyid.png)

### Create Item

![Create Item](create.png)

### Update Item

![Update Item](update.png)

### Delete Item

![Delete Item](delete.png)

## �🛠️ Development Setup (Local)

### Prasyarat

- Go 1.25+
- PostgreSQL 12+

### Setup Awal

```bash
# 1. Copy .env
cp .env.example .env

# 2. Update POSTGRES_DSN di .env sesuai lokasi PostgreSQL Anda
POSTGRES_DSN=postgres://postgres:password@localhost:5432/gbu_inventory?sslmode=disable

# 3. Run aplikasi
go run ./cmd/api

# API akan berjalan di http://localhost:8081
```

### Environment Variables

| Variable       | Default                                                                     | Deskripsi                  |
| -------------- | --------------------------------------------------------------------------- | -------------------------- |
| `APP_PORT`     | 8081                                                                        | Port aplikasi              |
| `POSTGRES_DSN` | `postgres://postgres:postgres@localhost:5432/gbu_inventory?sslmode=disable` | Database connection string |

### Database Schema

Table `inventory_items` dibuat otomatis saat startup jika belum ada.

```sql
CREATE TABLE inventory_items (
  id SERIAL PRIMARY KEY,
  sku VARCHAR(50) UNIQUE NOT NULL,
  name VARCHAR(255) NOT NULL,
  category VARCHAR(100),
  supplier_name VARCHAR(255),
  purchase_price DECIMAL(12,2),
  selling_price DECIMAL(12,2),
  stock INTEGER,
  reorder_level INTEGER,
  is_active BOOLEAN,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);
```

## 📚 File Pendukung

| File                      | Deskripsi                                      |
| ------------------------- | ---------------------------------------------- |
| `Dockerfile`              | Docker image configuration                     |
| `docker-compose.yml`      | Multi-container orchestration                  |
| `.dockerignore`           | Files excluded dari Docker build               |
| `requests.http`           | HTTP requests untuk REST client (VS Code, etc) |
| `postman_collection.json` | Postman collection untuk testing               |
| `postgres-seed.sql`       | Sample data untuk database                     |
| `DOCKER.md`               | Dokumentasi Docker lengkap                     |

## 🧪 Testing

### Menggunakan Postman

1. Import `postman_collection.json` ke Postman
2. Set base_url variable: `http://localhost:8081`
3. Run requests

### Menggunakan REST Client (VS Code)

Buka `requests.http` dan click "Send Request"

### Menggunakan cURL

Lihat contoh endpoint di section API Endpoints

## 🐛 Troubleshooting

### Port sudah digunakan

Edit port di `docker-compose.yml`:

```yaml
ports:
  - "8082:8081" # Ubah 8082 ke port kosong lainnya
```

### Database connection error

- Pastikan PostgreSQL container fully healthy: `docker compose logs postgres`
- Tunggu ~10 detik setelah `docker compose up`

### Rebuild dari scratch

```bash
docker compose down -v
docker compose up --build
```

## 📖 Notes

- Aplikasi auto-create table `inventory_items` saat startup
- Tabel sudah include created_at & updated_at fields
- Validasi SKU harus unique
- Response format selalu consistent dengan `success` & `data` fields
