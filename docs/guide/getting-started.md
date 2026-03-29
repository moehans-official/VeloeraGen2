# 快速开始

本文用于快速跑通 VeloeraGen2。

## 环境要求

- Go 1.25+
- Node.js 20+
- Bun 1+
- Docker / Docker Compose（推荐）

## 方式一：Docker Compose（推荐）

```bash
git clone https://github.com/moehans-official/VeloeraGen2.git
cd VeloeraGen2
# 生产环境务必修改 docker-compose.yml 内的默认密码/连接串
docker compose up -d
```

验证：

```bash
curl http://localhost:3000/api/status
```

## 方式二：本地开发

```bash
# 前端
cd web
bun install
bun run build

# 后端
cd ..
go run main.go
```

默认服务地址：`http://localhost:3000`
