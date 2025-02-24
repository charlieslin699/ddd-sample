CREATE SCHEMA `auth`;

CREATE TABLE `auth`.`accountPermission` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `accountUID` CHAR(21) NOT NULL,
  `permissionUID` CHAR(21) NOT NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `auth`.`account` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `uid` CHAR(21) NOT NULL UNIQUE,
  `username` VARCHAR(45) NOT NULL UNIQUE COMMENT '帳號',
  `password` VARCHAR(75) NOT NULL COMMENT '密碼',
  `status` INT UNSIGNED NOT NULL COMMENT '狀態',
  `secret` CHAR(21) NOT NULL COMMENT '帳號金鑰',
  `createTime` DATETIME NOT NULL COMMENT '建立時間',
  `updateTime` DATETIME NOT NULL COMMENT '更新時間',
  PRIMARY KEY (`id`));

CREATE TABLE `auth`.`loginRecord` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `accountUID` CHAR(21) NOT NULL,
  `token` VARCHAR(300) NOT NULL COMMENT '登入token',
  `createTime` DATETIME NOT NULL COMMENT '登入時間',
  PRIMARY KEY (`id`));

CREATE TABLE `auth`.`permission` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `uid` CHAR(21) NOT NULL UNIQUE,
  `group` VARCHAR(20) NOT NULL COMMENT '功能群組',
  `name` VARCHAR(20) NOT NULL COMMENT '功能名稱',
  `key` VARCHAR(20) NOT NULL UNIQUE COMMENT '功能代碼',
  `langIndex` VARCHAR(20) NOT NULL COMMENT '語系索引',
  PRIMARY KEY (`id`));

CREATE TABLE `auth`.`thirdPartyVerification` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `accountUID` CHAR(21) NOT NULL,
  `verificationType` INT UNSIGNED NOT NULL COMMENT '驗證類型(1:FOTP)',
  `createTime` DATETIME NOT NULL COMMENT '建立時間',
  PRIMARY KEY (`id`));
