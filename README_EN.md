<div align="center">
  <img src="https://cloudflareimg.cdn.sn/i/695929a70c0bc_1767451047.webp" alt="Veloera Gen2 Logo" width="180" />
  <h1>VeloeraGen2</h1>
  <p>Community-maintained Veloera, optimized for production operations</p>
  <p>
    <a href="./README.md">简体中文</a> | English
  </p>
  <p>
    <a href="./LICENSE"><img alt="License: AGPLv3" src="https://img.shields.io/badge/License-AGPLv3-blue.svg" /></a>
    <img alt="Go" src="https://img.shields.io/badge/Go-1.25+-00ADD8?logo=go" />
    <img alt="Node" src="https://img.shields.io/badge/Node-20+-339933?logo=node.js" />
    <img alt="Docker" src="https://img.shields.io/badge/Docker-Compose-2496ED?logo=docker" />
  </p>
</div>

## Overview

`VeloeraGen2 (Community Edition)` is a community-driven LLM API gateway evolved from `Veloera`.

Project goals:

- Sustain and maintain core upstream capabilities
- Improve reliability and operability in production
- Build a transparent, contributor-friendly community workflow

Upstream projects:

- Veloera: <https://github.com/Veloera/Veloera>
- new-api: <https://github.com/QuantumNous/new-api>

## Core Features

- Unified OpenAI-compatible API layer
- Multi-channel model routing
- Web admin console (channels, users, tokens, logs, billing)
- One-command Docker Compose deployment
- Open contribution and review process

## Tech Stack

- Backend: Go + Gin + GORM
- Frontend: React + Vite + Semi UI
- Database: MySQL / PostgreSQL / SQLite (configurable)
- Cache: Redis (optional)

## Quick Start

### Option 1: Docker Compose (Recommended)

1. Clone repository

```bash
git clone https://github.com/moehans-official/VeloeraGen2.git
cd VeloeraGen2
```

2. Edit `docker-compose.yml` (change default passwords/DSN for production)

3. Start services

```bash
docker compose up -d
```

4. Verify status

```bash
curl http://localhost:3000/api/status
```

Default URL: `http://localhost:3000`

### Option 2: Local Development

1. Build frontend assets (embedded by backend)

```bash
cd web
bun install
bun run build
cd ..
```

2. Start backend

```bash
go run main.go
```

3. Run frontend dev server (optional)

```bash
cd web
bun run dev
```

## Ops and Health Probes

- Status: `GET /api/status`

Common env vars (partial list, see `.env.example`):

- `SQL_DSN` / `LOG_SQL_DSN` / `SQLITE_PATH`
- `REDIS_CONN_STRING`
- `SESSION_SECRET` (must be customized in production)
- `SYNC_FREQUENCY`
- `RELAY_TIMEOUT` / `STREAMING_TIMEOUT`
- `TLS_INSECURE_SKIP_VERIFY`

## License

Licensed under [`GNU AGPLv3`](./LICENSE).
