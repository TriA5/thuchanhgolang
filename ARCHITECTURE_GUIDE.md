# 📚 HƯỚNG DẪN KIẾN TRÚC DỰ ÁN - THỰC HÀNH GOLANG

## 🏗️ KIẾN TRÚC DỰ ÁN - CLEAN ARCHITECTURE

Dự án này sử dụng **Clean Architecture** với pattern **Repository - Usecase - Delivery**.

---

## 📁 CẤU TRÚC FOLDER VÀ CHỨC NĂNG

### 1️⃣ **`cmd/` - Entry Point (Cửa vào chính)**

```
cmd/
  api/
    main.go  ← Điểm khởi đầu của ứng dụng
```

**Chức năng:**
- Khởi động ứng dụng
- Load config từ `.env`
- Kết nối MongoDB
- Khởi tạo dependencies (logger, database, repositories, usecases, handlers)
- Start HTTP server

**Ví dụ thực tế:** Giống như công tắc điện chính của nhà, bật lên thì mọi thứ hoạt động.

---

### 2️⃣ **`config/` - Quản lý cấu hình**

```
config/
  config.go  ← Load environment variables từ .env
```

**Chức năng:**
- Load file `.env` bằng `godotenv`
- Parse environment variables vào struct
- Cung cấp config cho toàn bộ app (MongoDB URI, Port, Logger level...)

**Ví dụ thực tế:** Giống như bảng điều khiển trung tâm chứa mọi cài đặt.

---

### 3️⃣ **`internal/` - Core Business Logic (QUAN TRỌNG NHẤT)**

Đây là phần **não** của ứng dụng, chia theo từng module/domain.

#### 📦 **`internal/appconfig/` - Cấu hình kết nối**

```
internal/appconfig/
  mongo/
    connect.go  ← Kết nối MongoDB
```

**Chức năng:**
- Setup connection tới MongoDB
- Ping để verify connection
- Disconnect khi thoát

---

#### 📦 **`internal/models/` - Data Models (Entities)**

```
internal/models/
  branch.go  ← Định nghĩa struct Branch
  scope.go   ← Định nghĩa Scope (user context)
```

**Chức năng:**
- Định nghĩa các struct đại diện cho data
- Map với MongoDB collections (dùng BSON tags)

**Code example:**
```go
type Branch struct {
    ID        primitive.ObjectID `bson:"_id"`
    Name      string             `bson:"name"`
    Code      string             `bson:"code"`
    CreatedAt time.Time          `bson:"created_at"`
}
```

**Ví dụ thực tế:** Giống như bản thiết kế của một sản phẩm (có gì, cấu trúc ra sao).

---

#### 📦 **`internal/branch/` - Module Branch (Clean Architecture)**

Đây là một **domain/module** hoàn chỉnh theo Clean Architecture:

```
internal/branch/
  ├── repo_interface.go      ← Interface của Repository
  ├── repo_types.go          ← Types/DTOs cho Repository
  ├── uc_interface.go        ← Interface của Usecase
  ├── uc_types.go            ← Types/DTOs cho Usecase
  ├── delivery/http/         ← Layer Delivery (HTTP Handler)
  ├── usecase/               ← Layer Business Logic
  └── repository/mongo/      ← Layer Data Access
```

##### **🔹 Layer 1: Delivery (HTTP Handlers)**

```
delivery/http/
  ├── new.go              ← Khởi tạo Handler
  ├── handlers.go         ← HTTP handlers (create, update, delete...)
  ├── routes.go           ← Đăng ký routes
  ├── presenters.go       ← Convert model → response JSON
  ├── process_request.go  ← Validate & parse request
  └── errors.go           ← Map lỗi business → HTTP status
```

**Chức năng:**
- Nhận HTTP request từ client
- Validate input
- Gọi Usecase để xử lý business logic
- Trả về HTTP response

**Flow:**
```
HTTP Request → Handler.create()
  → processCreateRequest() (validate input)
  → uc.Create() (gọi business logic)
  → newDetailResp() (format response)
  → response.OK() (trả JSON về client)
```

**Ví dụ thực tế:** Giống như nhân viên bán hàng, nhận yêu cầu từ khách → xử lý → trả kết quả.

---

##### **🔹 Layer 2: Usecase (Business Logic)**

```
usecase/
  ├── new.go       ← Khởi tạo Usecase
  └── usecase.go   ← Logic nghiệp vụ (create, update, delete...)
```

**Chức năng:**
- Xử lý business logic
- Gọi Repository để tương tác với database
- Validate business rules
- Transform data

**Code example:**
```go
func (uc implUsecase) Create(ctx context.Context, sc models.Scope, input branch.CreateInput) (models.Branch, error) {
    // Xử lý logic: tạo code, alias từ name
    branch, err := uc.repo.Create(ctx, sc, branch.CreateOptions{
        Name:  input.Name,
        Code:  util.BuildCode(input.Name),   // Business logic
        Alias: util.BuildAlias(input.Name),  // Business logic
    })
    return branch, nil
}
```

**Ví dụ thực tế:** Giống như phòng kế toán/kế hoạch, xử lý logic kinh doanh, không quan tâm database là gì.

---

##### **🔹 Layer 3: Repository (Data Access)**

```
repository/mongo/
  ├── new.go     ← Khởi tạo Repository
  └── branch.go  ← CRUD operations với MongoDB
```

**Chức năng:**
- Tương tác trực tiếp với database
- Insert, Update, Delete, Query
- Convert data giữa app và database

**Code example:**
```go
func (repo implRepository) Create(ctx context.Context, sc models.Scope, opts branch.CreateOptions) (models.Branch, error) {
    col := repo.db.Collection("branches")
    
    branch := models.Branch{
        ID:        repo.db.NewObjectID(),
        Name:      opts.Name,
        Code:      opts.Code,
        CreatedAt: time.Now(),
    }
    
    _, err := col.InsertOne(ctx, branch)
    return branch, err
}
```

**Ví dụ thực tế:** Giống như thủ kho, chỉ lo lưu trữ và lấy hàng từ kho (database).

---

#### 📦 **`internal/httpserver/` - HTTP Server Setup**

```
internal/httpserver/
  ├── new.go          ← Khởi tạo Gin server
  ├── httpserver.go   ← Run server, shutdown gracefully
  └── handlers.go     ← Map routes cho các modules
```

**Chức năng:**
- Khởi tạo Gin framework
- Đăng ký routes từ các module
- Start HTTP server trên port cấu hình
- Graceful shutdown khi nhận SIGTERM

**Ví dụ thực tế:** Giống như hệ thống điện trong tòa nhà, kết nối mọi thứ lại với nhau.

---

#### 📦 **`internal/middleware/` - Middleware**

**Chức năng:**
- Authentication (JWT verify)
- Authorization (check permissions)
- Logging requests
- CORS
- Rate limiting

**Ví dụ thực tế:** Giống như bảo vệ ở cửa, kiểm tra ai được vào, ai không.

---

### 4️⃣ **`pkg/` - Reusable Packages (Công cụ dùng chung)**

Các package này có thể dùng trong nhiều project khác.

#### 📦 **`pkg/mongo/` - MongoDB Wrapper**

```
pkg/mongo/
  ├── mongo.go   ← Wrapper cho mongo-driver
  ├── errors.go  ← Custom errors
  └── utils.go   ← Helper functions
```

**Chức năng:**
- Wrap MongoDB official driver
- Tạo interface để dễ mock khi test
- Custom behaviors

---

#### 📦 **`pkg/log/` - Logger**

```
pkg/log/
  ├── new.go  ← Logger interface
  └── zap.go  ← Zap logger implementation
```

**Chức năng:**
- Logging (Info, Warn, Error, Debug...)
- Dùng Uber Zap
- Support nhiều formats (JSON, Console)

---

#### 📦 **`pkg/response/` - HTTP Response Helper**

```
pkg/response/
  ├── response.go  ← Format JSON response
  └── time.go      ← Time utilities
```

**Chức năng:**
- Chuẩn hóa response format
- Helper: `OK()`, `Error()`, `Unauthorized()`

**Example response:**
```json
{
  "error_code": 0,
  "message": "Success",
  "data": {...}
}
```

---

#### 📦 **`pkg/jwt/` - JWT Authentication**

```
pkg/jwt/
  ├── jwt.go     ← Generate & verify JWT
  ├── scope.go   ← User scope/permissions
  ├── utils.go   ← Helper functions
  └── errors.go  ← JWT errors
```

**Chức năng:**
- Tạo JWT token khi login
- Verify token từ request header
- Extract user info từ token

---

#### 📦 **`pkg/errors/` - Error Handling**

```
pkg/errors/
  ├── http.go        ← Map errors → HTTP status
  └── validation.go  ← Validation errors
```

**Chức năng:**
- Custom error types
- Map business errors → HTTP status codes
- Validation error messages

---

#### 📦 **`pkg/util/` - Utilities**

```
pkg/util/
  ├── utils.go   ← Helper functions
  └── locale.go  ← I18n support
```

**Chức năng:**
- `BuildCode()` - Tạo code từ name
- `BuildAlias()` - Tạo alias từ name
- String manipulation
- Date/Time utilities

---

## 🔄 FLOW HOÀN CHỈNH - VÍ DỤ: TẠO BRANCH MỚI

```
1. CLIENT gửi POST request
   POST /api/branches
   Body: { "name": "Chi nhánh Hà Nội" }
         ↓
2. HTTP Server (Gin) nhận request
         ↓
3. delivery/http/handlers.go → handler.create()
   - Validate input (processCreateRequest)
   - Parse JSON → CreateRequest struct
         ↓
4. usecase/usecase.go → uc.Create()
   - Business logic: BuildCode(), BuildAlias()
   - Chuẩn bị data
         ↓
5. repository/mongo/branch.go → repo.Create()
   - Insert vào MongoDB collection "branches"
   - Trả về Branch object
         ↓
6. Quay lại handler → newDetailResp()
   - Format Branch → JSON response
         ↓
7. response.OK() → Trả về client
   Response: {
     "error_code": 0,
     "message": "Success",
     "data": {
       "id": "...",
       "name": "Chi nhánh Hà Nội",
       "code": "chi-nhanh-ha-noi"
     }
   }
```

---

## 📝 TÓM TẮT CHỨC NĂNG TỪNG FOLDER

| Folder | Chức năng | Ví dụ thực tế |
|--------|-----------|---------------|
| `cmd/` | Khởi động app | Công tắc điện chính |
| `config/` | Quản lý cấu hình | Bảng điều khiển |
| `internal/models/` | Data structures | Bản thiết kế sản phẩm |
| `internal/branch/delivery/` | Nhận request, trả response | Nhân viên bán hàng |
| `internal/branch/usecase/` | Business logic | Phòng kế hoạch |
| `internal/branch/repository/` | Thao tác database | Thủ kho |
| `internal/httpserver/` | HTTP server | Hệ thống điện tòa nhà |
| `pkg/mongo/` | MongoDB wrapper | Công cụ chuyên dụng |
| `pkg/log/` | Logging | Hệ thống camera giám sát |
| `pkg/response/` | Format response | Bộ đóng gói sản phẩm |
| `pkg/jwt/` | Authentication | Hệ thống thẻ từ |
| `pkg/errors/` | Error handling | Hệ thống báo lỗi |

---

## 🎯 QUY TRÌNH 1: KHỞI ĐỘNG ỨNG DỤNG

```
┌─────────────────────────────────────────────────────────────────────┐
│                        go run cmd/api/main.go                       │
└─────────────────────────────────────────────────────────────────────┘
                                  │
                                  ▼
┌─────────────────────────────────────────────────────────────────────┐
│  BƯỚC 1: Load Configuration                                         │
│  ────────────────────────────                                       │
│  • godotenv.Load() → Đọc file .env                                 │
│  • env.Parse() → Parse vào Config struct                           │
│  • Result: cfg.Mongo.URI, cfg.Mongo.DBName, cfg.HTTPServer.Port   │
└─────────────────────────────────────────────────────────────────────┘
                                  │
                                  ▼
┌─────────────────────────────────────────────────────────────────────┐
│  BƯỚC 2: Connect MongoDB                                            │
│  ─────────────────────────                                          │
│  • mongo.NewClient(uri) → Tạo client                               │
│  • client.Connect(ctx) → Kết nối TCP + TLS handshake              │
│  • client.Ping(ctx) → Verify connection + Authentication          │
│  • client.Database(dbName) → Lấy DB instance                      │
└─────────────────────────────────────────────────────────────────────┘
                                  │
                                  ▼
┌─────────────────────────────────────────────────────────────────────┐
│  BƯỚC 3: Initialize Logger                                          │
│  ───────────────────────────                                        │
│  • pkgLog.InitializeZapLogger()                                    │
│  • Config: Level (debug/info), Mode (dev/prod), Encoding (json)   │
└─────────────────────────────────────────────────────────────────────┘
                                  │
                                  ▼
┌─────────────────────────────────────────────────────────────────────┐
│  BƯỚC 4: Initialize Dependencies                                    │
│  ─────────────────────────────────────                              │
│  • Repository: branchRepo = mongo.New(db, logger)                 │
│  • Usecase: branchUC = usecase.New(branchRepo, logger)           │
│  • Handler: branchHandler = http.New(logger, branchUC)           │
└─────────────────────────────────────────────────────────────────────┘
                                  │
                                  ▼
┌─────────────────────────────────────────────────────────────────────┐
│  BƯỚC 5: Setup HTTP Server                                          │
│  ────────────────────────────                                       │
│  • httpserver.New(logger, config)                                  │
│  • gin.Default() → Khởi tạo Gin framework                         │
│  • mapHandlers() → Đăng ký routes                                 │
│    ├─ POST   /api/branches                                        │
│    ├─ GET    /api/branches/:id                                    │
│    ├─ PUT    /api/branches/:id                                    │
│    └─ DELETE /api/branches/:id                                    │
└─────────────────────────────────────────────────────────────────────┘
                                  │
                                  ▼
┌─────────────────────────────────────────────────────────────────────┐
│  BƯỚC 6: Start Server                                               │
│  ──────────────────────                                             │
│  • gin.Run(":8080") → Lắng nghe HTTP requests                     │
│  • Graceful shutdown: Listen SIGTERM/SIGINT                       │
│  ✅ Server đang chạy tại http://localhost:8080                    │
└─────────────────────────────────────────────────────────────────────┘
```

---

## 🔄 QUY TRÌNH 2: XỬ LÝ HTTP REQUEST - CREATE BRANCH

```
┌──────────────────────────────────────────────────────────────────┐
│  CLIENT: POST http://localhost:8080/api/branches                │
│  Headers: Content-Type: application/json                        │
│  Body: {                                                         │
│    "name": "Chi nhánh Hà Nội"                                   │
│  }                                                               │
└──────────────────────────────────────────────────────────────────┘
                            │
                            ▼
┌──────────────────────────────────────────────────────────────────┐
│  GIN FRAMEWORK: Route Matching                                   │
│  • Match route: POST /api/branches → branchHandler.create()    │
└──────────────────────────────────────────────────────────────────┘
                            │
                            ▼
┌──────────────────────────────────────────────────────────────────┐
│  DELIVERY LAYER: internal/branch/delivery/http/handlers.go      │
│  ═══════════════════════════════════════════════════════════════ │
│  func (h handler) create(c *gin.Context)                        │
│                                                                  │
│  BƯỚC 1: Process Request                                        │
│  ├─ h.processCreateRequest(c)                                   │
│  ├─ Bind JSON to struct                                         │
│  ├─ Validate fields:                                            │
│  │  • name required, min length                                │
│  ├─ Extract scope (user info from JWT token)                   │
│  └─ Return: CreateRequest{Name: "Chi nhánh Hà Nội"}           │
│                                                                  │
│  ⚠️  Nếu validation fail:                                       │
│      → mapError() → response.Error(c, err)                     │
│      → Return 400 Bad Request                                   │
└──────────────────────────────────────────────────────────────────┘
                            │
                            ▼
┌──────────────────────────────────────────────────────────────────┐
│  USECASE LAYER: internal/branch/usecase/usecase.go              │
│  ═══════════════════════════════════════════════════════════════ │
│  func (uc implUsecase) Create(ctx, sc, input)                  │
│                                                                  │
│  BƯỚC 2: Business Logic                                         │
│  ├─ util.BuildCode("Chi nhánh Hà Nội")                         │
│  │  → "chi-nhanh-ha-noi" (slug format)                         │
│  ├─ util.BuildAlias("Chi nhánh Hà Nội")                        │
│  │  → "Chi nhanh Ha Noi" (remove diacritics)                   │
│  ├─ Validate business rules (duplicate code, etc)              │
│  └─ Prepare CreateOptions                                       │
│                                                                  │
│  ⚠️  Nếu business rule fail:                                    │
│      → Return custom error                                      │
└──────────────────────────────────────────────────────────────────┘
                            │
                            ▼
┌──────────────────────────────────────────────────────────────────┐
│  REPOSITORY LAYER: internal/branch/repository/mongo/branch.go   │
│  ═══════════════════════════════════════════════════════════════ │
│  func (repo implRepository) Create(ctx, sc, opts)              │
│                                                                  │
│  BƯỚC 3: Database Operations                                    │
│  ├─ col := db.Collection("branches")                           │
│  ├─ Tạo Branch object:                                          │
│  │  {                                                           │
│  │    _id: ObjectID("..."),                                    │
│  │    name: "Chi nhánh Hà Nội",                                │
│  │    code: "chi-nhanh-ha-noi",                                │
│  │    alias: "Chi nhanh Ha Noi",                               │
│  │    created_at: 2025-12-29T10:30:00Z,                        │
│  │    updated_at: 2025-12-29T10:30:00Z                         │
│  │  }                                                           │
│  ├─ col.InsertOne(ctx, branch)                                 │
│  │  → Insert vào MongoDB collection "branches"                 │
│  └─ Return branch object                                        │
│                                                                  │
│  ⚠️  Nếu database error:                                        │
│      → Log error → Return error                                 │
└──────────────────────────────────────────────────────────────────┘
                            │
                            ▼
┌──────────────────────────────────────────────────────────────────┐
│  USECASE LAYER: Return về handler                               │
│  └─ Return branch object                                        │
└──────────────────────────────────────────────────────────────────┘
                            │
                            ▼
┌──────────────────────────────────────────────────────────────────┐
│  DELIVERY LAYER: Format Response                                │
│  ═══════════════════════════════════════════════════════════════ │
│  BƯỚC 4: Present Data                                           │
│  ├─ h.newDetailResp(branch)                                     │
│  │  → Convert Branch model → JSON response format              │
│  ├─ response.OK(c, data)                                        │
│  └─ Return HTTP 200 OK                                          │
└──────────────────────────────────────────────────────────────────┘
                            │
                            ▼
┌──────────────────────────────────────────────────────────────────┐
│  CLIENT nhận response:                                           │
│  Status: 200 OK                                                  │
│  Body: {                                                         │
│    "error_code": 0,                                             │
│    "message": "Success",                                        │
│    "data": {                                                    │
│      "id": "676abcd123456789",                                  │
│      "name": "Chi nhánh Hà Nội",                                │
│      "code": "chi-nhanh-ha-noi",                                │
│      "alias": "Chi nhanh Ha Noi",                               │
│      "created_at": "2025-12-29T10:30:00Z"                       │
│    }                                                             │
│  }                                                               │
└──────────────────────────────────────────────────────────────────┘
```

---

## 🔐 QUY TRÌNH 3: AUTHENTICATION VỚI JWT

### FLOW 1: LOGIN - Lấy JWT Token

```
CLIENT: POST /api/auth/login
Body: { "username": "admin", "password": "123456" }
          │
          ▼
    Login Handler
          │
          ├─ Verify username/password với database
          │
          ▼
    ✅ Valid → Generate JWT Token
          │
    pkg/jwt/jwt.go:
    ├─ jwt.GenerateToken(userID, role, permissions)
    ├─ Payload: {
    │    "user_id": "123",
    │    "role": "admin",
    │    "scope": ["branch:read", "branch:write"],
    │    "exp": 1735560000
    │  }
    ├─ Sign with SECRET_KEY
    │
          ▼
    Return: {
      "error_code": 0,
      "message": "Success",
      "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
        "expires_at": "2025-12-30T10:30:00Z"
      }
    }
```

### FLOW 2: Sử dụng JWT Token

```
CLIENT: POST /api/branches
Headers: Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
Body: { "name": "Chi nhánh Hà Nội" }
          │
          ▼
┌─────────────────────────────────────────────────────────────────┐
│  MIDDLEWARE: internal/middleware/auth.go                         │
│  ──────────────────────────────────────────────────────────────  │
│  func AuthMiddleware()                                           │
│                                                                  │
│  1. Extract token từ header                                     │
│     token := c.GetHeader("Authorization")                       │
│     token = strings.TrimPrefix(token, "Bearer ")                │
│                                                                  │
│  2. Verify JWT token                                            │
│     claims, err := pkg/jwt.VerifyToken(token, SECRET_KEY)       │
│     ├─ Verify signature                                         │
│     ├─ Check expiration                                         │
│     └─ Parse claims                                             │
│                                                                  │
│  3. Check permissions                                            │
│     if !hasPermission(claims.Scope, "branch:write") {           │
│       return 403 Forbidden                                      │
│     }                                                            │
│                                                                  │
│  4. Set user info vào context                                   │
│     c.Set("user_id", claims.UserID)                            │
│     c.Set("scope", claims.Scope)                               │
│                                                                  │
│  5. Next() → Chuyển sang handler tiếp theo                      │
│                                                                  │
│  ⚠️  Nếu invalid token:                                         │
│      → response.Unauthorized(c)                                 │
│      → Return 401 Unauthorized                                  │
└─────────────────────────────────────────────────────────────────┘
          │
          ▼
    Handler xử lý request bình thường
    (đã có user info trong context)
```

---

## 🆕 QUY TRÌNH 4: THÊM MODULE MỚI (VD: PRODUCT)

### BƯỚC 1: Tạo Model

**File:** `internal/models/product.go`

```go
type Product struct {
    ID          primitive.ObjectID `bson:"_id"`
    Name        string             `bson:"name"`
    Code        string             `bson:"code"`
    Price       float64            `bson:"price"`
    Description string             `bson:"description"`
    BranchID    primitive.ObjectID `bson:"branch_id"`
    CreatedAt   time.Time          `bson:"created_at"`
    UpdatedAt   time.Time          `bson:"updated_at"`
}
```

### BƯỚC 2: Tạo thư mục module

```
internal/product/
    ├── repo_interface.go       ← Interface Repository
    ├── repo_types.go           ← DTOs cho Repository
    ├── uc_interface.go         ← Interface Usecase
    ├── uc_types.go             ← DTOs cho Usecase
    ├── delivery/http/
    │   ├── new.go
    │   ├── handlers.go
    │   ├── routes.go
    │   ├── presenters.go
    │   ├── process_request.go
    │   └── errors.go
    ├── usecase/
    │   ├── new.go
    │   └── usecase.go
    └── repository/mongo/
        ├── new.go
        └── product.go
```

### BƯỚC 3: Implement Repository

**File:** `internal/product/repository/mongo/product.go`

```go
type implRepository struct {
    db    mongo.Database
    l     log.Logger
    clock func() time.Time
}

func (repo implRepository) Create(ctx, sc, opts) {
    col := repo.db.Collection("products")
    product := models.Product{
        ID:        repo.db.NewObjectID(),
        Name:      opts.Name,
        Code:      opts.Code,
        Price:     opts.Price,
        CreatedAt: repo.clock(),
    }
    _, err := col.InsertOne(ctx, product)
    return product, err
}
```

### BƯỚC 4: Implement Usecase

**File:** `internal/product/usecase/usecase.go`

```go
func (uc implUsecase) Create(ctx, sc, input) {
    // Business logic: validate price, check duplicate code
    if input.Price < 0 {
        return error("Price must be positive")
    }
    
    product, err := uc.repo.Create(ctx, sc, product.CreateOptions{
        Name:  input.Name,
        Code:  util.BuildCode(input.Name),
        Price: input.Price,
    })
    return product, err
}
```

### BƯỚC 5: Implement HTTP Handler

**File:** `internal/product/delivery/http/handlers.go`

```go
func (h handler) create(c *gin.Context) {
    req, sc, err := h.processCreateRequest(c)
    if err != nil {
        response.Error(c, h.mapError(err))
        return
    }
    
    p, err := h.uc.Create(ctx, sc, req.toInput())
    if err != nil {
        response.Error(c, h.mapError(err))
        return
    }
    
    response.OK(c, h.newDetailResp(p))
}
```

### BƯỚC 6: Đăng ký Routes

**File:** `internal/product/delivery/http/routes.go`

```go
func (h handler) RegisterRoutes(r *gin.RouterGroup) {
    products := r.Group("/products")
    {
        products.POST("", h.create)
        products.GET("/:id", h.getByID)
        products.PUT("/:id", h.update)
        products.DELETE("/:id", h.delete)
        products.GET("", h.list)
    }
}
```

### BƯỚC 7: Wire dependencies trong main.go

**File:** `cmd/api/main.go`

```go
// Initialize Product module
productRepo := productMongo.New(db, l)
productUC := productUsecase.New(productRepo, l)
productHandler := productHTTP.New(l, productUC)

// Register routes
srv := httpserver.New(l, httpserver.Config{
    Port:           cfg.HTTPServer.Port,
    Database:       db,
    ProductHandler: productHandler,
})
```

### BƯỚC 8: Test API

```bash
POST http://localhost:8080/api/products
Headers: Authorization: Bearer <token>
Body: {
  "name": "Laptop Dell XPS 13",
  "price": 25000000,
  "description": "Laptop cao cấp",
  "branch_id": "676abcd123456789"
}

Response: {
  "error_code": 0,
  "message": "Success",
  "data": {
    "id": "676xyz...",
    "name": "Laptop Dell XPS 13",
    "code": "laptop-dell-xps-13",
    "price": 25000000
  }
}
```

---

## ⚠️ QUY TRÌNH 5: XỬ LÝ LỖI (ERROR HANDLING)

### PHÂN CẤP LỖI TRONG HỆ THỐNG

#### 1. VALIDATION ERROR (Client lỗi)

**Nguồn:** Delivery Layer

- Missing required field
- Invalid format (email, phone)
- Out of range value

**Xử lý:**
- `pkg/errors/validation.go`
- Return 400 Bad Request
- Message: "name is required"

#### 2. BUSINESS LOGIC ERROR

**Nguồn:** Usecase Layer

- Duplicate code
- Insufficient balance
- Out of stock

**Xử lý:**
- Custom error types
- Return 400/409 Conflict
- Message: "Branch code already exists"

#### 3. AUTHENTICATION ERROR

**Nguồn:** Middleware

- Missing token
- Invalid token
- Token expired

**Xử lý:**
- Return 401 Unauthorized
- Message: "Invalid or expired token"

#### 4. AUTHORIZATION ERROR

**Nguồn:** Middleware

- Insufficient permissions
- Access denied

**Xử lý:**
- Return 403 Forbidden
- Message: "You don't have permission"

#### 5. DATABASE ERROR

**Nguồn:** Repository Layer

- Connection timeout
- Duplicate key
- Query error

**Xử lý:**
- Log chi tiết error
- Return 500 Internal Server Error
- Message: "Something went wrong"

#### 6. SYSTEM ERROR

**Nguồn:** Bất kỳ layer nào

- Out of memory
- Network error
- Panic/Crash

**Xử lý:**
- Recovery middleware
- Log stack trace
- Return 500 Internal Server Error
- Alert team (email/Slack)

### FLOW XỬ LÝ LỖI CHI TIẾT

```
Request → Handler
            │
            ├─ Validation error?
            │  └─ YES → mapError() → 400 Bad Request
            │
            ▼
          Usecase
            │
            ├─ Business error?
            │  └─ YES → return custom error → 400/409
            │
            ▼
          Repository
            │
            ├─ Database error?
            │  ├─ Log: repo.l.Errorf(ctx, "error: %v", err)
            │  └─ return error
            │
            ▼
          Usecase (nhận error)
            │
            ├─ Log: uc.l.Errorf(ctx, "error: %v", err)
            └─ return error
            │
            ▼
          Handler (nhận error)
            │
            ├─ mapError(err) → HTTP status + message
            ├─ Log: h.l.Warnf(ctx, "error: %v", err)
            └─ response.Error(c, mappedError)
            │
            ▼
          Client nhận response:
          {
            "error_code": 500,
            "message": "Something went wrong",
            "data": null
          }
```

### CODE EXAMPLE: Map Error

**File:** `internal/branch/delivery/http/errors.go`

```go
func (h handler) mapError(err error) response.Resp {
    switch {
    case errors.Is(err, branch.ErrDuplicateCode):
        return response.Resp{
            ErrorCode: 409,
            Message:   "Branch code already exists",
        }
    case errors.Is(err, branch.ErrNotFound):
        return response.Resp{
            ErrorCode: 404,
            Message:   "Branch not found",
        }
    case pkgErrors.IsValidationError(err):
        return response.Resp{
            ErrorCode: 400,
            Message:   err.Error(),
        }
    default:
        return response.Resp{
            ErrorCode: 500,
            Message:   "Something went wrong",
        }
    }
}
```

---

## 📊 QUY TRÌNH 6: TESTING WORKFLOW

### LEVEL 1: Unit Test Repository

**File:** `internal/branch/repository/mongo/branch_test.go`

```go
func TestCreate(t *testing.T) {
    // Mock MongoDB
    mockDB := &mocks.MockDatabase{}
    mockCol := &mocks.MockCollection{}
    
    mockDB.On("Collection", "branches").Return(mockCol)
    mockCol.On("InsertOne", mock.Anything, mock.Anything).
        Return(nil, nil)
    
    // Test
    repo := New(mockDB, logger)
    branch, err := repo.Create(ctx, sc, opts)
    
    // Assert
    assert.NoError(t, err)
    assert.Equal(t, "test", branch.Name)
    mockCol.AssertExpectations(t)
}
```

### LEVEL 2: Unit Test Usecase

**File:** `internal/branch/usecase/usecase_test.go`

```go
func TestCreate(t *testing.T) {
    // Mock Repository
    mockRepo := &mocks.MockRepository{}
    mockRepo.On("Create", mock.Anything, mock.Anything, mock.Anything).
        Return(models.Branch{Name: "test"}, nil)
    
    // Test
    uc := New(mockRepo, logger)
    branch, err := uc.Create(ctx, sc, input)
    
    // Assert
    assert.NoError(t, err)
    assert.Equal(t, "test", branch.Name)
    mockRepo.AssertExpectations(t)
}
```

### LEVEL 3: Integration Test Handler

**File:** `internal/branch/delivery/http/handlers_test.go`

```go
func TestCreateHandler(t *testing.T) {
    // Mock Usecase
    mockUC := &mocks.MockUsecase{}
    mockUC.On("Create", mock.Anything, mock.Anything, mock.Anything).
        Return(models.Branch{Name: "test"}, nil)
    
    // Setup Gin test
    gin.SetMode(gin.TestMode)
    r := gin.Default()
    handler := New(logger, mockUC)
    r.POST("/branches", handler.create)
    
    // Test request
    body := `{"name":"test"}`
    req := httptest.NewRequest("POST", "/branches", strings.NewReader(body))
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    
    // Assert
    assert.Equal(t, 200, w.Code)
    mockUC.AssertExpectations(t)
}
```

### LEVEL 4: E2E Test

**File:** `test/e2e/branch_test.go`

```go
func TestCreateBranchE2E(t *testing.T) {
    // Setup: Start real server + real MongoDB
    db := setupTestDB(t)
    defer cleanupTestDB(t, db)
    
    server := startTestServer(t, db)
    defer server.Close()
    
    // Test request
    resp, err := http.Post(
        server.URL+"/api/branches",
        "application/json",
        strings.NewReader(`{"name":"test"}`),
    )
    
    // Assert
    assert.NoError(t, err)
    assert.Equal(t, 200, resp.StatusCode)
    
    // Verify in database
    var branch models.Branch
    db.Collection("branches").FindOne(ctx, bson.M{"name": "test"}).
        Decode(&branch)
    assert.Equal(t, "test", branch.Name)
}
```

### CHẠY TESTS

```bash
# Chạy tất cả tests
go test ./...

# Chạy test một package
go test ./internal/branch/usecase/

# Chạy test với coverage
go test -cover ./...

# Chạy test với race detector
go test -race ./...

# Chạy E2E tests
go test -tags=e2e ./test/e2e/
```

---

## 🚀 TÓM TẮT CÁC QUY TRÌNH CHÍNH

| Quy trình | Mục đích | Thời gian ước tính |
|-----------|----------|-------------------|
| **Khởi động app** | Setup DB, routes, start server | 2-5 giây |
| **Xử lý request** | Nhận → Validate → Logic → DB → Response | 50-200ms |
| **Authentication** | Login → JWT token → Verify token | 10-50ms |
| **Thêm module mới** | Scaffold repository → usecase → handler | 30-60 phút |
| **Xử lý lỗi** | Catch → Log → Map → Response | Tự động |
| **Testing** | Unit → Integration → E2E | Theo độ phức tạp |

---

## 🎯 TẠI SAO PHÂN CHIA NHƯ VẬY?

### ✅ Ưu điểm Clean Architecture:

1. **Tách biệt concerns:**
   - Delivery: Lo HTTP
   - Usecase: Lo business logic
   - Repository: Lo database

2. **Dễ test:**
   - Mock từng layer độc lập
   - Test business logic không cần database thật

3. **Dễ thay đổi:**
   - Đổi MongoDB → PostgreSQL: Chỉ sửa Repository
   - Đổi HTTP → gRPC: Chỉ sửa Delivery
   - Business logic không đổi

4. **Tái sử dụng:**
   - `pkg/` dùng được cho nhiều project
   - Các module độc lập

5. **Mở rộng dễ:**
   - Thêm module mới (product, user...) theo pattern giống branch

---

## 📚 KẾT LUẬN

Đây là một kiến trúc rất chuyên nghiệp, dễ maintain và scale. Các quy trình được thiết kế để:

- ✅ Dễ maintain
- ✅ Dễ mở rộng
- ✅ Dễ test
- ✅ Nhất quán về coding style
- ✅ Tách biệt rõ ràng các concerns

---

**Tài liệu được tạo tự động - Ngày 29/12/2025**
