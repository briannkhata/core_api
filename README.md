# YathuERP - Complete Microservice Architecture

## 🎯 Project Overview

**YathuERP** is a comprehensive, production-ready microservice architecture for HR and Payroll management, following the exact MySQL schema structure with Go 1.22.

## 🏗️ Architecture Based on MySQL Schema

```
yathu_erp/                         # MAIN PROJECT ROOT
├── services/                         # All microservices
│   ├── employee-service/             # Employee Management (Port 8081)
│   ├── organization-service/         # Organization Management (Port 8082)
│   ├── attendance-service/            # Time & Attendance (Port 8083)
│   ├── leave-service/               # Leave Management (Port 8084)
│   ├── loan-service/                # Loan Management (Port 8085)
│   ├── payroll-service/              # Payroll Management (Port 8086)
│   ├── earnings-service/             # Earnings & Allowances (Port 8087)
│   ├── deductions-service/           # Deductions & Taxes (Port 8088)
│   ├── performance-service/          # Performance Management (Port 8089)
│   ├── banking-service/             # Banking & Benefits (Port 8090)
│   ├── auth-service/               # Authentication & Authorization (Port 8091)
│   ├── config-service/             # System Configuration (Port 8092)
│   └── audit-service/              # Audit & Settings (Port 8093)
├── shared/                           # Shared libraries
│   ├── config/                 # Configuration management
│   ├── database/               # Database utilities
│   ├── middleware/             # Auth, logging, tracing
│   ├── messaging/              # Event bus utilities
│   └── utils/                  # Common utilities
├── contracts/                        # API & event contracts
├── deploy/                           # Deployment configs
├── scripts/                          # Utility scripts
├── go.work                           # Go workspace
└── README.md                         # This file
```

## 🚀 Quick Start

### Prerequisites
- **Go 1.22+** (Latest version with improved performance and generics)
- Docker & Docker Compose
- PostgreSQL 14+
- Make (optional)

### Local Development

1. **Clone and Setup**
```bash
git clone <repository-url>
cd yathu_erp
go work sync
```

2. **Start Services**
```bash
# Start all services
docker-compose up -d

# Start individual service
cd services/employee-service
go run cmd/api/main.go
```

3. **Check Health**
```bash
curl http://localhost:8081/api/v1/employees/health
curl http://localhost:8082/api/v1/salary/health
curl http://localhost:8083/api/v1/earnings/health
```

## 🌐 API Documentation

### Employee Service (Port 8081)

#### Authentication
All endpoints require JWT authentication except health checks.

#### Endpoints

**Create Employee**
```bash
POST /api/v1/employees
Content-Type: application/json
Authorization: Bearer <jwt-token>

{
  "first_name": "John",
  "last_name": "Doe",
  "email": "john.doe@example.com",
  "phone": "+1234567890",
  "date_of_birth": "1990-01-01",
  "gender": "male",
  "address": "123 Main St",
  "city": "New York",
  "state": "NY",
  "country": "USA",
  "postal_code": "10001",
  "department_id": "uuid",
  "position_id": "uuid",
  "manager_id": "uuid",
  "salary": 75000.00
}
```

**Get All Employees**
```bash
GET /api/v1/employees?page=1&limit=10&search=john&department_id=uuid
Authorization: Bearer <jwt-token>
```

## 🔧 Configuration

### Environment Variables

```bash
# Server Configuration
PORT=8081
ENVIRONMENT=development
LOG_LEVEL=info
TIMEOUT_SECONDS=30

# Database Configuration
DATABASE_URL=postgres://user:password@localhost:5432/yathu_erp?sslmode=disable
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=yathu_erp
DB_SSL_MODE=disable

# Security
JWT_SECRET=your-super-secret-jwt-key
```

## 🐳 Docker Deployment

### Multi-Service Docker Compose

```yaml
version: '3.8'

services:
  postgres:
    image: postgres:15
    environment:
      POSTGRES_DB: yathu_erp
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: password
    volumes:
      - postgres_data:/var/lib/postgresql/data
    ports:
      - "5432:5432"
    networks:
      - yathu-erp-network

  employee-service:
    build: ./services/employee-service
    ports:
      - "8081:8081"
    environment:
      - DATABASE_URL: postgres://postgres:password@postgres:5432/yathu_erp?sslmode=disable
      - JWT_SECRET: your-super-secret-jwt-key
      - LOG_LEVEL: info
    depends_on:
      - postgres
    networks:
      - yathu-erp-network
    healthcheck:
      test: ["CMD", "wget", "--no-verbose", "--tries=1", "--spider", "http://localhost:8081/api/v1/employees/health", "||", "exit", "1"]
      interval: 30s
      timeout: 10s
      retries: 3

networks:
  yathu-erp-network:
    driver: bridge

volumes:
  postgres_data:
```

## 🔄 Development Workflow

### 1. Feature Development

```bash
# Create feature branch
git checkout -b feature/employee-management

# Make changes
# Run tests
go test ./services/employee-service/...

# Run service locally
cd services/employee-service && go run cmd/api/main.go
```

### 2. Testing

```bash
# Run unit tests
go test ./...

# Run integration tests
go test -tags=integration ./...

# Run API tests
go test -tags=api ./...
```

### 3. Code Quality

```bash
# Format code
go fmt ./...

# Lint code
golangci-lint run

# Run security scan
gosec ./...
```

## 📈 Benefits

### ✅ **Professional Architecture**
- Clean separation of concerns
- Industry-standard Go project structure
- Comprehensive testing strategy
- Production-ready deployment
- **Go 1.22** with latest features and performance improvements

### ✅ **Scalability**
- Independent service deployment
- Horizontal scaling capability
- Load balancing support

### ✅ **Maintainability**
- Clear code organization
- Comprehensive documentation
- Automated testing and deployment

### ✅ **Security**
- Authentication and authorization
- Input validation and sanitization
- Secure communication patterns

This architecture provides a **solid foundation** for a production-grade HR and payroll system using the latest Go 1.22!
