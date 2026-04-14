<div id="top"></div>

[![Go](https://github.com/rocboss/paopao-ce/actions/workflows/go.yml/badge.svg)](https://github.com/rocboss/paopao-ce/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/rocboss/paopao-ce)](https://goreportcard.com/report/github.com/rocboss/paopao-ce)
[![Forks](https://img.shields.io/github/forks/rocboss/paopao-ce?style=flat)](https://github.com/rocboss/paopao-ce/network/members)
[![Stars](https://img.shields.io/github/stars/rocboss/paopao-ce.svg?style=flat)](https://github.com/rocboss/paopao-ce/stargazers)
[![MIT License](https://img.shields.io/github/license/rocboss/paopao-ce.svg?style=flat)](https://github.com/rocboss/paopao-ce/blob/main/LICENSE)
[![Contributors](https://img.shields.io/github/contributors/rocboss/paopao-ce?style=flat)](https://github.com/rocboss/paopao-ce/graphs/contributors)
[![Sourcegraph](https://img.shields.io/badge/view%20on-Sourcegraph-brightgreen.svg)](https://sourcegraph.com/github.com/rocboss/paopao-ce)

<div align="center">
  <a href="https://github.com/rocboss/paopao-ce">
    <img src="https://assets.paopao.info/static/paopao-logo.png" alt="PaoPao logo" width="88" height="88">
  </a>

  <h1 align="center">PaoPao</h1>

  <p align="center">
    一个基于 Go 与 Vue 的开源微社区平台。
    <br />
    适合自部署社区产品、功能验证，以及在现有能力之上进行二次开发。
  </p>

  <p align="center">
    <a href="README.md">English</a>
    ·
    <a href="https://www.paopao.info/">在线演示</a>
    ·
    <a href="https://github.com/rocboss/paopao-ce/pulls">Pull Requests</a>
    ·
    <a href="https://www.yuque.com/rocs/paopao/about">项目笔记</a>
  </p>
</div>

---

## 项目简介

PaoPao 是一个完整的开源微社区系统，包含 Go 后端、Vue 3 Web 前端，以及可选的 Tauri 桌面端。项目通过模块化的特性开关体系，将存储、搜索、日志、可观测性与部署方式组合在一起，便于按需裁剪和扩展。

如果你希望快速搭建一个可运行的社区产品，或者基于成熟代码库继续定制开发，而不是从零开始，PaoPao 是一个很合适的起点。

## 为什么选择 PaoPao

- **全栈一体化**：后端、Web 前端、桌面端构建都在同一仓库中维护。
- **模块化运行能力**：通过 `Features` 套件启用不同能力组合，例如 `Default`、`Develop`、`Demo`、`Slim`。
- **基础设施灵活**：支持 MySQL、PostgreSQL、SQLite、Redis、Meilisearch 以及多种对象存储后端。
- **部署方式丰富**：可通过源码、Docker、Docker Compose 或 all-in-one 镜像运行。
- **适合自部署**：配置文件清晰，仓库内已提供部署与开发文档。

## 预览

### Web 端

[![明色主题预览](https://assets.paopao.info/static/paopao-light.jpeg)](https://www.paopao.info)

[![暗色主题预览](https://assets.paopao.info/static/paopao-dark.jpeg)](https://www.paopao.info)

### 桌面端

![](docs/proposal/.assets/000-00.jpg)

更多效果可前往 [paopao.info](https://www.paopao.info/) 查看。

## 架构概览

| 层级 | 主要技术 |
| --- | --- |
| 后端 | Go、Gin、Cobra、GORM、Mir |
| Web 前端 | Vue 3、Vite、Naive UI |
| 桌面端 | Tauri |
| 搜索 | Meilisearch |
| 缓存 | Redis |
| 对象存储 | Local OSS、MinIO、AliOSS、COS、Huawei OBS、S3 兼容存储 |
| 可观测性 | OpenTelemetry、Sentry、Pyroscope、Pprof |

## 仓库结构

| 路径 | 说明 |
| --- | --- |
| `cmd/`、`internal/`、`pkg/` | 后端应用与共享包 |
| `web/` | Vue 3 Web 应用与 Tauri 桌面端前端 |
| `docs/` | 部署文档、OpenAPI 文档、设计提案等 |
| `scripts/` | SQL 初始化脚本与辅助资源 |
| `config.yaml.sample` | 完整的运行配置模板 |

## 快速开始

### 方案 A：使用 Docker Compose 快速体验

这是本地体验项目的最快方式。

```sh
git clone https://github.com/rocboss/paopao-ce.git
cd paopao-ce
docker compose up -d
```

启动后可访问：

- `http://localhost:8008` - PaoPao
- `http://localhost:7700` - Meilisearch
- `http://localhost:8001` - RedisInsight

### 方案 B：从源码开发

#### 环境要求

- Go `1.24+`
- Node.js `20.19+` 或 `22.12+`
- MySQL `5.7+`
- Redis
- Meilisearch
- 如果需要构建桌面端，还需要安装 Rust 与 Tauri 依赖

#### 后端

1. 将 `scripts/paopao-mysql.sql` 导入 MySQL。
2. 复制示例配置，并只调整与你环境相关的启动关键项。
3. 启动后端服务。

```sh
cp config.yaml.sample config.yaml
make run
```

如果需要构建发布二进制：

```sh
make build
```

如果希望由 Go 服务直接内嵌并提供 Web UI，请先构建前端资源，再使用 `embed` 标签启动：

```sh
make build-web
make run TAGS='embed'
```

#### Web 前端

```sh
cd web
cp .env .env.local
yarn
yarn dev
```

构建静态资源：

```sh
yarn build
```

#### 桌面端

```sh
cd web
yarn tauri build
```

更完整的安装步骤、Docker 构建方式、桌面端依赖与 migration 说明，请参考 [INSTALL_ZH.md](INSTALL_ZH.md)。

## 配置与 Feature 套件

`config.yaml.sample` 现在是一个**最小化 bootstrap 模板**。运行时，PaoPao 会先加载内置默认配置，再叠加 `./custom/config.yaml` 或 `./config.yaml`，并优先使用先找到的文件。

这意味着外部配置文件可以有意保持精简。当前建议是：

- **继续放在 YAML 中**：端口、Feature 组合、数据库、Redis、JWT，以及 `AdminSettings.EncryptionKey`
- **尽量改到管理后台**（`/#/admin/settings`）：站点资料、应用行为阈值、全文索引、对象存储、短信/支付等已接入 registry 的配置
- **注意生效方式**：部分配置即时生效，部分会在后台中明确标记为“重启后生效”

`Features` 配置用于控制启用哪些能力组合：

```yaml
Features:
  Default: ["Base", "MySQL", "Option", "LocalOSS", "LoggerFile"]
  Develop: ["Base", "MySQL", "Option", "Sms", "AliOSS", "LoggerOtlp"]
  Demo: ["Base", "MySQL", "Option", "Sms", "MinIO", "LoggerOtlp"]
  Slim: ["Base", "Sqlite3", "LocalOSS", "LoggerFile"]
  Base: ["Zinc", "Redis", "Alipay"]
  Option: ["SimpleCacheIndex"]
  Sms: "SmsJuhe"
```

常见用法示例：

```sh
# 使用默认套件
release/paopao serve

# 仅使用 develop 中声明的功能集
release/paopao serve --no-default-features --features develop

# 在默认套件基础上增加 sms
release/paopao serve --features sms

# 手动显式指定功能项
release/paopao serve --no-default-features --features sqlite3,localoss,loggerfile,redis
```

各功能项的当前实现状态可参考 [features-status.md](features-status.md)。

## 文档导航

- [INSTALL_ZH.md](INSTALL_ZH.md) - 中文安装、Docker 使用与桌面端构建说明
- [INSTALL.md](INSTALL.md) - English installation guide
- [docs/README_ZH.md](docs/README_ZH.md) - 中文文档总索引
- [docs/README.md](docs/README.md) - English documentation index
- [docs/deploy/](docs/deploy/) - 部署相关文档
- [docs/openapi/](docs/openapi/) - 导出的 API 文档资源
- [docs/proposal/](docs/proposal/) - 设计说明与功能提案
- [deployed-sites.md](deployed-sites.md) - 已知部署站点列表
- [ROADMAP.md](ROADMAP.md) - 路线图与规划说明

## 分支策略

项目采用分阶段分支模型：

| 分支 | 作用 |
| --- | --- |
| `main` | 稳定生产分支，主要接收缺陷修复 |
| `beta` | 公测分支 |
| `alpha` | 内测分支 |
| `dev` | 主要开发分支，适合新功能开发 |
| `feature/*` | 独立功能分支 |
| `r/*` | 面向不同发行形态的发布分支 |

如果你准备提交新功能，通常应优先向 **`dev`** 分支发起贡献。

## 参与贡献

欢迎提交 Pull Request。建议的贡献流程：

1. Fork 当前仓库。
2. 从 `dev` 分支拉出自己的功能分支。
3. 保持改动聚焦，并补充必要文档。
4. 提交 PR 时说明问题背景、实现方式与验证结果。

如果你已经部署了自己的实例，也欢迎将站点补充到 [deployed-sites.md](deployed-sites.md)。

[![Star History Chart](https://api.star-history.com/svg?repos=rocboss/paopao-ce&type=Date)](https://star-history.com/#rocboss/paopao-ce&Date)

## License

项目基于 MIT License 发布，详见 [LICENSE](LICENSE)。

<p align="right">(<a href="#top">回到顶部</a>)</p>
