# GRID Admin

一个使用 Go 构建的简单后台管理系统，专为前端开发者学习后端开发而设计。

[English](./README.md)

## 技术栈

- **Web 框架**: [Gin](https://github.com/gin-gonic/gin)
- **ORM**: [GORM](https://gorm.io/)
- **数据库**: PostgreSQL
- **认证方式**: JWT (JSON Web Token)
- **密码加密**: bcrypt

## 项目结构

```
go-admin/
├── main.go                 # 入口文件
├── go.mod                  # Go 模块定义
├── config/
│   └── config.go           # 配置管理
├── database/
│   └── db.go               # 数据库连接
├── models/
│   └── user.go             # 用户模型
├── handlers/
│   └── auth.go             # 认证处理器
├── middleware/
│   └── jwt.go              # JWT 中间件
├── routes/
│   └── routes.go           # 路由定义
└── utils/
    └── response.go         # 统一响应格式
```

## 快速开始

### 环境要求

- Go 1.20+
- PostgreSQL 12+

### 数据库配置

创建名为 `go_admin` 的数据库：

```sql
CREATE DATABASE go_admin;
```

### 环境变量配置

| 变量名 | 说明 | 默认值 |
|--------|------|--------|
| DB_HOST | 数据库地址 | localhost |
| DB_PORT | 数据库端口 | 5432 |
| DB_USER | 数据库用户名 | postgres |
| DB_PASSWORD | 数据库密码 | postgres |
| DB_NAME | 数据库名称 | go_admin |
| JWT_SECRET | JWT 签名密钥 | your-secret-key |
| SERVER_PORT | 服务端口 | 8080 |

### 启动服务

```bash
cd go-admin
go run main.go
```

服务将在 `http://localhost:8080` 启动

## API 接口文档

### 公开接口（无需认证）

#### 用户注册

```
POST /api/auth/register
```

**请求体：**
```json
{
  "username": "admin",
  "password": "123456",
  "email": "admin@example.com"
}
```

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "username": "admin",
    "email": "admin@example.com",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

#### 用户登录

```
POST /api/auth/login
```

**请求体：**
```json
{
  "username": "admin",
  "password": "123456"
}
```

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user_info": {
      "id": 1,
      "username": "admin",
      "email": "admin@example.com"
    }
  }
}
```

### 受保护接口（需要 JWT 认证）

#### 获取用户信息

```
GET /api/user/profile
Authorization: Bearer <token>
```

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "username": "admin",
    "email": "admin@example.com",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

## 使用 cURL 快速测试

```bash
# 注册新用户
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"123456","email":"admin@example.com"}'

# 登录
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"123456"}'

# 获取用户信息（将 <token> 替换为登录返回的 token）
curl http://localhost:8080/api/user/profile \
  -H "Authorization: Bearer <token>"
```

## 许可证

MIT
