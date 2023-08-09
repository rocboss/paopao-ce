| 编号 | 作者 | 发表时间 | 变更时间 | 版本 | 状态 |
| ----- | ----- | ----- | ----- | ----- | ----- |
| 22121409 | 北野 | 2022-12-14 | 2022-01-09 | v1.1 | 提议 |

### 关于Lightship功能项的设计
Lightship(开放模式)功能提供完全公开的推文分享服务，有别于[Friendship](002-关于Friendship功能项的设计.md "关于Friendship功能项的设计")、[Followship](003-关于Followship功能项的设计.md "关于Followship功能项的设计")，使用Lightship用户模式部署paopao-ce，用户发布的所有推文都是公开可访问的，广场推文列表展示的是全站所有公开推文的Timeline Tweets。

### 场景
一般用于非常小的站点，或者推文更新不频繁的站点。
* 推文更新不频繁站点；
* 个人推文站点；
* 单纯想要所有推文公开访问；

### 需求
* 广场推文列表只展示公开可访问推文；
* 推文访问权限只能是 `公开/私密`；

### 方案

#### 实现细节
* 广场推文列表只展示公开可访问推文 - 前端/后端   
* 推文访问权限只能是 `公开/私密` - 前端/后端  

#### 参考实现(PR):  
[add Lightship feature support #198](https://github.com/rocboss/paopao-ce/pull/198)  

### 疑问

1. 公开模式为什么命名为Lightship？   
Lightship有灯塔的意思，灯塔是为航行船只指明方向的公共设施，这里取其公有设施的属性指喻 开放模式下paopao-ce不需要授权公开访问推 文的意思。    
1. 如何开启这个功能？     
在配置文件config.yaml中的`Features`中添加`Lightship`功能项开启该功能：
```yaml
...
# features中加上 Lightship
Features:
  Default: ["Meili", "LoggerMeili", "Base", "Sqlite3", "BigCacheIndex", "MinIO", "Lightship"]
  Base: ["Redis", "PhoneBind"]
...
```

### 更新记录
#### v1.0(2022-12-14) - 北野
* 初始文档

#### v1.1(2022-01-09) - 北野
* 添加参考实现PR信息
