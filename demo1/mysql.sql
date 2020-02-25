CREATE DATABASE IF NOT EXISTS demo DEFAULT CHARSET utf8 COLLATE utf8_general_ci;

use demo;

DROP TABLE IF EXISTS `demo`;

CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '设置主键自增',
  `name` varchar(25) NOT NULL COMMENT '设置名称',
  `age` tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '设置年龄',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `users` (`name`, `age`) VALUES
('张三', 25),
('李四', 22),
('田七', 25);