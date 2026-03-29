# 部署指南

本文档提供 VeloeraGen2 的部署建议，以 Docker Compose 为主。

## 1. 单机部署（Docker Compose）

### 步骤

1. 克隆仓库并进入目录

```bash
git clone https://github.com/moehans-official/VeloeraGen2.git
cd VeloeraGen2
```

2. 准备环境变量

```bash
# 生产环境建议使用 .env 存放敏感配置（如 SQL_DSN/SESSION_SECRET 等）
# cp .env.example .env
```

3. 修改 `docker-compose.yml`（生产环境务必修改默认密码/连接串）

4. 启动

```bash
docker compose up -d
```

5. 检查容器与健康状态

```bash
docker compose ps
docker compose logs -f new-api
curl http://localhost:3000/api/status
```

## 2. 关键配置项

常用环境变量示例：

- `SQL_DSN`：主数据库连接串（MySQL/PostgreSQL）
- `REDIS_CONN_STRING`：Redis 连接串
- `SESSION_SECRET`：会话密钥（生产环境必须自定义）
- `SYNC_FREQUENCY`：数据库同步频率（秒）
- `RELAY_TIMEOUT` / `STREAMING_TIMEOUT`：超时配置（秒）
- `TLS_INSECURE_SKIP_VERIFY`：跳过 TLS 验证（谨慎使用）

## 3. 生产部署建议

- 使用反向代理（Nginx/Caddy）处理 TLS
- 将 `SESSION_SECRET`、数据库密码放入安全密钥管理
- 为数据库与数据目录启用备份策略
- 通过 `/api/status` 接入监控/探针
- 开启 CI，确保每次变更可构建与可回归

## 4. 升级建议

- 升级前先备份数据库与配置
- 先在测试环境验证迁移与启动
- 生产变更采用分批/低峰发布
