CREATE DATABASE IF NOT EXISTS test;
use test;

CREATE TABLE IF NOT EXISTS activity(
    id int not null auto_increment primary key,
    activity_group_id int,
    title varchar(255),
    is_active tinyint,
    "priority" varchar(255),
    created_at datetime,
    updated_at datetime
);

CREATE TABLE IF NOT EXISTS todos(
    id int not null auto_increment primary key,
    title varchar(255),
    email varchar(255),
    created_at datetime,
    updated_at datetime
);