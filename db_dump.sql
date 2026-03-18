-- Adminer 4.8.1 MySQL 8.0.29 dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `company`;
CREATE TABLE `company`
(
    `id`              binary(16)   NOT NULL DEFAULT (uuid_to_bin(uuid())),
    `name`            varchar(15)  NOT NULL,
    `description`     varchar(3000)         DEFAULT NULL,
    `total_employees` int unsigned NOT NULL,
    `is_registered`   tinyint      NOT NULL,
    `type_id`         int          NOT NULL,
    PRIMARY KEY (`id`),
    KEY `type_id` (`type_id`),
    CONSTRAINT `company_ibfk_1` FOREIGN KEY (`type_id`) REFERENCES `company_type` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

INSERT INTO `company` (`id`, `name`, `description`, `total_employees`, `is_registered`, `type_id`)
VALUES (UNHEX('330542E99A5B11EDAE360242C0A83002'), 'Company3', NULL, 100, 1, 1),
       (UNHEX('CAA09E5C9A5A11EDAE360242C0A83002'), 'Company1', 'Company1 description', 10, 1, 2),
       (UNHEX('D1C5428C9A5A11EDAE360242C0A83002'), 'Company2', 'Company2 description', 50, 0, 3);

DROP TABLE IF EXISTS `company_type`;
CREATE TABLE `company_type`
(
    `id`   int         NOT NULL AUTO_INCREMENT,
    `name` varchar(64) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

INSERT INTO `company_type` (`id`, `name`)
VALUES (1, 'Corporations'),
       (2, 'NonProfit'),
       (3, 'Cooperative'),
       (4, 'SoleProprietorship');

-- 2023-01-22 13:47:25