DROP TABLE IF EXISTS `secret`;
CREATE TABLE `credencials` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `name` varchar(200) NOT NULL,
  `account` varchar(200) NOT NULL,
  `password` varchar(320) NOT NULL,

  PRIMARY KEY (`id`),
  KEY `idx_user_deleted_at` (`deleted_at`),
  KEY `credencialsname` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `secret`;
CREATE TABLE `secret` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `name` varchar(200) NOT NULL,
  `secret` text NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_deleted_at` (`deleted_at`),
  KEY `secretname` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;