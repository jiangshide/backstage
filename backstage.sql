/*
 Navicat MySQL Data Transfer

 Source Server         : localhost
 Source Server Version : 50712
 Source Host           : localhost
 Source Database       : pack

 Target Server Version : 50712
 File Encoding         : utf-8

 Date: 01/19/2018 23:57:42 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

DROP DATABASE IF EXISTS `zd112_pack`;
CREATE  DATABASE `zd112_pack`;

USE `zd112_pack`;

-- ----------------------------
--  Table structure for `zd_uc_user`
-- ----------------------------
DROP TABLE IF EXISTS `zd_uc_user`;
CREATE TABLE `zd_uc_user` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
  `role_ids` varchar(255) NOT NULL DEFAULT '0' COMMENT '角色id字符串，如：2,3,4',
  `salt` char(10) NOT NULL DEFAULT '' COMMENT '密码盐',
  `last_login` int(11) NOT NULL DEFAULT '0' COMMENT '最后登录时间',
  `last_ip` char(15) NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态，1-正常 0禁用',
  `create_id` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='管理员表';

-- ----------------------------
--  Records of `zd_uc_admin`
-- ----------------------------
BEGIN;
INSERT INTO `zd_uc_user` VALUES ('1','admin', 'c1875edcd37820e1346f0e3be812dd41', '0', 'e5Ps', '1515904905', '[', '1', '0', '0', '0', '1506128438');
COMMIT;

-- ----------------------------
--  Table structure for `zd_uc_auth`
-- ----------------------------
DROP TABLE IF EXISTS `zd_uc_auth`;
CREATE TABLE `zd_uc_auth` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `pid` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '上级ID，0为顶级',
  `auth_name` varchar(64) NOT NULL DEFAULT '0' COMMENT '权限名称',
  `auth_url` varchar(255) NOT NULL DEFAULT '0' COMMENT 'URL地址',
  `sort` int(1) unsigned NOT NULL DEFAULT '999' COMMENT '排序，越小越前',
  `icon` varchar(255) NOT NULL,
  `is_show` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否显示，0-隐藏，1-显示',
  `user_id` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '操作者ID',
  `create_id` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态，1-正常，0-删除',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=utf8mb4 COMMENT='权限因子';

-- ----------------------------
--  Records of `zd_uc_auth`
-- ----------------------------
BEGIN;
INSERT INTO `zd_uc_auth` VALUES ('1', '0', '所有权限', '/', '1', '', '0', '1', '1', '1', '1', '1505620970', '1505620970'), ('2', '1', '权限管理', '/', '999', 'fa-id-card', '1', '1', '0', '1', '1', '0', '1505622360'), ('3', '2', '管理员', '/admin/list', '1', 'fa-user-o', '1', '1', '1', '1', '1', '1505621186', '1505621186'), ('4', '2', '角色管理', '/role/list', '2', 'fa-user-circle-o', '1', '1', '0', '1', '1', '0', '1505621852'), ('5', '3', '新增', '/admin/add', '1', '', '0', '1', '0', '1', '1', '0', '1505621685'), ('6', '3', '修改', '/admin/edit', '2', '', '0', '1', '0', '1', '1', '0', '1505621697'), ('7', '3', '删除', '/admin/ajaxdel', '3', '', '0', '1', '1', '1', '1', '1505621756', '1505621756'), ('8', '4', '新增', '/role/add', '1', '', '1', '1', '0', '1', '1', '0', '1505698716'), ('9', '4', '修改', '/role/edit', '2', '', '0', '1', '1', '1', '1', '1505621912', '1505621912'), ('10', '4', '删除', '/role/ajaxdel', '3', '', '0', '1', '1', '1', '1', '1505621951', '1505621951'), ('11', '2', '权限因子', '/auth/list', '3', 'fa-list', '1', '1', '1', '1', '1', '1505621986', '1505621986'), ('12', '11', '新增', '/auth/add', '1', '', '0', '1', '1', '1', '1', '1505622009', '1505622009'), ('13', '11', '修改', '/auth/edit', '2', '', '0', '1', '1', '1', '1', '1505622047', '1505622047'), ('14', '11', '删除', '/auth/ajaxdel', '3', '', '0', '1', '1', '1', '1', '1505622111', '1505622111'), ('15', '1', '个人中心', 'profile/edit', '1001', 'fa-user-circle-o', '1', '1', '0', '1', '1', '0', '1506001114'),('16', '15', '资料修改', '/user/edit', '1', 'fa-edit', '1', '1', '0', '1', '1', '0', '1506057468');
COMMIT;

-- ----------------------------
--  Table structure for `zd_uc_role`
-- ----------------------------
DROP TABLE IF EXISTS `zd_uc_role`;
CREATE TABLE `zd_uc_role` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `role_name` varchar(32) NOT NULL DEFAULT '0' COMMENT '角色名称',
  `detail` varchar(255) NOT NULL DEFAULT '0' COMMENT '备注',
  `create_id` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '修改这ID',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态1-正常，0-删除',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='角色表';

-- ----------------------------
--  Records of `zd_uc_role`
-- ----------------------------
BEGIN;
INSERT INTO `zd_uc_role` VALUES ('2', '系统管理员', '系统管理员', '0', '0', '1', '1506124114', '1506124114');
COMMIT;

-- ----------------------------
--  Table structure for `zd_uc_role_auth`
-- ----------------------------
DROP TABLE IF EXISTS `zd_uc_role_auth`;
CREATE TABLE `zd_uc_role_auth` (
  `role_id` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '角色ID',
  `auth_id` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '权限ID',
  PRIMARY KEY (`role_id`,`auth_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='权限和角色关系表';

-- ----------------------------
--  Records of `zd_uc_role_auth`
-- ----------------------------
BEGIN;
INSERT INTO `zd_uc_role_auth` VALUES ('2', '0'), ('2', '1'), ('2', '15'), ('2', '20'), ('2', '21'), ('2', '22'), ('2', '23'), ('2', '24');
COMMIT;


DROP TABLE IF EXISTS `zd_uc_user_profile`; #用户详情
CREATE TABLE `zd_uc_user_profile` (
  `id`            BIGINT UNSIGNED  NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id`       BIGINT UNSIGNED  NOT NULL DEFAULT '0' COMMENT '用户ID',
  `icon`          VARCHAR(200)     NOT NULL DEFAULT ''  COMMENT '用户头像',
  `level`         TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '级别:0~普通...',
  `score`         INT              NOT NULL DEFAULT '0' COMMENT '积分:',
  `name`          VARCHAR(50)      NOT NULL DEFAULT ''  COMMENT '用户真实名',
  `phone`         VARCHAR(20)      NOT NULL DEFAULT '0' COMMENT '手机号码',
  `motto`         VARCHAR(50)      NOT NULL DEFAULT ''  COMMENT '格言',
  `sex`           TINYINT(1)       NOT NULL DEFAULT '0' COMMENT '性别,0:保密,1:男,2:女',
  `email`         VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '邮箱',
  `weixin`        VARCHAR(20)      NOT NULL DEFAULT '' COMMENT '微信',
  `qq`            VARCHAR(20)      NOT NULL DEFAULT '' COMMENT 'qq',
  `weibo`         VARCHAR(30)      NOT NULL DEFAULT '' COMMENT '微博',
  `id_card`         VARCHAR(20)      NOT NULL DEFAULT '' COMMENT '身份证',
  `id_card_pic_front` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '身份证前照',
  `id_card_pic_behind` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '身份证后照',
  `create_id`     BIGINT UNSIGNED  NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id`     BIGINT UNSIGNED  NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time`   INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time`   INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='用户详情表';

DROP TABLE IF EXISTS `zd_uc_user_device`; #银行卡账户
CREATE TABLE `zd_uc_user_device` (
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `user_id`     BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
  `name`      VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '使用设备名称:ios:android...',
  `mac`         VARCHAR(50)      NOT NULL DEFAULT '' COMMENT '客户端MAC地址',
  `latitude`    VARCHAR(12)      NOT NULL DEFAULT '' COMMENT '经度',
  `longitude`   VARCHAR(12)      NOT NULL DEFAULT '' COMMENT '纬度',
  `mode`        VARCHAR(20)      NOT NULL DEFAULT '' COMMENT '机型',
  `arch`        VARCHAR(10)      NOT NULL DEFAULT '' COMMENT 'OS架构',
  `sdk_version` VARCHAR(20)      NOT NULL DEFAULT '' COMMENT '系统sdk版本',
  `app_version` VARCHAR(10)      NOT NULL DEFAULT '' COMMENT '客户端应用访问版本',
  `create_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '创建者ID',
  `update_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '修改者ID',
  `create_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0' COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_name` (`mac`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='银行卡账户管理表';

DROP TABLE IF EXISTS `zd_uc_user_brand`; #银行卡账户
CREATE TABLE `zd_uc_user_brand` (
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `user_id`     BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
  `icon` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '银行logo',
  `name`        VARCHAR(50)     NOT NULL  DEFAULT '' COMMENT '银行卡名称',
  `code`        VARCHAR(20)              NOT NULL  DEFAULT '0' COMMENT '银行卡卡号',
  `create_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '创建者ID',
  `update_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '修改者ID',
  `create_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0' COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_brand_name` (`name`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='银行卡账户管理表';

DROP TABLE IF EXISTS `zd_uc_user_address`; #用户邮寄地址
CREATE TABLE `zd_uc_user_address` (
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `user_id`     BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '用户ID',
  `name`        VARCHAR(200)     NOT NULL  DEFAULT ''  COMMENT '邮件地址',
  `code`        INT              NOT NULL  DEFAULT '0' COMMENT '邮政编码',
  `create_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '创建者ID',
  `update_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '修改者ID',
  `create_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0' COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_address_name` (`name`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='用户地址管理表';


DROP TABLE IF EXISTS `zd_app`;#App应用管理
CREATE TABLE `zd_app` (
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `name`        VARCHAR(50)     NOT NULL  DEFAULT '' COMMENT '应用名称',
  `pkg`        VARCHAR(100)              NOT NULL  DEFAULT '' COMMENT '应用包名',
  `friend_id` VARCHAR(50) NOT NULL DEFAULT '' COMMENT 'friendID',
  `channel` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '渠道名称',
  `version` VARCHAR(10) NOT NULL DEFAULT '' COMMENT '版本',
  `version_code`       INT       NOT NULL DEFAULT '0' COMMENT '版本号',
  `environment` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '环境:production,preproduction,stage1,stage2,stage3',
  `status` TINYINT(4)       NOT NULL DEFAULT '1' COMMENT '状态:1:打包中,2:打包成功,3:打包失败,4,无',
  `cmd` text NOT NULL COMMENT '打包脚本',
  `log` MEDIUMTEXT NOT NULL COMMENT '打包日志详情',
  `err_msg` text NOT NULL COMMENT '错误详情',
  `qr_url` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '本地二维码:可扫描下载',
  `app_url` VARCHAR(200) NOT NULL DEFAULT '' COMMENT 'app本地静态地址',
  `platform` TINYINT(4)       NOT NULL DEFAULT '1' COMMENT '平台:0~无限制, 1~android, 2~ios, 3~web',
  `create_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '创建者ID',
  `update_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '修改者ID',
  `create_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0' COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='App应用管理表';

DROP TABLE IF EXISTS `zd_app_channel`;#App渠道管理
CREATE TABLE `zd_app_channel` (
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `name`        VARCHAR(50)     NOT NULL  DEFAULT '' COMMENT '应用名称',
  `pkg`        VARCHAR(100)              NOT NULL  DEFAULT '' COMMENT '应用包名',
  `friend_id` VARCHAR(50) NOT NULL DEFAULT '' COMMENT 'friendID',
  `channel` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '渠道名称',
  `platform` TINYINT(4)       NOT NULL DEFAULT '1' COMMENT '平台:0~无限制, 1~android, 2~ios, 3~web',
  `status` TINYINT(4)       NOT NULL DEFAULT '1' COMMENT '状态:1:正常,2:下架,3:异常',
  `create_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '创建者ID',
  `update_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '修改者ID',
  `create_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0' COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='App渠道管理表';

DROP TABLE IF EXISTS `zd_app_update`;#App更新管理
CREATE TABLE `zd_app_update` (
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `name`        VARCHAR(50)     NOT NULL  DEFAULT '' COMMENT '应用名称',
  `platform` TINYINT(4)       NOT NULL DEFAULT '1' COMMENT '平台:0~无限制, 1~android, 2~ios, 3~web',
  `channel` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '渠道名称',
  `version` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '版本',
  `content` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '更新内容提示',
  `url` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '下载地址',
  `status` TINYINT(4)       NOT NULL DEFAULT '1' COMMENT '更新状态:0~普通提示更新,1~提示强制更新,2～后台自动下载后更新(非静默更新),3~静默更新',
  `count`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '更新次数',
  `create_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '创建者ID',
  `update_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '修改者ID',
  `create_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0' COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='App更新管理';

DROP TABLE IF EXISTS `zd_app_stop`;#App停服管理
CREATE TABLE `zd_app_stop` (
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `name`        VARCHAR(50)     NOT NULL  DEFAULT '' COMMENT '应用名称',
  `channel` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '渠道名称',
  `version` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '版本',
  `url` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '下载地址',
  `platform` TINYINT(4)       NOT NULL DEFAULT '1' COMMENT '平台:0~无限制, 1~android, 2~ios, 3~web',
  `count`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '更新次数',
  `create_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '创建者ID',
  `update_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '修改者ID',
  `create_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0' COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='App停服管理';

DROP TABLE IF EXISTS `zd_app_advert`;#App广告管理
CREATE TABLE `zd_app_advert` (
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `name`        VARCHAR(50)     NOT NULL  DEFAULT '' COMMENT '广告名称',
  `type` TINYINT(4)       NOT NULL DEFAULT '2' COMMENT '广告类型:1~开机广告,2~app启动广告,3～app第一安装导航广告,4~其它',
  `platform` TINYINT(4)       NOT NULL DEFAULT '1' COMMENT '平台:0~无限制, 1~android, 2~ios, 3~web',
  `price` INT NOT NULL DEFAULT '0' COMMENT '价格',
  `times` INT UNSIGNED NOT NULL DEFAULT '0' COMMENT '展示次数',
  `path` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '内容据对路径',
  `url` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '内容下载地址',
  `status` TINYINT(4)       NOT NULL DEFAULT '1' COMMENT '状态:1:正常,2:下架,3:异常',
  `count`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '更新次数',
  `create_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '创建者ID',
  `update_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '修改者ID',
  `create_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0' COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='App启动管理';


DROP TABLE IF EXISTS `zd_app_web_banner`;#Web~Banner
CREATE TABLE `zd_app_web_banner`(
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `name`        VARCHAR(50)     NOT NULL  DEFAULT '' COMMENT '应用名称',
  `platform` TINYINT(4)       NOT NULL DEFAULT '1' COMMENT '平台:0~无限制, 1~android, 2~ios, 3~web',
  `channel` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '渠道名称',
  `version` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '版本',
  `content` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '更新内容提示',
  `url` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '下载地址',
  `status` TINYINT(4)       NOT NULL DEFAULT '1' COMMENT '更新状态:0~普通提示更新,1~提示强制更新,2～后台自动下载后更新(非静默更新),3~静默更新',
  `count`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '更新次数',
  `create_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '创建者ID',
  `update_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0' COMMENT '修改者ID',
  `create_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0' COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='App更新管理';

SET FOREIGN_KEY_CHECKS = 1;
