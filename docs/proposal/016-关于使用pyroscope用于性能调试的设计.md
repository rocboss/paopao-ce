| 编号 | 作者 | 发表时间 | 变更时间 | 版本 | 状态 |
| ----- | ----- | ----- | ----- | ----- | ----- |
| 016| 北野 | 2023-02-15 | 2023-02-15 | v1.0 | 提议 |

### 概述  
Pyroscope 是一个开源的持续性能剖析平台。它能够帮你：
* 找出源代码中的性能问题和瓶颈
* 解决 CPU 利用率高的问题
* 理解应用程序的调用树（call tree）
* 追踪随一段时间内变化的情况

### 需求   
* 开发环境下启用Pyroscope，但是部署环境下禁用。

### 方案  
* 使用//go:build pyroscope 可选编译app集成Pyroscope功能
* config.yaml中添加`Pyroscope` 功能来启用Pyroscope

### 疑问

1. 为什么要引入Pyroscope？      
用于开发环境下对paopao-ce进行性能优化。

2. 如何开启这个功能？          
在配置文件config.yaml中的`Features`中添加`Pyroscope`功能项开启该功能：
    ```yaml
    ...
    # features中加上 Friendship
    Features:
      Default: ["Meili", "LoggerMeili", "Base", "Sqlite3", "BigCacheIndex", "MinIO", "Pyroscope"]
      Base: ["Redis", "PhoneBind"]
    ...
    ```



### 参考文档
* [pyroscope](https://github.com/pyroscope-io/pyroscope)
* [pyroscope client](https://github.com/pyroscope-io/client)

### 更新记录
#### v1.0(2023-02-15) - 北野
* 初始文档
