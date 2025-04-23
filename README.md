# Guide

## How to run Dev

- Create .env file with this config in root directory.

```bash
# Echo app
APP_ENV=development
TIMEZONE=UTC
LOCAL_TIMEZONE=Asia/Bangkok

# Postgres
DATABASE_HOST=postgres-payment
DATABASE_USER=postgres
DATABASE_PASSWORD=postgres
DATABASE_NAME=payment_db
DATABASE_HOST_PORT=5436
DATABASE_DOCKER_PORT=5432

# NATS
NATS_URL=nats://nats:4222 # Docker service name:port
NATS_TOKEN=#platong1234

# Docker
COMPOSE_PROJECT_NAME=demo-payment-service
APP_BUILD_CONTEXT=../../
```

- For first time.

```bash
cd scripts && chmod +x dev-start.sh && ./dev-start.sh
```

- Later

```bash
cd scripts && ./dev-start.sh
```

## Database CLI

```bash
pgcli postgres://postgres:postgres@127.0.0.1:5436/payment_db
```
