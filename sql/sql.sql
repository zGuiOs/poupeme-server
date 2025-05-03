CREATE DATABASE IF NOT EXISTS poupeme;

USE poupeme;

DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id int auto_increment primary key,
    uuid varchar(36) not null unique,
    name varchar(50) not null,
    email varchar(255) not null unique,
    password varchar(100) not null,
    status enum ('ATIVO', 'DESATIVADO', 'SUSPENSO') not null default 'ATIVO',
    created_at timestamp default current_timestamp(),
    updated_at timestamp default current_timestamp() on update current_timestamp(),
    INDEX idx_users_name (name),
    INDEX idx_users_email (email)
) ENGINE = INNODB;

DROP TABLE IF EXISTS types;

CREATE TABLE types (
    id int auto_increment primary key,
    name varchar(50) not null,
    user_id int,
    created_at timestamp default current_timestamp(),
    foreign key (user_id) references users (id),
    unique index idx_types_name_user_id (name, user_id)
) ENGINE = INNODB;

INSERT INTO
    types (name, user_id, created_at)
VALUES
    ('Receita', NULL, CURRENT_TIMESTAMP()),
    ('Despesa', NULL, CURRENT_TIMESTAMP());

DROP TABLE IF EXISTS categories;

CREATE TABLE categories (
    id int auto_increment primary key,
    name varchar(50) not null,
    user_id int,
    created_at timestamp default current_timestamp(),
    foreign key (user_id) references users (id),
    unique index idx_categories_name_user_id (name, user_id)
) ENGINE = INNODB;

INSERT INTO
    categories (name, user_id, created_at)
VALUES
    ('Alimentação', NULL, CURRENT_TIMESTAMP()),
    ('Transporte', NULL, CURRENT_TIMESTAMP()),
    ('Moradia', NULL, CURRENT_TIMESTAMP()),
    ('Salário', NULL, CURRENT_TIMESTAMP()),
    ('Lazer', NULL, CURRENT_TIMESTAMP()),
    ('Saúde', NULL, CURRENT_TIMESTAMP());

DROP TABLE IF EXISTS transactions;

CREATE TABLE transactions (
    id int auto_increment primary key,
    uuid varchar(36) not null unique,
    user_id int not null,
    title varchar(50) not null,
    description varchar(100),
    amount bigint not null,
    date datetime not null,
    type_id int not null,
    category_id int not null,
    status enum ('PENDENTE', 'REALIZADA', 'CANCELADA') not null default 'REALIZADA',
    is_recurring boolean not null default false,
    method enum (
        'DINHEIRO',
        'DEBITO',
        'CREDITO',
        'PIX',
        'TRANSFERENCIA'
    ) not null,
    created_at timestamp default current_timestamp(),
    updated_at timestamp default current_timestamp() on update current_timestamp(),
    foreign key (user_id) references users (id),
    foreign key (type_id) references types (id),
    foreign key (category_id) references categories (id),
    index idx_transactions_amount (amount),
    index idx_transactions_date (date),
    index idx_transactions_status (status),
    index idx_transactions_is_recurring (is_recurring),
    index idx_transactions_method (method)
) ENGINE = INNODB;
