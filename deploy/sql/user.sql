CREATE DATABASE IF NOT EXISTS `user` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE `user`;
CREATE TABLE `users` (
    `id` VARCHAR(24) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户ID，唯一标识',
    `avatar` VARCHAR(191) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户头像URL',
    `nickname` VARCHAR(191) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户昵称',
    `phone` VARCHAR(29) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户手机号',
    `password` VARCHAR(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户密码，哈希存储',
    `status` TINYINT COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户状态（0:禁用, 1:启用）',
    `gender` TINYINT COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '性别（0:未知, 1:男, 2:女）',
    `created_at` TIMESTAMP NULL DEFAULT NULL COMMENT '创建时间',
    `updated_at` TIMESTAMP NULL DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';
