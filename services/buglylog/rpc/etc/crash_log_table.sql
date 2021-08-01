/*
 Navicat Premium Data Transfer

 Source Server         : root
 Source Server Type    : MySQL
 Source Server Version : 80022
 Source Host           : localhost:3306
 Source Schema         : bugly

 Target Server Type    : MySQL
 Target Server Version : 80022
 File Encoding         : 65001

 Date: 01/08/2021 13:00:07
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for app_log
-- ----------------------------
DROP TABLE IF EXISTS `app_log`;
CREATE TABLE `app_log` (
  `id` int NOT NULL COMMENT '数据唯一 id',
  `mobile` varchar(255) DEFAULT NULL COMMENT '日志 id',
  `user_id` varchar(255) DEFAULT NULL,
  `log_url` varchar(255) DEFAULT NULL,
  `message` varchar(255) DEFAULT NULL,
  `log_create_time` timestamp NULL DEFAULT NULL COMMENT '日志上传的时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;
