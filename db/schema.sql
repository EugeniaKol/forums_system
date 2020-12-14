CREATE SCHEMA IF NOT EXISTS `forums_sys` DEFAULT CHARACTER SET utf8 ;

-- Schema forums
USE `forums_sys` ;

-- Create tables.
DROP TABLE IF EXISTS `forums`;
CREATE TABLE `forums`
(
    `id`   INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(50) NOT NULL UNIQUE,
    `topicKeyword` VARCHAR(50) NOT NULL
);


DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    `id`   INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `nickname` VARCHAR(50) NOT NULL UNIQUE
);

DROP TABLE IF EXISTS `interests`;
CREATE TABLE `interests`
(
    `id`   INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `users_id` INT NOT NULL,
    `forums_id` INT NOT NULL,
    INDEX `fk_interests_users_idx` (`users_id` ASC) VISIBLE,
	INDEX `fk_interests_forums_idx` (`forums_id` ASC) VISIBLE,
	CONSTRAINT `fk_interests_users`
    FOREIGN KEY (`users_id`)
    REFERENCES `users` (`id`),
	CONSTRAINT `fk_interests_forums`
    FOREIGN KEY (`forums_id`)
    REFERENCES `forums` (`id`)

);

-- Insert demo data.
INSERT INTO `forums` (`name`, `topicKeyword`) VALUES ('Політика в Україні', 'ukraine-politics');
INSERT INTO `users` (`nickname`) VALUES ('user1');
