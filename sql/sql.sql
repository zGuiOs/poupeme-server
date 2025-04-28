CREATE DATABASE IF NOT EXISTS poupeme;

USE poupeme;

DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id int auto_increment primary key,
    uuid varchar(36) not null unique,
    name varchar(50) not null,
    email varchar(50) not null unique,
    password varchar(100) not null,
    created_at timestamp default current_timestamp(),
    updated_at timestamp default current_timestamp() on update current_timestamp(),
    INDEX idx_users_name (name)
) ENGINE = INNODB;
