SET FOREIGN_KEY_CHECKS = 0;

CREATE TABLE `user`
(
    `id`                   int(11) NOT NULL AUTO_INCREMENT,
    `firstname`            varchar(255) NOT NULL,
    `lastname`             varchar(255) NOT NULL,
    `email`                varchar(255) NOT NULL,
    `password`             varchar(255) NOT NULL,
    `category_type_ref_id` int(11) NOT NULL,
    `created_by`           int(11) DEFAULT NULL,
    `last_updated_by`      int(11) DEFAULT NULL,
    `created_at`           timestamp    NOT NULL DEFAULT current_timestamp,
    `last_updated_at`      timestamp NULL DEFAULT NULL ON UPDATE current_timestamp,
    `is_active`            bool         NOT NULL DEFAULT TRUE,
    `reset_token`          varchar(255)          DEFAULT NULL,
    `address`              varchar(255)          DEFAULT NULL,
    `birthday`             date                  DEFAULT NULL,
    `gender`               char(1)               DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `email` (`email`),
    KEY                    `fk_category_type` (`category_type_ref_id`),
    KEY                    `fk_created_by` (`created_by`),
    KEY                    `fk_created_last_updated_by` (`last_updated_by`),
    CONSTRAINT `fk_category_type` FOREIGN KEY (`category_type_ref_id`) REFERENCES `category` (`id`) ON DELETE RESTRICT,
    CONSTRAINT `fk_created_by` FOREIGN KEY (`created_by`) REFERENCES `user` (`id`) ON DELETE RESTRICT,
    CONSTRAINT `fk_created_last_updated_by` FOREIGN KEY (`last_updated_by`) REFERENCES `user` (`id`) ON DELETE RESTRICT
);

SET FOREIGN_KEY_CHECKS = 1;
