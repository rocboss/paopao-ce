# Changelog

All notable changes to paopao-ce are documented in this file.

## 0.3.0+dev ([`dev`](https://github.com/rocboss/paopao-ce/tree/dev))

### Added

- add custom comment sort strategy support [#243](https://github.com/rocboss/paopao-ce/pull/243)
- add `RedisCacheIndex` feature [#250](https://github.com/rocboss/paopao-ce/pull/250)
- add `Sentry` feature [#258](https://github.com/rocboss/paopao-ce/pull/258)
- add simple tweet share feature(just copy tweet link to clipboard now) support [#264](https://github.com/rocboss/paopao-ce/pull/264)
- add default tweet max length configure in web/.env support. [&a1160ca](https://github.com/rocboss/paopao-ce/commit/a1160ca79380445157146d9eae1710543c153cce 'commit a1160ca')   
 Set the value of `VITE_DEFAULT_TWEET_MAX_LENGTH` in file web/.env to change the tweet max default length.
  ```
  # file: web/.env or web/.env.local
  ...
  # 局部参数
  VITE_DEFAULT_TWEET_MAX_LENGTH=300
  ...
  ```

### Fixed

- fixed sql ddl p_contact's column `is_delete` define error (change to `is_del`) in scripts/paopao-mysql.sql [&afd8fe1](https://github.com/rocboss/paopao-ce/commit/afd8fe18d2dce08a4af846c2f822379d99a3d3b3 'commit afd8fe1')

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

### Removed

- remove `Deprecated:OldWeb` feature  [#256](https://github.com/rocboss/paopao-ce/pull/256)   

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