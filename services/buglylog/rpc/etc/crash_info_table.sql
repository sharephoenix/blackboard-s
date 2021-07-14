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

 Date: 14/07/2021 22:14:22
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for crash_info
-- ----------------------------
DROP TABLE IF EXISTS `crash_info`;
CREATE TABLE `crash_info` (
  `id_t` int NOT NULL AUTO_INCREMENT,
  `upload_log_num` int DEFAULT NULL,
  `issue_id` varchar(255) DEFAULT NULL,
  `error_type` varchar(255) DEFAULT NULL,
  `app_version` varchar(255) DEFAULT NULL,
  `app_name` varchar(255) DEFAULT NULL,
  `crash_times` varchar(255) DEFAULT NULL,
  `crash_device_num` varchar(255) DEFAULT NULL,
  `stack_flag` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `crash_description` varchar(255) DEFAULT NULL,
  `last_crash_time` varchar(255) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `process_person` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id_t`),
  UNIQUE KEY `issue_id` (`issue_id`)
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
