-- Database for nku_treehole
-- Set the default charset and collation for the database
CREATE DATABASE IF NOT EXISTS `nku_treehole` CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `nku_treehole`;

-- Create comments table
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT "ID",
  `reply_to` int NOT NULL COMMENT "Reply to post_id",
  `user_id` bigint NOT NULL COMMENT "User ID",
  `content` varchar(10000) NOT NULL DEFAULT '' COMMENT "Reply content",
  `created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT "Creation time",
  `deleted_at` datetime NULL COMMENT "Deletion time",
  PRIMARY KEY (`id`),
  INDEX `idx_replyto_post_id`(`reply_to`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT "Comments Table";

-- Create posts table
DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT "ID",
  `user_id` bigint NOT NULL COMMENT "User ID",
  `content` varchar(10000) NOT NULL DEFAULT '' COMMENT "Post content",
  `created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT "Creation time",
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT "Update time",
  `deleted_at` datetime NULL COMMENT "Deletion time",
  PRIMARY KEY (`id`),
  INDEX `idx_post_uid`(`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT "Posts Table";

-- Create sessions table
DROP TABLE IF EXISTS `sessions`;
CREATE TABLE `sessions` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT "ID",
  `user_id` bigint NOT NULL COMMENT "User ID",
  `token` varchar(255) NOT NULL DEFAULT '' COMMENT "Token",
  `expired_at` datetime NOT NULL COMMENT "Expiration time",
  `created_at` datetime NOT NULL COMMENT "Creation time",
  `deleted_at` datetime DEFAULT NULL COMMENT "Deletion time",
  PRIMARY KEY (`id`),
  INDEX `idx_token`(`token`) USING BTREE,
  INDEX `idx_user_token`(`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT "Sessions Table";


-- Create users table
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint NOT NULL COMMENT "User ID",
  `level` int unsigned NOT NULL COMMENT "User permission level 0--Regular User 1--Normal Administrator",
  `name` varchar(32) DEFAULT NULL COMMENT "Name",
  `sex` int DEFAULT NULL COMMENT "Gender 0: Male 1: Female",
  `birthday` date DEFAULT NULL COMMENT "Birthday",
  `email` varchar(50) NOT NULL COMMENT "Email",
  `avatar` varchar(255) DEFAULT NULL COMMENT "Avatar",
  `phone` varchar(32) NOT NULL COMMENT "Registered Phone Number",
  `password` varchar(50) NOT NULL COMMENT "Password",
  `created_at` datetime NOT NULL COMMENT "Creation time",
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT "Update time",
  `deleted_at` datetime DEFAULT NULL COMMENT "Deletion time",
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_email` (`email`),
  KEY `idx_email` (`email`) USING BTREE,
  KEY `idx_phone` (`phone`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT "Users Table";
