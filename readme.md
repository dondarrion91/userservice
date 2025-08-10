# 🚀 user service

Service builded with go, performing crud handler with DAO pattern and dependency injection.

## 📋 Prerequisites

Before you start, make sure you have installed:

- [Go](https://go.dev/dl/) **>= 1.21**
- [Git](https://git-scm.com/)

## 🛠 Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/dondarrion91/userservice.git
   cd userservice

2. Download dependencies
```bash
go mod tidy
```

3. Set up environment variables:
```bash
vim main.sh
```
Set environment variables in main.sh shell script
```bash
export DB_HOST=<DB_HOST>
export DB_NAME=user_service
export DB_USER=<DB_USER>
export DB_PASS=<DB_PASS>
export DB_PORT=3306
export PORT=<PORT>

go run cmd/userService/main.go
```

4. Migrate db
```bash
./automigrate.sh
```

5. Execute main script
```bash
main.sh
```
