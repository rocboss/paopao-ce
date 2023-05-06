## Features Status
功能项状态说明。

## 目前支持的功能集合
#### 子服务
* `OldWeb` 开启旧的Web服务(目前状态: 默认)
    * [ ] 提按文档
    * [x] 服务初始化逻辑
    * [x] 接口定义
    * [x] 业务逻辑实现
* `Web` 开启Web服务(目前状态: WIP)
    * [ ] 提按文档
    * [x] 服务初始化逻辑
    * [x] 接口定义
    * [x] 业务逻辑实现
* `Admin` 开启Admin后台运维服务(目前状态: WIP)
    * [ ] 提按文档
    * [x] 服务初始化逻辑
    * [ ] 接口定义
    * [ ] 业务逻辑实现
* `SpaceX` 开启SpaceX服务(目前状态: WIP)
    * [ ] 提按文档
    * [x] 服务初始化逻辑
    * [ ] 接口定义
    * [ ] 业务逻辑实现
* `Bot` 开启Bot服务(目前状态: WIP)
    * [ ] 提按文档
    * [x] 服务初始化逻辑
    * [ ] 接口定义
    * [ ] 业务逻辑实现
* `NativeOBS` 开启NativeOBS服务(目前状态: WIP)
    * [ ] 提按文档
    * [ ] 服务初始化逻辑
    * [ ] 接口定义
    * [ ] 业务逻辑实现
* `Docs` 开启NativeOBS服务(目前状态: WIP)
    * [ ] 提按文档
    * [x] 服务初始化逻辑
* `Frontend:Web` 开启独立前端服务(目前状态: 内测)
    * [ ] 提按文档
    * [x] 服务初始化逻辑
* `Frontend:EmbedWeb` 开启内嵌于后端Web API服务中的前端服务(目前状态: 内测)
    * [ ] 提按文档
    * [x] 服务初始化逻辑
* `Deprecated:OldWeb`  开启旧的Web服务(目前状态: 已弃，不可用)
    * [ ] 提按文档
    * [x] 服务初始化逻辑

#### 数据库:
* `Gorm` + `MySQL`/`Sqlite3`/`PostgreSQL`
使用[gorm](https://github.com/go-gorm/gorm)作为数据库的ORM，默认使用 `Grom` + `MySQL`组合(目前状态：稳定，默认，推荐使用)； 
    * [ ] 提按文档
    * [x] 接口定义
    * [x] 业务逻辑实现  
* `Sqlx` + `MySQL`/`PostgreSQL`
使用[sqlx](https://github.com/jmoiron/sqlx)作为数据库的ORM(目前状态：WIP)；
    * [ ] 提按文档
    * [ ] 接口定义
    * [ ] 业务逻辑实现  
    
#### 对象存储: 
* `AliOSS` 阿里云对象存储服务；
    * [ ] 提按文档  
    * [x] 接口定义
    * [x] 业务逻辑实现    
* `COS` 腾讯云对象存储服务；
    * [ ] 提按文档  
    * [x] 接口定义
    * [x] 业务逻辑实现   
* `HuaweiOBS` 华为云对象存储服务；
    * [ ] 提按文档  
    * [x] 接口定义
    * [x] 业务逻辑实现   
* `MinIO` [MinIO](https://github.com/minio/minio)对象存储服务；
    * [ ] 提按文档  
    * [x] 接口定义
    * [x] 业务逻辑实现   
* `S3` AWS S3兼容的对象存储服务；
    * [ ] 提按文档  
    * [x] 接口定义
    * [x] 业务逻辑实现     
* `LocalOSS` 提供使用本地目录文件作为对象存储的功能，仅用于开发调试环境；
    * [ ] 提按文档  
    * [x] 接口定义
    * [x] 业务逻辑实现 
* `OSS:Retention` 基于对象存储系统的对象过期自动删除特性实现 先创建临时对象再持久化的功能(目前状态: 内测阶段)；  
    * [ ] 提按文档  
    * [x] 接口定义
    * [x] 业务逻辑实现 
* `OSS:TempDir` 基于对象存储系统的对象拷贝/移动特性实现 先创建临时对象再持久化的功能(目前状态: 内测阶段)；
    * [ ] 提按文档  
    * [x] 接口定义
    * [x] 业务逻辑实现   

#### 缓存:
* `Redis` Redis缓存
    * [ ] 提按文档  
    * [x] 接口定义
    * [x] 业务逻辑实现 
* `SimpleCacheIndex` 提供简单的 广场推文列表 的缓存功能能(目前状态: Deprecated)； 
    * [ ] 提按文档  
    * [x] 接口定义
    * [x] 业务逻辑实现   
* `BigCacheIndex` 使用[BigCache](https://github.com/allegro/bigcache)缓存 广场推文列表，缓存每个用户每一页，简单做到千人千面(推荐使用)；  
    * [ ] 提按文档  
    * [x] 接口定义
    * [x] 业务逻辑实现 
* `RedisCacheIndex` 使用Redis缓存 广场推文列表，缓存每个用户每一页，简单做到千人千面(目前状态: 内测阶段，推荐使用)；  
    * [ ] 提按文档  
    * [x] 接口定义
    * [x] 业务逻辑实现 

#### 搜索:
* `Zinc` 基于[Zinc](https://github.com/zinclabs/zinc)搜索引擎提供推文搜索服务(目前状态: 稳定，推荐使用)；  
    * [ ] 提按文档  
    * [x] 接口定义
    * [x] 业务逻辑实现 
* `Meili` 基于[Meilisearch](https://github.com/meilisearch/meilisearch)搜索引擎提供推文搜索服务(目前状态: 稳定，推荐使用); 
    * [ ] 提按文档  
    * [x] 接口定义
    * [x] 业务逻辑实现  
* `Bleve` 基于[Bleve](https://github.com/blevesearch/bleve)搜索引擎提供推文搜索服务(目前状态: WIP); 
    * [ ] 提按文档  
    * [ ] 接口定义
    * [ ] 业务逻辑实现  

#### 日志:
* `LoggerFile` 使用文件写日志(目前状态: 稳定); 
    * [ ] 提按文档  
    * [x] 接口定义
    * [x] 业务逻辑实现   
* `LoggerZinc` 使用[Zinc](https://github.com/zinclabs/zinc)写日志(目前状态: 稳定，推荐使用);   
    * [ ] 提按文档  
    * [x] 接口定义
    * [x] 业务逻辑实现  
* `LoggerMeili` 使用[Meilisearch](https://github.com/meilisearch/meilisearch)写日志(目前状态: 内测阶段);  
    * [ ] 提按文档  
    * [x] 接口定义
    * [x] 业务逻辑实现 

#### 监控:
* `Sentry` 使用Sentry进行错误跟踪与性能监控(目前状态: 内测); 
    * [x] [提按文档](docs/proposal/23040412-关于使用sentry用于错误追踪与性能检测的设计.md)  
    * [x] 接口定义
    * [x] 业务逻辑实现   

#### 关系模式:    
* `Friendship` 弱关系好友模式，类似微信朋友圈(目前状态: 内测); 
    * [x] [提按文档](docs/proposal/002-关于Friendship功能项的设计.md)
    * [x] 接口定义
    * [x] 业务逻辑实现     
* `Followship` 关注者模式，类似Twitter的Follow模式(目前状态: WIP);  
    * [ ] [提按文档](docs/proposal/003-关于Followship功能项的设计.md)
    * [ ] 接口定义
    * [ ] 业务逻辑实现   
* `Lightship` 开放模式，所有推文都公开可见(目前状态: 内测、默认);  
    * [x] [提按文档](docs/proposal/011-关于Lightship功能项的设计.md)
    * [x] 接口定义
    * [x] 业务逻辑实现 

### 支付:  
* `Alipay`  开启基于[支付宝开放平台](https://open.alipay.com/)的钱包功能； 
    * [ ] 提按文档  
    * [x] 接口定义
    * [x] 业务逻辑实现   

### 短信验证: 
* SmsJuhe(需要开启sms)  
    * [ ] 提按文档  
    * [x] 接口定义
    * [x] 业务逻辑实现 
* `Sms` 开启短信验证码功能，用于手机绑定验证手机是否注册者的；功能如果没有开启，手机绑定时任意短信验证码都可以绑定手机；
    * [ ] 提按文档  
    * [x] 接口定义
    * [x] 业务逻辑实现 

### 开发文档:  
* `Docs:OpenAPI` 开启openapi文档功能，提供web api文档说明(visit http://127.0.0.1:8008/docs/openapi);  
    * [ ] 提按文档  
    * [x] 接口定义
    * [x] 业务逻辑实现   

### 性能优化
* [`Pyroscope`](docs/proposal/016-关于使用pyroscope用于性能调试的设计.md) 开启Pyroscope功能用于性能调试(目前状态: 内测);  
    * [x] 提按文档  
    * [x] 业务逻辑实现  
 
### 其他:    
* `PhoneBind` 手机绑定功能； 
    * [ ] 提按文档  
    * [x] 接口定义
    * [x] 业务逻辑实现    

### 功能特性:
* `Web:DisallowUserRegister` 不允许用户注册；       
    * [ ] 提按文档  
    * [x] 接口定义
    * [x] 业务逻辑实现   
