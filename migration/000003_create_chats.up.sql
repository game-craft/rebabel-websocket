CREATE TABLE `chats` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `worlds_id` INT NOT NULL,
    `chats_content` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP NOT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

ALTER TABLE chats
ADD CONSTRAINT fk_chats_worlds
FOREIGN KEY (worlds_id)
REFERENCES worlds(id);