CREATE TABLE `worlds` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `worlds_name` VARCHAR(255) NOT NULL,
    `worlds_status` VARCHAR(255) NOT NULL,
    `worlds_count_people` INT NOT NULL,
    `created_at` TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP NOT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;