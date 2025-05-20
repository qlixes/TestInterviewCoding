create database morning_glory;

use morning_glory;

create table categories(
    id bigint unsigned auto_increment,
    name varchar(128) not null unique,
    primary key(id)
);

create table users(
    id bigint unsigned auto_increment,
    user_type enum('supplier','customer') not null,
    name varchar(128) not null unique,
    primary key(id)
);

create table products(
    id bigint unsigned auto_increment,
    category_id bigint unsigned not null,
    name varchar(128),
    stock int default 0,
    price decimal(10,2),
    primary key(id),
    constraint fk_product_category_id foreign key (category_id) references categories(id)
);

create table transactions(
    id bigint unsigned auto_increment,
    user_id bigint unsigned not null,
    product_id bigint unsigned not null,
    quantity int default 0,
    price decimal(10,2),
    primary key(id),
    constraint fk_transaction_product_id foreign key (product_id) references products(id),
    constraint fk_transaction_user_id foreign key (user_id) references users(id)
);
