create table users (
    id int auto_increment primary key,
    name varchar(256) not null,
    email varchar(256) not null,
    password varchar(255) not null,
    role tinyint(1) default 1 not null
);

create table posts (
    id int auto_increment primary key,
    title varchar(256) not null,
    content text not null,
    user_id int not null
);

alter table posts
	add constraint posts_users_id_fk
		foreign key (user_id) references users (id)
			on update cascade on delete cascade;
