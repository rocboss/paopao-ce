# Deployment Documentation

English | [简体中文](README_ZH.md)

This directory collects deployment-oriented documentation for PaoPao across local environments, cloud platforms, and Kubernetes-based setups.

For the higher-level setup flow, start with:

- [../../INSTALL.md](../../INSTALL.md) - installation and runtime setup guide
- [../README.md](../README.md) - overall documentation index

## Deployment Sections

| Section | Description |
| --- | --- |
| [core/](core/) | Core deployment concepts and configuration references |
| [local/](local/) | Local deployment and local dependency setup notes |
| [aliyun/](aliyun/) | Alibaba Cloud deployment documentation |
| [huawei/](huawei/) | Huawei Cloud deployment documentation |
| [tencent/](tencent/) | Tencent Cloud deployment documentation |
| [k8s/](k8s/) | Kubernetes deployment references |

## Suggested Reading Order

### For local self-hosting

1. Read [../../INSTALL.md](../../INSTALL.md)
2. Continue with [local/README.md](local/README.md)
3. Use [core/](core/) if you need additional configuration detail

### For cloud deployment

1. Read [../../INSTALL.md](../../INSTALL.md)
2. Review [core/](core/)
3. Open the provider-specific guide for your target platform:
   - [aliyun/README.md](aliyun/README.md)
   - [huawei/README.md](huawei/README.md)
   - [tencent/README.md](tencent/README.md)

### For Kubernetes-based deployment

1. Read [../../INSTALL.md](../../INSTALL.md)
2. Review [core/](core/)
3. Continue with [k8s/README.md](k8s/README.md)

## Notes

- Some platform directories primarily contain a single README with environment-specific instructions.
- `local/` includes additional assets such as local dependency notes and example config material.
- When deployment behavior depends on enabled runtime features, align your infrastructure choices with `config.yaml` and the selected `Features` suite.
