# Changelog

All notable changes to paopao-ce are documented in this file.
## 0.5.0+dev ([`dev`](https://github.com/rocboss/paopao-ce/tree/dev))
### Added
- add `LoggerOpenObserve` feature use OpenObserve to collect log.[#370](https://github.com/rocboss/paopao-ce/pull/370)    
  add `LoggerOpenObserve` to `conf.yaml` 's `Features` section to enable this feature like below:
  ```yaml
  # file config.yaml
  ...
  Features:
    Default: ["Base", "Postgres", "Meili", "LocalOSS", "LoggerOpenObserve", "BigCacheIndex", "web"]
  LoggerOpenObserve: # 使用OpenObserve写日志
  Host: 127.0.0.1:5080
  Organization: paopao-ce
  Stream: default
  User: root@paopao.info
  Password: tiFEI8UeJWuYA7kN
  Secure: False
  MinWorker: 5               # 最小后台工作者, 设置范围[5, 100], 默认5
  MaxLogBuffer: 100          # 最大log缓存条数, 设置范围[10, 10000], 默认100
  ...
  ```
- Added friend tweets bar feature support in home page. [#377](https://github.com/rocboss/paopao-ce/pull/377)    
- web: add custom `Friendship` feature support. To custom setup `Friendship` use below configure in `web/.env` or `web/.env.local`
  ```
  # 功能特性开启
  VITE_USE_FRIENDSHIP=true

  # 模块开启
  VITE_ENABLE_FRIENDS_BAR=true
  ```
- add Newest/Hots/Following tweets support in friend bar feature. 
  mirgration database first(sql ddl file in `scripts/migration/**/*_home_timeline.up.sql`):
    ```sql
    CREATE TABLE `p_post_metric` (
	  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
	  `post_id` bigint unsigned NOT NULL,
	  `rank_score` bigint unsigned NOT NULL DEFAULT 0,
	  `incentive_score` int unsigned NOT NULL DEFAULT 0,
	  `decay_factor` int unsigned NOT NULL DEFAULT 0,
	  `motivation_factor` int unsigned NOT NULL DEFAULT 0,
	  `is_del` tinyint NOT NULL DEFAULT 0, -- 是否删除, 0否, 1是
	  `created_on` bigint unsigned NOT NULL DEFAULT '0',
	  `modified_on` bigint unsigned NOT NULL DEFAULT '0',
	  `deleted_on` bigint unsigned NOT NULL DEFAULT '0',
	  PRIMARY KEY (`id`) USING BTREE,
	  KEY `idx_post_metric_post_id_rank_score` (`post_id`,`rank_score`) USING BTREE
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

    INSERT INTO p_post_metric (post_id, rank_score, created_on) 
    SELECT id AS post_id, 
	    comment_count + upvote_count*2 + collection_count*4 AS rank_score,
	    created_on
    FROM p_post
    WHERE is_del=0;

    -- 原来的可见性: 0公开 1私密 2好友可见 3关注可见
    -- 现在的可见性: 0私密 10充电可见 20订阅可见 30保留 40保留 50好友可见 60关注可见 70保留 80保留 90公开
    UPDATE p_post a, p_post b 
    SET a.visibility = (
	    CASE b.visibility 
		    WHEN 0 THEN 90 
		    WHEN 1 THEN 0 
		    WHEN 2 THEN 50 
		    WHEN 3 THEN 60 
		    ELSE 0
	    END 
    )
    WHERE a.ID = b.ID;
    ```
- add cache support for index/home etc. page.
- add hots comments support for post detail page.
- add highlight comments support for post detail page.
    mirgration database first(sql ddl file in `scripts/migration/**/*_comment_esence.up.sql`):
    ```sql
      ALTER TABLE `p_comment` ADD COLUMN `is_essence` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '是否精选';
    ```
- add follow/unfollow user support in index/home/collecion/message/post page.
- add simple prometheus metrics support.
    add `Metrics` to `conf.yaml` 's `Features` section to enable this feature like below:
    ```yaml
    # file config.yaml
    ...
    Features:
      Default: ["Base", "Postgres", "Meili", "LocalOSS", "Metrics", "web"]
    JobManager: # Cron Job理器的配置参数
      MaxOnlineInterval: "@every 5m"       # 更新最大在线人数，默认每5分钟更新一次
      UpdateMetricsInterval: "@every 5m"   # 更新Prometheus指标，默认每5分钟更新一次
    MetricsServer: # Prometheus Metrics服务
      RunMode: debug
      HttpIp: 0.0.0.0
      HttpPort: 6080
      ReadTimeout: 60
      WriteTimeout: 60
    ...
    ```
- add full support for tweet hots comment logic and add cache support for tweet comments.
    mirgration database first(sql ddl file in `scripts/migration/**/*_rank_metrics.up.sql`):
    ```sql
    ALTER TABLE `p_comment` ADD COLUMN `reply_count` int unsigned NOT NULL DEFAULT 0 COMMENT '回复数';

    UPDATE p_comment comment 
    SET reply_count = (
      SELECT count(*) FROM p_comment_reply reply WHERE reply.comment_id=comment.id AND reply.is_del=0
    )
    WHERE is_del=0;

    CREATE TABLE `p_comment_metric` (
	    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  	  `comment_id` bigint unsigned NOT NULL,
	    `rank_score` bigint unsigned NOT NULL DEFAULT 0,
	    `incentive_score` int unsigned NOT NULL DEFAULT 0,
	    `decay_factor` int unsigned NOT NULL DEFAULT 0,
	    `motivation_factor` int unsigned NOT NULL DEFAULT 0,
	    `is_del` tinyint NOT NULL DEFAULT 0,
	    `created_on` bigint unsigned NOT NULL DEFAULT 0,
	    `modified_on` bigint unsigned NOT NULL DEFAULT 0,
	    `deleted_on` bigint unsigned NOT NULL DEFAULT 0,
	    PRIMARY KEY (`id`) USING BTREE,
	    KEY `idx_comment_metric_comment_id_rank_score` (`comment_id`, `rank_score`) USING BTREE
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

    INSERT INTO p_comment_metric (comment_id, rank_score, created_on) 
    SELECT id AS comment_id, 
	    reply_count*2 + thumbs_up_count*4 - thumbs_down_count AS rank_score,
	    created_on
    FROM p_comment
    WHERE is_del=0;

    CREATE TABLE `p_user_metric` (
	    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
	    `user_id` bigint unsigned NOT NULL,
	    `tweets_count` int unsigned NOT NULL DEFAULT 0,
      `latest_trends_on` bigint unsigned NOT NULL DEFAULT 0 COMMENT '最新动态时间',
	    `is_del` tinyint NOT NULL DEFAULT 0,
	    `created_on` bigint unsigned NOT NULL DEFAULT 0,
	    `modified_on` bigint unsigned NOT NULL DEFAULT 0,
	    `deleted_on` bigint unsigned NOT NULL DEFAULT 0,
	    PRIMARY KEY (`id`) USING BTREE,
	    KEY `idx_user_metric_user_id_tweets_count_trends` (`user_id`, `tweets_count`, `latest_trends_on`) USING BTREE
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

    INSERT INTO p_user_metric (user_id, tweets_count) 
    SELECT user_id, count(*) AS tweets_count
    FROM p_post
    WHERE is_del=0
    GROUP BY user_id;
    ```
- add message filter support for message page.
- add read all unread message and display unread message count support for message page.
- add support follow user embed to index trends enable navigation user tweets by slide bar.
  mirgration database first(sql ddl file in `scripts/migration/**/*_user_relation.up.sql`):
  ```sql
  CREATE VIEW p_user_relation AS 
  SELECT user_id, friend_id he_uid, 5 AS style 
  FROM p_contact WHERE status=2 AND is_del=0
  UNION
  SELECT user_id, follow_id he_uid, 10 AS style 
  FROM p_following WHERE is_del=0;
  ```
    
## 0.4.2
### Fixed
- fixed remove multi-objects no effects and occurs resource leak error when use Minio as OSS(Object Storage System).[#371](https://github.com/rocboss/paopao-ce/pull/371) [#372](https://github.com/rocboss/paopao-ce/pull/372)  

## 0.4.1
### Changed
- infinite scrolling instead of pagination for Home/User/Profile page

## 0.4.0
### Added
- add `pprof` feature support [#327](https://github.com/rocboss/paopao-ce/pull/327)  
- use compiler profile-guided optimization (PGO) to further optimize builds. [#327](https://github.com/rocboss/paopao-ce/pull/327)  
- frontend: re-add stars page embed to  profile page. [#339](https://github.com/rocboss/paopao-ce/pull/339)  
- simple support for user posts filter by style(post/comment/media/star). [#345](https://github.com/rocboss/paopao-ce/pull/345)  
  migration database first(sql ddl file in `scripts/migration/**/*_create_view_post_filter.up.sql`):
  ```sql
  CREATE VIEW p_post_by_media AS SELECT post.*FROM (SELECT DISTINCT post_id FROM p_post_content WHERE (TYPE=3 OR TYPE=4 OR TYPE=7 OR TYPE=8) AND is_del=0) media JOIN p_post post ON media.post_id=post.ID WHERE post.is_del=0;
  CREATE VIEW p_post_by_comment AS SELECT P.*,C.user_id comment_user_id FROM (SELECT post_id,user_id FROM p_comment WHERE is_del=0 UNION SELECT post_id,reply.user_id user_id FROM p_comment_reply reply JOIN p_comment COMMENT ON reply.comment_id=COMMENT.ID WHERE reply.is_del=0 AND COMMENT.is_del=0) C JOIN p_post P ON C.post_id=P.ID WHERE P.is_del=0;
  ```
- add user highlight tweet support include custom tweet set to highlight and list in user/profile page.
- add cli subcommand to start paopao-ce serve or other task. [#354](https://github.com/rocboss/paopao-ce/pull/354)  
- add `Followship` feature . [#355](https://github.com/rocboss/paopao-ce/pull/355)
  migration database first(sql ddl file in `scripts/migration/**/*_user_following.up.sql`):
  ```sql
  DROP TABLE IF EXISTS p_following; 
  CREATE TABLE p_following (ID BIGSERIAL PRIMARY KEY,user_id BIGINT NOT NULL,follow_id BIGINT NOT NULL,is_del SMALLINT NOT NULL DEFAULT 0,created_on BIGINT NOT NULL DEFAULT 0,modified_on BIGINT NOT NULL DEFAULT 0,deleted_on BIGINT NOT NULL DEFAULT 0); 
  CREATE INDEX idx_following_user_follow ON p_following USING btree (user_id,follow_id);
  ```

### Changed
- change man content width to 600px and optimize tweet/comment/replay text length. [#333](https://github.com/rocboss/paopao-ce/pull/333) 
- optimize embed web ui to paopao execute binary file logic. [#354](https://github.com/rocboss/paopao-ce/pull/354)   
  ```sh
  # embed web ui to execute file default
  make build 
  # use slim model to disable embed web ui to exectute file
  make build TAGS='slim embed'
  ```
- frontend: optimize user profile page route path to domain/#/u/?s=username. [&c857142](https://github.com/rocboss/paopao-ce/commit/c857142565f0c28294344c7abc5c2df4e363b04c
- change the `Friendship` feature and `Followship` feature as builtin feature. [#362](https://github.com/rocboss/paopao-ce/pull/362)  
- deprecated/remove `Lightship` feature. [#362](https://github.com/rocboss/paopao-ce/pull/362)  
 - optimize Followship feature allow follow/unfollow user in follow page. [&fd5e54b](https://github.com/rocboss/paopao-ce/commit/fd5e54bd31108bb5a879d8b7afe0d2cbd56b3395)
- use Meilisearch as default search in docker-compose.yaml.

### Fixed
- fixed JWT valide error whent get user by id failed.[&51fd972](https://github.com/rocboss/paopao-ce/commit/51fd9724b96bbc8cc64405b36a5fd7ae69b7a52c)
  
## 0.3.1
### Fixed
- fixed: video player assets cdn error. [&caff8c0](https://github.com/rocboss/paopao-ce/commit/caff8c052be6c8d59576011192f830fd98e17ab3 'commit caff8c0')  

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