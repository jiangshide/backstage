-- MySQL dump 10.13  Distrib 5.7.18, for macos10.12 (x86_64)
--
-- Host: localhost    Database: backstage
-- ------------------------------------------------------
-- Server version	5.7.18

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `zd_app`
--

DROP TABLE IF EXISTS `zd_app`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_app` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '应用名称',
  `pkg` varchar(100) NOT NULL DEFAULT '' COMMENT '应用包名',
  `friend_id` varchar(50) NOT NULL DEFAULT '' COMMENT 'friendID',
  `channel` varchar(50) NOT NULL DEFAULT '' COMMENT '渠道名称',
  `version` varchar(10) NOT NULL DEFAULT '' COMMENT '版本',
  `version_code` int(11) NOT NULL DEFAULT '0' COMMENT '版本号',
  `environment` varchar(20) NOT NULL DEFAULT '' COMMENT '环境:production,preproduction,stage1,stage2,stage3',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态:1:打包中,2:打包成功,3:打包失败,4,无',
  `cmd` text NOT NULL COMMENT '打包脚本',
  `log` mediumtext NOT NULL COMMENT '打包日志详情',
  `err_msg` text NOT NULL COMMENT '错误详情',
  `qr_url` varchar(200) NOT NULL DEFAULT '' COMMENT '本地二维码:可扫描下载',
  `app_url` varchar(200) NOT NULL DEFAULT '' COMMENT 'app本地静态地址',
  `platform` tinyint(4) NOT NULL DEFAULT '1' COMMENT '平台:0~无限制, 1~android, 2~ios, 3~web',
  `create_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='App应用管理表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `zd_app_advert`
--

DROP TABLE IF EXISTS `zd_app_advert`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_app_advert` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '广告名称',
  `type` tinyint(4) NOT NULL DEFAULT '2' COMMENT '广告类型:1~开机广告,2~app启动广告,3～app第一安装导航广告,4~其它',
  `platform` tinyint(4) NOT NULL DEFAULT '1' COMMENT '平台:0~无限制, 1~android, 2~ios, 3~web',
  `price` int(11) NOT NULL DEFAULT '0' COMMENT '价格',
  `times` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '展示次数',
  `path` varchar(200) NOT NULL DEFAULT '' COMMENT '内容据对路径',
  `url` varchar(200) NOT NULL DEFAULT '' COMMENT '内容下载地址',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态:1:正常,2:下架,3:异常',
  `count` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '更新次数',
  `create_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='App启动管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `zd_app_channel`
--

DROP TABLE IF EXISTS `zd_app_channel`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_app_channel` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '应用名称',
  `pkg` varchar(100) NOT NULL DEFAULT '' COMMENT '应用包名',
  `friend_id` varchar(50) NOT NULL DEFAULT '' COMMENT 'friendID',
  `channel` varchar(50) NOT NULL DEFAULT '' COMMENT '渠道名称',
  `platform` tinyint(4) NOT NULL DEFAULT '1' COMMENT '平台:0~无限制, 1~android, 2~ios, 3~web',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态:1:正常,2:下架,3:异常',
  `create_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8mb4 COMMENT='App渠道管理表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `zd_app_stop`
--

DROP TABLE IF EXISTS `zd_app_stop`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_app_stop` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '名称',
  `channel` varchar(50) NOT NULL DEFAULT '' COMMENT '渠道名称',
  `version` varchar(20) NOT NULL DEFAULT '' COMMENT '版本',
  `url` varchar(200) NOT NULL DEFAULT '' COMMENT '下载地址',
  `platform` tinyint(4) NOT NULL DEFAULT '1' COMMENT '平台:0~无限制, 1~android, 2~ios, 3~web',
  `count` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '更新次数',
  `create_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='App停服管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `zd_app_update`
--

DROP TABLE IF EXISTS `zd_app_update`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_app_update` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '应用名称',
  `platform` tinyint(4) NOT NULL DEFAULT '1' COMMENT '平台:0~无限制, 1~android, 2~ios, 3~web',
  `channel` varchar(50) NOT NULL DEFAULT '' COMMENT '渠道名称',
  `version` varchar(20) NOT NULL DEFAULT '' COMMENT '版本',
  `content` varchar(200) NOT NULL DEFAULT '' COMMENT '更新内容提示',
  `url` varchar(200) NOT NULL DEFAULT '' COMMENT '下载地址',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '更新状态:0~普通提示更新,1~提示强制更新,2～后台自动下载后更新(非静默更新),3~静默更新',
  `count` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '更新次数',
  `create_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='App更新管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `zd_app_web_banner`
--

DROP TABLE IF EXISTS `zd_app_web_banner`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_app_web_banner` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '应用名称',
  `platform` tinyint(4) NOT NULL DEFAULT '1' COMMENT '平台:0~无限制, 1~android, 2~ios, 3~web',
  `channel` varchar(50) NOT NULL DEFAULT '' COMMENT '渠道名称',
  `version` varchar(20) NOT NULL DEFAULT '' COMMENT '版本',
  `content` varchar(200) NOT NULL DEFAULT '' COMMENT '更新内容提示',
  `url` varchar(200) NOT NULL DEFAULT '' COMMENT '下载地址',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '更新状态:0~普通提示更新,1~提示强制更新,2～后台自动下载后更新(非静默更新),3~静默更新',
  `count` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '更新次数',
  `create_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='App更新管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `zd_uc_auth`
--

DROP TABLE IF EXISTS `zd_uc_auth`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_uc_auth` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `pid` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '上级ID，0为顶级',
  `auth_name` varchar(64) NOT NULL DEFAULT '0' COMMENT '权限名称',
  `auth_url` varchar(255) NOT NULL DEFAULT '0' COMMENT 'URL地址',
  `sort` int(1) unsigned NOT NULL DEFAULT '999' COMMENT '排序，越小越前',
  `icon` varchar(255) NOT NULL,
  `is_show` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否显示，0-隐藏，1-显示',
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '操作者ID',
  `create_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态，1-正常，0-删除',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COMMENT='权限因子';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `zd_uc_role`
--

DROP TABLE IF EXISTS `zd_uc_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_uc_role` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `role_name` varchar(32) NOT NULL DEFAULT '0' COMMENT '角色名称',
  `detail` varchar(255) NOT NULL DEFAULT '0' COMMENT '备注',
  `create_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '修改这ID',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态1-正常，0-删除',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='角色表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `zd_uc_role_auth`
--

DROP TABLE IF EXISTS `zd_uc_role_auth`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_uc_role_auth` (
  `role_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
  `auth_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '权限ID',
  PRIMARY KEY (`role_id`,`auth_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='权限和角色关系表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `zd_uc_user`
--

DROP TABLE IF EXISTS `zd_uc_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_uc_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
  `role_ids` varchar(255) NOT NULL DEFAULT '0' COMMENT '角色id字符串，如：2,3,4',
  `salt` char(10) NOT NULL DEFAULT '' COMMENT '密码盐',
  `last_login` int(11) NOT NULL DEFAULT '0' COMMENT '最后登录时间',
  `last_ip` char(15) NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态，1-正常 0禁用',
  `create_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='管理员表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `zd_uc_user_address`
--

DROP TABLE IF EXISTS `zd_uc_user_address`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_uc_user_address` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `name` varchar(200) NOT NULL DEFAULT '' COMMENT '邮件地址',
  `code` int(11) NOT NULL DEFAULT '0' COMMENT '邮政编码',
  `create_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_address_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户地址管理表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `zd_uc_user_brand`
--

DROP TABLE IF EXISTS `zd_uc_user_brand`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_uc_user_brand` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `icon` varchar(200) NOT NULL DEFAULT '' COMMENT '银行logo',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '银行卡名称',
  `code` varchar(20) NOT NULL DEFAULT '0' COMMENT '银行卡卡号',
  `create_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_brand_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='银行卡账户管理表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `zd_uc_user_device`
--

DROP TABLE IF EXISTS `zd_uc_user_device`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_uc_user_device` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '使用设备名称:ios:android...',
  `mac` varchar(50) NOT NULL DEFAULT '' COMMENT '客户端MAC地址',
  `latitude` varchar(12) NOT NULL DEFAULT '' COMMENT '经度',
  `longitude` varchar(12) NOT NULL DEFAULT '' COMMENT '纬度',
  `mode` varchar(20) NOT NULL DEFAULT '' COMMENT '机型',
  `arch` varchar(10) NOT NULL DEFAULT '' COMMENT 'OS架构',
  `sdk_version` varchar(20) NOT NULL DEFAULT '' COMMENT '系统sdk版本',
  `app_version` varchar(10) NOT NULL DEFAULT '' COMMENT '客户端应用访问版本',
  `create_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_name` (`mac`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='银行卡账户管理表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `zd_uc_user_profile`
--

DROP TABLE IF EXISTS `zd_uc_user_profile`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_uc_user_profile` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `icon` varchar(200) NOT NULL DEFAULT '' COMMENT '用户头像',
  `level` tinyint(4) NOT NULL DEFAULT '0' COMMENT '级别:0~普通...',
  `score` int(11) NOT NULL DEFAULT '0' COMMENT '积分:',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '用户真实名',
  `phone` varchar(20) NOT NULL DEFAULT '0' COMMENT '手机号码',
  `motto` varchar(50) NOT NULL DEFAULT '' COMMENT '格言',
  `sex` tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别,0:保密,1:男,2:女',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
  `weixin` varchar(20) NOT NULL DEFAULT '' COMMENT '微信',
  `qq` varchar(20) NOT NULL DEFAULT '' COMMENT 'qq',
  `weibo` varchar(30) NOT NULL DEFAULT '' COMMENT '微博',
  `id_card` varchar(20) NOT NULL DEFAULT '' COMMENT '身份证',
  `id_card_pic_front` varchar(200) NOT NULL DEFAULT '' COMMENT '身份证前照',
  `id_card_pic_behind` varchar(200) NOT NULL DEFAULT '' COMMENT '身份证后照',
  `create_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户详情表';
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-08-20 14:03:08
