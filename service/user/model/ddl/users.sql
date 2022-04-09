CREATE TABLE `users`
(
    `id`           bigint(20) NOT NULL /*T![auto_rand] AUTO_RANDOM(5) */,
    `account`      varchar(30) COLLATE utf8mb4_0900_ai_ci   DEFAULT '',
    `name`         varchar(30) COLLATE utf8mb4_0900_ai_ci   DEFAULT '',
    `avatar`       varchar(255) COLLATE utf8mb4_0900_ai_ci  DEFAULT '',
    `password`     varchar(255) COLLATE utf8mb4_0900_ai_ci  DEFAULT '',
    `gender`       varchar(10) COLLATE utf8mb4_0900_ai_ci   DEFAULT '',
    `status`       varchar(20) COLLATE utf8mb4_0900_ai_ci   DEFAULT '',
    `level`        varchar(20) COLLATE utf8mb4_0900_ai_ci   DEFAULT '',
    `birthday`     varchar(10) COLLATE utf8mb4_0900_ai_ci   DEFAULT '',
    `contact`      varchar(255) COLLATE utf8mb4_0900_ai_ci  DEFAULT '',
    `introduction` varchar(1000) COLLATE utf8mb4_0900_ai_ci DEFAULT '',
    `create_time`  datetime                                 DEFAULT CURRENT_TIMESTAMP,
    `update_time`  datetime                                 DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) /*T![clustered_index] CLUSTERED */,
    UNIQUE KEY `account` (`account`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;