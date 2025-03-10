# Gin+Gorm SQLite Web服务

这是一个基于Golang的简单Web服务器，使用Gin框架和Gorm ORM，提供用户表的增删改查功能，数据库采用SQLite。

## 项目结构

```
├── controllers/       # 控制器层，处理HTTP请求
├── database/          # 数据库配置
├── models/            # 数据模型定义
├── gin-gorm.db        # SQLite数据库文件（运行后自动生成）
├── go.mod             # Go模块依赖
├── main.go            # 主程序入口
└── README.md          # 项目说明
```

## 功能特性

- 用户信息的增删改查操作
- RESTful API设计
- SQLite本地数据存储
- 错误处理和参数验证

## 安装运行

1. 确保已安装Go环境（要求Go 1.16+）

2. 获取依赖并运行

```bash
go mod tidy
go run main.go
```

服务器将在 http://localhost:8081 启动

## API接口

### 创建用户

- **URL**: `/api/users/`
- **方法**: `POST`
- **请求体**:
```json
{
  "username": "张三",
  "email": "zhangsan@example.com",
  "age": 28,
  "address": "北京市海淀区"
}
```
- **成功响应**: `201 Created`

### 获取所有用户

- **URL**: `/api/users/`
- **方法**: `GET`
- **成功响应**: `200 OK`

### 获取单个用户

- **URL**: `/api/users/:id`
- **方法**: `GET`
- **成功响应**: `200 OK`

### 更新用户

- **URL**: `/api/users/:id`
- **方法**: `PUT`
- **请求体**:
```json
{
  "username": "李四",
  "email": "lisi@example.com",
  "age": 30,
  "address": "上海市浦东新区"
}
```
- **成功响应**: `200 OK`

### 删除用户

- **URL**: `/api/users/:id`
- **方法**: `DELETE`
- **成功响应**: `200 OK`

## 测试API

可以使用curl、Postman或其他API测试工具进行测试。

示例（使用curl）:

```bash
# 创建用户
curl -X POST http://localhost:8081/api/users/ \
  -H "Content-Type: application/json" \
  -d '{"username":"张三","email":"zhangsan@example.com","age":28,"address":"北京市海淀区"}'

# 获取所有用户
curl -X GET http://localhost:8081/api/users/

# 获取单个用户（ID为1）
curl -X GET http://localhost:8081/api/users/1

# 更新用户（ID为1）
curl -X PUT http://localhost:8081/api/users/1 \
  -H "Content-Type: application/json" \
  -d '{"username":"李四","email":"lisi@example.com","age":30,"address":"上海市浦东新区"}'

# 删除用户（ID为1）
curl -X DELETE http://localhost:8081/api/users/1
```