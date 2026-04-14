# 部署文档索引

[English](README.md) | 简体中文

本目录汇总了 PaoPao 在本地环境、云平台以及 Kubernetes 场景下的部署文档。

如果你需要先了解完整安装流程，建议先阅读：

- [../../INSTALL_ZH.md](../../INSTALL_ZH.md) - 安装与运行配置指南
- [../README_ZH.md](../README_ZH.md) - 文档总索引

## 部署分区

| 分区 | 说明 |
| --- | --- |
| [core/](core/) | 部署基础概念与核心配置说明 |
| [local/](local/) | 本地部署与本地依赖环境说明 |
| [aliyun/](aliyun/) | 阿里云部署文档 |
| [huawei/](huawei/) | 华为云部署文档 |
| [tencent/](tencent/) | 腾讯云部署文档 |
| [k8s/](k8s/) | Kubernetes 部署参考 |

## 推荐阅读顺序

### 本地自部署

1. 先阅读 [../../INSTALL_ZH.md](../../INSTALL_ZH.md)
2. 再查看 [local/README.md](local/README.md)
3. 如果需要更细的配置说明，再进入 [core/](core/)

### 云平台部署

1. 先阅读 [../../INSTALL_ZH.md](../../INSTALL_ZH.md)
2. 再查看 [core/](core/)
3. 最后进入对应平台文档：
   - [aliyun/README.md](aliyun/README.md)
   - [huawei/README.md](huawei/README.md)
   - [tencent/README.md](tencent/README.md)

### Kubernetes 部署

1. 先阅读 [../../INSTALL_ZH.md](../../INSTALL_ZH.md)
2. 再查看 [core/](core/)
3. 最后阅读 [k8s/README.md](k8s/README.md)

## 说明

- 某些平台目录主要只有一个 README 文件，集中说明该平台的部署步骤。
- `local/` 目录中除了说明文档外，还包含本地依赖环境说明与示例配置材料。
- 如果部署方式依赖特定 Feature 组合，请确保基础设施与 `config.yaml` 中启用的 `Features` 保持一致。
