# 安装指南

[English](INSTALL.md) | 简体中文

本文档介绍 PaoPao 在本地体验、开发调试和自部署场景下的推荐安装方式。项目整体说明请参考 [README_ZH.md](README_ZH.md)。

## 选择安装方式

| 场景 | 推荐方式 |
| --- | --- |
| 快速本地体验 | [Docker Compose](#docker-compose) |
| 基于容器部署 | [Docker 镜像](#docker-images) |
| 后端或前端开发 | [源码运行](#run-from-source) |
| 构建桌面端 | [桌面应用](#desktop-app) |

## 环境要求

### 源码开发所需环境

- Go `1.24+`
- Node.js `20.19+` 或 `22.12+`
- Yarn `1.x`
- 若使用 MySQL 方案，需要 MySQL `5.7+`
- Redis
- Meilisearch
- 如果需要构建桌面端，还需要安装 Rust 以及 Tauri 对应平台依赖

### 关键文件

- `config.yaml.sample` - 标准配置模板
- `scripts/paopao-mysql.sql` - MySQL 初始化脚本
- `scripts/paopao-postgres.sql` - PostgreSQL 初始化脚本
- `scripts/paopao-sqlite3.sql` - SQLite 初始化脚本
- `docker-compose.yaml` - 本地多服务启动配置

<a id="docker-compose"></a>

## 方案 1：Docker Compose（推荐用于快速体验）

这是最快速的本地启动方式，适合先把主流程跑起来。

```sh
git clone https://github.com/rocboss/paopao-ce.git
cd paopao-ce
docker compose up -d
```

默认会启动以下服务：

- `http://localhost:8008` - PaoPao 应用
- `http://localhost:7700` - Meilisearch
- `http://localhost:8001` - RedisInsight
- `http://localhost:3306` - MySQL

说明：

- 后端容器默认会将 `./config.yaml.sample` 挂载为运行配置。
- 持久化数据默认保存在 `./custom/` 目录下。
- `docker-compose.yaml` 中还预留了 MinIO、OpenObserve、Pyroscope、phpMyAdmin 等可选服务，但默认是注释状态。

如果需要使用自定义配置文件，可将 `docker-compose.yaml` 中的挂载改为：

```yaml
backend:
  volumes:
    - ./config.yaml:/app/paopao-ce/config.yaml
    - ./custom:/app/paopao-ce/custom
```

<a id="docker-images"></a>

## 方案 2：Docker 镜像

### 后端镜像

```sh
# 默认构建：内嵌 Web UI，并使用默认 API host 逻辑
docker build -t your/paopao-ce:tag .

# 内嵌 Web UI，并指定 API host
docker build -t your/paopao-ce:tag --build-arg API_HOST=http://api.paopao.info .

# 内嵌 Web UI，并沿用本地 web/.env 中的 API host
docker build -t your/paopao-ce:tag --build-arg USE_API_HOST=no .

# 使用本地预编译的 web/dist 构建
docker build -t your/paopao-ce:tag --build-arg USE_DIST=yes .

# 仅构建后端，不内嵌 Web UI
docker build -t your/paopao-ce:tag --build-arg EMBED_UI=no .
```

运行本地构建镜像：

```sh
mkdir -p custom
docker run -d -p 8008:8008 \
  -v ${PWD}/custom:/app/paopao-ce/custom \
  -v ${PWD}/config.yaml.sample:/app/paopao-ce/config.yaml \
  your/paopao-ce:tag
```

或者直接使用已发布镜像：

```sh
mkdir -p custom
docker run -d -p 8008:8008 \
  -v ${PWD}/custom:/app/paopao-ce/custom \
  -v ${PWD}/config.yaml.sample:/app/paopao-ce/config.yaml \
  bitbus/paopao-ce:latest
```

### Web 镜像

```sh
cd web

# 默认构建
docker build -t your/paopao-ce:web .

# 自定义 API host
docker build -t your/paopao-ce:web --build-arg API_HOST=http://api.paopao.info .

# 使用本地预编译 dist 构建
docker build -t your/paopao-ce:web --build-arg USE_DIST=yes .

# 运行
docker run -d -p 8010:80 your/paopao-ce:web
```

### All-in-one 镜像

```sh
# 构建
docker buildx build --build-arg USE_DIST=yes -t your/paopao-ce:all-in-one-latest -f Dockerfile.allinone .

# 运行本地镜像
docker run --name paopao-ce-allinone -d -p 8000:8008 -p 7700:7700 \
  -v ./data/custom:/app/custom \
  -v ./data/meili_data:/app/meili_data \
  your/paopao-ce:all-in-one-latest

# 运行已发布镜像
docker run --name paopao-ce-allinone -d -p 8000:8008 -p 7700:7700 \
  -v ./data/custom:/app/custom \
  -v ./data/meili_data:/app/meili_data \
  bitbus/paopao-ce:all-in-one-latest
```

如果挂载了自定义 `config.yaml`，请确保其中的 `Meili.ApiKey` 与容器内的 `MEILI_MASTER_KEY` 保持一致。默认值为 `paopao-meilisearch`。

<a id="run-from-source"></a>

## 方案 3：从源码运行

### 后端

1. 按照所选数据库导入对应 SQL 初始化脚本。
2. 复制配置模板。
3. 只调整与你环境相关的启动关键配置。
4. 启动或构建后端。

```sh
cp config.yaml.sample config.yaml
make run
```

构建发布二进制：

```sh
make build
```

生成的二进制默认位于 `release/paopao`。

### Web 前端

```sh
cd web
cp .env .env.local
yarn
yarn dev
```

构建 Web 静态资源：

```sh
yarn build
```

### 内嵌 Web UI

如果希望由 Go 服务直接提供 Web 前端，请先构建前端资源，再使用 `embed` 标签运行：

```sh
make build-web
make run TAGS='embed'
```

<a id="desktop-app"></a>

### 桌面应用

桌面端位于 `web/` 目录，使用 Tauri 构建：

```sh
cd web
cp .env .env.local
yarn
yarn build
yarn tauri build
```

执行 Tauri 构建前，请先安装对应操作系统的官方前置依赖。

## 常用构建标签

| 标签 | 作用 |
| --- | --- |
| `embed` | 将 Web 前端打包进 Go 二进制，由后端直接提供服务 |
| `migration` | 在后端二进制中包含 migration 支持 |
| `docs` | 启用开发文档 / OpenAPI 服务 |

示例：

```sh
make build TAGS='migration'
make run TAGS='docs'
make run TAGS='embed'
```

## 配置基础

启动时，PaoPao 会按以下顺序读取配置：

1. `./custom/config.yaml`
2. `./config.yaml`

先找到哪个文件，就使用哪个文件。

注意：外部配置文件现在不再要求承载全部运行参数。PaoPao 会先加载内置默认配置，再叠加你本地的配置文件。

建议的职责划分：

- **Bootstrap YAML**：端口、Feature 组合、数据库、Redis、JWT、`AdminSettings.EncryptionKey`
- **管理后台**（`/#/admin/settings`）：大部分站点、搜索、存储、短信、支付以及应用行为类配置

如果某个配置项在后台中被标记为**重启后生效**，表示它会先持久化保存，但需要重启进程后才会真正切换到新值。

`Features` 用于控制不同能力组合：

```yaml
Features:
  Default: ["Web", "Frontend:EmbedWeb", "Meili", "LocalOSS", "MySQL", "BigCacheIndex", "LoggerFile"]
  Develop: ["Base", "MySQL", "BigCacheIndex", "Meili", "Sms", "AliOSS", "LoggerMeili", "OSS:Retention"]
  Demo: ["Base", "MySQL", "Option", "Zinc", "Sms", "MinIO", "LoggerZinc", "Migration"]
  Slim: ["Base", "Sqlite3", "LocalOSS", "LoggerFile", "OSS:TempDir"]
```

常见命令：

```sh
# 使用默认套件
release/paopao serve

# 仅使用 develop 套件
release/paopao serve --no-default-features --features develop

# 在默认套件基础上增加一个功能
release/paopao serve --features sms

# 手动显式指定功能项
release/paopao serve --no-default-features --features sqlite3,localoss,loggerfile,redis
```

功能项成熟度与支持状态请参考 [features-status.md](features-status.md)。

## 可选基础设施服务

当前更推荐的默认组合是 **Meilisearch**、**Redis**，以及 **LocalOSS / MinIO / 云对象存储** 三选一。其他集成按需启用即可。

### Meilisearch（推荐搜索引擎）

```sh
mkdir -p data/meili/data
docker run -d --name meili \
  -v ${PWD}/data/meili/data:/meili_data \
  -p 7700:7700 \
  -e MEILI_MASTER_KEY=paopao-meilisearch \
  getmeili/meilisearch:v0.29.0
```

对应配置示例：

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

对应配置示例：

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

### Zinc（遗留 / 可选）

仓库中仍保留了 Zinc 相关代码与 Feature 定义，但当前默认推荐的搜索方案是 Meilisearch。只有在你明确需要兼容旧方案时，再考虑启用 Zinc。

## 本地启用 API 文档

在配置中加入 Docs 套件，并使用 `docs` 标签运行：

```yaml
Features:
  Default: ["Base", "MySQL", "Option", "LocalOSS", "LoggerFile", "Docs"]
  Docs: ["Docs:OpenAPI"]
```

```sh
make run TAGS='docs'
```

然后访问：

- `http://127.0.0.1:8011/docs/openapi`

## 更多部署文档

如果需要平台化或生产化部署参考，请继续阅读：

- [docs/deploy/README.md](docs/deploy/README.md)
- [docs/deploy/core/](docs/deploy/core/)
- [docs/deploy/local/](docs/deploy/local/)
- [docs/deploy/k8s/](docs/deploy/k8s/)
- [docs/deploy/aliyun/](docs/deploy/aliyun/)
- [docs/deploy/huawei/](docs/deploy/huawei/)
- [docs/deploy/tencent/](docs/deploy/tencent/)

## 运维建议

- 对于长期运行环境，建议使用进程守护工具管理后端服务，并通过 Nginx 做反向代理。
- 示例配置中的短信通道使用 Juhe；如果不适合你的部署场景，可以替换为其他兼容服务商。
- 项目支持多种运行组合，请确保 `Features` 与你实际部署的基础设施保持一致。
