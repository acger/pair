CREATE TABLE `chat_histories`
(
    `id`          bigint(20) NOT NULL /*T![auto_rand] AUTO_RANDOM(5) */,
    `create_time` datetime   DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime   DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `uid`         bigint(20) DEFAULT '0',
    `to_uid`      bigint(20) DEFAULT '0',
    PRIMARY KEY (`id`) /*T![clustered_index] CLUSTERED */,
    KEY `idx_chat_histories_uid` (`uid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

