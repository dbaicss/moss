/*
 Navicat MySQL Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50710
 Source Host           : localhost:3306
 Source Schema         : moss

 Target Server Type    : MySQL
 Target Server Version : 50710
 File Encoding         : 65001

 Date: 11/11/2019 10:43:49
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for services
-- ----------------------------
DROP TABLE IF EXISTS `services`;
CREATE TABLE `services` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT '' COMMENT '服务名称',
  `namespace` varchar(50) DEFAULT '' COMMENT '所在空间',
  `description` varchar(200) DEFAULT '' COMMENT '服务描述',
  `owner_id` varchar(80) DEFAULT '' COMMENT '服务所有者(邮箱账号)',
  `port_tcp` varchar(200) DEFAULT '' COMMENT 'tcp端口',
  `port_udp` varchar(200) DEFAULT '' COMMENT 'ucp端口',
  `replicas` int(10) unsigned DEFAULT '1' COMMENT '实例个数',
  `limit_cpu` varchar(10) DEFAULT '' COMMENT 'Limits资源cpu',
  `limit_mem` varchar(10) DEFAULT '' COMMENT 'Limits资源mem',
  `request_cpu` varchar(10) DEFAULT '' COMMENT 'Request资源cpu',
  `request_mem` varchar(10) DEFAULT '' COMMENT 'Request资源mem',
  `status` tinyint(3) unsigned DEFAULT '0' COMMENT '状态 0-启用、1-禁用、2-删除',
  `creator` varchar(80) DEFAULT '' COMMENT '创建者(邮箱账号)',
  `created_time` datetime DEFAULT NULL COMMENT '创建时间',
  `modifier` varchar(50) DEFAULT '' COMMENT '修改者(邮箱账号)',
  `modified_time` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of services
-- ----------------------------
BEGIN;
INSERT INTO `services` VALUES (1, 'icx-blog', 'default', '碳云博客系统', 'huangyunqing@icarbonx.com', '{\"8080\":\"8080\",\"8090\":\"8000\"}', '', 1, '1', '2', '0.2', '0.5', 0, 'huangyunqing', '2019-11-11 18:21:56', 'huangyunqing', '2019-11-12 18:22:03');
INSERT INTO `services` VALUES (2, 'moderation-service', 'health-buddy', '测试项目', 'huangyunqing@icarbonx.com', '{\"8080\":\"8080\",\"8090\":\"8000\"}', '', 1, '1', '2', '0.2', '0.5', 0, 'huangyunqing', '2019-11-11 18:21:56', 'huangyunqing', '2019-11-12 18:22:03');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
