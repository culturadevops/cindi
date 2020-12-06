DROP TABLE IF EXISTS `credential`;
CREATE TABLE `credential` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `owner` varchar(50) NOT NULL,
  `name` varchar(100) NOT NULL,
  `account` varchar(200) NOT NULL,
  `password` varchar(320) NOT NULL,

  PRIMARY KEY (`id`),
  KEY `idx_user_deleted_at` (`deleted_at`),
  KEY `credentialsname` (`name`),
    KEY `credentialowner` (`owner`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `secret`;
CREATE TABLE `secret` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `owner` varchar(50) NOT NULL,
  `name` varchar(100) NOT NULL,
  `secret` json NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_deleted_at` (`deleted_at`),
  KEY `secretname` (`name`),
  KEY `secretowner` (`owner`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;