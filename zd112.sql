SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

#--------the user---start------#

DROP TABLE IF EXISTS `zd_permission_role`;
CREATE TABLE `zd_permission_role` (
  `id`          BIGINT UNSIGNED  NOT NULL AUTO_INCREMENT
  COMMENT '主键ID',
  `name`        VARCHAR(32)      NOT NULL DEFAULT '0'
  COMMENT '角色名称',
  `detail`      VARCHAR(255)     NOT NULL DEFAULT '0'
  COMMENT '备注',
  `create_id`   BIGINT UNSIGNED  NOT NULL DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT UNSIGNED  NOT NULL DEFAULT '0'
  COMMENT '修改ID',
  `create_time` INT(11) UNSIGNED NOT NULL DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED NOT NULL DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = Innodb
  DEFAULT CHARSET = utf8
  COMMENT ='角色表';

INSERT INTO `zd_permission_role` VALUES (1, '超级管理员', '拥有所有权限', 1, 0, 1514260694, 0);

DROP TABLE IF EXISTS `zd_permission_auth`;
CREATE TABLE `zd_permission_auth` (
  `id`          BIGINT UNSIGNED     NOT NULL AUTO_INCREMENT
  COMMENT '主键ID',
  `pid`         TINYINT(1) UNSIGNED NOT NULL DEFAULT '0'
  COMMENT '上级ID:0~顶级菜单;',
  `name`        VARCHAR(20)         NOT NULL DEFAULT '0'
  COMMENT '权限名称',
  `action`      VARCHAR(255)        NOT NULL DEFAULT '0'
  COMMENT '权限Action',
  `icon`        VARCHAR(255)        NOT NULL DEFAULT '0'
  COMMENT '图标',
  `is_show`     TINYINT(1) UNSIGNED NOT NULL DEFAULT '0'
  COMMENT '是否显示,0:隐藏,1:显示',
  `create_id`   BIGINT UNSIGNED     NOT NULL DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT UNSIGNED     NOT NULL DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11) UNSIGNED    NOT NULL DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED    NOT NULL DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='权限因子';

BEGIN;
INSERT INTO `zd_permission_auth` VALUES (1, 0, '超级权限', '/', '', 1, 1, 1, 1514260694, 0),
  (2, 1, '权限管理', '/permission', 'fa-id-card', 1, 1, 0, 1514260694, 0),
  (3, 2, '用户管理', '/permission/user', 'fa-user-o', 1, 1, 0, 1514260694, 0),
  (4, 2, '角色管理', '/permission/role', 'fa-user-circle-o', 1, 1, 0, 1514260694, 0),
  (5, 2, '权限因子', '/permission/auth', 'fa-tasks', 1, 1, 0, 1514260694, 0),
  (6, 1, '区域管理', '/area', 'fa-area-chart', 1, 1, 0, 1514260694, 0),
  (7, 6, '区域列表', '/area/list', 'fa-pie-chart', 1, 1, 0, 1514260694, 0),
  (8, 1, '民族管理', '/nation', 'fa-flag', 1, 1, 0, 1514260694, 0),
  (9, 8, '民族列表', '/nation/list', 'fa-list', 1, 1, 0, 1514260694, 0),
  (10, 1, 'Web管理', '/web/', 'fa-cloud', 1, 1, 0, 1514260694, 0),
  (11, 10, '首页Banner', '/web/banner', 'fa-windows', 1, 1, 0, 1514260694, 0),
  (12, 1, '工具管理', '/tool', 'fa-briefcase', 1, 1, 0, 1514260694, 0),
  (13, 12, '二维码', '/tool/qrcode', 'fa-qrcode', 1, 1, 0, 1514260694, 0),
  (14, 12, '压缩', '/tool/compress', 'fa-archive', 1, 1, 0, 1514260694, 0),
  (15, 12, '文件格式', '/tool/format', 'fa-file', 1, 1, 0, 1514260694, 0),
  (16, 10, '导航栏管理', '/web/nav', 'fa-list', 1, 1, 0, 1514260694, 0);
COMMIT;

DROP TABLE IF EXISTS `zd_permission_user`;
CREATE TABLE `zd_permission_user` (
  `id`          BIGINT UNSIGNED  NOT NULL AUTO_INCREMENT
  COMMENT '主键ID',
  `role_id`     VARCHAR(500)     NOT NULL DEFAULT '0'
  COMMENT '角色ID',
  `name`        VARCHAR(20)      NOT NULL DEFAULT ''
  COMMENT '用户名',
  `password`    CHAR(32)         NOT NULL DEFAULT ''
  COMMENT '密码',
  `salt`        CHAR(10)         NOT NULL DEFAULT ''
  COMMENT '密码加盐',
  `status`      TINYINT(4)       NOT NULL DEFAULT '0'
  COMMENT '状态,0:未激活,1:已激活,2:禁用',
  `create_id`   BIGINT UNSIGNED  NOT NULL DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT UNSIGNED  NOT NULL DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11) UNSIGNED NOT NULL DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED NOT NULL DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_name` (`name`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='用户表';

INSERT INTO `zd_permission_user`
VALUES (1, '0', 'jiangshide', '5c7914423a3bcf0ab7932fe3e7482308', 'Pgfg7mgihF', 1, 1, 0, 1514260694, 0);

DROP TABLE IF EXISTS `zd_user_location`; #用户实时地址
CREATE TABLE `zd_user_location` (
  `id`          BIGINT UNSIGNED  NOT NULL AUTO_INCREMENT
  COMMENT '主键ID',
  `user_id`     BIGINT UNSIGNED  NOT NULL DEFAULT '0'
  COMMENT '用户ID',
  `ip`          CHAR(15)         NOT NULL DEFAULT ''
  COMMENT '最后登录IP',
  `mac`         VARCHAR(50)      NOT NULL DEFAULT ''
  COMMENT '客户端MAC地址',
  `latitude`    VARCHAR(12)      NOT NULL DEFAULT ''
  COMMENT '经度',
  `longitude`   VARCHAR(12)      NOT NULL DEFAULT ''
  COMMENT '纬度',
  `device`      VARCHAR(50)      NOT NULL DEFAULT ''
  COMMENT '使用设备名称:ios:android...',
  `mode`        VARCHAR(20)      NOT NULL DEFAULT ''
  COMMENT '机型',
  `arch`        VARCHAR(10)      NOT NULL DEFAULT ''
  COMMENT 'OS架构',
  `sdk_version` VARCHAR(20)      NOT NULL DEFAULT ''
  COMMENT '系统sdk版本',
  `app_version` VARCHAR(10)      NOT NULL DEFAULT ''
  COMMENT '',
  `create_id`   BIGINT UNSIGNED  NOT NULL DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT UNSIGNED  NOT NULL DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11) UNSIGNED NOT NULL DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED NOT NULL DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='用户位置表';

DROP TABLE IF EXISTS `zd_user_profile`; #用户简介
CREATE TABLE `zd_user_profile` (
  `id`            BIGINT UNSIGNED  NOT NULL AUTO_INCREMENT
  COMMENT '主键ID',
  `user_id`       BIGINT UNSIGNED  NOT NULL DEFAULT '0'
  COMMENT '用户ID',
  `psw_intensity` TINYINT(1)       NOT NULL DEFAULT '0'
  COMMENT '密码强度:0~危险,1~普通,2~中等,3~安全',
  `icon`          VARCHAR(500)     NOT NULL DEFAULT ''
  COMMENT '用户头像',
  `level`         TINYINT(4)       NOT NULL DEFAULT '0'
  COMMENT '级别:0~普通;1~管理员,2~10文化币一级:',
  `score`         INT              NOT NULL DEFAULT '0'
  COMMENT '积分:',
  `name`          VARCHAR(50)      NOT NULL DEFAULT ''
  COMMENT '用户真实名',
  `phone`         VARCHAR(20)      NOT NULL DEFAULT '0'
  COMMENT '手机号码',
  `motto`         VARCHAR(50)      NOT NULL DEFAULT ''
  COMMENT '格言',
  `sex`           TINYINT(1)       NOT NULL DEFAULT '0'
  COMMENT '性别,0:保密,1:男,1:女',
  `email`         VARCHAR(50)      NOT NULL DEFAULT ''
  COMMENT '邮箱',
  `weixin`        VARCHAR(20)      NOT NULL DEFAULT ''
  COMMENT '微信',
  `qq`            VARCHAR(20)      NOT NULL DEFAULT ''
  COMMENT 'qq',
  `weibo`         VARCHAR(30)      NOT NULL DEFAULT ''
  COMMENT 'WEIBO',
  `address`       VARCHAR(200)     NOT NULL DEFAULT ''
  COMMENT '常住详细地址',
  `code`          INT              NOT NULL DEFAULT '0'
  COMMENT '常住地址邮政编码',
  `create_id`     BIGINT UNSIGNED  NOT NULL DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`     BIGINT UNSIGNED  NOT NULL DEFAULT '0'
  COMMENT '修改者ID',
  `create_time`   INT(11) UNSIGNED NOT NULL DEFAULT '0'
  COMMENT '创建时间',
  `update_time`   INT(11) UNSIGNED NOT NULL DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='用户简介表';

DROP TABLE IF EXISTS `zd_user_address`; #用户邮寄地址
CREATE TABLE `zd_user_address` (
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `user_id`     BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '用户ID',
  `name`        VARCHAR(200)     NOT NULL  DEFAULT ''
  COMMENT '邮件地址',
  `code`        INT              NOT NULL  DEFAULT '0'
  COMMENT '邮政编码',
  `create_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_address_name` (`name`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='用户地址管理表';

#--------the user---end------#


#--------the comm---start------#

DROP TABLE IF EXISTS `zd_classify`; #分类
CREATE TABLE `zd_classify` (
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `name`        VARCHAR(20)      NOT NULL  DEFAULT ''
  COMMENT '分类名称',
  `descript`    VARCHAR(255)     NOT NULL  DEFAULT ''
  COMMENT '说明',
  `create_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='分类';

DROP TABLE IF EXISTS `zd_age`; #年代
CREATE TABLE `zd_age` (
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `name`        VARCHAR(20)      NOT NULL  DEFAULT ''
  COMMENT '年代名称',
  `descript`    VARCHAR(255)     NOT NULL  DEFAULT ''
  COMMENT '说明',
  `create_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='年代';

DROP TABLE IF EXISTS `zd_auth`; #作者
CREATE TABLE `zd_auth` (
  `id`              BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `icon`            VARCHAR(255)     NOT NULL  DEFAULT ''
  COMMENT '头像',
  `name`            VARCHAR(50)      NOT NULL  DEFAULT ''
  COMMENT '作者名称',
  `constellationId` BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '星座ID',
  `constellation`   VARCHAR(10)      NOT NULL  DEFAULT ''
  COMMENT '星座名称',
  `impression`      VARCHAR(255)     NOT NULL  DEFAULT ''
  COMMENT '大众印象:',
  `workCount`       INT              NOT NULL  DEFAULT '0'
  COMMENT '著作数量',
  `descript`        VARCHAR(255)     NOT NULL  DEFAULT ''
  COMMENT '说明',
  `create_id`       BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`       BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '修改者ID',
  `create_time`     INT(11) UNSIGNED NOT NULL  DEFAULT '0'
  COMMENT '创建时间',
  `update_time`     INT(11) UNSIGNED NOT NULL  DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='作者属性';

DROP TABLE IF EXISTS `zd_work`; #著作
CREATE TABLE `zd_work` (
  `id`           BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `authId`       BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '作者ID',
  `classifyId`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '分类ID',
  'classifyName' VARCHAR(20)      NOT NULL  DEFAULT '0'
  COMMENT '分类名称',
  `ageId`        BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '年代ID',
  `cover`        VARCHAR(255)     NOT NULL  DEFAULT ''
  COMMENT '封面',
  `name`         VARCHAR(20)      NOT NULL  DEFAULT ''
  COMMENT '年代名称',
  `title`        VARCHAR(50)      NOT NULL  DEFAULT ''
  COMMENT '标题',
  `tag`          VARCHAR(255)     NOT NULL  DEFAULT ''
  COMMENT '标签:描写花,高中,委婉派',
  `content`      TEXT COMMENT '具体内容',
  `appreciation` TEXT COMMENT '作品赏析',
  `like`         INT(11)          NOT NULL  DEFAULT '0'
  COMMENT '喜欢',
  `follow`       INT(11)          NOT NULL  DEFAULT '0'
  COMMENT '关注',
  `commentId`    BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '评论ID',
  `commentCount` INT(11) UNSIGNED NULL      DEFAULT '0'
  COMMENT '评论总数',
  `descript`     VARCHAR(255)     NOT NULL  DEFAULT ''
  COMMENT '说明',
  `create_id`    BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`    BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '修改者ID',
  `create_time`  INT(11) UNSIGNED NOT NULL  DEFAULT '0'
  COMMENT '创建时间',
  `update_time`  INT(11) UNSIGNED NOT NULL  DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='著作';

DROP TABLE IF EXISTS `zd_work_comment`; #针对著作的评论
CREATE TABLE `zd_work_comment` (
  `id`            BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `workId`        BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '著作ID',
  `owner_user_id` BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '发表评论的用户ID',
  `taget_user_id` BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '评论的目标用户ID',
  `content`       TEXT COMMENT '评论内容',
  `praise_count`  INT UNSIGNED     NOT NULL  DEFAULT '0'
  COMMENT '点赞次数',
  `low_count`     INT UNSIGNED     NOT NULL  DEFAULT '0'
  COMMENT '差评次数',
  `create_id`     BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`     BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '修改者ID',
  `create_time`   INT(11) UNSIGNED NOT NULL  DEFAULT '0'
  COMMENT '创建时间',
  `update_time`   INT(11) UNSIGNED NOT NULL  DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='评论';

DROP TABLE IF EXISTS `zd_work_recomment`; #回复评论
CREATE TABLE `zd_work_recomment` (
  `id`            BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `commentId`     BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '评论ID',
  `owner_user_id` BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '发表评论的用户ID',
  `taget_user_id` BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '评论的目标用户ID',
  `content`       TEXT COMMENT '评论内容',
  `praise_count`  INT UNSIGNED     NOT NULL  DEFAULT '0'
  COMMENT '点赞次数',
  `low_count`     INT UNSIGNED     NOT NULL  DEFAULT '0'
  COMMENT '差评次数',
  `create_id`     BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`     BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '修改者ID',
  `create_time`   INT(11) UNSIGNED NOT NULL  DEFAULT '0'
  COMMENT '创建时间',
  `update_time`   INT(11) UNSIGNED NOT NULL  DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='回复评论';

#https://www.xzw.com/astro/
DROP TABLE IF EXISTS `zd_constellation`; #星座
CREATE TABLE `zd_constellation` (
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `icon`        VARCHAR(255)     NOT NULL  DEFAULT ''
  COMMENT '图标',
  `name`        VARCHAR(10)      NOT NULL  DEFAULT ''
  COMMENT '星座名称',
  `datestr`     VARCHAR(20)      NOT NULL  DEFAULT ''
  COMMENT '区间日期',
  `five_line`   VARCHAR(10)      NOT NULL  DEFAULT ''
  COMMENT '五行:水,木,金,火,土',
  `trait`       VARCHAR(20)      NOT NULL  DEFAULT ''
  COMMENT '特点:',
  `descript`    VARCHAR(255)     NOT NULL  DEFAULT ''
  COMMENT '说明',
  `create_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='星座';

#http://www.xinglai.com/articles/astrology/tutorial/house/
DROP TABLE IF EXISTS `zd_house`; #宫位
CREATE TABLE `zd_house` (
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `icon`        VARCHAR(255)     NOT NULL  DEFAULT ''
  COMMENT '图标',
  `name`        VARCHAR(10)      NOT NULL  DEFAULT ''
  COMMENT '宫位名称',
  `live`        VARCHAR(100)     NOT NULL  DEFAULT ''
  COMMENT '生命领域',
  `base`        VARCHAR(255)     NOT NULL  DEFAULT ''
  COMMENT '基本意义',
  `planet_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '行星表ID',
  `planet`      VARCHAR(20)      NOT NULL  DEFAULT ''
  COMMENT '行星名称',
  `descript`    VARCHAR(255)     NOT NULL  DEFAULT ''
  COMMENT '说明',
  `create_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='星座';

DROP TABLE IF EXISTS `zd_house_planet`; #行星落入宫位
CREATE TABLE `zd_house_planet` (
  `id`          BIGINT UNSIGNED  NOT NULL  AUTO_INCREMENT,
  `icon`        VARCHAR(255)     NOT NULL  DEFAULT ''
  COMMENT '图标',
  `name`        VARCHAR(20)      NOT NULL  DEFAULT ''
  COMMENT '行星名称',
  `trait`       TEXT COMMENT '基本特点',
  `descript`    VARCHAR(255)     NOT NULL  DEFAULT ''
  COMMENT '详细说明',
  `create_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT UNSIGNED  NOT NULL  DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED NOT NULL  DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='星座落入宫位表';

#--------the comm---end------#


#--------the set---start------#

DROP TABLE IF EXISTS `zd_set_group`;
CREATE TABLE `zd_set_group` (
  `id`          BIGINT UNSIGNED     NOT NULL AUTO_INCREMENT
  COMMENT '自增',
  `name`        VARCHAR(50)         NOT NULL DEFAULT ''
  COMMENT '组名',
  `detail`      VARCHAR(255)        NOT NULL DEFAULT ''
  COMMENT '说明',
  `status`      TINYINT(1) UNSIGNED NOT NULL DEFAULT '0'
  COMMENT '状态:1:正常,0:删除',
  `create_id`   BIGINT              NOT NULL DEFAULT '0'
  COMMENT '用户ID',
  `update_id`   BIGINT              NOT NULL DEFAULT '0'
  COMMENT '更新者ID',
  `create_time` INT(11)             NOT NULL DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)             NOT NULL DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_create_id` (`create_id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `zd_set_evn`;
CREATE TABLE `zd_set_evn` (
  `id`          BIGINT UNSIGNED     NOT NULL AUTO_INCREMENT,
  `name`        VARCHAR(50)         NOT NULL DEFAULT ''
  COMMENT '环境名称',
  `env_host`    VARCHAR(255)        NOT NULL DEFAULT '0'
  COMMENT '主机',
  `detail`      VARCHAR(255)        NOT NULL DEFAULT '0'
  COMMENT '备注',
  `status`      TINYINT(4) UNSIGNED NOT NULL DEFAULT '0'
  COMMENT '状态,1:正常,0:禁用',
  `create_id`   BIGINT UNSIGNED     NOT NULL DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT UNSIGNED     NOT NULL DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11) UNSIGNED    NOT NULL DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED    NOT NULL DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_env_name` (`name`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='环境分组表';

DROP TABLE IF EXISTS `zd_set_code`;
CREATE TABLE `zd_set_code` (
  `id`          BIGINT UNSIGNED     NOT NULL AUTO_INCREMENT,
  `code`        VARCHAR(50)         NOT NULL DEFAULT '0'
  COMMENT '状态码',
  `descript`    VARCHAR(255)        NOT NULL DEFAULT '0'
  COMMENT '描述',
  `detail`      VARCHAR(255)        NOT NULL DEFAULT '0'
  COMMENT '备注',
  `status`      TINYINT(4) UNSIGNED NOT NULL DEFAULT '0'
  COMMENT '状态,1:正常,0:禁用',
  `create_id`   BIGINT UNSIGNED     NOT NULL DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT UNSIGNED     NOT NULL DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11) UNSIGNED    NOT NULL DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED    NOT NULL DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_env_name` (`code`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='状态代码表';

#--------the set---end------#

#--------the api---start------#

DROP TABLE IF EXISTS `zd_api_src`;
CREATE TABLE `zd_api_src` (
  `id`          BIGINT UNSIGNED     NOT NULL AUTO_INCREMENT,
  `group_id`    BIGINT              NOT NULL DEFAULT '0'
  COMMENT '分组ID',
  `name`        VARCHAR(50)         NOT NULL DEFAULT '0'
  COMMENT '接口名称',
  `status`      TINYINT(1) UNSIGNED NOT NULL DEFAULT '1'
  COMMENT '状态,1:审核通过,0:暂停使用,2:草稿,3:审核中',
  `audit_id`    BIGINT UNSIGNED     NOT NULL DEFAULT '0'
  COMMENT '审核人ID',
  `create_id`   BIGINT UNSIGNED     NOT NULL DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT UNSIGNED     NOT NULL DEFAULT '0'
  COMMENT '更新者ID',
  `create_time` INT(11) UNSIGNED    NOT NULL DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED    NOT NULL DEFAULT '0'
  COMMENT '更新时间',
  `audit_time`  INT(11) UNSIGNED    NOT NULL DEFAULT '0'
  COMMENT '审核时间',
  PRIMARY KEY (`id`),
  KEY `idx_group_id` (`group_id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='API主表';

DROP TABLE IF EXISTS `zd_api_param`;
CREATE TABLE `zd_api_param` (
  `id`          BIGINT UNSIGNED     NOT NULL AUTO_INCREMENT,
  `detail_id`   BIGINT UNSIGNED     NOT NULL DEFAULT '0'
  COMMENT '附表ID',
  `api_key`     VARCHAR(100)        NOT NULL DEFAULT '0'
  COMMENT '参数名',
  `api_type`    VARCHAR(100)        NOT NULL DEFAULT '0'
  COMMENT '类型',
  `api_value`   VARCHAR(500)        NOT NULL DEFAULT '0'
  COMMENT '参数值',
  `api_detail`  VARCHAR(500)        NOT NULL DEFAULT '0'
  COMMENT '参数说明',
  `is_null`     VARCHAR(10)         NOT NULL DEFAULT '0'
  COMMENT '是否必填',
  `status`      TINYINT(1) UNSIGNED NOT NULL DEFAULT '1'
  COMMENT '状态,1:正常,0:删除',
  `create_id`   BIGINT UNSIGNED     NOT NULL DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT UNSIGNED     NOT NULL DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11) UNSIGNED    NOT NULL DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED    NOT NULL DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_main_id` (`detail_id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='API参数表';

DROP TABLE IF EXISTS `zd_api_detail`;
CREATE TABLE `zd_api_detail` (
  `id`            BIGINT UNSIGNED     NOT NULL AUTO_INCREMENT,
  `src_id`        BIGINT              NOT NULL DEFAULT '0'
  COMMENT '主表ID',
  `method`        TINYINT(1)          NOT NULL DEFAULT '1'
  COMMENT '方法名称,1:GET,2:POST,3:PUT,4:PATCH,5:DELETE',
  `name`          VARCHAR(100)        NOT NULL DEFAULT '0'
  COMMENT '接口名称',
  `api_url`       VARCHAR(100)        NOT NULL DEFAULT '0'
  COMMENT '接口地址',
  `protocol_type` VARCHAR(20)         NOT NULL DEFAULT '1'
  COMMENT '协议类型,1:http,2:https',
  `result`        TEXT COMMENT '返回结果,正确或者错误',
  `example`       TEXT COMMENT '接口示例',
  `detail`        VARCHAR(1000)       NOT NULL DEFAULT '0'
  COMMENT '注意事项',
  `audit_time`    INT(11) UNSIGNED    NOT NULL DEFAULT '0'
  COMMENT '审核时间',
  `audit_id`      BIGINT              NOT NULL DEFAULT '0'
  COMMENT '审核人ID',
  `status`        TINYINT(1) UNSIGNED NOT NULL DEFAULT '1'
  COMMENT '状态,0:暂停 使用,1:正在审核,2:审核通过',
  `create_id`     BIGINT              NOT NULL DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`     BIGINT              NOT NULL DEFAULT '0'
  COMMENT '修改者ID',
  `create_time`   INT(11)             NOT NULL DEFAULT '0'
  COMMENT '创建时间',
  `update_time`   INT(11)             NOT NULL DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_main_id` (`src_id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='API附表';

#--------the api---end------#

#--------the nation---start-----#

DROP TABLE IF EXISTS `zd_nation`;
CREATE TABLE `zd_nation` (
  `id`          BIGINT UNSIGNED      NOT NULL AUTO_INCREMENT,
  `name`        VARCHAR(25) UNIQUE   NOT NULL DEFAULT ''
  COMMENT '名称',
  `icon`        VARCHAR(100)         NOT NULL DEFAULT ''
  COMMENT 'LOGO',
  `create_id`   BIGINT               NOT NULL DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT               NOT NULL DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)              NOT NULL DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)              NOT NULL DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='民族名称';
#--------the nation---end-----#


#--------the test---start-----#

DROP TABLE IF EXISTS `zd_area`; #区域管理
CREATE TABLE `zd_area` (
  `id`              BIGINT UNSIGNED    NOT NULL      AUTO_INCREMENT,
  `parent_id`       BIGINT UNSIGNED    NOT NULL      DEFAULT '0'
  COMMENT '所属ID:0~洲/洋,1~区域,2~国家,3~省,4~市,5~地区,6~县,7~镇,8~乡,9~村,10~组,11~队',
  `name`            VARCHAR(50) UNIQUE NOT NULL      DEFAULT ''
  COMMENT '名称:中国',
  `name_en`         VARCHAR(50) UNIQUE NOT NULL      DEFAULT ''
  COMMENT '代码名称~外文:china',
  `name_ab`         VARCHAR(10)        NOT NULL      DEFAULT ''
  COMMENT '名称英文简写:CN',
  `icon`            VARCHAR(100)       NOT NULL NULL DEFAULT ''
  COMMENT 'LOGO',
  `zone_code`       INT                NOT NULL      DEFAULT '0'
  COMMENT '区号:86',
  `zip_code`        INT                NOT NULL      DEFAULT '0'
  COMMENT '邮编:999001',
  `area_code`       VARCHAR(20)        NOT NULL      DEFAULT ''
  COMMENT '区域代码～北京:1000',
  `time_difference` VARCHAR(10)        NOT NULL      DEFAULT ''
  COMMENT '时差:0',
  `create_id`       BIGINT             NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`       BIGINT             NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time`     INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time`     INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='区域管理';

#--------the test---end-----#

#--------the web---start-----#
DROP TABLE IF EXISTS `zd_web_banner`; #焦点图
CREATE TABLE `zd_web_banner` (
  `id`          BIGINT UNSIGNED    NOT NULL      AUTO_INCREMENT,
  `title`       VARCHAR(50)        NOT NULL      DEFAULT ''
  COMMENT '标题',
  `name`        VARCHAR(50) UNIQUE NOT NULL      DEFAULT ''
  COMMENT '名称',
  `link`        TEXT COMMENT '跳转新页面',
  `icon`        VARCHAR(100)       NOT NULL NULL DEFAULT ''
  COMMENT '图片展示',
  `descript`    VARCHAR(500)       NOT NULL      DEFAULT '0'
  COMMENT '描述',
  `create_id`   BIGINT             NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT             NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       BIGINT             NOT NULL      DEFAULT '0'
  COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='焦点图管理';

DROP TABLE IF EXISTS `zd_web_nav`; #栏目
CREATE TABLE `zd_web_nav` (
  `id`          BIGINT UNSIGNED    NOT NULL      AUTO_INCREMENT,
  `name`        VARCHAR(50) UNIQUE NOT NULL      DEFAULT ''
  COMMENT '名称',
  `action`      VARCHAR(50)        NOT NULL      DEFAULT ''
  COMMENT '跳转事件',
  `descript`    TEXT COMMENT '描述',
  `create_id`   BIGINT             NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT             NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       BIGINT             NOT NULL      DEFAULT '0'
  COMMENT '当前点击次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='栏目';

DROP TABLE IF EXISTS `zd_web_nav_university`; #栏目子项
CREATE TABLE `zd_web_nav_university` (
  `id`          BIGINT UNSIGNED NOT NULL      AUTO_INCREMENT,
  `group_id`    BIGINT UNSIGNED NOT NULL      DEFAULT '0'
  COMMENT '所属组',
  `name`        VARCHAR(50)     NOT NULL      DEFAULT ''
  COMMENT '名称',
  `action`      VARCHAR(50)     NOT NULL      DEFAULT ''
  COMMENT '跳转事件',
  `descript`    TEXT COMMENT '描述',
  `create_id`   BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)         NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)         NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '当前点击次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='栏目子项';

#--------the web---end-----#

#--------the area---start-----#

#--------the area---end-----#


#--------the tool---start-----#

DROP TABLE IF EXISTS `zd_tools_qrcode`; #二维码管理
CREATE TABLE `zd_tools_qrcode` (
  `id`          BIGINT UNSIGNED    NOT NULL      AUTO_INCREMENT,
  `name`        VARCHAR(50) UNIQUE NOT NULL      DEFAULT ''
  COMMENT '名称',
  `content`     TEXT COMMENT '内容',
  `url`         TEXT COMMENT '地址',
  `descript`    TEXT COMMENT '描述',
  `create_id`   BIGINT             NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT             NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='二维码管理';

DROP TABLE IF EXISTS `zd_tools_format_type`; #文件格式类型
CREATE TABLE `zd_tools_format_type` (
  `id`          BIGINT UNSIGNED    NOT NULL      AUTO_INCREMENT,
  `name`        VARCHAR(50) UNIQUE NOT NULL      DEFAULT ''
  COMMENT '文件类型名称',
  `descript`    TEXT COMMENT '文件类型描述',
  `create_id`   BIGINT             NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT             NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='文件类型管理';

DROP TABLE IF EXISTS `zd_tools_format`; #文件格式
CREATE TABLE `zd_tools_format` (
  `id`          BIGINT UNSIGNED      NOT NULL      AUTO_INCREMENT,
  `parent_id`   BIGINT UNSIGNED      NOT NULL      DEFAULT '0'
  COMMENT '所格式类型',
  `name`        VARCHAR(50) UNIQUE   NOT NULL      DEFAULT ''
  COMMENT '格式民称',
  `descript`    TEXT COMMENT '格式描述',
  `create_id`   BIGINT               NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT               NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)              NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)              NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='文件格式管理';

DROP TABLE IF EXISTS `zd_tools_compress`; #压缩工具
CREATE TABLE `zd_tools_compress` (
  `id`          BIGINT UNSIGNED     NOT NULL      AUTO_INCREMENT,
  `name`        VARCHAR(50)         NOT NULL      DEFAULT ''
  COMMENT '文件名称',
  `url`         TEXT COMMENT '下载地址',
  `type`        TINYINT(1) UNSIGNED NOT NULL      DEFAULT '0'
  COMMENT '0:图片,1:文件',
  `format`      VARCHAR(10)         NOT NULL      DEFAULT ''
  COMMENT '以文件后缀为准',
  `descript`    VARCHAR(500)        NOT NULL      DEFAULT '0'
  COMMENT '文件描述',
  `size`        INT(11)             NOT NULL      DEFAULT '0'
  COMMENT '文件原大小',
  `re_size`     INT(11)             NOT NULL      DEFAULT '0'
  COMMENT '文件压缩后大小',
  `compress`    INT                 NOT NULL      DEFAULT '0'
  COMMENT '压缩次数',
  `downs`       INT                 NOT NULL      DEFAULT '0'
  COMMENT '下载次数',
  `create_id`   BIGINT              NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT              NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)             NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)             NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       BIGINT              NOT NULL      DEFAULT '0'
  COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='文件压缩管理';
#--------the tool---end-----#

#--------the app---start-----#
DROP TABLE IF EXISTS `zd_app_channel`; #应用渠道名称
CREATE TABLE `zd_app_channer` (
  `id`          BIGINT UNSIGNED    NOT NULL      AUTO_INCREMENT,
  `name`        VARCHAR(20) UNIQUE NOT NULL      DEFAULT ''
  COMMENT '渠道名称',
  `friend_id`   VARCHAR(30) UNIQUE NOT NULL      DEFAULT ''
  COMMENT '关联ID',
  `drescript`   TEXT
  COMMENT '描述',
  `create_id`   BIGINT             NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT             NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       BIGINT             NOT NULL      DEFAULT '0'
  COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='APP渠道管理';

DROP TABLE IF EXISTS `zd_app_name`; #应用名称
CREATE TABLE `zd_app_name` (
  `id`          BIGINT UNSIGNED    NOT NULL      AUTO_INCREMENT,
  `name`        VARCHAR(20) UNIQUE NOT NULL      DEFAULT ''
  COMMENT '应用名称',
  `descript`    TEXT
  COMMENT '描述',
  `create_id`   BIGINT             NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT             NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       BIGINT             NOT NULL      DEFAULT '0'
  COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='应用名称管理';

DROP TABLE IF EXISTS `zd_app_pkgs`; #应用包管理
CREATE TABLE `zd_app_pkgs` (
  `id`          BIGINT UNSIGNED    NOT NULL      AUTO_INCREMENT,
  `NAME`        VARCHAR(50) UNIQUE NOT NULL      DEFAULT ''
  COMMENT '应用包',
  `descript`    TEXT
  COMMENT '描述',
  `create_id`   BIGINT             NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT             NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       BIGINT             NOT NULL      DEFAULT '0'
  COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='应用包管理';

DROP TABLE IF EXISTS `zd_app_version`; #应用版本管理
CREATE TABLE `zd_app_version` (
  `id`          BIGINT UNSIGNED NOT NULL      AUTO_INCREMENT,
  `name`        VARCHAR(20)     NOT NULL      DEFAULT ''
  COMMENT '应用版本',
  `descript`    TEXT
  COMMENT '描述',
  `create_id`   BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)         NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)         NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='应用版本管理';

DROP TABLE IF EXISTS `zd_app_code`; #应用版本号管理
CREATE TABLE `zd_app_code` (
  `id`          BIGINT UNSIGNED NOT NULL      AUTO_INCREMENT,
  `code`        INT             NOT NULL      DEFAULT '0'
  COMMENT '应用版本号',
  `descript`    TEXT
  COMMENT '描述',
  `create_id`   BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)         NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)         NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='应用版本号管理';

DROP TABLE IF EXISTS `zd_app_env`; #应用环境管理
CREATE TABLE `zd_app_env` (
  `id`          BIGINT UNSIGNED NOT NULL      AUTO_INCREMENT,
  `name`        VARCHAR(20)     NOT NULL      DEFAULT ''
  COMMENT '应用环境',
  `descript`    TEXT
  COMMENT '描述',
  `create_id`   BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)         NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)         NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='应用环境管理';

DROP TABLE IF EXISTS `zd_app_build`; #构建类型
CREATE TABLE `zd_app_build` (
  `id`          BIGINT UNSIGNED NOT NULL      AUTO_INCREMENT,
  `name`        VARCHAR(10)     NOT NULL      DEFAULT ''
  COMMENT '构建类型:debug...',
  `descript`    TEXT
  COMMENT '描述',
  `create_id`   BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)         NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)         NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='应用构建类型管理';

DROP TABLE IF EXISTS `zd_app_type`; #应用平台类型
CREATE TABLE `zd_app_type` (
  `id`          BIGINT UNSIGNED NOT NULL      AUTO_INCREMENT,
  `name`        VARCHAR(30)     NOT NULL      DEFAULT ''
  COMMENT '应用平台类型,如:android',
  `descript`    TEXT
  COMMENT '描述',
  `create_id`   BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)         NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)         NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='应用类型管理';

DROP TABLE IF EXISTS `zd_app`; #渠道包管理
CREATE TABLE `zd_app` (
  `id`             BIGINT UNSIGNED NOT NULL      AUTO_INCREMENT,
  `project_id`     BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '项目ID',
  `test_id`        BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '测试ID',
  `icon`           VARCHAR(100)    NOT NULL NULL DEFAULT ''
  COMMENT 'Logo',
  `type_id`        BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '应用平台类型ID',
  `application_id` BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '应用名称ID',
  `pkg_id`         BIGINT UNIQUE   NOT NULL      DEFAULT '0'
  COMMENT '应用包ID',
  `version_id`     BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '应用版本ID',
  `code_id`        BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '应用版本号ID',
  `env_id`         BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '应用环境ID',
  `build_id`       BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '构建类型ID',
  `channel_id`     BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '渠道ID',
  `descript`       TEXT
  COMMENT '描述',
  `status`         TINYINT(4)      NOT NULL      DEFAULT '0'
  COMMENT '当前状态:0~无,-1~打包失败,1~打包中,2~打包成功,-3~测试失败,3~测试中,4~测试完成',
  `times`          INT             NOT NULL      DEFAULT '0'
  COMMENT '打包次数',
  `url`            VARCHAR(100)    NOT NULL      DEFAULT ''
  COMMENT '下载地址',
  `qr_img`         VARCHAR(100)    NOT NULL      DEFAULT ''
  COMMENT '二维码地址',
  `downs`          INT             NOT NULL      DEFAULT '0'
  COMMENT '下载次数',
  `create_id`      BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`      BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time`    INT(11)         NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time`    INT(11)         NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`          BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='Android应用包管理';

#--------the app---end-----#

#--------the test---start-----#

DROP TABLE IF EXISTS `zd_test_environment`; #系统环境设置
CREATE TABLE `zd_test_environment` (
  `id`          BIGINT UNSIGNED NOT NULL      AUTO_INCREMENT,
  `name`        VARCHAR(30)     NOT NULL      DEFAULT ''
  COMMENT '环境类型',
  `jdk`         VARCHAR(100)    NOT NULL      DEFAULT ''
  COMMENT 'jdk安装路径',
  `git`         VARCHAR(100)    NOT NULL      DEFAULT ''
  COMMENT 'Git安装路径',
  `gradle`      VARCHAR(100)    NOT NULL      DEFAULT ''
  COMMENT 'Gradle安装路径',
  `adb`         VARCHAR(100)    NOT NULL      DEFAULT ''
  COMMENT 'Adb安装路径',
  `create_id`   BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)         NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)         NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='系统环境管理';

DROP TABLE IF EXISTS `zd_test_project`; #GIT工程目录
CREATE TABLE `zd_test_project` (
  `id`          BIGINT UNSIGNED    NOT NULL      AUTO_INCREMENT,
  `parent_id`   BIGINT             NOT NULL      DEFAULT '0'
  COMMENT '系统环境ID',
  `name`        VARCHAR(50) UNIQUE NOT NULL      DEFAULT ''
  COMMENT '项目名称',
  `icon`        VARCHAR(100)       NOT NULL NULL DEFAULT ''
  COMMENT 'Logo',
  `address`     TEXT COMMENT '项目地址:git,svn,...',
  `account`     VARCHAR(30)        NOT NULL      DEFAULT ''
  COMMENT '账号名称',
  `psw`         VARCHAR(50)        NOT NULL      DEFAULT ''
  COMMENT '账号密码',
  `branch`      VARCHAR(50)        NOT NULL      DEFAULT ''
  COMMENT '分支',
  `tag`         VARCHAR(50)        NOT NULL      DEFAULT ''
  COMMENT '标签tag',
  `descript`    TEXT
  COMMENT '描述',
  `create_id`   BIGINT             NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT             NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       BIGINT             NOT NULL      DEFAULT '0'
  COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='项目工程管理';

DROP TABLE IF EXISTS `zd_test_test`; #测试类型与方法
CREATE TABLE `zd_test_test` (
  `id`          BIGINT UNSIGNED NOT NULL      AUTO_INCREMENT,
  `name`        VARCHAR(50)     NOT NULL      DEFAULT ''
  COMMENT '测试名称',
  `type`        TINYINT(4)      NOT NULL      DEFAULT '0'
  COMMENT '测试类型',
  `descript`    TEXT
  COMMENT '描述',
  `create_id`   BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)         NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)         NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       BIGINT          NOT NULL      DEFAULT '0'
  COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='测试状态管理';

#--------the test---end-----#

#--------the api---start-----#

DROP TABLE IF EXISTS `zd_api_key`; #api认证key管理
CREATE TABLE `zd_api_key` (
  `id`          BIGINT UNSIGNED     NOT NULL      AUTO_INCREMENT,
  `name`        VARCHAR(20) UNIQUE  NOT NULL      DEFAULT ''
  COMMENT '申请名称',
  `key`         VARCHAR(500) UNIQUE NOT NULL      DEFAULT ''
  COMMENT '申请的key',
  `descript`    TEXT
  COMMENT '描述',
  `create_id`   BIGINT              NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT              NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)             NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)             NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       BIGINT              NOT NULL      DEFAULT '0'
  COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='api申请key管理';

#--------the api---end-----#

#--------the web---start-----#

DROP TABLE IF EXISTS `zd_web_nav`; #官网菜单管理
CREATE TABLE `zd_web_nav` (
  `id`          BIGINT UNSIGNED     NOT NULL      AUTO_INCREMENT,
  `name`        VARCHAR(20) UNIQUE  NOT NULL      DEFAULT ''
  COMMENT '名称',
  `action`      VARCHAR(200) UNIQUE NOT NULL      DEFAULT ''
  COMMENT '触发事件',
  `menu`        TEXT COMMENT '二级菜单',
  `descript`    TEXT
  COMMENT '描述',
  `create_id`   BIGINT              NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT              NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)             NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)             NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       BIGINT              NOT NULL      DEFAULT '0'
  COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='官网导航菜单管理';

#--------the web---end-----#

#--------the other---start-----#

DROP TABLE IF EXISTS `zd_upload`; #文件上传
CREATE TABLE `zd_upload` (
  `id`          BIGINT UNSIGNED    NOT NULL      AUTO_INCREMENT,
  `cover`       VARCHAR(200)       NOT NULL      DEFAULT ''
  COMMENT '封面',
  `name`        VARCHAR(50) UNIQUE NOT NULL      DEFAULT ''
  COMMENT '文件名称',
  `descript`    VARCHAR(200)       NOT NULL      DEFAULT ''
  COMMENT '文件描述',
  `type`        INT                NOT NULL      DEFAULT '0'
  COMMENT '文件类型',
  `format`      VARCHAR(20)        NOT NULL      DEFAULT ''
  COMMENT '文件格式',
  `size`        BIGINT             NOT NULL      DEFAULT '0'
  COMMENT '文件尺寸',
  `width`       INT                NOT NULL      DEFAULT '0'
  COMMENT '图片宽度',
  `height`      INT                NOT NULL      DEFAULT '0'
  COMMENT '图片高度',
  `length`      INT                NOT NULL      DEFAULT '0'
  COMMENT '视频时长',
  `path`        VARCHAR(200)       NOT NULL      DEFAULT ''
  COMMENT '保存路径',
  `url`         VARCHAR(200)       NOT NULL      DEFAULT ''
  COMMENT '访问地址',
  `times`       INT                NOT NULL      DEFAULT '0'
  COMMENT '上传次数',
  `create_id`   BIGINT             NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   BIGINT             NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='文件上传管理';

#--------the other---end-----#


SET FOREIGN_KEY_CHECKS = 1;