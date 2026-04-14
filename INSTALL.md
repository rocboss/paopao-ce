# Installation Guide

English | [简体中文](INSTALL_ZH.md)

This guide covers the recommended ways to run PaoPao in development, evaluation, and self-hosted deployments. For a project overview, see [README.md](README.md).

## Choose an Installation Path

| Scenario | Recommended path |
| --- | --- |
| Quick local evaluation | [Docker Compose](#docker-compose) |
| Container-based deployment | [Docker images](#docker-images) |
| Backend or frontend development | [Run from source](#run-from-source) |
| Desktop client build | [Desktop app](#desktop-app) |

## Requirements

### For source-based development

- Go `1.24+`
- Node.js `20.19+` or `22.12+`
- Yarn `1.x`
- MySQL `5.7+` if using the MySQL path
- Redis
- Meilisearch
- Rust and the platform prerequisites required by Tauri if you want to build the desktop client

### Helpful repository files

- `config.yaml.sample` - canonical configuration template
- `scripts/paopao-mysql.sql` - MySQL bootstrap schema
- `scripts/paopao-postgres.sql` - PostgreSQL bootstrap schema
- `scripts/paopao-sqlite3.sql` - SQLite bootstrap schema
- `docker-compose.yaml` - local multi-service environment

<a id="docker-compose"></a>

## Option 1: Docker Compose (recommended for quick evaluation)

This is the fastest way to start a local environment with the main dependencies wired together.

```sh
git clone https://github.com/rocboss/paopao-ce.git
cd paopao-ce
docker compose up -d
```

The default compose setup starts these services:

- `http://localhost:8008` - PaoPao application
- `http://localhost:7700` - Meilisearch
- `http://localhost:8001` - RedisInsight
- `http://localhost:3306` - MySQL

Notes:

- The backend container mounts `./config.yaml.sample` as its runtime config by default.
- Persistent data is stored under `./custom/`.
- Optional services such as MinIO, OpenObserve, Pyroscope, and phpMyAdmin are present in `docker-compose.yaml` but commented out by default.

If you want to use a custom config file, replace the mounted file in `docker-compose.yaml`:

```yaml
backend:
  volumes:
    - ./config.yaml:/app/paopao-ce/config.yaml
    - ./custom:/app/paopao-ce/custom
```

<a id="docker-images"></a>

## Option 2: Docker Images

### Backend image

```sh
# Default build: embeds the web UI and uses the default API host behavior
docker build -t your/paopao-ce:tag .

# Embed the web UI and set a custom API host
docker build -t your/paopao-ce:tag --build-arg API_HOST=http://api.paopao.info .

# Embed the web UI and keep the API host from local web/.env
docker build -t your/paopao-ce:tag --build-arg USE_API_HOST=no .

# Build with a precompiled local web/dist
docker build -t your/paopao-ce:tag --build-arg USE_DIST=yes .

# Build backend only, without embedded web UI
docker build -t your/paopao-ce:tag --build-arg EMBED_UI=no .
```

Run a locally built image:

```sh
mkdir -p custom
docker run -d -p 8008:8008 \
  -v ${PWD}/custom:/app/paopao-ce/custom \
  -v ${PWD}/config.yaml.sample:/app/paopao-ce/config.yaml \
  your/paopao-ce:tag
```

Or use the published image:

```sh
mkdir -p custom
docker run -d -p 8008:8008 \
  -v ${PWD}/custom:/app/paopao-ce/custom \
  -v ${PWD}/config.yaml.sample:/app/paopao-ce/config.yaml \
  bitbus/paopao-ce:latest
```

### Web image

```sh
cd web

# Default build
docker build -t your/paopao-ce:web .

# Build with a custom API host
docker build -t your/paopao-ce:web --build-arg API_HOST=http://api.paopao.info .

# Build with a precompiled local dist
docker build -t your/paopao-ce:web --build-arg USE_DIST=yes .

# Run
docker run -d -p 8010:80 your/paopao-ce:web
```

### All-in-one image

```sh
# Build
docker buildx build --build-arg USE_DIST=yes -t your/paopao-ce:all-in-one-latest -f Dockerfile.allinone .

# Run a local image
docker run --name paopao-ce-allinone -d -p 8000:8008 -p 7700:7700 \
  -v ./data/custom:/app/custom \
  -v ./data/meili_data:/app/meili_data \
  your/paopao-ce:all-in-one-latest

# Run the published image
docker run --name paopao-ce-allinone -d -p 8000:8008 -p 7700:7700 \
  -v ./data/custom:/app/custom \
  -v ./data/meili_data:/app/meili_data \
  bitbus/paopao-ce:all-in-one-latest
```

If you mount a custom `config.yaml`, make sure `Meili.ApiKey` matches the container's `MEILI_MASTER_KEY`. The default key is `paopao-meilisearch`.

<a id="run-from-source"></a>

## Option 3: Run from Source

### Backend

1. Initialize your database with the matching SQL file for your chosen database engine.
2. Copy the configuration template.
3. Adjust only the bootstrap-critical settings for your environment.
4. Run or build the backend.

```sh
cp config.yaml.sample config.yaml
make run
```

Build a release binary:

```sh
make build
```

The binary is written to `release/paopao`.

### Web frontend

```sh
cd web
cp .env .env.local
yarn
yarn dev
```

Build the web bundle:

```sh
yarn build
```

### Embedded web UI

If you want the Go service to serve the web frontend directly, build the frontend assets first and then run with the `embed` tag:

```sh
make build-web
make run TAGS='embed'
```

<a id="desktop-app"></a>

### Desktop app

The desktop client is built from `web/` using Tauri:

```sh
cd web
cp .env .env.local
yarn
yarn build
yarn tauri build
```

Before running the Tauri build, install the platform prerequisites from the official Tauri documentation for your OS.

## Common Build Tags

| Tag | Purpose |
| --- | --- |
| `embed` | Serve the built web frontend from the Go binary |
| `migration` | Include migration support in the backend binary |
| `docs` | Enable the developer docs / OpenAPI service |

Examples:

```sh
make build TAGS='migration'
make run TAGS='docs'
make run TAGS='embed'
```

## Configuration Basics

At startup, PaoPao reads either:

1. `./custom/config.yaml`
2. `./config.yaml`

The first file found is used.

Important: the external file is no longer expected to carry every runtime knob. PaoPao loads embedded defaults first, then overlays your local config file.

Recommended split:

- **Bootstrap YAML**: ports, feature selection, database, Redis, JWT, `AdminSettings.EncryptionKey`
- **Admin UI (`/#/admin/settings`)**: most site, search, storage, SMS, payment, and app-behavior settings

If a setting is marked **restart required** in the admin page, it is persisted immediately but only becomes active after a process restart.

The `Features` section controls which capability bundles are enabled:

```yaml
Features:
  Default: ["Web", "Frontend:EmbedWeb", "Meili", "LocalOSS", "MySQL", "BigCacheIndex", "LoggerFile"]
  Develop: ["Base", "MySQL", "BigCacheIndex", "Meili", "Sms", "AliOSS", "LoggerMeili", "OSS:Retention"]
  Demo: ["Base", "MySQL", "Option", "Zinc", "Sms", "MinIO", "LoggerZinc", "Migration"]
  Slim: ["Base", "Sqlite3", "LocalOSS", "LoggerFile", "OSS:TempDir"]
```

Useful commands:

```sh
# Use the default suite
release/paopao serve

# Use only the develop suite
release/paopao serve --no-default-features --features develop

# Add one extra feature on top of default
release/paopao serve --features sms

# Specify features explicitly
release/paopao serve --no-default-features --features sqlite3,localoss,loggerfile,redis
```

For feature maturity and support status, see [features-status.md](features-status.md).

## Optional Infrastructure Services

The default modern stack is centered on **Meilisearch**, **Redis**, and either **LocalOSS**, **MinIO**, or a cloud object store. Optional integrations can be started separately when needed.

### Meilisearch (recommended search engine)

```sh
mkdir -p data/meili/data
docker run -d --name meili \
  -v ${PWD}/data/meili/data:/meili_data \
  -p 7700:7700 \
  -e MEILI_MASTER_KEY=paopao-meilisearch \
  getmeili/meilisearch:v0.29.0
```

Matching config example:

```yaml
Meili:
  Host: 127.0.0.1:7700
  Index: paopao-data
  ApiKey: paopao-meilisearch
  Secure: False
```

### MinIO

```sh
mkdir -p data/minio/data
docker run -d --name minio \
  -v ${PWD}/data/minio/data:/data \
  -p 9000:9000 -p 9001:9001 \
  -e MINIO_ROOT_USER=minio-root-user \
  -e MINIO_ROOT_PASSWORD=minio-root-password \
  -e MINIO_DEFAULT_BUCKETS=paopao:public \
  bitnami/minio:latest
```

Matching config example:

```yaml
MinIO:
  AccessKey: Q3AM3UQ867SPQQA43P2F
  SecretKey: zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG
  Secure: False
  Endpoint: 127.0.0.1:9000
  Bucket: paopao
  Domain: 127.0.0.1:9000
```

### OpenObserve

```sh
mkdir -p data/openobserve
docker run -v ${PWD}/data/openobserve:/data \
  -e ZO_DATA_DIR=/data \
  -p 5080:5080 \
  -e ZO_ROOT_USER_EMAIL=root@paopao.info \
  -e ZO_ROOT_USER_PASSWORD=paopao-ce \
  public.ecr.aws/zinclabs/openobserve:latest
```

### Pyroscope

```sh
docker run -it -p 4040:4040 pyroscope/pyroscope:latest server
```

### Zinc (legacy / optional)

Zinc still appears in the repository and feature definitions, but the current default stack is Meilisearch-based. Use it only if you intentionally want the legacy search path.

## Enable API Documentation Locally

Add the Docs feature suite and run with the `docs` build tag:

```yaml
Features:
  Default: ["Base", "MySQL", "Option", "LocalOSS", "LoggerFile", "Docs"]
  Docs: ["Docs:OpenAPI"]
```

```sh
make run TAGS='docs'
```

Then visit:

- `http://127.0.0.1:8011/docs/openapi`

## Additional Deployment Docs

For platform-specific or production-oriented deployment references, see:

- [docs/deploy/README.md](docs/deploy/README.md)
- [docs/deploy/core/](docs/deploy/core/)
- [docs/deploy/local/](docs/deploy/local/)
- [docs/deploy/k8s/](docs/deploy/k8s/)
- [docs/deploy/aliyun/](docs/deploy/aliyun/)
- [docs/deploy/huawei/](docs/deploy/huawei/)
- [docs/deploy/tencent/](docs/deploy/tencent/)

## Operational Notes

- For long-running deployments, it is reasonable to run the backend under a process manager and place Nginx in front of the application.
- The SMS implementation currently references Juhe in the sample configuration. If that provider is not suitable for your deployment, replace it with another compatible service.
- The repository includes multiple runtime combinations; keep your selected `Features` set aligned with the infrastructure you actually provision.
