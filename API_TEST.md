# Golang 认证系统 API 测试指南

本文档展示了如何测试新实现的认证系统功能。

## 前置条件

1. 启动服务器：
```bash
go run main.go
```

服务器将在 `http://localhost:8080` 启动

## 公开接口测试

### 1. 健康检查
```bash
curl -X GET http://localhost:8080/api/public/health
```

期望响应：
```json
{
  "message": "OK",
  "data": {
    "status": "healthy",
    "service": "golang-auth-system"
  }
}
```

### 2. 版本信息
```bash
curl -X GET http://localhost:8080/api/public/version
```

期望响应：
```json
{
  "message": "OK", 
  "data": {
    "version": "1.0.0",
    "build_time": "2024-01-01",
    "go_version": "1.19"
  }
}
```

## 认证接口测试

### 3. 用户登录
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin@qq.com",
    "password": "admin123"
  }'
```

期望响应：
```json
{
  "message": "登录成功",
  "data": {
    "token": "eyJhbGci...",
    "user": {
      "id": 1,
      "name": "张三",
      "email": "admin@qq.com",
      "age": 18,
      "address": "杭州",
      "phone": 132000000000,
      "created_at": "2024-01-01 12:00:00",
      "updated_at": "2024-01-01 12:00:00"
    }
  }
}
```

### 4. 用户注册
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "新用户",
    "email": "newuser@example.com",
    "password": "NewPass123!",
    "age": 25,
    "address": "上海",
    "phone": 139000000000
  }'
```

### 5. 获取当前用户信息（需要Token）
```bash
curl -X GET http://localhost:8080/api/auth/me \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

### 6. 用户注销（需要Token）
```bash
curl -X POST http://localhost:8080/api/auth/logout \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

### 7. 修改密码（需要Token）
```bash
curl -X PUT http://localhost:8080/api/auth/password \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "old_password": "admin123",
    "new_password": "NewPassword123!"
  }'
```

### 8. 获取所有用户（需要管理员权限）
```bash
curl -X GET http://localhost:8080/api/auth/users \
  -H "Authorization: Bearer ADMIN_TOKEN_HERE"
```

## 受保护的用户接口测试

### 9. 获取用户信息（需要Token）
```bash
curl -X GET http://localhost:8080/api/getUser \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

### 10. 添加用户（需要管理员权限）
```bash
curl -X POST http://localhost:8080/api/addUser \
  -H "Authorization: Bearer ADMIN_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "测试用户",
    "email": "test@example.com",
    "password": "TestPass123!",
    "age": 30,
    "address": "深圳",
    "phone": 135000000000
  }'
```

## 错误场景测试

### 11. 无Token访问受保护接口
```bash
curl -X GET http://localhost:8080/api/getUser
```

期望响应：
```json
{
  "code": 40103,
  "message": "用户未登录"
}
```

### 12. 错误的登录凭据
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin@qq.com",
    "password": "wrongpassword"
  }'
```

期望响应：
```json
{
  "code": 40101,
  "message": "用户名或密码错误"
}
```

### 13. 无效Token
```bash
curl -X GET http://localhost:8080/api/auth/me \
  -H "Authorization: Bearer invalid_token"
```

期望响应：
```json
{
  "code": 40102,
  "message": "Token无效或已过期"
}
```

## 预置测试用户

系统预置了以下测试用户：

1. **管理员用户**
   - 用户名：`admin@qq.com` 或 `张三`
   - 密码：`admin123`
   - 权限：管理员权限

2. **普通用户**
   - 用户名：`user@qq.com` 或 `李四`
   - 密码：`user123`
   - 权限：普通用户权限

## PowerShell 测试示例

如果在 Windows PowerShell 中测试，请使用以下格式：

```powershell
# 登录
$response = Invoke-RestMethod -Uri "http://localhost:8080/api/auth/login" -Method POST -ContentType "application/json" -Body '{"username":"admin@qq.com","password":"admin123"}'

# 保存Token
$token = $response.data.token

# 使用Token访问受保护接口
Invoke-RestMethod -Uri "http://localhost:8080/api/auth/me" -Method GET -Headers @{"Authorization"="Bearer $token"}
```

## 安全注意事项

1. **密码要求**：密码必须包含大小写字母、数字和特殊字符，且长度至少8位
2. **Token管理**：Token有24小时有效期，注销后会被加入黑名单
3. **权限控制**：管理员接口需要admin@qq.com用户权限
4. **CORS支持**：系统已配置CORS中间件支持跨域请求

## 运行单元测试

```bash
# 测试JWT工具
go test ./utils -v

# 测试所有模块
go test ./... -v
```