# E-Wallet Framework (Golang)

![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat-square&logo=go&logoColor=white)
![Architecture](https://img.shields.io/badge/Architecture-Clean-orange?style=flat-square)
![Microservices](https://img.shields.io/badge/Microservices-UMS%20%7C%20Wallet%20%7C%20Tx-blue?style=flat-square)

A production-ready Backend Framework for an E-Wallet application built in Go. This repository serves as a live progress tracker for building a robust, scalable, and enterprise-grade backend system.

## 🏗️ System Architecture

Arsitektur aplikasi E-Wallet ini memisahkan layanan menjadi beberapa *microservices* internal. Komunikasi ke *Client* menggunakan HTTP REST, sedangkan komunikasi antar layanan dan validasi token menggunakan gRPC.

```mermaid
flowchart LR
    Client([📱 Client])
    
    subgraph "Core Services (HTTP API)"
        UMS["User Management Service (UMS)"]
        Wallet["Wallet Service"]
        Transaction["Transaction Service"]
    end
    
    subgraph "Databases"
        DB_UMS[("MySQL (Users)")]
        DB_Wallet[("MySQL (Wallets)")]
        DB_Tx[("MySQL (Transactions)")]
    end
    
    subgraph "Internal Services (gRPC)"
        Notification["Notification Service"]
    end
    
    SMTP["📧 Google SMTP"]
    
    Client -- "HTTP" --> UMS
    Client -- "HTTP" --> Wallet
    Client -- "HTTP" --> Transaction
    
    UMS --> DB_UMS
    Wallet --> DB_Wallet
    Transaction --> DB_Tx
    
    Wallet -- "Validate Token (gRPC)" --> UMS
    Transaction -- "Validate Token (gRPC)" --> UMS
    
    Wallet -- "Send Email (gRPC)" --> Notification
    Transaction -- "Send Email (gRPC)" --> Notification
    
    Notification -- "SMTP" --> SMTP
```

## 🛠️ Tech Stack

- **Language:** Go 1.25+
- **HTTP Router:** [Gin](https://github.com/gin-gonic/gin)
- **RPC:** gRPC
- **Database ORM:** [GORM](https://gorm.io/)
- **Database Engine:** MySQL
- **Configuration:** [Koanf](https://github.com/knadh/koanf)
- **Logging:** [Zap](https://github.com/uber-go/zap)

---

## 📈 Progress Tracker

### ✅ Yang Sudah Dibuat (Done)

**Core Architecture & Bootstrapping**
- [x] Hierarchical Config Management (`.env` & OS Env)
- [x] Structured Logging Setup (Zap)
- [x] Database Connection Pooling & Ping Verification
- [x] HTTP Server (Gin) & gRPC Server Implementation
- [x] Graceful Shutdown (Signal, Context Cancellation, WaitGroup)
- [x] Health Check Endpoint (HTTP GET /health) with DB Ping

**User Management Service (UMS)**
- [x] DB Migration, Register & Password Hashing (Bcrypt)
- [x] Login JWT & Session DB Storage
- [x] Logout Mechanism & Auth Middleware
- [x] Refresh Token Sync & Session Sync
- [x] Internal Token Validation Server (gRPC)

---

### 🎯 Target Selanjutnya (Up Next)

**Wallet Service**
- [ ] Setup Repository & DB Connection
- [ ] API Create Wallet (Dipanggil UMS Pasca-Register)
- [ ] API Credit Balance
- [ ] API Debit Balance
- [ ] API Get Balance
- [ ] API Wallet History (Log Transaksi)

---

## 🚀 Cara Menjalankan Aplikasi Lokal

1. Pastikan mesin MySQL sudah menyala di lokal.
2. Buat database bernama `ewallet` (`mysql -u root -e "CREATE DATABASE ewallet;"`).
3. Buat file `.env` dari `.env.example` dan isi `DB_URI` dengan kredensial aslimu.
4. Jalankan aplikasi:
   ```bash
   go run cmd/api/*.go
   ```
