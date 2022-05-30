<div id="top"></div>


<!-- PROJECT SHIELDS -->
[![Go Report Card][goreport-shield]][goreport-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![MIT License][license-shield]][license-url]
[![Contributors][contributors-shield]][contributors-url]
[![Issues][issues-shield]][issues-url]

<!-- PROJECT LOGO -->
<div align="center">
  <a href="https://github.com/rocboss/paopao-ce">
    <img src="https://cdn.rocs.me/static/paopao-logo.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">PaoPao</h3>

  <p align="center">
    🔥一个清新文艺的微社区
    <br />
    <a href="https://www.paopao.info/">View Demo</a>
    ·
    <a href="https://github.com/rocboss/paopao-ce/issues">Report Bug</a>
    ·
    <a href="https://www.rocs.me/archives/paopao.html">Features</a>
  </p>
</div>



## 1. 截图预览

[![明色主题][product-light-screenshot]](https://www.paopao.info)

[![暗色主题][product-dark-screenshot]](https://www.paopao.info)

更多演示请前往[官网](https://www.paopao.info)体验（谢绝灌水）

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- GETTING STARTED -->
## 2. 快速开始

### 2.1 技术栈

PaoPao主要由以下优秀的开源项目/工具构建

* [Go](https://go.dev/)
* [Gin](https://gin-gonic.com/)
* [Zinc](https://zinclabs.io/)
* [Naive UI](https://www.naiveui.com/)
* [Vue.js](https://vuejs.org/)
* [Vite.js](https://vitejs.dev/)


### 2.2 环境要求

- Go (1.17+)
- Node.js (14+)
- MySQL (5.7+)
- Redis
- Zinc

\* Zinc是一款轻量级全文搜索引擎，可以查阅 https://zincsearch.com/ 安装

以上环境版本为PaoPao官方的开发版本，仅供参考，其他版本的环境未进行充分测试

<br />

### 3. 安装说明

***宝塔安装***

我们为宝塔用户提供了超详细安装教程 [点此查看](https://www.rocs.me/archives/paopao_bt_install.html)

***普通安装***

克隆代码库
   ```sh
   git clone https://github.com/rocboss/paopao-ce.git
   ```
#### 3.1 后端

1. 导入项目根目录下的 `paopao.sql` 文件至MySQL数据库
2. 拷贝项目根目录下 `config.yaml.sample` 文件至 `config.yaml`，按照注释完成配置编辑
3. 编译后端
    ```sh
    go mod download
    go build .
    ```
4. 启动后端
    ```sh
    chmod +x paopao-api
    ./paopao-api
    ```

#### 3.2 前端

1. 进入前端目录 `web`，编辑 `.env` 文件中后端服务地址，下载依赖包
    ```sh
    cd ./web
    vim .env
    yarn
    ```
2. 编译前端
    ```sh
    yarn build
    ```

    build完成后，可以在dist目录获取编译产出，配置nginx指向至该目录即可

#### 3.3 其他
建议后端服务使用 `supervisor` 守护进程，并通过nginx反向代理后，提供API给前端服务调用。

短信通道使用的聚合数据(juhe.cn)，如果申请不下来，可以考虑替换其他服务商。

代码结构比较简单，很方便扩展，喜欢的朋友欢迎给个Star、贡献PR。

<!-- MARKDOWN LINKS & IMAGES -->
[contributors-shield]: https://img.shields.io/github/contributors/rocboss/paopao-ce?style=flat
[contributors-url]: https://github.com/rocboss/paopao-ce/graphs/contributors
[goreport-shield]: https://goreportcard.com/badge/github.com/rocboss/paopao-ce
[goreport-url]: https://goreportcard.com/report/github.com/rocboss/paopao-ce
[forks-shield]: https://img.shields.io/github/forks/rocboss/paopao-ce?style=flat
[forks-url]: https://github.com/rocboss/paopao-ce/network/members
[stars-shield]: https://img.shields.io/github/stars/rocboss/paopao-ce.svg?style=flat
[stars-url]: https://github.com/rocboss/paopao-ce/stargazers
[issues-shield]: https://img.shields.io/github/issues/rocboss/paopao-ce.svg?style=flat
[issues-url]: https://github.com/rocboss/paopao-ce/issues
[license-shield]: https://img.shields.io/github/license/rocboss/paopao-ce.svg?style=flat
[license-url]: https://github.com/rocboss/paopao-ce/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat&logo=linkedin&colorB=555
[product-light-screenshot]: https://cdn.rocs.me/static/paopao-light.jpeg
[product-dark-screenshot]: https://cdn.rocs.me/static/paopao-dark.jpeg