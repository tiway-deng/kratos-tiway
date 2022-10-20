/*
 Navicat Premium Data Transfer

 Source Server         : local-docker
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : mysql:3306
 Source Schema         : kratos-tiway

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 20/10/2022 15:21:23
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户名',
  `nickname` varchar(150) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '昵称',
  `avatar` varchar(150) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '头像',
  `password` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '密码',
  `salt` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '加密盐',
  `email` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `mobile` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '手机号',
  `status` tinyint(0) NULL DEFAULT NULL COMMENT '状态  0：禁用   1：正常',
  `created_at` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `updated_at` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '更细时间',
  `deleted_at` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, '天天', 'tiwayD', '', '617533d3b5eb8aea9312d5bfd7e964fa4229177bc2ce2f4508476550ecbd4193', 'D1f8KM1c5kFmO1kA', '5011125@qq.com', '13536999999', 1, '2022-10-20 07:15:20', '2022-10-20 07:16:11', '0000-00-00 00:00:00');

SET FOREIGN_KEY_CHECKS = 1;
