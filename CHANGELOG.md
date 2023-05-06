# Changelog

All notable changes to paopao-ce are documented in this file.
## 0.4.0+dev ([`dev`](https://github.com/rocboss/paopao-ce/tree/dev))

## 0.3.0
### Added

- add custom comment sort strategy support [#243](https://github.com/rocboss/paopao-ce/pull/243)
- add `RedisCacheIndex` feature [#250](https://github.com/rocboss/paopao-ce/pull/250)
- add `Sentry` feature [#258](https://github.com/rocboss/paopao-ce/pull/258)
- add simple tweet share feature(just copy tweet link to clipboard now) support [#264](https://github.com/rocboss/paopao-ce/pull/264)
- add default tweet max length configure in web/.env support. [&a1160ca](https://github.com/rocboss/paopao-ce/commit/a1160ca79380445157146d9eae1710543c153cce 'commit a1160ca')   
 Set the value of `VITE_DEFAULT_TWEET_MAX_LENGTH` in file web/.env to change the tweet max default length.
 - add custom whether provide user register configure in web/.env support. [#267](https://github.com/rocboss/paopao-ce/pull/267)   
    Set the value of `VITE_ALLOW_USER_REGISTER` in file web/.env to custom whether provide user register feature.
    ```
    # file: web/.env or web/.env.local
    ...
    # 局部参数
    VITE_ALLOW_USER_REGISTER=true
    ...
    ```
    and disallow user register in backend(add `Web:DisallowUserRegister` feature in `config.yaml`):
    ```yaml
    # file config.yaml
    ...
    Features:
      Default: ["Base", "Postgres", "Zinc", "LocalOSS", "LoggerZinc", "BigCacheIndex", "Friendship", "Service", "Web:DisallowUserRegister"]
    ...
    ```
- add topic follow feature support [#273](https://github.com/rocboss/paopao-ce/pull/273)   
  mirgration database first(sql ddl file in `scripts/migration/**/*_topic_follow.up.sql`):
    ```sql
      CREATE TABLE `p_topic_user` (
	      `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	      `topic_id` BIGINT UNSIGNED NOT NULL COMMENT '标签ID',
	      `user_id` BIGINT UNSIGNED NOT NULL COMMENT '创建者ID',
	      `alias_name` VARCHAR ( 255 ) COMMENT '别名',
	      `remark` VARCHAR ( 512 ) COMMENT '备注',
	      `quote_num` BIGINT UNSIGNED COMMENT '引用数',
	      `is_top` TINYINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否置顶 0 为未置顶、1 为已置顶',
	      `created_on` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
	      `modified_on` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '修改时间',
	      `deleted_on` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '删除时间',
	      `is_del` TINYINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
	      `reserve_a` VARCHAR ( 255 ) COMMENT '保留字段a',
	      `reserve_b` VARCHAR ( 255 ) COMMENT '保留字段b',
	      PRIMARY KEY ( `id` ) USING BTREE,
        UNIQUE KEY `idx_topic_user_uid_tid` ( `topic_id`, `user_id` ) USING BTREE 
      ) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户话题';
    ```
- add tweet comment thumbs up/down feature support [#275](https://github.com/rocboss/paopao-ce/pull/275)   
  mirgration database first(sql ddl file in `scripts/migration/**/*_comment_thumbs.up.sql`):
- add load more comments feature support [&60b217b](https://github.com/rocboss/paopao-ce/commit/60b217bcd950c69ba45cebcaa17efdf8048d5a4f 'commit 60b217b')  

### Fixed

- fixed sql ddl p_contact's column `is_delete` define error (change to `is_del`) in scripts/paopao-mysql.sql [&afd8fe1](https://github.com/rocboss/paopao-ce/commit/afd8fe18d2dce08a4af846c2f822379d99a3d3b3 'commit afd8fe1')
- fixed cache index not expire in delete/add tweet error [#266](https://github.com/rocboss/paopao-ce/pull/266)
- fixed tweet's owner not allow star/collection action when tweet is private error [#274](https://github.com/rocboss/paopao-ce/pull/274)
- fixed user not list owner's collectioned private tweet error [#274](https://github.com/rocboss/paopao-ce/pull/274)
- fixed comments thumbs up/down state incorrect error [#283](https://github.com/rocboss/paopao-ce/pull/283)

### Changed

- use [github.com/rueian/rueidis](https://github.com/rueian/rueidis) as Redis client [#249](https://github.com/rocboss/paopao-ce/pull/249)   
    the **Old** redis client configure field
    ```yaml
    ...
    Redis:
      Host: redis:6379
      Password: 
      DB: 
    ```
    the **New** redis client configure field
    ```yaml
    ...
    Redis:
      InitAddress:
      - redis:6379
      Username:
      Password:
      SelectDB:
      ConnWriteTimeout: 60   # 连接写超时时间 多少秒 默认 60秒
    ```
- optimize web frontend dark theme [&b082a8f](https://github.com/rocboss/paopao-ce/commit/b082a8fa5e43dd6dacf459df93fa7e243dd901ea 'commit b082a8f')
- change web frontend main content layout default size to 544px [&b082a8f](https://github.com/rocboss/paopao-ce/commit/b082a8fa5e43dd6dacf459df93fa7e243dd901ea 'commit b082a8f')
- optimize web frontend in mobile environment use Drawer to display menu [#265](https://github.com/rocboss/paopao-ce/pull/265)   
- optimize Dockerfile use pre-build builder/runner image to prevent network latency problem (`bitbus/paopao-ce-backend-builder` `bitbus/paopao-ce-backend-runner`) [#265](https://github.com/rocboss/paopao-ce/pull/265)  
- optimize web ui in mobile environment [#280](https://github.com/rocboss/paopao-ce/pull/280)  
- optimize upload zip attachment compatible with different browsers for uploading zip mimetype [#286](https://github.com/rocboss/paopao-ce/pull/286)  
- adapte meilisearch to version v1.1 [#288](https://github.com/rocboss/paopao-ce/pull/288)  

### Removed

- remove `Deprecated:OldWeb` feature  [#256](https://github.com/rocboss/paopao-ce/pull/256)   

## 0.2.5

### Changed

- fixed sql ddl error for contact table [#281](https://github.com/rocboss/paopao-ce/pull/281)

## 0.2.4

### Added

-  add PWA support for web frontend [#242](https://github.com/rocboss/paopao-ce/pull/242)


## 0.2.3

### Added

-  add PostgreSQL DDL file [#229](https://github.com/rocboss/paopao-ce/pull/229)

### Changed

- optimize MySQL DDL file [#229](https://github.com/rocboss/paopao-ce/pull/229)
- optimize Sqlite3 DDL file [#229](https://github.com/rocboss/paopao-ce/pull/229)

## 0.2.2

### Fixed

- fixed add star to tweet error [#222](https://github.com/rocboss/paopao-ce/pull/222)

## 0.2.1

### Changed

- optimize docker-compose.yaml use bitbus/paopao-ce:latest as release image [#217](https://github.com/rocboss/paopao-ce/pull/217)

### Fixed

- fixed sql ddl in scripts/paopao-mysql.sql and scripts/paopao-sqlite3.sql [#217](https://github.com/rocboss/paopao-ce/pull/217)

## 0.2.0

### Added

- add `Friendship` feature [#192](https://github.com/rocboss/paopao-ce/pull/192)
- add `Lightship` feature [#198](https://github.com/rocboss/paopao-ce/pull/198)
- add `Pyroscope` feature [#199](https://github.com/rocboss/paopao-ce/pull/199) 
- add new `Web` service [#196](https://github.com/rocboss/paopao-ce/pull/196)  
- add `Frontend:Web` feature [#196](https://github.com/rocboss/paopao-ce/pull/196)  
- add `Deprecated:OldWeb` feature [#196](https://github.com/rocboss/paopao-ce/pull/196)  

### Changes

- support run multiple service in single paopao-ce instance [#196](https://github.com/rocboss/paopao-ce/pull/196)  
- use [go-mir](https://github.com/alimy/mir) optimize paopao-ce source code architecture [#196](https://github.com/rocboss/paopao-ce/pull/196)  

### Fixed

- some other features optimize and bug fix 

---

**Older change logs can be found on [GitHub](https://github.com/rocboss/paopao-ce/releases?after=v0.2.0).**