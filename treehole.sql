CREATE DATABASE IF NOT EXISTS `nku_treehole`;
USE `nku_treehole`;
DROP TABLE IF EXISTS `comments`,`posts`,`sessions`,`users`;
CREATE TABLE `comments`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `reply_to` int NOT NULL COMMENT '回复对象 post_id',
  `user_id` bigint NOT NULL COMMENT '用户ID',
  `content` varchar(10000) NOT NULL DEFAULT '' COMMENT '回复内容',
  `created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `deleted_at` datetime NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  INDEX `idx_replyto_post_id`(`reply_to`) USING BTREE
);

CREATE TABLE `posts`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL COMMENT '用户ID',
  `content` varchar(10000) NOT NULL DEFAULT '' COMMENT '帖子内容',
  `created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  `deleted_at` datetime NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  INDEX `idx_post_uid`(`user_id`) USING BTREE
);

CREATE TABLE `sessions`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL COMMENT '用户ID',
  `token` varchar(255) NOT NULL DEFAULT '' COMMENT 'token',
  `expired_at` datetime NOT NULL COMMENT '过期时间',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  INDEX `idx_token`(`token`) USING BTREE,
  INDEX `idx_user_token` (`user_id`) USING BTREE
);

CREATE TABLE `users`  (
  `id` bigint NOT NULL COMMENT '用户ID',
  `level` int unsigned NOT NULL COMMENT '用户权限等级 0--普通用户 1--普通管理员',
  `name` varchar(32) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '姓名',
  `sex` int DEFAULT NULL COMMENT '性别 0:男 1:女',
  `birthday` date DEFAULT NULL COMMENT '生日',
  `email` varchar(50) COLLATE utf8mb4_general_ci NOT NULL COMMENT '邮箱',
  `avatar` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '头像',
  `phone` varchar(32) COLLATE utf8mb4_general_ci NOT NULL COMMENT '注册手机号码',
  `password` varchar(50) COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_email` (`email`),
  KEY `idx_email` (`email`) USING BTREE,
  KEY `idx_phone` (`phone`) USING BTREE
);

