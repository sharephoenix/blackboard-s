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

 Date: 14/07/2021 22:15:00
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for crash_detail
-- ----------------------------
DROP TABLE IF EXISTS `crash_detail`;
CREATE TABLE `crash_detail` (
  `id_t` int NOT NULL AUTO_INCREMENT,
  `upload_log_num` int DEFAULT NULL,
  `crash_hash` varchar(255) DEFAULT NULL,
  `issue_id` varchar(255) DEFAULT NULL,
  `crash_id` varchar(255) DEFAULT NULL,
  `user_id` varchar(255) DEFAULT NULL,
  `device_id` varchar(255) DEFAULT NULL,
  `upload_time` varchar(255) DEFAULT NULL,
  `crash_time` varchar(255) DEFAULT NULL,
  `app_bundle_id` varchar(255) DEFAULT NULL,
  `app_version` varchar(255) DEFAULT NULL,
  `device_model` varchar(255) DEFAULT NULL,
  `system_version` varchar(255) DEFAULT NULL,
  `rom_detail` varchar(255) DEFAULT NULL,
  `cpu_architecture` varchar(255) DEFAULT NULL,
  `is_jump` varchar(255) DEFAULT NULL,
  `memory_size` varchar(255) DEFAULT NULL,
  `store_sizse` varchar(255) DEFAULT NULL,
  `sd_size` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id_t`),
  UNIQUE KEY `crash_id` (`crash_id`)
) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
