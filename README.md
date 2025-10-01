# Audio Upload Service (Go + Gin + Postgres)

A simple Go service that demonstrates how to build an ingestion pipeline for audio (or any file type) with metadata storage.

This project uses:

[Gin](https://github.com/gin-gonic/gin)
 – lightweight HTTP framework

[PostgreSQL](https://www.postgresql.org/)
 – relational DB for upload metadata

Pluggable storage layer – local filesystem by default, easily extendable to S3

It supports:

- Uploading files (with device/location metadata)
- Persisting metadata into Postgres
- Fetching metadata by ID

## 🚀 Features

Upload an audio file with metadata (device ID, location, timestamp)

Store the file via an abstract FileStorer (local disk now, S3 later)

Save metadata (id, path, device_id, location, uploaded_at) in Postgres

Retrieve metadata via GET /uploads/:id

Clean service layering (Handler → Service → Repository + Storer)

## Flow:
Handler → calls `FileService` → orchestrates `FileRepository` (DB) + `FileStorer` (storage).

## 🛠 Setup

1. Clone & Install
git clone <https://github.com/rmschick/audio-upload-service.git>
cd audio-upload-service
go mod tidy

2. Database

Create a Postgres DB:

```bash
createdb audio_db
```

Minimal schema:

```sql
CREATE TABLE uploads (
    id TEXT PRIMARY KEY,
    path TEXT NOT NULL,
    device_id TEXT,
    location TEXT,
    uploaded_at TIMESTAMPTZ DEFAULT now()
);
```

3. Config

Set your DB connection string (local example):

```bash
export DATABASE_URL="user=postgres password=postgres dbname=audio_db sslmode=disable"
```

4. Run
go run main.go

## 📡 API

Upload File

POST /upload

Request (multipart/form-data):

```bash
curl -X POST <http://localhost:8080/upload> \
  -F "file=@/Users/ryan/documents/test.mp3" \
  -F "device_id=device-123" \
  -F "location=Dallas,TX"
```

Response:

```json
{
  "id": "c7c1e5c9-8bb2-4cc7-8c72-10d73f97d11a",
  "path": "./uploads/c7c1e5c9_test.mp3",
  "device_id": "device-123",
  "location": "Dallas,TX",
  "uploaded_at": "2025-09-30T20:00:00Z"
}
```

Get Upload Metadata

GET /uploads/:id

```bash
curl <http://localhost:8080/uploads/c7c1e5c9-8bb2-4cc7-8c72-10d73f97d11a>
```

Response:

```json
{
  "id": "c7c1e5c9-8bb2-4cc7-8c72-10d73f97d11a",
  "path": "./uploads/c7c1e5c9_test.mp3",
  "device_id": "device-123",
  "location": "Dallas,TX",
  "uploaded_at": "2025-09-30T20:00:00Z"
}
```

## Future Improvements

Add an S3Storer implementation to store files in S3 instead of local disk.

Add background processing (extract duration, file size, codec from audio).

Add structured logging + metrics (Prometheus, OpenTelemetry).

Add Dockerfile & docker-compose for one-click setup (service + Postgres).

## 📖 License

MIT
