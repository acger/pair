CREATE TABLE `elements`
(
    `id`          bigint(20) NOT NULL /*T![auto_rand] AUTO_RANDOM(5) */,
    `create_time` datetime                                 DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime                                 DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `uid`         bigint(20)                               DEFAULT '0',
    `skill`       varchar(6000) COLLATE utf8mb4_0900_ai_ci DEFAULT '',
    `skill_need`  varchar(6000) COLLATE utf8mb4_0900_ai_ci DEFAULT '',
    `star`        tinyint(1)                               DEFAULT '0',
    `boost`       bigint(20)                               DEFAULT '0',
    PRIMARY KEY (`id`) /*T![clustered_index] CLUSTERED */,
    UNIQUE KEY `uid` (`uid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;
