<div align="center">
  <img src="https://cloudflareimg.cdn.sn/i/695929a70c0bc_1767451047.webp" alt="Veloera Gen2 Logo" width="180" />
  <h1>Veloera Gen2</h1>
  <p>Veloera Gen2，面向生产可用</p>
  <p>
    简体中文 | <a href="./README_EN.md">English</a>
  </p>
  <p>
    <a href="./LICENSE"><img alt="License: AGPLv3" src="https://img.shields.io/badge/License-AGPLv3-blue.svg" /></a>
    <img alt="Go" src="https://img.shields.io/badge/Go-1.25+-00ADD8?logo=go" />
    <img alt="Node" src="https://img.shields.io/badge/Node-20+-339933?logo=node.js" />
    <img alt="Docker" src="https://img.shields.io/badge/Docker-Compose-2496ED?logo=docker" />
  </p>
</div>

## 项目简介

`VeloeraGen2` 是基于 `Veloera` 继续演进的LLM API 网关，目标是：

- 持续维护原项目核心能力
- 提升生产环境稳定性与可运维性
- 通过开放协作持续迭代

上游项目：

- Veloera: <https://github.com/Veloera/Veloera>
- new-api: <https://github.com/QuantumNous/new-api>

## 核心能力

- 统一的 OpenAI 兼容接口层
- 多渠道接入与模型路由
- 管理后台（渠道、用户、令牌、日志、计费等）
- Docker Compose 一键部署
- 社区治理与公开 PR 流程

## 技术栈

- 后端：Go + Gin + GORM
- 前端：React + Vite + Semi UI
- 存储：MySQL / PostgreSQL / SQLite（按配置）
- 缓存：Redis（可选）

## 快速开始

### 方式一：Docker Compose（推荐）

1. 克隆仓库

```bash
git clone https://github.com/moehans-official/VeloeraGen2.git
cd VeloeraGen2
```

2. 修改 `docker-compose.yml`（生产环境务必修改默认密码/连接串）

3. 启动服务

```bash
docker compose up -d
```

4. 检查状态

```bash
curl http://localhost:3000/api/status
```

默认访问地址：`http://localhost:3000`

### 方式二：本地开发

1. 前端构建（生产构建用于嵌入后端）

```bash
cd web
bun install
bun run build
cd ..
```

2. 启动后端

```bash
go run main.go
```

3. 前端开发模式（可选）

```bash
cd web
bun run dev
```

## 运维与健康检查

- 健康状态：`GET /api/status`

常用环境变量（节选，详见 `.env.example`）：

- `SQL_DSN` / `LOG_SQL_DSN` / `SQLITE_PATH`
- `REDIS_CONN_STRING`
- `SESSION_SECRET`（生产环境必须自定义）
- `SYNC_FREQUENCY`
- `RELAY_TIMEOUT` / `STREAMING_TIMEOUT`
- `TLS_INSECURE_SKIP_VERIFY`

## License

本项目采用 [`GNU AGPLv3`](./LICENSE) 许可证。
