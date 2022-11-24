<div id="top"></div>

<!-- PROJECT SHIELDS -->
[![Go Report Card][goreport-shield]][goreport-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![MIT License][license-shield]][license-url]
[![Contributors][contributors-shield]][contributors-url]

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
    <a href="https://github.com/rocboss/paopao-ce/pulls">Pull Request</a>
    ·
    <a href="https://www.yuque.com/rocs/paopao/about">Features</a>
  </p>
</div>

---

## 预览
Web端：  
[![明色主题][product-light-screenshot]](https://www.paopao.info)

[![暗色主题][product-dark-screenshot]](https://www.paopao.info)

更多演示请前往[官网](https://www.paopao.info)体验（谢绝灌水）  

桌面端：  
![](docs/proposal/.assets/00-00.png)

<p align="right">(<a href="#top">back to top</a>)</p>

## 🛠 技术栈

PaoPao主要由以下优秀的开源项目/工具构建

* [Go](https://go.dev/)
* [Gin](https://gin-gonic.com/)
* [Zinc](https://zinclabs.io/)
* [Naive UI](https://www.naiveui.com/)
* [Vue.js](https://vuejs.org/)
* [Vite.js](https://vitejs.dev/)

<!-- GETTING STARTED -->
## 🏗 快速开始

### 环境要求

* Go (1.17+)
* Node.js (14+)
* MySQL (5.7+)
* Redis
* Zinc

\* Zinc是一款轻量级全文搜索引擎，可以查阅 <https://zincsearch.com/> 安装

以上环境版本为PaoPao官方的开发版本，仅供参考，其他版本的环境未进行充分测试

### 安装说明


### 方式一. 手动安装（推荐）

克隆代码库

   ```sh
   git clone https://github.com/rocboss/paopao-ce.git
   ```

#### 后端

1. 导入项目根目录下的 `scripts/paopao.sql` 文件至MySQL数据库
2. 拷贝项目根目录下 `config.yaml.sample` 文件至 `config.yaml`，按照注释完成配置编辑
3. 编译后端    
    编译api服务:
    ```sh
    make build
    ```
    编译api服务、内嵌web前端ui:
    ```sh
    make build TAGS='embed'
    ```
    编译后在`release`目录可以找到对应可执行文件。
    ```sh
    release/paopao-ce
    ```

4. 直接运行后端    
    运行api服务:
    ```sh
    make run
    ```
    运行api服务、web前端ui服务:
    ```sh
    make run TAGS='embed'
    ```
    提示: 如果需要内嵌web前端ui，请先构建web前端(建议设置web/.env为VITE_HOST="")。

5. 使用内置的Migrate机制自动升级维护SQL DDL:
    ```sh
    # 添加 Migration 功能到 Features 中 开启migrate功能
    vim config.yaml
    # file: config.yaml
    # Features:
    #   Default: ["Base", "MySQL", "Zinc", "MinIO", "LoggerZinc", "Migration"]
   
    # 编译时加入migration tag编译出支持migrate功能的可执行文件
    make build TAGS='migration'
    release/paopao-ce

    # 或者 带上migration tag直接运行
    make run TAGS='migration'
    ```
    > 注意：默认编译出来的可执行文件是不内置migrate功能，需要编译时带上migration tag才能内置支持migrage功能。


#### 前端

1. 进入前端目录 `web`，拷贝`.env` 到 `.env.local`，编辑 `.env.local ` 文件中后端服务地址及其他配置项，下载依赖包

    ```sh
    cd ./web && cp .env .env.local
    vim .env.local
    yarn
    ```

2. 编译前端

    ```sh
    yarn build
    ```

    build完成后，可以在dist目录获取编译产出，配置nginx指向至该目录即可

#### 桌面端

1. 进入前端目录 `web`，拷贝`.env` 到 `.env.local`，编辑 `.env.local ` 文件中后端服务地址及其他配置项，下载依赖包

    ```sh
    cd ./web && cp .env .env.local
    vim .env.local
    yarn
    ```

2. 编译前端

    ```sh
    yarn build
    ```
   
3. 构建桌面端
   ```sh
   yarn tauri build
   ```
   桌面端是使用[Rust](https://www.rust-lang.org/) + [tauri](https://github.com/tauri-apps/tauri)编写
   的，需要安装tauri的依赖，具体参考[https://tauri.studio/v1/guides/getting-started/prerequisites](https://tauri.studio/v1/guides/getting-started/prerequisites).


### 方式二. 使用Docker构建、运行
  * 后端:
  ```sh
  # 默认参数构建, 默认内嵌web ui并设置api host为空
  docker build -t your/paopao-ce:tag .

  # 内嵌web ui并且自定义API host参数
  docker build -t your/paopao-ce:tag --build-arg API_HOST=http://api.paopao.info .

  # 内嵌web ui并且使用本地web/.env中的API host
  docker build -t your/paopao-ce:tag --build-arg USE_API_HOST=no .

  # 内嵌web ui并且使用本地编译的web/dist构建
  docker build -t your/paopao-ce:tag --build-arg USE_DIST=yes .

  # 只编译api server
  docker build -t your/paopao-ce:tag --build-arg EMBED_UI=no .

  # 运行
  docker run -d -p 8008:8008 -v ${PWD}/config.yaml.sample:/app/paopao-ce/config.yaml your/paopao-ce:tag
  ```

  * 前端:
  ```sh
  cd web

  # 默认参数构建
  docker build -t your/paopao-ce:web .

  # 自定义API host 参数构建
  docker build -t your/paopao-ce:web --build-arg API_HOST=http://api.paopao.info .

  # 使用本地编译的dist构建
  docker build -t your/paopao-ce:web --build-arg USE_DIST=yes .

  # 运行
  docker run -d -p 8010:80 your/paopao-ce:web
  ```

### 方式三. 使用 docker-compose 运行
```sh
git clone https://github.com/rocboss/paopao-ce.git
docker compose up --build
# visit paopao-ce(http://127.0.0.1:8008) and phpMyadmin(http://127.0.0.1:8080)
```

默认是使用config.yaml.sample的配置，如果需要自定义配置，请拷贝默认配置文件(比如config.yaml)，修改后再同步配置到docker-compose.yaml如下：

```
# file: docker-compose.yaml
...
  backend:
    build:
      context: .
    restart: always
    depends_on:
      - db
      - redis
      - zinc
    # modify below to reflect your custom configure
    volumes:
      - ./config.yaml:/app/paopao-ce/config.yaml
    ports:
      - 8008:8008
    networks:
      - paopao-network
....
```

> 注意：默认提供的 docker-compose.yaml 初衷是搭建本机开发调试环境，如果需要产品部署供外网访问，请自行调优配置参数或使用其他方式部署。

### API 文档
* `config.yaml` 添加 `Docs:OpenAPI` 功能项:
```yaml
...
Features:
  Default: ["Base", "MySQL", "Option", "LocalOSS", "LoggerFile", "Docs"]
  Docs: ["Docs:OpenAPI"]
...
```

* 构建时将 `docs` 添加到TAGS中:
```sh
make run TAGS='docs'

# visit http://127.0.0.1:8008/docs/openapi
```

### 配置说明

`config.yaml.sample` 是一份完整的配置文件模版，paopao-ce启动时会读取`./custom/config.yaml`、`./config.yaml`任意一份配置文件（优先读取最先找到的文件）。

```sh
cp config.yaml.sample config.yaml
vim config.yaml # 修改参数
paopao-ce
```

配置文件中的 `Features` 小节是声明paopao-ce运行时开启哪些功能项:

```yaml
...

Features:
  Default: ["Base", "MySQL", "Option", "LocalOSS", "LoggerFile"]
  Develop: ["Base", "MySQL", "Option", "Sms", "AliOSS", "LoggerZinc"]
  Demo: ["Base", "MySQL", "Option", "Sms", "MinIO", "LoggerZinc"]
  Slim: ["Base", "Sqlite3", "LocalOSS", "LoggerFile"]
  Base: ["Zinc", "Redis", "Alipay",]
  Option: ["SimpleCacheIndex"]
  Sms: "SmsJuhe"

...
```

如上： 
Default/Develop/Demo/Slim 是不同 功能集套件(Features Suite)， Base/Option 是子功能套件， Sms是关于短信验证码功能的参数选项。

这里 `Default`套件 代表的意思是： 使用`Base/Option` 中的功能，外加 `MySQL/LocalOSS/LoggerFile`功能，也就是说开启了`Zinc/Redis/Alipay/SimpleCacheIndex/MySQL/LocalOSS/LoggerFile` 7项功能； 
`Develop`套件依例类推。 

使用Feautures:

```sh
release/paopao-ce --help
Usage of release/paopao-ce:
  -features value
        use special features
  -no-default-features
        whether use default features

# 默认使用 Default 功能套件
release/paopao-ce 

# 不包含 default 中的功能集，仅仅使用 develop 中声明的功能集
release/paopao-ce --no-default-features --features develop 

# 使用 default 中的功能集，外加 sms 功能
release/paopao-ce --features sms  

# 手动指定需要开启的功能集
release/paopao-ce --no-default-features --features sqlite3,localoss,loggerfile,redis 
```

目前支持的功能集合:
* 数据库: MySQL/Sqlite3/PostgreSQL   
  `Gorm` + `MySQL`/`Sqlite3`/`PostgreSQL` 使用[gorm](https://github.com/go-gorm/gorm)作为数据库的ORM，默认使用 `Grom` + `MySQL`组合(目前状态：稳定，默认，推荐使用)；    
  `Sqlx` + `MySQL`/`PostgreSQL` 使用[sqlx](https://github.com/jmoiron/sqlx)作为数据库的ORM(目前状态：WIP)；      
* 对象存储: AliOSS/COS/HuaweiOBS/MinIO/LocalOSS      
  `AliOSS` 阿里云对象存储服务；    
  `COS` 腾讯云对象存储服务；  
  `HuaweiOBS` 华为云对象存储服务；  
  `MinIO` [MinIO](https://github.com/minio/minio)对象存储服务；  
  `S3` AWS S3兼容的对象存储服务；    
  `LocalOSS` 提供使用本地目录文件作为对象存储的功能，仅用于开发调试环境；  
* 缓存: Redis/SimpleCacheIndex/BigCacheIndex     
  `SimpleCacheIndex` 提供简单的 广场推文列表 的缓存功能；   
  `BigCacheIndex` 使用[BigCache](https://github.com/allegro/bigcache)缓存 广场推文列表，缓存每个用户每一页，简单做到千人千面(推荐使用)；  
* 搜索: Zinc/Meili   
  `Zinc` 基于[Zinc](https://github.com/zinclabs/zinc)搜索引擎提供推文搜索服务(目前状态: 稳定，推荐使用)；  
  `Meili` 基于[Meilisearch](https://github.com/meilisearch/meilisearch)搜索引擎提供推文搜索服务(目前状态: 稳定，推荐使用);  
* 日志: LoggerFile/LoggerZinc/LoggerMeili  
  `LoggerFile` 使用文件写日志(目前状态: 稳定);   
  `LoggerZinc` 使用[Zinc](https://github.com/zinclabs/zinc)写日志(目前状态: 稳定，推荐使用);    
  `LoggerMeili` 使用[Meilisearch](https://github.com/meilisearch/meilisearch)写日志(目前状态: 内测阶段);  
* 用户关系模式: Friendship/Followship    
  `Friendship` 弱关系好友模式，类似微信朋友圈(目前状态: 开发阶段);    
  `Followship` 关注者模式，类似Twitter的Follow模式(目前状态: WIP);    
* 支付: Alipay   
  `Alipay`  开启基于[支付宝开放平台](https://open.alipay.com/)的钱包功能；   
* 短信验证码: SmsJuhe(需要开启sms)    
  `Sms` 开启短信验证码功能，用于手机绑定验证手机是否注册者的；功能如果没有开启，手机绑定时任意短信验证码都可以绑定手机；
* 开发文档: Docs:OpenAPI     
  `Docs:OpenAPI` 开启openapi文档功能，提供web api文档说明(visit http://127.0.0.1:8008/docs/openapi)；        
* 其他: PhoneBind/OSS:Retention/OSS:TempDir     
  `PhoneBind` 手机绑定功能；     
  `OSS:Retention` 基于对象存储系统的对象过期自动删除特性实现 先创建临时对象再持久化的功能(目前状态: 内测阶段)；  
  `OSS:TempDir` 基于对象存储系统的对象拷贝/移动特性实现 先创建临时对象再持久化的功能(目前状态: 内测阶段)；     

### 搭建依赖环境
#### [Zinc](https://github.com/zinclabs/zinc) 搜索引擎:
* Zinc运行
```sh
# 创建用于存放zinc数据的目录
mkdir -p data/zinc/data

# 使用Docker运行zinc
docker run -d --name zinc --user root -v ${PWD}/data/zinc/data:/data -p 4080:4080 -e ZINC_FIRST_ADMIN_USER=admin -e ZINC_FIRST_ADMIN_PASSWORD=admin -e DATA_PATH=/data public.ecr.aws/zinclabs/zinc:latest

# 查看zinc运行状态
docker ps
CONTAINER ID   IMAGE                                COMMAND                  CREATED        STATUS        PORTS                    NAMES
41465feea2ff   getmeili/meilisearch:v0.27.0         "tini -- /bin/sh -c …"   20 hours ago   Up 20 hours   0.0.0.0:7700->7700/tcp   paopao-ce-meili-1
7daf982ca062   public.ecr.aws/prabhat/zinc:latest   "/go/bin/zinc"           3 weeks ago    Up 6 days     0.0.0.0:4080->4080/tcp   zinc

# 使用docker compose运行
docker compose up -d zinc
# visit http://localhost:4080 打开自带的ui管理界面
```

* 修改Zinc配置
```yaml
# features中加上 Zinc 和 LoggerZinc
Features:
  Default: ["Zinc", "LoggerZinc", "Base", "Sqlite3", "BigCacheIndex","MinIO"]
...
LoggerZinc: # 使用Zinc写日志
  Host: 127.0.0.1:4080  # 这里的host就是paopao-ce能访问到的zinc主机
  Index: paopao-log
  User: admin
  Password: admin
  Secure: False         # 如果使用https访问zinc就设置为True
...
Zinc: # Zinc搜索配置
  Host: 127.0.0.1:4080
  Index: paopao-data
  User: admin
  Password: admin
  Secure: False
```

#### [Meilisearch](https://github.com/meilisearch/meilisearch) 搜索引擎:
* Meili运行
```sh
mkdir -p data/meili/data

# 使用Docker运行
docker run -d --name meili -v ${PWD}/data/meili/data:/meili_data -p 7700:7700 -e MEILI_MASTER_KEY=paopao-meilisearch getmeili/meilisearch:v0.29.0
# visit http://localhost:7700 打开自带的搜索前端ui

# 使用docker compose运行，需要删除docker-compose.yaml中关于meili的注释
docker compose up -d meili

# 使用docker运行meilisearch的ui管理前端
docker run -d --name uirecord -p 7701:3000 bitriory/uirecord
# visit http://localhost:7701

# 使用docker compose运行meilisearch的ui管理前端，需要删除docker-compose.yaml中关于uirecord的注释
docker compose up -d uirecord
# visit http://loclahost:7701

# 查看meili运行状态
docker compose ps
NAME                   COMMAND                  SERVICE             STATUS              PORTS
paopao-ce-meili-1      "tini -- /bin/sh -c …"   meili               running             0.0.0.0:7700->7700/tcp
paopao-ce-uirecord-1   "docker-entrypoint.s…"   uirecord            running             0.0.0.0:7701->3000/tcp
```

* 修改Meili配置
```yaml
# features中加上 Meili 和 LoggerMeili
Features:
  Default: ["Meili", "LoggerMeili", "Base", "Sqlite3", "BigCacheIndex","MinIO"]
...
LoggerMeili: # 使用Meili写日志
  Host: 127.0.0.1:7700
  Index: paopao-log
  ApiKey: paopao-meilisearch
  Secure: False
  MinWorker: 5               # 最小后台工作者, 设置范围[5, 100], 默认5
  MaxLogBuffer: 100          # 最大log缓存条数, 设置范围[10, 10000], 默认100
...
Meili: # Meili搜索配置
  Host: 127.0.0.1:7700      # 这里的host就是paopao-ce能访问到的meili主机
  Index: paopao-data
  ApiKey: paopao-meilisearch
  Secure: False             # 如果使用https访问meili就设置为True
```

#### [MinIO](https://github.com/minio/minio) 对象存储服务
* MinIO运行
```sh
mkdir -p data/minio/data

# 使用Docker运行
docker run -d --name minio -v ${PWD}/data/minio/data:/data -p 9000:9000 -p 9001:9001 -e MINIO_ROOT_USER=minio-root-user -e  MINIO_ROOT_PASSWORD=minio-root-password -e MINIO_DEFAULT_BUCKETS=paopao:public bitnami/minio:latest

# 使用docker compose运行， 需要删除docker-compose.yaml中关于minio的注释
docker compose up -d minio
```

* 修改Minio配置
```yaml
# features中加上 MinIO
Features:
  Default: ["MinIO", "Meili", "LoggerMeili", "Base", "Sqlite3", "BigCacheIndex"]
...
MinIO: # MinIO 存储配置
  AccessKey: Q3AM3UQ867SPQQA43P2F      # AccessKey/SecretKey 需要登入minio管理界面手动创建，管理界面地址: http://127.0.0.1:9001
  SecretKey: zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG
  Secure: False
  Endpoint: 127.0.0.1:9000             # 根据部署的minio主机修改对应地址
  Bucket: paopao                       # 如上，需要在管理界面创建bucket并赋予外部可读写权限
  Domain: 127.0.0.1:9000               # minio外网访问的地址(如果想让外网访问，这里需要设置为外网可访问到的minio主机地址)
...
```

### 其他说明

建议后端服务使用 `supervisor` 守护进程，并通过 `nginx` 反向代理后，提供API给前端服务调用。

短信通道使用的[聚合数据](https://www.juhe.cn/)，如果申请不下来，可以考虑替换其他服务商。

代码结构比较简单，很方便扩展，开发文档请参阅[docs](docs '开发文档').

## 👯‍♀️ 贡献

喜欢的朋友欢迎给个Star、贡献PR。

[![Star History Chart](https://api.star-history.com/svg?repos=rocboss/paopao-ce&type=Date)](https://star-history.com/#rocboss/paopao-ce&Date)

## License

Distributed under the MIT License. See `LICENSE` for more information.

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
