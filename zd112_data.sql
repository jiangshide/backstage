-- MySQL dump 10.13  Distrib 5.7.20, for osx10.12 (x86_64)
--
-- Host: localhost    Database: zd112
-- ------------------------------------------------------
-- Server version	5.7.20

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
-- Table structure for table `zd_api_detail`
--

DROP TABLE IF EXISTS `zd_api_detail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_api_detail` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `src_id` int(11) NOT NULL DEFAULT '0' COMMENT '主表ID',
  `method` tinyint(1) NOT NULL DEFAULT '1' COMMENT '方法名称,1:GET,2:POST,3:PUT,4:PATCH,5:DELETE',
  `name` varchar(100) NOT NULL DEFAULT '0' COMMENT '接口名称',
  `api_url` varchar(100) NOT NULL DEFAULT '0' COMMENT '接口地址',
  `protocol_type` varchar(20) NOT NULL DEFAULT '1' COMMENT '协议类型,1:http,2:https',
  `result` text COMMENT '返回结果,正确或者错误',
  `example` text COMMENT '接口示例',
  `detail` varchar(1000) NOT NULL DEFAULT '0' COMMENT '注意事项',
  `audit_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '审核时间',
  `audit_id` int(11) NOT NULL DEFAULT '0' COMMENT '审核人ID',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态,0:暂停 使用,1:正在审核,2:审核通过',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_main_id` (`src_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='API附表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_api_detail`
--

LOCK TABLES `zd_api_detail` WRITE;
/*!40000 ALTER TABLE `zd_api_detail` DISABLE KEYS */;
/*!40000 ALTER TABLE `zd_api_detail` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_api_key`
--

DROP TABLE IF EXISTS `zd_api_key`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_api_key` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '申请名称',
  `key` varchar(500) NOT NULL DEFAULT '' COMMENT '申请的key',
  `descript` text COMMENT '描述',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `views` int(11) NOT NULL DEFAULT '0' COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `key` (`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='api申请key管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_api_key`
--

LOCK TABLES `zd_api_key` WRITE;
/*!40000 ALTER TABLE `zd_api_key` DISABLE KEYS */;
/*!40000 ALTER TABLE `zd_api_key` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_api_param`
--

DROP TABLE IF EXISTS `zd_api_param`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_api_param` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `detail_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '附表ID',
  `api_key` varchar(100) NOT NULL DEFAULT '0' COMMENT '参数名',
  `api_type` varchar(100) NOT NULL DEFAULT '0' COMMENT '类型',
  `api_value` varchar(500) NOT NULL DEFAULT '0' COMMENT '参数值',
  `api_detail` varchar(500) NOT NULL DEFAULT '0' COMMENT '参数说明',
  `is_null` varchar(10) NOT NULL DEFAULT '0' COMMENT '是否必填',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态,1:正常,0:删除',
  `create_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_main_id` (`detail_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='API参数表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_api_param`
--

LOCK TABLES `zd_api_param` WRITE;
/*!40000 ALTER TABLE `zd_api_param` DISABLE KEYS */;
/*!40000 ALTER TABLE `zd_api_param` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_api_src`
--

DROP TABLE IF EXISTS `zd_api_src`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_api_src` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `group_id` int(11) NOT NULL DEFAULT '0' COMMENT '分组ID',
  `name` varchar(50) NOT NULL DEFAULT '0' COMMENT '接口名称',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态,1:审核通过,0:暂停使用,2:草稿,3:审核中',
  `audit_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '审核人ID',
  `create_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新者ID',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `audit_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '审核时间',
  PRIMARY KEY (`id`),
  KEY `idx_group_id` (`group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='API主表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_api_src`
--

LOCK TABLES `zd_api_src` WRITE;
/*!40000 ALTER TABLE `zd_api_src` DISABLE KEYS */;
/*!40000 ALTER TABLE `zd_api_src` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_app`
--

DROP TABLE IF EXISTS `zd_app`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_app` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `project_id` int(11) NOT NULL DEFAULT '0' COMMENT '项目ID',
  `test_id` int(11) NOT NULL DEFAULT '0' COMMENT '测试ID',
  `icon` varchar(100) DEFAULT '' COMMENT 'Logo',
  `type_id` int(11) NOT NULL DEFAULT '0' COMMENT '应用平台类型ID',
  `application_id` int(11) NOT NULL DEFAULT '0' COMMENT '应用名称ID',
  `pkg_id` int(11) NOT NULL DEFAULT '0' COMMENT '应用包ID',
  `version_id` int(11) NOT NULL DEFAULT '0' COMMENT '应用版本ID',
  `code_id` int(11) NOT NULL DEFAULT '0' COMMENT '应用版本号ID',
  `env_id` int(11) NOT NULL DEFAULT '0' COMMENT '应用环境ID',
  `build_id` int(11) NOT NULL DEFAULT '0' COMMENT '构建类型ID',
  `channel_id` int(11) NOT NULL DEFAULT '0' COMMENT '渠道ID',
  `descript` text COMMENT '描述',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '当前状态:0~无,-1~打包失败,1~打包中,2~打包成功,-3~测试失败,3~测试中,4~测试完成',
  `times` int(11) NOT NULL DEFAULT '0' COMMENT '打包次数',
  `url` varchar(100) NOT NULL DEFAULT '' COMMENT '下载地址',
  `qr_img` varchar(100) NOT NULL DEFAULT '' COMMENT '二维码地址',
  `downs` int(11) NOT NULL DEFAULT '0' COMMENT '下载次数',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `views` int(11) NOT NULL DEFAULT '0' COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`),
  UNIQUE KEY `pkg_id` (`pkg_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Android应用包管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_app`
--

LOCK TABLES `zd_app` WRITE;
/*!40000 ALTER TABLE `zd_app` DISABLE KEYS */;
/*!40000 ALTER TABLE `zd_app` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_app_build`
--

DROP TABLE IF EXISTS `zd_app_build`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_app_build` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(10) NOT NULL DEFAULT '' COMMENT '构建类型:debug...',
  `descript` text COMMENT '描述',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `views` int(11) NOT NULL DEFAULT '0' COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='应用构建类型管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_app_build`
--

LOCK TABLES `zd_app_build` WRITE;
/*!40000 ALTER TABLE `zd_app_build` DISABLE KEYS */;
/*!40000 ALTER TABLE `zd_app_build` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_app_channer`
--

DROP TABLE IF EXISTS `zd_app_channer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_app_channer` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '渠道名称',
  `friend_id` varchar(30) NOT NULL DEFAULT '' COMMENT '关联ID',
  `drescript` text COMMENT '描述',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `views` int(11) NOT NULL DEFAULT '0' COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `friend_id` (`friend_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='APP渠道管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_app_channer`
--

LOCK TABLES `zd_app_channer` WRITE;
/*!40000 ALTER TABLE `zd_app_channer` DISABLE KEYS */;
/*!40000 ALTER TABLE `zd_app_channer` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_app_code`
--

DROP TABLE IF EXISTS `zd_app_code`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_app_code` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `code` int(11) NOT NULL DEFAULT '0' COMMENT '应用版本号',
  `descript` text COMMENT '描述',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `views` int(11) NOT NULL DEFAULT '0' COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='应用版本号管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_app_code`
--

LOCK TABLES `zd_app_code` WRITE;
/*!40000 ALTER TABLE `zd_app_code` DISABLE KEYS */;
/*!40000 ALTER TABLE `zd_app_code` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_app_env`
--

DROP TABLE IF EXISTS `zd_app_env`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_app_env` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '应用环境',
  `descript` text COMMENT '描述',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `views` int(11) NOT NULL DEFAULT '0' COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='应用环境管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_app_env`
--

LOCK TABLES `zd_app_env` WRITE;
/*!40000 ALTER TABLE `zd_app_env` DISABLE KEYS */;
/*!40000 ALTER TABLE `zd_app_env` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_app_name`
--

DROP TABLE IF EXISTS `zd_app_name`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_app_name` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '应用名称',
  `descript` text COMMENT '描述',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `views` int(11) NOT NULL DEFAULT '0' COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='应用名称管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_app_name`
--

LOCK TABLES `zd_app_name` WRITE;
/*!40000 ALTER TABLE `zd_app_name` DISABLE KEYS */;
/*!40000 ALTER TABLE `zd_app_name` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_app_pkgs`
--

DROP TABLE IF EXISTS `zd_app_pkgs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_app_pkgs` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `NAME` varchar(50) NOT NULL DEFAULT '' COMMENT '应用包',
  `descript` text COMMENT '描述',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `views` int(11) NOT NULL DEFAULT '0' COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`),
  UNIQUE KEY `NAME` (`NAME`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='应用包管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_app_pkgs`
--

LOCK TABLES `zd_app_pkgs` WRITE;
/*!40000 ALTER TABLE `zd_app_pkgs` DISABLE KEYS */;
/*!40000 ALTER TABLE `zd_app_pkgs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_app_type`
--

DROP TABLE IF EXISTS `zd_app_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_app_type` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '应用平台类型,如:android',
  `descript` text COMMENT '描述',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `views` int(11) NOT NULL DEFAULT '0' COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='应用类型管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_app_type`
--

LOCK TABLES `zd_app_type` WRITE;
/*!40000 ALTER TABLE `zd_app_type` DISABLE KEYS */;
/*!40000 ALTER TABLE `zd_app_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_app_version`
--

DROP TABLE IF EXISTS `zd_app_version`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_app_version` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '应用版本',
  `descript` text COMMENT '描述',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `views` int(11) NOT NULL DEFAULT '0' COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='应用版本管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_app_version`
--

LOCK TABLES `zd_app_version` WRITE;
/*!40000 ALTER TABLE `zd_app_version` DISABLE KEYS */;
/*!40000 ALTER TABLE `zd_app_version` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_area_city`
--

DROP TABLE IF EXISTS `zd_area_city`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_area_city` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '所属省',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '市名',
  `icon` varchar(100) DEFAULT '' COMMENT 'LOGO',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='所属市';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_area_city`
--

LOCK TABLES `zd_area_city` WRITE;
/*!40000 ALTER TABLE `zd_area_city` DISABLE KEYS */;
INSERT INTO `zd_area_city` VALUES (1,1,'dad','/static/upload/jpg/2963b57e46f5ce04de82e55fff81b5a9_1512640332.jpg',1,0,1512640332,0);
/*!40000 ALTER TABLE `zd_area_city` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_area_continent`
--

DROP TABLE IF EXISTS `zd_area_continent`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_area_continent` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '洲名',
  `icon` varchar(100) DEFAULT '' COMMENT 'LOGO',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COMMENT='所属洲';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_area_continent`
--

LOCK TABLES `zd_area_continent` WRITE;
/*!40000 ALTER TABLE `zd_area_continent` DISABLE KEYS */;
INSERT INTO `zd_area_continent` VALUES (9,'亚洲','/static/upload/png/efd602163ac41e0753d391e66e1a1e57_1513071656.png',0,1,0,1513071668),(10,'dada','/static/upload/png/066d5be101db6e14cc91658aecb85d3f_1513071719.png',1,0,1513071724,0);
/*!40000 ALTER TABLE `zd_area_continent` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_area_country`
--

DROP TABLE IF EXISTS `zd_area_country`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_area_country` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '所属镇',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '乡名',
  `icon` varchar(100) DEFAULT '' COMMENT 'LOGO',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='所属乡';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_area_country`
--

LOCK TABLES `zd_area_country` WRITE;
/*!40000 ALTER TABLE `zd_area_country` DISABLE KEYS */;
/*!40000 ALTER TABLE `zd_area_country` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_area_county`
--

DROP TABLE IF EXISTS `zd_area_county`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_area_county` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '所属地区',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '县名',
  `icon` varchar(100) DEFAULT '' COMMENT 'LOGO',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='所属县';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_area_county`
--

LOCK TABLES `zd_area_county` WRITE;
/*!40000 ALTER TABLE `zd_area_county` DISABLE KEYS */;
/*!40000 ALTER TABLE `zd_area_county` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_area_group`
--

DROP TABLE IF EXISTS `zd_area_group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_area_group` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '所属村',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '组名',
  `icon` varchar(100) DEFAULT '' COMMENT 'LOGO',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='所属组';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_area_group`
--

LOCK TABLES `zd_area_group` WRITE;
/*!40000 ALTER TABLE `zd_area_group` DISABLE KEYS */;
/*!40000 ALTER TABLE `zd_area_group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_area_province`
--

DROP TABLE IF EXISTS `zd_area_province`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_area_province` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '所属国家',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '省名',
  `icon` varchar(100) DEFAULT '' COMMENT 'LOGO',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='所属省';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_area_province`
--

LOCK TABLES `zd_area_province` WRITE;
/*!40000 ALTER TABLE `zd_area_province` DISABLE KEYS */;
INSERT INTO `zd_area_province` VALUES (1,8,'dasa','/static/upload/jpg/595aab7ef0554e07faed7c779912f986_1512546676.jpg',1,1,1512546677,1512550208);
/*!40000 ALTER TABLE `zd_area_province` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_area_region`
--

DROP TABLE IF EXISTS `zd_area_region`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_area_region` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '所属城市',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '地区名',
  `icon` varchar(100) DEFAULT '' COMMENT 'LOGO',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='所属地区';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_area_region`
--

LOCK TABLES `zd_area_region` WRITE;
/*!40000 ALTER TABLE `zd_area_region` DISABLE KEYS */;
INSERT INTO `zd_area_region` VALUES (1,1,'sdaada','/static/upload/jpg/df128977cb3ea1d6a04777f2eb1c2949_1512640356.jpg',1,0,1512640357,0);
/*!40000 ALTER TABLE `zd_area_region` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_area_state`
--

DROP TABLE IF EXISTS `zd_area_state`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_area_state` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '所属洲',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '国名',
  `icon` varchar(100) DEFAULT '' COMMENT 'LOGO',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COMMENT='所属国家';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_area_state`
--

LOCK TABLES `zd_area_state` WRITE;
/*!40000 ALTER TABLE `zd_area_state` DISABLE KEYS */;
INSERT INTO `zd_area_state` VALUES (7,7,'dada','/static/upload/jpg/1db9e1b87f52ed597e3d86958ab39022_1512543243.jpg',1,1,1512543243,1512548329),(8,5,'asdc','/static/upload/jpg/cabdad9d9e91925ca1af1c72a3ac9b6e_1512544000.jpg',1,1,1512544001,1512550184);
/*!40000 ALTER TABLE `zd_area_state` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_area_team`
--

DROP TABLE IF EXISTS `zd_area_team`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_area_team` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '所属组',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '队名',
  `icon` varchar(100) DEFAULT '' COMMENT 'LOGO',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='所属队';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_area_team`
--

LOCK TABLES `zd_area_team` WRITE;
/*!40000 ALTER TABLE `zd_area_team` DISABLE KEYS */;
/*!40000 ALTER TABLE `zd_area_team` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_area_town`
--

DROP TABLE IF EXISTS `zd_area_town`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_area_town` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '所属县',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '镇名',
  `icon` varchar(100) DEFAULT '' COMMENT 'LOGO',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='所属镇';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_area_town`
--

LOCK TABLES `zd_area_town` WRITE;
/*!40000 ALTER TABLE `zd_area_town` DISABLE KEYS */;
/*!40000 ALTER TABLE `zd_area_town` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_area_village`
--

DROP TABLE IF EXISTS `zd_area_village`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_area_village` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '所属乡',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '村名',
  `icon` varchar(100) DEFAULT '' COMMENT 'LOGO',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='所属村';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_area_village`
--

LOCK TABLES `zd_area_village` WRITE;
/*!40000 ALTER TABLE `zd_area_village` DISABLE KEYS */;
/*!40000 ALTER TABLE `zd_area_village` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_nation`
--

DROP TABLE IF EXISTS `zd_nation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_nation` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(25) NOT NULL DEFAULT '' COMMENT '名称',
  `icon` varchar(100) NOT NULL DEFAULT '' COMMENT 'LOGO',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=57 DEFAULT CHARSET=utf8mb4 COMMENT='民族名称';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_nation`
--

LOCK TABLES `zd_nation` WRITE;
/*!40000 ALTER TABLE `zd_nation` DISABLE KEYS */;
INSERT INTO `zd_nation` VALUES (1,'阿昌族','/static/upload/img/f1e5b6df6871b0d58f417f510de0f75b.jpg',1,0,1512023950,0),(2,'白族','/static/upload/img/4aa1e2207a90b1a6698d4b173acb5860.jpg',1,0,1512023950,0),(3,'保安族','/static/upload/img/3fe8ddb04353b0a76e0aa08436631104.jpg',1,0,1512023950,0),(4,'布朗族','/static/upload/img/aae92384871153617f6dc2d361ab3214.jpg',1,0,1512023950,0),(5,'布依族','/static/upload/img/da6fd3e096782a03306711f6f9d3ce9d.jpg',1,0,1512023950,0),(6,'藏族','/static/upload/img/b871b354c91b5cac4c24f49982657790.jpg',1,0,1512023950,0),(7,'朝鲜族','/static/upload/img/01e5500c79e1d97947bf0ee1365c0858.jpg',1,0,1512023950,0),(8,'达翰尔族','/static/upload/img/1f2c585705ae27829f5a02dc6a609515.jpg',1,0,1512023950,0),(9,'傣族','/static/upload/img/4138e4310c568175f4ef84f861cd5f8d.jpg',1,0,1512023950,0),(10,'昂德族','/static/upload/img/3c90ac2109fc60dd05cbd6a5762acabf.jpg',1,0,1512023951,0),(11,'东乡族','/static/upload/img/098b644d4766f5c5265371a0d16b957c.jpg',1,0,1512023951,0),(12,'侗族','/static/upload/img/8ecec961ee5e60610bc4b5ea03c96e8a.jpg',1,0,1512023951,0),(13,'独龙族','/static/upload/img/0c6178babfc520f3154053143551260b.jpg',1,0,1512023951,0),(14,'俄罗斯族','/static/upload/img/9251b1a157bdffea643d27ca8bd9a8da.jpg',1,0,1512023951,0),(15,'鄂伦春族','/static/upload/img/f0ba330349b722fa964d3cb0729118be.jpg',1,0,1512023951,0),(16,'鄂温克族','/static/upload/img/71acd17ef91189c0e0030ee5166c420e.jpg',1,0,1512023951,0),(17,'高山族','/static/upload/img/697b01a0c3a21960b9a56176b1f2cb5c.jpg',1,0,1512023951,0),(18,'哈尼族','/static/upload/img/c0e6291f694c418de8c55dc171c92351.jpg',1,0,1512023951,0),(19,'哈萨克族','/static/upload/img/3b893a759d9eca61b52094e44e3add7c.jpg',1,0,1512023951,0),(20,'汉族','/static/upload/img/ae90992346cad02e18075a3116cd57e9.jpg',1,0,1512023951,0),(21,'赫哲族','/static/upload/img/ab77ba517ffe8e016316cbf83266d68a.jpg',1,0,1512023951,0),(22,'回族','/static/upload/img/8078bcde76b66133874d7c59c11c6d5f.jpg',1,0,1512023951,0),(23,'基诺族','/static/upload/img/4ee726f0c777ee5df1271d6a9d28e73f.jpg',1,0,1512023951,0),(24,'京族','/static/upload/img/e20e66b23234ee7c0bed843fede41858.jpg',1,0,1512023951,0),(25,'景颇族','/static/upload/img/8cb4f0889580ec7e4a527a352237772e.jpg',1,0,1512023951,0),(26,'柯尔克孜族','/static/upload/img/f82c328095cad09a97b71fd742415a50.jpg',1,0,1512023951,0),(27,'拉祜族','/static/upload/img/47d5ab9ff041ada9a753ea4c2e78a012.jpg',1,0,1512023952,0),(28,'黎族','/static/upload/img/3ef820d7dcd0382b20a65989bbf17cc4.jpg',1,0,1512023952,0),(29,'傈僳族','/static/upload/img/e3011d8ed0db68fe4474399a23f0a2f8.jpg',1,0,1512023952,0),(30,'珞巴族','/static/upload/img/54728b95cf4f94216c5436fac9584402.jpg',1,0,1512023952,0),(31,'满族','/static/upload/img/ee017da4b58cad24ccbb9f27405a29aa.jpg',1,0,1512023952,0),(32,'毛南族','/static/upload/img/895636a8a452d5f19a79bb55b627e60a.jpg',1,0,1512023952,0),(33,'门巴族','/static/upload/img/36f2f61f3cc8ef4da0aaaa4746166460.jpg',1,0,1512023952,0),(34,'蒙古族','/static/upload/img/5b45e72079f144de753a73f84b8a9171.jpg',1,0,1512023952,0),(35,'苗族','/static/upload/img/cbf109f5edfe3a569648a2d5a248ab88.jpg',1,0,1512023953,0),(36,'仫佬族','/static/upload/img/001cca129549444d424bfa3d528f4d74.jpg',1,0,1512023953,0),(37,'纳西族','/static/upload/img/5813658cf74a18cdbe745a7a81892830.jpg',1,0,1512023953,0),(38,'怒族','/static/upload/img/3f4274841885d431c54c518c21502214.jpg',1,0,1512023953,0),(39,'普米族','/static/upload/img/edd6bfc005381e63e3b94f74a4ab17c9.jpg',1,0,1512023954,0),(40,'羌族','/static/upload/img/1fc98f55a64c6ef6f21a02e9022dfb80.jpg',1,0,1512023954,0),(41,'撒拉族','/static/upload/img/e6a53264115153f2b3dabcfb8cc33f65.jpg',1,0,1512023954,0),(42,'畲族','/static/upload/img/c3d125c08a35027a19644dca9bbba598.jpg',1,0,1512023954,0),(43,'水族','/static/upload/img/2c4376f2622ce88886a36749acdd436b.jpg',1,0,1512023954,0),(44,'塔吉克族','/static/upload/img/9283a9b1c46cd05694bea8c79c03e564.jpg',1,0,1512023955,0),(45,'塔塔尔族','/static/upload/img/90988b13924c6abb18a2cd8259472cf6.jpg',1,0,1512023955,0),(46,'土家族','/static/upload/img/ffe807fc5a4d18ae9cb0f55764ef3c78.jpg',1,0,1512023955,0),(47,'图族','/static/upload/img/01a2dfd9e53c7c34abc3a85d5a4d35ae.jpg',1,0,1512023955,0),(48,'佤族','/static/upload/img/67ea573f3cc281c7a408313b266f4959.jpg',1,0,1512023955,0),(49,'维吾尔族','/static/upload/img/698af781f1a7c2ceb821a6729c7ec1f2.jpg',1,0,1512023955,0),(50,'乌孜别克族','/static/upload/img/3d70efc8460fe3fc8e211ece04b48adf.jpg',1,0,1512023955,0),(51,'锡伯族','/static/upload/img/6a3aaaf558ff1bb83d136e525539031f.jpg',1,1,1512023955,1512368221),(52,'瑶族','/static/upload/img/fe3ea6602bd00c8227c5faa2875d21fc.jpg',1,0,1512023955,0),(53,'彝族','/static/upload/img/5e33f33d61748be897a1f1769594f2cb.jpg',1,0,1512023955,0),(54,'仡佬族','/static/upload/img/614489e11cfe52ed57b3fe7d4de22a17.jpg',1,0,1512023955,0),(55,'裕固族','/static/upload/img/4f5f9f7f12074583dbee53339d45a6d6.jpg',1,0,1512023955,0),(56,'壮族','/static/upload/img/07069875e0b7bd68022e24774b0c4dca.jpg',1,0,1512023956,0);
/*!40000 ALTER TABLE `zd_nation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_set_code`
--

DROP TABLE IF EXISTS `zd_set_code`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_set_code` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `code` varchar(50) NOT NULL DEFAULT '0' COMMENT '状态码',
  `descript` varchar(255) NOT NULL DEFAULT '0' COMMENT '描述',
  `detail` varchar(255) NOT NULL DEFAULT '0' COMMENT '备注',
  `status` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '状态,1:正常,0:禁用',
  `create_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_env_name` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='状态代码表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_set_code`
--

LOCK TABLES `zd_set_code` WRITE;
/*!40000 ALTER TABLE `zd_set_code` DISABLE KEYS */;
/*!40000 ALTER TABLE `zd_set_code` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_set_evn`
--

DROP TABLE IF EXISTS `zd_set_evn`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_set_evn` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '环境名称',
  `env_host` varchar(255) NOT NULL DEFAULT '0' COMMENT '主机',
  `detail` varchar(255) NOT NULL DEFAULT '0' COMMENT '备注',
  `status` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '状态,1:正常,0:禁用',
  `create_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_env_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='环境分组表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_set_evn`
--

LOCK TABLES `zd_set_evn` WRITE;
/*!40000 ALTER TABLE `zd_set_evn` DISABLE KEYS */;
/*!40000 ALTER TABLE `zd_set_evn` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_set_group`
--

DROP TABLE IF EXISTS `zd_set_group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_set_group` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '组名',
  `detail` varchar(255) NOT NULL DEFAULT '' COMMENT '说明',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态:1:正常,0:删除',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '更新者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_create_id` (`create_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_set_group`
--

LOCK TABLES `zd_set_group` WRITE;
/*!40000 ALTER TABLE `zd_set_group` DISABLE KEYS */;
/*!40000 ALTER TABLE `zd_set_group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_test_environment`
--

DROP TABLE IF EXISTS `zd_test_environment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_test_environment` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '环境类型',
  `jdk` varchar(100) NOT NULL DEFAULT '' COMMENT 'jdk安装路径',
  `git` varchar(100) NOT NULL DEFAULT '' COMMENT 'Git安装路径',
  `gradle` varchar(100) NOT NULL DEFAULT '' COMMENT 'Gradle安装路径',
  `adb` varchar(100) NOT NULL DEFAULT '' COMMENT 'Adb安装路径',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `views` int(11) NOT NULL DEFAULT '0' COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='系统环境管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_test_environment`
--

LOCK TABLES `zd_test_environment` WRITE;
/*!40000 ALTER TABLE `zd_test_environment` DISABLE KEYS */;
INSERT INTO `zd_test_environment` VALUES (1,'系统环境','/usr/bin/java\n','/usr/local/bin/git\n','/Users/etongdai/developer/gradle/gradle-4.1/bin/gradle\n','/Users/etongdai/Library/Android/sdk/platform-tools/adb\n',0,0,0,0,0);
/*!40000 ALTER TABLE `zd_test_environment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_test_project`
--

DROP TABLE IF EXISTS `zd_test_project`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_test_project` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '系统环境ID',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '项目名称',
  `icon` varchar(100) DEFAULT '' COMMENT 'Logo',
  `address` text COMMENT '项目地址:git,svn,...',
  `account` varchar(30) NOT NULL DEFAULT '' COMMENT '账号名称',
  `psw` varchar(50) NOT NULL DEFAULT '' COMMENT '账号密码',
  `branch` varchar(50) NOT NULL DEFAULT '' COMMENT '分支',
  `tag` varchar(50) NOT NULL DEFAULT '' COMMENT '标签tag',
  `descript` text COMMENT '描述',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `views` int(11) NOT NULL DEFAULT '0' COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='项目工程管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_test_project`
--

LOCK TABLES `zd_test_project` WRITE;
/*!40000 ALTER TABLE `zd_test_project` DISABLE KEYS */;
/*!40000 ALTER TABLE `zd_test_project` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_test_test`
--

DROP TABLE IF EXISTS `zd_test_test`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_test_test` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '测试名称',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '测试类型',
  `descript` text COMMENT '描述',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `views` int(11) NOT NULL DEFAULT '0' COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='测试状态管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_test_test`
--

LOCK TABLES `zd_test_test` WRITE;
/*!40000 ALTER TABLE `zd_test_test` DISABLE KEYS */;
/*!40000 ALTER TABLE `zd_test_test` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_tools_compress`
--

DROP TABLE IF EXISTS `zd_tools_compress`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_tools_compress` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '文件名称',
  `url` text COMMENT '下载地址',
  `type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '0:图片,1:文件',
  `format` varchar(10) NOT NULL DEFAULT '' COMMENT '以文件后缀为准',
  `descript` varchar(500) NOT NULL DEFAULT '0' COMMENT '文件描述',
  `size` int(11) NOT NULL DEFAULT '0' COMMENT '文件原大小',
  `re_size` int(11) NOT NULL DEFAULT '0' COMMENT '文件压缩后大小',
  `compress` int(11) NOT NULL DEFAULT '0' COMMENT '压缩次数',
  `downs` int(11) NOT NULL DEFAULT '0' COMMENT '下载次数',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `views` int(11) NOT NULL DEFAULT '0' COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文件压缩管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_tools_compress`
--

LOCK TABLES `zd_tools_compress` WRITE;
/*!40000 ALTER TABLE `zd_tools_compress` DISABLE KEYS */;
/*!40000 ALTER TABLE `zd_tools_compress` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_tools_format`
--

DROP TABLE IF EXISTS `zd_tools_format`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_tools_format` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '所格式类型',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '格式民称',
  `descript` text COMMENT '格式描述',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=129 DEFAULT CHARSET=utf8mb4 COMMENT='文件格式管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_tools_format`
--

LOCK TABLES `zd_tools_format` WRITE;
/*!40000 ALTER TABLE `zd_tools_format` DISABLE KEYS */;
INSERT INTO `zd_tools_format` VALUES (1,1,'','',1,0,1512694558,0),(2,1,'mpeg','',1,0,1512694558,0),(3,1,'mpeg-1','',1,0,1512694558,0),(4,1,'mpeg1','',1,0,1512694558,0),(5,1,'mpeg-2','',1,0,1512694558,0),(6,1,'mpeg2','',1,0,1512694558,0),(7,1,'mpeg-4','',1,0,1512694558,0),(8,1,'mpeg4','',1,0,1512694558,0),(9,1,'mpg','',1,0,1512694558,0),(10,1,'mpe','',1,0,1512694558,0),(11,1,'dat','',1,0,1512694558,0),(12,1,'mp4','',1,0,1512694558,0),(13,1,'m4v','',1,0,1512694558,0),(14,1,'3gp','',1,0,1512694558,0),(15,1,'3g2','',1,0,1512694558,0),(16,1,'avi','',1,0,1512694558,0),(17,1,'navi','',1,0,1512694558,0),(18,1,'mov','',1,0,1512694558,0),(19,1,'asf','',1,0,1512694558,0),(20,1,'asx','',1,0,1512694558,0),(21,1,'wmv','',1,0,1512694558,0),(22,1,'wmvhd','',1,0,1512694558,0),(23,1,'mkv','',1,0,1512694558,0),(24,1,'flv','',1,0,1512694559,0),(25,1,'f4v','',1,0,1512694559,0),(26,1,'rmvb','',1,0,1512694559,0),(27,1,'webm','',1,0,1512694559,0),(28,1,'qsv','',1,0,1512694559,0),(29,1,'ogg','',1,0,1512694559,0),(30,1,'vob','',1,0,1512694559,0),(31,1,'swf','',1,0,1512694559,0),(32,1,'xv','',1,0,1512694559,0),(33,1,'rm','',1,0,1512694559,0),(34,1,'vcd','',1,0,1512694559,0),(35,1,'svcd','',1,0,1512694559,0),(36,1,'divx','',1,0,1512694559,0),(37,1,'xvid','',1,0,1512694559,0),(38,1,'dvd','',1,0,1512694559,0),(39,1,'dv','',1,0,1512694559,0),(40,1,'mts','',1,0,1512694559,0),(41,1,'ra','',1,0,1512694559,0),(42,1,'ram','',1,0,1512694559,0),(43,1,'wma','',1,0,1512694559,0),(44,1,'mpa','',1,0,1512694559,0),(45,1,'mod','',1,0,1512694559,0),(46,1,'wav','',1,0,1512694559,0),(47,1,'au','',1,0,1512694559,0),(48,1,'dif','',1,0,1512694559,0),(49,1,'ape','',1,0,1512694559,0),(50,1,'avs','',1,0,1512694559,0),(51,1,'avc','',1,0,1512694559,0),(52,1,'mux\n','',1,0,1512694559,0),(55,2,'mp3','',1,0,1512694559,0),(56,2,'mp3pro','',1,0,1512694559,0),(57,2,'midi','',1,0,1512694559,0),(60,2,'md','',1,0,1512694559,0),(61,2,'cda','',1,0,1512694559,0),(62,2,'sacd','',1,0,1512694559,0),(63,2,'quicktime','',1,0,1512694559,0),(64,2,'vqf','',1,0,1512694559,0),(65,2,'dvdaudio','',1,0,1512694559,0),(66,2,'realaudio','',1,0,1512694559,0),(67,2,'voc','',1,0,1512694559,0),(69,2,'aiff','',1,0,1512694559,0),(70,2,'amiga','',1,0,1512694559,0),(71,2,'mac','',1,0,1512694559,0),(72,2,'s48','',1,0,1512694559,0),(73,2,'aac\n','',1,0,1512694559,0),(75,3,'bmp','',1,0,1512694559,0),(76,3,'gif','',1,0,1512694559,0),(77,3,'jpeg','',1,0,1512694559,0),(78,3,'jpg','',1,0,1512694559,0),(79,3,'png','',1,0,1512694559,0),(80,3,'psd','',1,0,1512694559,0),(81,3,'tga','',1,0,1512694559,0),(82,3,'pcx','',1,0,1512694559,0),(83,3,'tiff','',1,0,1512694559,0),(84,3,'exif','',1,0,1512694559,0),(85,3,'fpx','',1,0,1512694559,0),(86,3,'svg','',1,0,1512694559,0),(87,3,'cdr','',1,0,1512694559,0),(88,3,'pcd','',1,0,1512694559,0),(89,3,'dxf','',1,0,1512694559,0),(90,3,'ufo','',1,0,1512694559,0),(91,3,'eps','',1,0,1512694559,0),(92,3,'ai','',1,0,1512694559,0),(93,3,'hdri','',1,0,1512694559,0),(94,3,'raw','',1,0,1512694559,0),(95,3,'wmf','',1,0,1512694559,0),(96,3,'lic','',1,0,1512694559,0),(97,3,'fli','',1,0,1512694559,0),(98,3,'flc','',1,0,1512694559,0),(99,3,'emf','',1,0,1512694559,0),(100,3,'dif\n','',1,0,1512694559,0),(102,4,'doc','',1,0,1512694559,0),(103,4,'docx','',1,0,1512694559,0),(104,4,'txt','',1,0,1512694559,0),(105,4,'pdf','',1,0,1512694559,0),(106,4,'wps','',1,0,1512694559,0),(107,4,'wpt','',1,0,1512694559,0),(108,4,'dot','',1,0,1512694559,0),(109,4,'rtf','',1,0,1512694559,0),(110,4,'dotx','',1,0,1512694559,0),(111,4,'docm','',1,0,1512694559,0),(112,4,'dotm','',1,0,1512694559,0),(113,4,'xls','',1,0,1512694559,0),(114,4,'xlt','',1,0,1512694559,0),(115,4,'xltx','',1,0,1512694559,0),(116,4,'xltm','',1,0,1512694559,0),(117,4,'xlsx','',1,0,1512694559,0),(118,4,'xlsm','',1,0,1512694559,0),(119,4,'xml','',1,0,1512694559,0),(120,4,'html','',1,0,1512694559,0),(121,4,'htm','',1,0,1512694559,0),(122,4,'mhtml','',1,0,1512694559,0),(123,4,'mht','',1,0,1512694559,0),(125,4,'csv','',1,0,1512694559,0),(126,4,'chm','',1,0,1512694559,0),(127,4,'wdl','',1,0,1512694559,0),(128,4,'ppt\n','',1,0,1512694559,0);
/*!40000 ALTER TABLE `zd_tools_format` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_tools_format_type`
--

DROP TABLE IF EXISTS `zd_tools_format_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_tools_format_type` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '文件类型名称',
  `descript` text COMMENT '文件类型描述',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='文件类型管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_tools_format_type`
--

LOCK TABLES `zd_tools_format_type` WRITE;
/*!40000 ALTER TABLE `zd_tools_format_type` DISABLE KEYS */;
INSERT INTO `zd_tools_format_type` VALUES (1,'视频','',1,0,1512694558,0),(2,'音频','',1,0,1512694559,0),(3,'图片','',1,0,1512694559,0),(4,'文字','',1,0,1512694559,0);
/*!40000 ALTER TABLE `zd_tools_format_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_tools_qrcode`
--

DROP TABLE IF EXISTS `zd_tools_qrcode`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_tools_qrcode` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '名称',
  `content` text COMMENT '内容',
  `url` text COMMENT '地址',
  `descript` text COMMENT '描述',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='二维码管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_tools_qrcode`
--

LOCK TABLES `zd_tools_qrcode` WRITE;
/*!40000 ALTER TABLE `zd_tools_qrcode` DISABLE KEYS */;
/*!40000 ALTER TABLE `zd_tools_qrcode` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_uc_admin`
--

DROP TABLE IF EXISTS `zd_uc_admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_uc_admin` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '用户名',
  `real_name` varchar(32) NOT NULL DEFAULT '0' COMMENT '真实姓名',
  `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
  `role_ids` varchar(255) NOT NULL DEFAULT '0' COMMENT '角色id字符串,eg:2,3,4,5',
  `phone` varchar(20) NOT NULL DEFAULT '0' COMMENT '手机号码',
  `motto` text COMMENT '个人格言',
  `sex` tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别,0:男,1:女',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
  `salt` char(10) NOT NULL DEFAULT '' COMMENT '密码加盐',
  `last_login` int(11) NOT NULL DEFAULT '0' COMMENT '最后登录时间',
  `last_ip` char(15) NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态,1:正常,0:禁用',
  `create_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COMMENT='管理员表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_uc_admin`
--

LOCK TABLES `zd_uc_admin` WRITE;
/*!40000 ALTER TABLE `zd_uc_admin` DISABLE KEYS */;
INSERT INTO `zd_uc_admin` VALUES (1,'jiangshide','jiangshide','c233ae5551f2413b3a1249cf65ccb635','1','18311271399','dadsadsa',0,'18311271399@163.com','F8aUvXWk7H',1513578890,'127.0.0.1',0,0,0,1511493641,0),(16,'litian123','litian','84931bb6cf7ad6e282d3c2db4c465015','2,1','18311271399','',0,'18311271399@163.com','339xXAn87w',1511514804,'127.0.0.1',1,1,1,1511514783,1511514783);
/*!40000 ALTER TABLE `zd_uc_admin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_uc_auth`
--

DROP TABLE IF EXISTS `zd_uc_auth`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_uc_auth` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `pid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '上级ID,0为顶级',
  `name` varchar(64) NOT NULL DEFAULT '0' COMMENT '权限名称',
  `url` varchar(255) NOT NULL DEFAULT '0' COMMENT 'URL地址',
  `sort` int(1) unsigned NOT NULL DEFAULT '999' COMMENT '排序,从小到大',
  `icon` varchar(255) NOT NULL DEFAULT '0' COMMENT '图标',
  `is_show` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否显示,0:隐藏,1:显示',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '操作者ID',
  `create_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态,1:正常,0:删除',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=80 DEFAULT CHARSET=utf8mb4 COMMENT='权限因子';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_uc_auth`
--

LOCK TABLES `zd_uc_auth` WRITE;
/*!40000 ALTER TABLE `zd_uc_auth` DISABLE KEYS */;
INSERT INTO `zd_uc_auth` VALUES (1,0,'所有权限','/backstage',1,'',0,1,1,1,1,1505620970,1505620970),(2,1,'权限管理','/backstage',999,'fa-id-card',1,1,0,1,1,0,1505622360),(3,2,'管理员','/backstage/admin/list',1,'fa-user-o',1,1,1,1,1,1505621186,1505621186),(4,2,'角色管理','/backstage/role/list',2,'fa-user-circle-o',1,1,0,1,1,0,1505621852),(5,3,'新增','/backstage/admin/add',1,'',0,1,0,1,1,0,1505621685),(6,3,'修改','/backstage/admin/edit',2,'',0,1,0,1,1,0,1505621697),(7,3,'删除','/backstage/admin/ajaxdel',3,'',0,1,1,1,1,1505621756,1505621756),(8,4,'新增','/backstage/role/add',1,'',1,1,0,1,1,0,1505698716),(9,4,'修改','/backstage/role/edit',2,'',0,1,1,1,1,1505621912,1505621912),(10,4,'删除','/backstage/role/ajaxdel',3,'',0,1,1,1,1,1505621951,1505621951),(11,2,'权限因子','/backstage/auth/list',3,'fa-list',1,1,1,1,1,1505621986,1505621986),(12,11,'新增','/backstage/auth/add',1,'',0,1,1,1,1,1505622009,1505622009),(13,11,'修改','/backstage/auth/edit',2,'',0,1,1,1,1,1505622047,1505622047),(14,11,'删除','/backstage/auth/ajaxdel',3,'',0,1,1,1,1,1505622111,1505622111),(15,1,'个人中心','/backstage/profile/edit',1001,'fa-user-circle-o',1,1,0,1,1,0,1506001114),(16,1,'API管理','/backstage',1,'fa-cubes',1,0,0,0,1,0,1506125698),(17,16,'API接口','/backstage/api/list',1,'fa-link',1,1,1,1,1,1505622447,1505622447),(19,16,'API监控','/backstage/apimonitor/list',3,'fa-bar-chart',1,1,0,1,1,0,1507700851),(20,1,'基础设置','/backstage/',2,'fa-cogs',1,1,1,1,1,1505622601,1505622601),(21,20,'分组设置','/backstage/group/list',1,'fa-object-ungroup',1,1,1,1,1,1505622645,1505622645),(22,20,'环境设置','/backstage/env/list',2,'fa-tree',1,1,1,1,1,1505622681,1505622681),(23,20,'状态码设置','/backstage/code/list',3,'fa-code',1,1,1,1,1,1505622728,1505622728),(24,15,'资料修改','/backstage/user/edit',1,'fa-edit',1,1,0,1,1,0,1506057468),(25,21,'新增','/backstage/group/add',1,'n',1,0,0,0,1,1506229739,1506229739),(26,21,'修改','/backstage/group/edit',2,'fa',0,0,0,0,1,1506237920,1506237920),(27,21,'删除','/backstage/group/ajaxdel',3,'fa',0,0,0,0,1,1506237948,1506237948),(28,22,'新增','/backstage/env/add',1,'fa',0,0,0,0,1,1506316506,1506316506),(29,22,'修改','/backstage/env/edit',2,'fa',0,0,0,0,1,1506316532,1506316532),(30,22,'删除','/backstage/env/ajaxdel',3,'fa',0,0,0,0,1,1506316567,1506316567),(31,23,'新增','/backstage/code/add',1,'fa',0,0,0,0,1,1506327812,1506327812),(32,23,'修改','/backstage/code/edit',2,'fa',0,0,0,0,1,1506327831,1506327831),(33,23,'删除','/backstage/code/ajadel',3,'fa',0,0,0,0,1,1506327857,1506327857),(34,17,'新增资源','/backstage/api/add',1,'fa-link',1,1,0,1,1,0,1507436029),(35,17,'修改资源','/backstage/api/edit',2,'fa-link',1,1,0,1,1,0,1507436042),(36,17,'删除资源','/backstage/api/ajaxdel',3,'fa-link',1,1,0,1,1,0,1507436052),(37,17,'新增接口','/backstage/api/addapi',4,'',0,1,1,1,1,1507436014,1507436014),(38,17,'修改接口','/backstage/api/editapi',5,'',0,1,0,1,1,0,1507705049),(52,1,'工具管理','/backstage',1,'fa-briefcase',1,1,0,1,1,0,1511838566),(55,52,'二维码','/backstage/tools/qrcode',1,'fa-qrcode',1,1,1,1,1,1511836676,1511836676),(56,1,'区域管理','/backstage',1,'fa-area-chart',1,1,1,1,1,1511838063,1511838063),(57,1,'民族管理','/backstage',1,'fa-flag',1,1,1,1,1,1511838220,1511838220),(58,1,'测试管理','/backstage',1,'fa-bug',1,1,1,1,1,1511838538,1511838538),(59,58,'环境配置','/backstage/test/environment',1,'fa-usb',1,1,0,1,1,0,1511839774),(62,57,'名族列表','/backstage/nation',1,'fa-list',1,1,1,1,1,1511840283,1511840283),(63,56,'洲','/backstage/area/continent',1,'fa-pie-chart',1,1,1,1,1,1512350590,1512350590),(64,56,'国家','/backstage/area/state',1,'fa-institution',1,1,1,1,1,1512350771,1512350771),(65,56,'省','/backstage/area/province',1,'fa-flag-checkered',1,1,1,1,1,1512350897,1512350897),(66,56,'市','/backstage/area/city',1,'fa-id-card',1,1,1,1,1,1512350980,1512350980),(67,56,'地区','/backstage/area/region',1,'fa-industry',1,1,1,1,1,1512351063,1512351063),(68,56,'县','/backstage/area/county',1,'fa-bar-chart',1,1,1,1,1,1512351165,1512351165),(69,56,'镇','/backstage/area/town',1,'fa-university',1,1,1,1,1,1512351218,1512351218),(70,56,'乡','/backstage/area/country',1,'fa-car',1,1,1,1,1,1512351438,1512351438),(71,56,'村','/backstage/area/village',1,'fa-tree',1,1,1,1,1,1512351591,1512351591),(72,56,'组','/backstage/area/group',1,'fa-server',1,1,1,1,1,1512351707,1512351707),(73,56,'队','/backstage/area/team',1,'fa-bicycle',1,1,1,1,1,1512351765,1512351765),(74,1,'WEB管理','/backstage/web/',1,'fa-cloud',1,1,1,1,1,1512551278,1512551278),(75,74,'焦点图','/backstage/web/banner',1,'fa-windows',1,1,1,1,1,1512551694,1512551694),(76,52,'压缩','/backstage/tools/compress',1,'fa-archive',1,1,1,1,1,1512613323,1512613323),(77,52,'文件格式','/backstage/tools/format',1,'fa-file',1,1,1,1,1,1512626324,1512626324),(78,58,'app列表','/backstage/test/app',1,'fa-envira',1,1,1,1,1,1512985040,1512985040),(79,58,'项目列表','/backstage/test/project',1,'fa-gitlab',1,1,1,1,1,1513056533,1513056533);
/*!40000 ALTER TABLE `zd_uc_auth` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_uc_role`
--

DROP TABLE IF EXISTS `zd_uc_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_uc_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(32) NOT NULL DEFAULT '0' COMMENT '角色名称',
  `detail` varchar(255) NOT NULL DEFAULT '0' COMMENT '备注',
  `create_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改ID',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态1:正常,0:删除',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='角色表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_uc_role`
--

LOCK TABLES `zd_uc_role` WRITE;
/*!40000 ALTER TABLE `zd_uc_role` DISABLE KEYS */;
INSERT INTO `zd_uc_role` VALUES (1,'API管理员','拥有API所有权限',0,2,1,1505874156,1505874156),(2,'系统管理员','系统管理员',0,0,1,1506124114,1506124114);
/*!40000 ALTER TABLE `zd_uc_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_uc_role_auth`
--

DROP TABLE IF EXISTS `zd_uc_role_auth`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_uc_role_auth` (
  `role_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
  `auth_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '权限ID',
  PRIMARY KEY (`role_id`,`auth_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='权限和角色关系管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_uc_role_auth`
--

LOCK TABLES `zd_uc_role_auth` WRITE;
/*!40000 ALTER TABLE `zd_uc_role_auth` DISABLE KEYS */;
INSERT INTO `zd_uc_role_auth` VALUES (1,16),(1,17),(1,18),(1,19),(2,0),(2,1),(2,15),(2,20),(2,21),(2,22),(2,23),(2,24);
/*!40000 ALTER TABLE `zd_uc_role_auth` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zd_web_banner`
--

DROP TABLE IF EXISTS `zd_web_banner`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zd_web_banner` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '名称',
  `link` text COMMENT '跳转新页面',
  `icon` varchar(100) DEFAULT '' COMMENT '图片展示',
  `descript` varchar(500) NOT NULL DEFAULT '0' COMMENT '描述',
  `clicks` int(11) NOT NULL DEFAULT '0' COMMENT '点击次数',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `views` int(11) NOT NULL DEFAULT '0' COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='焦点图管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zd_web_banner`
--

LOCK TABLES `zd_web_banner` WRITE;
/*!40000 ALTER TABLE `zd_web_banner` DISABLE KEYS */;
INSERT INTO `zd_web_banner` VALUES (2,'活动2','http://www.baidu.com','/static/upload/jpg/381f923fa4d0bcdd9902ff5b7e649958_1512914908.jpg','说三道四的',0,1,1,1512618080,1512914935,0),(3,'活动3','http://www.baidu.com','/static/upload/jpg/7475b8166d6fd5dd4181db58bf03b207_1512914869.jpg','大大的',0,1,1,1512618098,1512914883,0),(4,'活动4','http://www.baidu.com','/static/upload/jpg/035356d5d378638009da7aeab8b7773b_1512914834.jpg','是大大的',0,1,1,1512618134,1512914838,0);
/*!40000 ALTER TABLE `zd_web_banner` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-12-22 15:12:35
