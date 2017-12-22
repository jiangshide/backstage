SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

#--------the user---start------#

DROP TABLE IF EXISTS `zd_uc_role_auth`;
CREATE TABLE `zd_uc_role_auth` (
  `role_id` INT(11) UNSIGNED NOT NULL DEFAULT '0'
  COMMENT '角色ID',
  `auth_id` INT(11) UNSIGNED NOT NULL DEFAULT '0'
  COMMENT '权限ID',
  PRIMARY KEY (`role_id`, `auth_id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='权限和角色关系管理';

BEGIN;
INSERT INTO `zd_uc_role_auth`
VALUES ('1', '16'), ('1', '17'), ('1', '18'), ('1', '19'), ('2', '0'), ('2', '1'), ('2', '15'), ('2', '20'),
  ('2', '21'), ('2', '22'), ('2', '23'), ('2', '24');
COMMIT;

DROP TABLE IF EXISTS `zd_uc_role`;
CREATE TABLE `zd_uc_role` (
  `id`          INT(11) UNSIGNED    NOT NULL AUTO_INCREMENT
  COMMENT '主键ID',
  `name`        VARCHAR(32)         NOT NULL DEFAULT '0'
  COMMENT '角色名称',
  `detail`      VARCHAR(255)        NOT NULL DEFAULT '0'
  COMMENT '备注',
  `create_id`   INT(11) UNSIGNED    NOT NULL DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11) UNSIGNED    NOT NULL DEFAULT '0'
  COMMENT '修改ID',
  `status`      TINYINT(1) UNSIGNED NOT NULL DEFAULT '1'
  COMMENT '状态1:正常,0:删除',
  `create_time` INT(11) UNSIGNED    NOT NULL DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11) UNSIGNED    NOT NULL DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = Innodb
  DEFAULT CHARSET = utf8
  COMMENT ='角色表';

BEGIN;
INSERT INTO `zd_uc_role` VALUES ('1', 'API管理员', '拥有API所有权限', '0', '2', '1', '1505874156', '1505874156'),
  ('2', '系统管理员', '系统管理员', '0', '0', '1', '1506124114', '1506124114');
COMMIT;

DROP TABLE IF EXISTS `zd_uc_auth`;
CREATE TABLE `zd_uc_auth` (
  `id`          INT(11) UNSIGNED    NOT NULL AUTO_INCREMENT
  COMMENT '自增ID',
  `pid`         INT(11) UNSIGNED    NOT NULL DEFAULT '0'
  COMMENT '上级ID,0为顶级',
  `name`        VARCHAR(64)         NOT NULL DEFAULT '0'
  COMMENT '权限名称',
  `url`         VARCHAR(255)        NOT NULL DEFAULT '0'
  COMMENT 'URL地址',
  `sort`        INT(1) UNSIGNED     NOT NULL DEFAULT '999'
  COMMENT '排序,从小到大',
  `icon`        VARCHAR(255)        NOT NULL DEFAULT '0'
  COMMENT '图标',
  `is_show`     TINYINT(1) UNSIGNED NOT NULL DEFAULT '0'
  COMMENT '是否显示,0:隐藏,1:显示',
  `user_id`     INT(11) UNSIGNED    NOT NULL DEFAULT '0'
  COMMENT '操作者ID',
  `create_id`   INT(11) UNSIGNED    NOT NULL DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11) UNSIGNED    NOT NULL DEFAULT '0'
  COMMENT '修改者ID',
  `status`      TINYINT(1) UNSIGNED NOT NULL DEFAULT '0'
  COMMENT '状态,1:正常,0:删除',
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
INSERT INTO `zd_uc_auth`
VALUES ('1', '0', '所有权限', '/backstage', '1', '', '0', '1', '1', '1', '1', '1505620970', '1505620970'),
  ('2', '1', '权限管理', '/backstage', '999', 'fa-id-card', '1', '1', '0', '1', '1', '0', '1505622360'),
  ('3', '2', '管理员', '/backstage/admin/list', '1', 'fa-user-o', '1', '1', '1', '1', '1', '1505621186', '1505621186'),
  ('4', '2', '角色管理', '/backstage/role/list', '2', 'fa-user-circle-o', '1', '1', '0', '1', '1', '0', '1505621852'),
  ('5', '3', '新增', '/backstage/admin/add', '1', '', '0', '1', '0', '1', '1', '0', '1505621685'),
  ('6', '3', '修改', '/backstage/admin/edit', '2', '', '0', '1', '0', '1', '1', '0', '1505621697'),
  ('7', '3', '删除', '/backstage/admin/ajaxdel', '3', '', '0', '1', '1', '1', '1', '1505621756', '1505621756'),
  ('8', '4', '新增', '/backstage/role/add', '1', '', '1', '1', '0', '1', '1', '0', '1505698716'),
  ('9', '4', '修改', '/backstage/role/edit', '2', '', '0', '1', '1', '1', '1', '1505621912', '1505621912'),
  ('10', '4', '删除', '/backstage/role/ajaxdel', '3', '', '0', '1', '1', '1', '1', '1505621951', '1505621951'),
  ('11', '2', '权限因子', '/backstage/auth/list', '3', 'fa-list', '1', '1', '1', '1', '1', '1505621986', '1505621986'),
  ('12', '11', '新增', '/backstage/auth/add', '1', '', '0', '1', '1', '1', '1', '1505622009', '1505622009'),
  ('13', '11', '修改', '/backstage/auth/edit', '2', '', '0', '1', '1', '1', '1', '1505622047', '1505622047'),
  ('14', '11', '删除', '/backstage/auth/ajaxdel', '3', '', '0', '1', '1', '1', '1', '1505622111', '1505622111'),
  ('15', '1', '个人中心', '/backstage/profile/edit', '1001', 'fa-user-circle-o', '1', '1', '0', '1', '1', '0', '1506001114'),
  ('16', '1', 'API管理', '/backstage', '1', 'fa-cubes', '1', '0', '0', '0', '1', '0', '1506125698'),
  ('17', '16', 'API接口', '/backstage/api/list', '1', 'fa-link', '1', '1', '1', '1', '1', '1505622447', '1505622447'),
  ('19', '16', 'API监控', '/backstage/apimonitor/list', '3', 'fa-bar-chart', '1', '1', '0', '1', '1', '0', '1507700851'),
  ('20', '1', '基础设置', '/backstage/', '2', 'fa-cogs', '1', '1', '1', '1', '1', '1505622601', '1505622601'),
  ('21', '20', '分组设置', '/backstage/group/list', '1', 'fa-object-ungroup', '1', '1', '1', '1', '1', '1505622645', '1505622645'),
  ('22', '20', '环境设置', '/backstage/env/list', '2', 'fa-tree', '1', '1', '1', '1', '1', '1505622681', '1505622681'),
  ('23', '20', '状态码设置', '/backstage/code/list', '3', 'fa-code', '1', '1', '1', '1', '1', '1505622728', '1505622728'),
  ('24', '15', '资料修改', '/backstage/user/edit', '1', 'fa-edit', '1', '1', '0', '1', '1', '0', '1506057468'),
  ('25', '21', '新增', '/backstage/group/add', '1', 'n', '1', '0', '0', '0', '1', '1506229739', '1506229739'),
  ('26', '21', '修改', '/backstage/group/edit', '2', 'fa', '0', '0', '0', '0', '1', '1506237920', '1506237920'),
  ('27', '21', '删除', '/backstage/group/ajaxdel', '3', 'fa', '0', '0', '0', '0', '1', '1506237948', '1506237948'),
  ('28', '22', '新增', '/backstage/env/add', '1', 'fa', '0', '0', '0', '0', '1', '1506316506', '1506316506'),
  ('29', '22', '修改', '/backstage/env/edit', '2', 'fa', '0', '0', '0', '0', '1', '1506316532', '1506316532'),
  ('30', '22', '删除', '/backstage/env/ajaxdel', '3', 'fa', '0', '0', '0', '0', '1', '1506316567', '1506316567'),
  ('31', '23', '新增', '/backstage/code/add', '1', 'fa', '0', '0', '0', '0', '1', '1506327812', '1506327812'),
  ('32', '23', '修改', '/backstage/code/edit', '2', 'fa', '0', '0', '0', '0', '1', '1506327831', '1506327831'),
  ('33', '23', '删除', '/backstage/code/ajadel', '3', 'fa', '0', '0', '0', '0', '1', '1506327857', '1506327857'),
  ('34', '17', '新增资源', '/backstage/api/add', '1', 'fa-link', '1', '1', '0', '1', '1', '0', '1507436029'),
  ('35', '17', '修改资源', '/backstage/api/edit', '2', 'fa-link', '1', '1', '0', '1', '1', '0', '1507436042'),
  ('36', '17', '删除资源', '/backstage/api/ajaxdel', '3', 'fa-link', '1', '1', '0', '1', '1', '0', '1507436052'),
  ('37', '17', '新增接口', '/backstage/api/addapi', '4', '', '0', '1', '1', '1', '1', '1507436014', '1507436014'),
  ('38', '17', '修改接口', '/backstage/api/editapi', '5', '', '0', '1', '0', '1', '1', '0', '1507705049');
COMMIT;

DROP TABLE IF EXISTS `zd_uc_admin`;
CREATE TABLE `zd_uc_admin` (
  `id`          INT(11) UNSIGNED NOT NULL AUTO_INCREMENT
  COMMENT '自增ID',
  `name`        VARCHAR(20)      NOT NULL DEFAULT ''
  COMMENT '用户名',
  `real_name`   VARCHAR(32)      NOT NULL DEFAULT '0'
  COMMENT '真实姓名',
  `password`    CHAR(32)         NOT NULL DEFAULT ''
  COMMENT '密码',
  `role_ids`    VARCHAR(255)     NOT NULL DEFAULT '0'
  COMMENT '角色id字符串,eg:2,3,4,5',
  `phone`       VARCHAR(20)      NOT NULL DEFAULT '0'
  COMMENT '手机号码',
  `motto`       TEXT COMMENT '个人格言',
  `sex`         TINYINT(1)       NOT NULL DEFAULT '0'
  COMMENT '性别,0:男,1:女',
  `email`       VARCHAR(50)      NOT NULL DEFAULT ''
  COMMENT '邮箱',
  `salt`        CHAR(10)         NOT NULL DEFAULT ''
  COMMENT '密码加盐',
  `last_login`  INT(11)          NOT NULL DEFAULT '0'
  COMMENT '最后登录时间',
  `last_ip`     CHAR(15)         NOT NULL DEFAULT ''
  COMMENT '最后登录IP',
  `status`      TINYINT(4)       NOT NULL DEFAULT '0'
  COMMENT '状态,1:正常,0:禁用',
  `create_id`   INT(11) UNSIGNED NOT NULL DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11) UNSIGNED NOT NULL DEFAULT '0'
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
  COMMENT ='管理员表';

#--------the user---end------#

#--------the set---start------#

DROP TABLE IF EXISTS `zd_set_group`;
CREATE TABLE `zd_set_group` (
  `id`          INT(11) UNSIGNED    NOT NULL AUTO_INCREMENT
  COMMENT '自增',
  `name`        VARCHAR(50)         NOT NULL DEFAULT ''
  COMMENT '组名',
  `detail`      VARCHAR(255)        NOT NULL DEFAULT ''
  COMMENT '说明',
  `status`      TINYINT(1) UNSIGNED NOT NULL DEFAULT '0'
  COMMENT '状态:1:正常,0:删除',
  `create_id`   INT(11)             NOT NULL DEFAULT '0'
  COMMENT '用户ID',
  `update_id`   INT(11)             NOT NULL DEFAULT '0'
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
  `id`          INT(11) UNSIGNED    NOT NULL AUTO_INCREMENT,
  `name`        VARCHAR(50)         NOT NULL DEFAULT ''
  COMMENT '环境名称',
  `env_host`    VARCHAR(255)        NOT NULL DEFAULT '0'
  COMMENT '主机',
  `detail`      VARCHAR(255)        NOT NULL DEFAULT '0'
  COMMENT '备注',
  `status`      TINYINT(4) UNSIGNED NOT NULL DEFAULT '0'
  COMMENT '状态,1:正常,0:禁用',
  `create_id`   INT(11) UNSIGNED    NOT NULL DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11) UNSIGNED    NOT NULL DEFAULT '0'
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
  `id`          INT(11) UNSIGNED    NOT NULL AUTO_INCREMENT,
  `code`        VARCHAR(50)         NOT NULL DEFAULT '0'
  COMMENT '状态码',
  `descript`    VARCHAR(255)        NOT NULL DEFAULT '0'
  COMMENT '描述',
  `detail`      VARCHAR(255)        NOT NULL DEFAULT '0'
  COMMENT '备注',
  `status`      TINYINT(4) UNSIGNED NOT NULL DEFAULT '0'
  COMMENT '状态,1:正常,0:禁用',
  `create_id`   INT(11) UNSIGNED    NOT NULL DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11) UNSIGNED    NOT NULL DEFAULT '0'
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
  `id`          INT(11) UNSIGNED    NOT NULL AUTO_INCREMENT,
  `group_id`    INT(11)             NOT NULL DEFAULT '0'
  COMMENT '分组ID',
  `name`        VARCHAR(50)         NOT NULL DEFAULT '0'
  COMMENT '接口名称',
  `status`      TINYINT(1) UNSIGNED NOT NULL DEFAULT '1'
  COMMENT '状态,1:审核通过,0:暂停使用,2:草稿,3:审核中',
  `audit_id`    INT(11) UNSIGNED    NOT NULL DEFAULT '0'
  COMMENT '审核人ID',
  `create_id`   INT(11) UNSIGNED    NOT NULL DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11) UNSIGNED    NOT NULL DEFAULT '0'
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
  `id`          INT(11) UNSIGNED    NOT NULL AUTO_INCREMENT,
  `detail_id`   INT(11) UNSIGNED    NOT NULL DEFAULT '0'
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
  `create_id`   INT(11) UNSIGNED    NOT NULL DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11) UNSIGNED    NOT NULL DEFAULT '0'
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
  `id`            INT(11) UNSIGNED    NOT NULL AUTO_INCREMENT,
  `src_id`        INT(11)             NOT NULL DEFAULT '0'
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
  `audit_id`      INT(11)             NOT NULL DEFAULT '0'
  COMMENT '审核人ID',
  `status`        TINYINT(1) UNSIGNED NOT NULL DEFAULT '1'
  COMMENT '状态,0:暂停 使用,1:正在审核,2:审核通过',
  `create_id`     INT(11)             NOT NULL DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`     INT(11)             NOT NULL DEFAULT '0'
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
  `id`          INT(11) UNSIGNED     NOT NULL AUTO_INCREMENT,
  `name`        VARCHAR(25) UNIQUE   NOT NULL DEFAULT ''
  COMMENT '名称',
  `icon`        VARCHAR(100)         NOT NULL DEFAULT ''
  COMMENT 'LOGO',
  `create_id`   INT(11)              NOT NULL DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11)              NOT NULL DEFAULT '0'
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
DROP TABLE IF EXISTS `zd_area_continent`; #洲
CREATE TABLE `zd_area_continent` (
  `id`          INT(11) UNSIGNED   NOT NULL      AUTO_INCREMENT,
  `name`        VARCHAR(50) UNIQUE NOT NULL      DEFAULT ''
  COMMENT '洲名',
  `icon`        VARCHAR(100)       NOT NULL NULL DEFAULT ''
  COMMENT 'LOGO',
  `create_id`   INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='所属洲';

DROP TABLE IF EXISTS `zd_area_state`; #国家
CREATE TABLE `zd_area_state` (
  `id`          INT(11) UNSIGNED   NOT NULL      AUTO_INCREMENT,
  `parent_id`   INT(11) UNSIGNED   NOT NULL      DEFAULT '0'
  COMMENT '所属洲',
  `name`        VARCHAR(50) UNIQUE NOT NULL      DEFAULT ''
  COMMENT '国名',
  `icon`        VARCHAR(100)       NOT NULL NULL DEFAULT ''
  COMMENT 'LOGO',
  `create_id`   INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='所属国家';

DROP TABLE IF EXISTS `zd_area_province`; #省
CREATE TABLE `zd_area_province` (
  `id`          INT(11) UNSIGNED NOT NULL      AUTO_INCREMENT,
  `parent_id`   INT(11) UNSIGNED NOT NULL      DEFAULT '0'
  COMMENT '所属国家',
  `name`        VARCHAR(50)      NOT NULL      DEFAULT ''
  COMMENT '省名',
  `icon`        VARCHAR(100)     NOT NULL NULL DEFAULT ''
  COMMENT 'LOGO',
  `create_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='所属省';

DROP TABLE IF EXISTS `zd_area_city`; #市
CREATE TABLE `zd_area_city` (
  `id`          INT(11) UNSIGNED NOT NULL      AUTO_INCREMENT,
  `parent_id`   INT(11) UNSIGNED NOT NULL      DEFAULT '0'
  COMMENT '所属省',
  `name`        VARCHAR(50)      NOT NULL      DEFAULT ''
  COMMENT '市名',
  `icon`        VARCHAR(100)     NOT NULL NULL DEFAULT ''
  COMMENT 'LOGO',
  `create_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='所属市';

DROP TABLE IF EXISTS `zd_area_region`; #地区
CREATE TABLE `zd_area_region` (
  `id`          INT(11) UNSIGNED NOT NULL      AUTO_INCREMENT,
  `parent_id`   INT(11) UNSIGNED NOT NULL      DEFAULT '0'
  COMMENT '所属城市',
  `name`        VARCHAR(50)      NOT NULL      DEFAULT ''
  COMMENT '地区名',
  `icon`        VARCHAR(100)     NOT NULL NULL DEFAULT ''
  COMMENT 'LOGO',
  `create_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='所属地区';

DROP TABLE IF EXISTS `zd_area_county`; #县
CREATE TABLE `zd_area_county` (
  `id`          INT(11) UNSIGNED NOT NULL      AUTO_INCREMENT,
  `parent_id`   INT(11) UNSIGNED NOT NULL      DEFAULT '0'
  COMMENT '所属地区',
  `name`        VARCHAR(50)      NOT NULL      DEFAULT ''
  COMMENT '县名',
  `icon`        VARCHAR(100)     NOT NULL NULL DEFAULT ''
  COMMENT 'LOGO',
  `create_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='所属县';

DROP TABLE IF EXISTS `zd_area_town`; #镇
CREATE TABLE `zd_area_town` (
  `id`          INT(11) UNSIGNED NOT NULL      AUTO_INCREMENT,
  `parent_id`   INT(11) UNSIGNED NOT NULL      DEFAULT '0'
  COMMENT '所属县',
  `name`        VARCHAR(50)      NOT NULL      DEFAULT ''
  COMMENT '镇名',
  `icon`        VARCHAR(100)     NOT NULL NULL DEFAULT ''
  COMMENT 'LOGO',
  `create_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='所属镇';

DROP TABLE IF EXISTS `zd_area_country`; #乡
CREATE TABLE `zd_area_country` (
  `id`          INT(11) UNSIGNED NOT NULL      AUTO_INCREMENT,
  `parent_id`   INT(11) UNSIGNED NOT NULL      DEFAULT '0'
  COMMENT '所属镇',
  `name`        VARCHAR(50)      NOT NULL      DEFAULT ''
  COMMENT '乡名',
  `icon`        VARCHAR(100)     NOT NULL NULL DEFAULT ''
  COMMENT 'LOGO',
  `create_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='所属乡';

DROP TABLE IF EXISTS `zd_area_village`; #村
CREATE TABLE `zd_area_village` (
  `id`          INT(11) UNSIGNED NOT NULL      AUTO_INCREMENT,
  `parent_id`   INT(11) UNSIGNED NOT NULL      DEFAULT '0'
  COMMENT '所属乡',
  `name`        VARCHAR(50)      NOT NULL      DEFAULT ''
  COMMENT '村名',
  `icon`        VARCHAR(100)     NOT NULL NULL DEFAULT ''
  COMMENT 'LOGO',
  `create_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='所属村';

DROP TABLE IF EXISTS `zd_area_group`; #组
CREATE TABLE `zd_area_group` (
  `id`          INT(11) UNSIGNED NOT NULL      AUTO_INCREMENT,
  `parent_id`   INT(11) UNSIGNED NOT NULL      DEFAULT '0'
  COMMENT '所属村',
  `name`        VARCHAR(50)      NOT NULL      DEFAULT ''
  COMMENT '组名',
  `icon`        VARCHAR(100)     NOT NULL NULL DEFAULT ''
  COMMENT 'LOGO',
  `create_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='所属组';

DROP TABLE IF EXISTS `zd_area_team`; #队
CREATE TABLE `zd_area_team` (
  `id`          INT(11) UNSIGNED NOT NULL      AUTO_INCREMENT,
  `parent_id`   INT(11) UNSIGNED NOT NULL      DEFAULT '0'
  COMMENT '所属组',
  `name`        VARCHAR(50)      NOT NULL      DEFAULT ''
  COMMENT '队名',
  `icon`        VARCHAR(100)     NOT NULL NULL DEFAULT ''
  COMMENT 'LOGO',
  `create_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='所属队';

#--------the test---end-----#

#--------the web---start-----#
DROP TABLE IF EXISTS `zd_web_banner`; #焦点图
CREATE TABLE `zd_web_banner` (
  `id`          INT(11) UNSIGNED   NOT NULL      AUTO_INCREMENT,
  `name`        VARCHAR(50) UNIQUE NOT NULL      DEFAULT ''
  COMMENT '名称',
  `link`        TEXT COMMENT '跳转新页面',
  `icon`        VARCHAR(100)       NOT NULL NULL DEFAULT ''
  COMMENT '图片展示',
  `descript`    VARCHAR(500)       NOT NULL      DEFAULT '0'
  COMMENT '描述',
  `clicks`      INT                NOT NULL      DEFAULT '0'
  COMMENT '点击次数',
  `create_id`   INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='焦点图管理';

#--------the web---end-----#

#--------the area---start-----#

#--------the area---end-----#


#--------the tool---start-----#

DROP TABLE IF EXISTS `zd_tools_qrcode`; #二维码管理
CREATE TABLE `zd_tools_qrcode` (
  `id`          INT(11) UNSIGNED   NOT NULL      AUTO_INCREMENT,
  `name`        VARCHAR(50) UNIQUE NOT NULL      DEFAULT ''
  COMMENT '名称',
  `content`     TEXT COMMENT '内容',
  `url`         TEXT COMMENT '地址',
  `descript`    TEXT COMMENT '描述',
  `create_id`   INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11)            NOT NULL      DEFAULT '0'
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
  `id`          INT(11) UNSIGNED   NOT NULL      AUTO_INCREMENT,
  `name`        VARCHAR(50) UNIQUE NOT NULL      DEFAULT ''
  COMMENT '文件类型名称',
  `descript`    TEXT COMMENT '文件类型描述',
  `create_id`   INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11)            NOT NULL      DEFAULT '0'
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
  `id`          INT(11) UNSIGNED     NOT NULL      AUTO_INCREMENT,
  `parent_id`   INT(11) UNSIGNED     NOT NULL      DEFAULT '0'
  COMMENT '所格式类型',
  `name`        VARCHAR(50) UNIQUE   NOT NULL      DEFAULT ''
  COMMENT '格式民称',
  `descript`    TEXT COMMENT '格式描述',
  `create_id`   INT(11)              NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11)              NOT NULL      DEFAULT '0'
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
  `id`          INT(11) UNSIGNED    NOT NULL      AUTO_INCREMENT,
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
  `create_id`   INT(11)             NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11)             NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)             NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)             NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       INT(11)             NOT NULL      DEFAULT '0'
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
  `id`          INT(11) UNSIGNED   NOT NULL      AUTO_INCREMENT,
  `name`        VARCHAR(20) UNIQUE NOT NULL      DEFAULT ''
  COMMENT '渠道名称',
  `friend_id`   VARCHAR(30) UNIQUE NOT NULL      DEFAULT ''
  COMMENT '关联ID',
  `drescript`   TEXT
  COMMENT '描述',
  `create_id`   INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='APP渠道管理';

DROP TABLE IF EXISTS `zd_app_name`; #应用名称
CREATE TABLE `zd_app_name` (
  `id`          INT(11) UNSIGNED   NOT NULL      AUTO_INCREMENT,
  `name`        VARCHAR(20) UNIQUE NOT NULL      DEFAULT ''
  COMMENT '应用名称',
  `descript`    TEXT
  COMMENT '描述',
  `create_id`   INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='应用名称管理';

DROP TABLE IF EXISTS `zd_app_pkgs`; #应用包管理
CREATE TABLE `zd_app_pkgs` (
  `id`          INT(11) UNSIGNED   NOT NULL      AUTO_INCREMENT,
  `NAME`        VARCHAR(50) UNIQUE NOT NULL      DEFAULT ''
  COMMENT '应用包',
  `descript`    TEXT
  COMMENT '描述',
  `create_id`   INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='应用包管理';

DROP TABLE IF EXISTS `zd_app_version`; #应用版本管理
CREATE TABLE `zd_app_version` (
  `id`          INT(11) UNSIGNED NOT NULL      AUTO_INCREMENT,
  `name`        VARCHAR(20)      NOT NULL      DEFAULT ''
  COMMENT '应用版本',
  `descript`    TEXT
  COMMENT '描述',
  `create_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='应用版本管理';

DROP TABLE IF EXISTS `zd_app_code`; #应用版本号管理
CREATE TABLE `zd_app_code` (
  `id`          INT(11) UNSIGNED NOT NULL      AUTO_INCREMENT,
  `code`        INT              NOT NULL      DEFAULT '0'
  COMMENT '应用版本号',
  `descript`    TEXT
  COMMENT '描述',
  `create_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='应用版本号管理';

DROP TABLE IF EXISTS `zd_app_env`; #应用环境管理
CREATE TABLE `zd_app_env` (
  `id`          INT(11) UNSIGNED NOT NULL      AUTO_INCREMENT,
  `name`        VARCHAR(20)      NOT NULL      DEFAULT ''
  COMMENT '应用环境',
  `descript`    TEXT
  COMMENT '描述',
  `create_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='应用环境管理';

DROP TABLE IF EXISTS `zd_app_build`; #构建类型
CREATE TABLE `zd_app_build` (
  `id`          INT(11) UNSIGNED NOT NULL      AUTO_INCREMENT,
  `name`        VARCHAR(10)      NOT NULL      DEFAULT ''
  COMMENT '构建类型:debug...',
  `descript`    TEXT
  COMMENT '描述',
  `create_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='应用构建类型管理';

DROP TABLE IF EXISTS `zd_app_type`; #应用平台类型
CREATE TABLE `zd_app_type` (
  `id`          INT(11) UNSIGNED NOT NULL      AUTO_INCREMENT,
  `name`        VARCHAR(30)      NOT NULL      DEFAULT ''
  COMMENT '应用平台类型,如:android',
  `descript`    TEXT
  COMMENT '描述',
  `create_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='应用类型管理';

DROP TABLE IF EXISTS `zd_app`; #渠道包管理
CREATE TABLE `zd_app` (
  `id`             INT(11) UNSIGNED NOT NULL      AUTO_INCREMENT,
  `project_id`     INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '项目ID',
  `test_id`        INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '测试ID',
  `icon`           VARCHAR(100)     NOT NULL NULL DEFAULT ''
  COMMENT 'Logo',
  `type_id`        INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '应用平台类型ID',
  `application_id` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '应用名称ID',
  `pkg_id`         INT(11) UNIQUE   NOT NULL      DEFAULT '0'
  COMMENT '应用包ID',
  `version_id`     INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '应用版本ID',
  `code_id`        INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '应用版本号ID',
  `env_id`         INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '应用环境ID',
  `build_id`       INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '构建类型ID',
  `channel_id`     INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '渠道ID',
  `descript`       TEXT
  COMMENT '描述',
  `status`         TINYINT(4)       NOT NULL      DEFAULT '0'
  COMMENT '当前状态:0~无,-1~打包失败,1~打包中,2~打包成功,-3~测试失败,3~测试中,4~测试完成',
  `times`          INT              NOT NULL      DEFAULT '0'
  COMMENT '打包次数',
  `url`            VARCHAR(100)     NOT NULL      DEFAULT ''
  COMMENT '下载地址',
  `qr_img`         VARCHAR(100)     NOT NULL      DEFAULT ''
  COMMENT '二维码地址',
  `downs`          INT              NOT NULL      DEFAULT '0'
  COMMENT '下载次数',
  `create_id`      INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`      INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time`    INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time`    INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`          INT(11)          NOT NULL      DEFAULT '0'
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
  `id`          INT(11) UNSIGNED NOT NULL      AUTO_INCREMENT,
  `name`        VARCHAR(30)      NOT NULL      DEFAULT ''
  COMMENT '环境类型',
  `jdk`         VARCHAR(100)     NOT NULL      DEFAULT ''
  COMMENT 'jdk安装路径',
  `git`         VARCHAR(100)     NOT NULL      DEFAULT ''
  COMMENT 'Git安装路径',
  `gradle`      VARCHAR(100)     NOT NULL      DEFAULT ''
  COMMENT 'Gradle安装路径',
  `adb`         VARCHAR(100)     NOT NULL      DEFAULT ''
  COMMENT 'Adb安装路径',
  `create_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='系统环境管理';

DROP TABLE IF EXISTS `zd_test_project`; #GIT工程目录
CREATE TABLE `zd_test_project` (
  `id`          INT(11) UNSIGNED   NOT NULL      AUTO_INCREMENT,
  `parent_id`   INT(11)            NOT NULL      DEFAULT '0'
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
  `create_id`   INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       INT(11)            NOT NULL      DEFAULT '0'
  COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='项目工程管理';

DROP TABLE IF EXISTS `zd_test_test`; #测试类型与方法
CREATE TABLE `zd_test_test` (
  `id`          INT(11) UNSIGNED NOT NULL      AUTO_INCREMENT,
  `name`        VARCHAR(50)      NOT NULL      DEFAULT ''
  COMMENT '测试名称',
  `type`        TINYINT(4)       NOT NULL      DEFAULT '0'
  COMMENT '测试类型',
  `descript`    TEXT
  COMMENT '描述',
  `create_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='测试状态管理';

#--------the test---end-----#

#--------the api---start-----#

DROP TABLE IF EXISTS `zd_api_key`;#api认证key管理
CREATE TABLE `zd_api_key` (
  `id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(20) UNIQUE NOT NULL DEFAULT '' COMMENT '申请名称',
  `key` VARCHAR(500) UNIQUE NOT NULL DEFAULT '' COMMENT '申请的key',
  `descript`    TEXT
  COMMENT '描述',
  `create_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建者ID',
  `update_id`   INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '修改者ID',
  `create_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '创建时间',
  `update_time` INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '更新时间',
  `views`       INT(11)          NOT NULL      DEFAULT '0'
  COMMENT '当前页面展示次数',
  PRIMARY KEY (`id`)
)ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='api申请key管理';

#--------the api---end-----#

#--------the other---start-----#

#--------the other---end-----#


SET FOREIGN_KEY_CHECKS = 1;