CREATE TABLE `chats`
(
    `id`          bigint(20) NOT NULL /*T![auto_rand] AUTO_RANDOM(5) */,
    `create_time` datetime                                 DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime                                 DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `uid`         bigint(20)                               DEFAULT '0',
    `to_uid`      bigint(20)                               DEFAULT '0',
    `message`     varchar(3000) COLLATE utf8mb4_0900_ai_ci DEFAULT '',
    `status`      tinyint(1)                               DEFAULT '0',
    PRIMARY KEY (`id`) /*T![clustered_index] CLUSTERED */,
    KEY `idx_chats_uid` (`uid`),
    KEY `idx_chats_to_uid` (`to_uid`),
    KEY `idx_chats_create_time` (`create_time`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

