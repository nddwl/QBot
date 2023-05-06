USE q_bot;

CREATE TABLE `artwork`(
                          `id` BIGINT UNSIGNED AUTO_INCREMENT ,
                          `created_at` DATETIME DEFAULT NULL,
                          `updated_at` DATETIME DEFAULT NULL,
                          `deleted_at` DATETIME DEFAULT NULL,
                          `pid` VARCHAR(15) CHARSET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL UNIQUE ,
                          PRIMARY KEY (`id`)
)ENGINE = InnoDB DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;

CREATE TABLE `artwork_url`(
                               `id` BIGINT UNSIGNED AUTO_INCREMENT ,
                               `created_at` DATETIME DEFAULT NULL,
                               `updated_at` DATETIME DEFAULT NULL,
                               `deleted_at` DATETIME DEFAULT NULL,
                               `artwork_id` BIGINT UNSIGNED ,
                               `thumb_mini` VARCHAR(255) CHARSET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                               `small` VARCHAR(255)  CHARSET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                               `regular` VARCHAR(255)  CHARSET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                               `original` VARCHAR(255)  CHARSET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                               `cq_code_small` VARCHAR(255)  CHARSET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                               `cq_code_original` VARCHAR(255)  CHARSET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                               PRIMARY KEY (`id`),
                               KEY `artwork_id` (`artwork_id`) USING BTREE
)ENGINE = InnoDB DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;

CREATE TABLE `artwork_tag`(
                              `id` BIGINT UNSIGNED AUTO_INCREMENT ,
                              `created_at` DATETIME DEFAULT NULL,
                              `updated_at` DATETIME DEFAULT NULL,
                              `deleted_at` DATETIME DEFAULT NULL,
                              `tag` VARCHAR(50) CHARSET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL ,
                              `tag_translation` VARCHAR(50) CHARSET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                              PRIMARY KEY (`id`),
                              UNIQUE KEY `tag` (`tag`) USING BTREE
)ENGINE = InnoDB DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;

CREATE TABLE `artwork_tag_association`(
                                          `id` BIGINT UNSIGNED AUTO_INCREMENT ,
                                          `created_at` DATETIME DEFAULT NULL,
                                          `updated_at` DATETIME DEFAULT NULL,
                                          `deleted_at` DATETIME DEFAULT NULL,
                                          `artwork_id` BIGINT UNSIGNED NOT NULL ,
                                          `artwork_tag_id` BIGINT UNSIGNED NOT NULL,
                                          PRIMARY KEY (`id`),
                                          KEY `artwork_id` (artwork_id) USING BTREE,
                                          KEY `artwork_tag_id` (artwork_tag_id) USING BTREE
)ENGINE = InnoDB DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;

