# Docker Setup Guide

## Prasyarat

- Docker dan Docker Compose sudah terinstall

## Cara Menjalankan

### 1. Build dan Start Services

```bash
docker-compose up -d
```

Perintah ini akan:

- Build Docker image untuk aplikasi Go
- Start PostgreSQL container
- Start API container
- Inisialisasi database dengan seed data

### 2. Verifikasi Services Berjalan

```bash
docker-compose ps
```

### 3. Test API

```bash
# Health check
curl http://localhost:8081/health

# Get all inventory items
curl http://localhost:8081/api/inventory-items
```

## Melihat Logs

```bash
# Semua services
docker-compose logs -f

# Hanya API
docker-compose logs -f api

# Hanya PostgreSQL
docker-compose logs -f postgres
```

## Stop Services

```bash
docker-compose down
```

## Menghapus Data (Volume)

```bash
docker-compose down -v
```

Akan menghapus PostgreSQL data volume agar database kosong untuk run berikutnya.

## Rebuild Image

```bash
docker-compose up --build -d
```

## Akses PostgreSQL dari Host

```bash
psql -h localhost -U postgres -d gbu_inventory
```

Password: `postgres`

## Troubleshooting

### API can't connect to database

Pastikan PostgreSQL fully healthy dengan menunggu ~15 detik atau check logs:

```bash
docker-compose logs postgres
```

### Port sudah terpakai

Ubah port di `docker-compose.yml`:

```yaml
ports:
  - "8082:8081" # ubah 8082 ke port yang kosong
```

### Rebuild dari scratch

```bash
docker-compose down -v
docker-compose up --build -d
```
