| 编号 | 作者 | 发表时间 | 变更时间 | 版本 | 状态 |
| ----- | ----- | ----- | ----- | ----- | ----- |
| 23062905| 北野 | 2023-06-29 | 2023-06-29 | v1.0 | 提议 |

### 概述  
使用net/http/pprof获取Profile信息。添加`Pprof`功能特性用于开启PprofServer服务。

### 需求   
- 只在开发环境下获取profile信息
- 可以在配置文件中开启获取profile的服务

### 方案  
- 添加`Pprof`功能特性
- 添加`PprofServer`服务

#### 设计细节 
- 配置文件中添加`PprofServer`项用于配制http server
- 添加`Pprof`功能特性
- 添加`PprofServer`服务
- 使用`//go:build pprof`按需编译profile功能
- `PprofServer` 不能和其他服务共用一个端口

### 疑问

1. 为什么要添加`pprof`功能特性？   
使用 net/http/pprof 在线获取CPU profile信息，用于 [PGO](https://go.dev/doc/pgo) 编译优化。   

2. 如何开启`Pprof`服务？   
* 构建时将 `pprof` 添加到TAGS中:
    ```sh
    make run TAGS='pprof'   
    ```          
* 在配置文件config.yaml中的`Features`中添加`Pprof`功能项开启该功能：
    ```yaml
    ...
    # features中加上 Friendship
    Features:
      Default: ["Meili", "LoggerMeili", "Base", "Sqlite3", "BigCacheIndex", "MinIO", "Pprof"]
      Base: ["Redis", "PhoneBind"]
    ...
    ```

3. 常见命令  

#### Merging profiles
The pprof tool can merge multiple profiles like this:
```sh
$ go tool pprof -proto a.pprof b.pprof > merged.pprof
```
This merge is effectively a straightforward sum of samples in the input, regardless of wall duration of the profile. As a result, when profiling a small time slice of an application (e.g., a server that runs indefinitely), you likely want to ensure that all profiles have the same wall duration (i.e., all profiles are collected for 30s). Otherwise, profiles with longer wall duration will be overrepresented in the merged profile.

#### look at a 30-second CPU profile:
```sh
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
```

### 参考文档
* [PGO](https://go.dev/doc/pgo)
* [net/http/pprof](https://pkg.go.dev/net/http/pprof)
* [runtime/pprof](https://pkg.go.dev/runtime/pprof)  

### 参考实现
- [#327](https://github.com/rocboss/paopao-ce/pull/327)

### 更新记录
#### v1.0(2023-06-20) - 北野
* 初始文档
