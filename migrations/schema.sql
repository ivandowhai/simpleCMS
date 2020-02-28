create table posts (
    id int auto_increment primary key,
    title varchar(256) not null,
    content text not null
);

create table users (
    id int auto_increment primary key,
    name varchar(256) not null,
    email varchar(256) not null,
    password varchar(255) not null,
    role tinyint(1) default 1 not null
);
