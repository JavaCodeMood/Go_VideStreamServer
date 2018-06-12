CREATE DATABASE IF NOT EXISTS video_server default charset utf8 COLLATE utf8_general_ci;

create table `users`(
	`id` int not null auto_increment primary key,
	`login_name` varchar(64) not null ,
	`pwd` text not null
);

create table `video_info`(
	`id` varchar(64) not null primary key,
	`author_id` int ,
	`name` text,
	`display_ctime` text,
	`create_time` timestamp not null default current_timestamp
);

create table `comments`(
	`id` varchar(64) not null primary key,
	`video_id` varchar(64) ,
	`author_id` int ,
	`content` text,
	`time` timestamp not null default current_timestamp
);

create table `sessions`(
	`session_id` varchar(64) not null primary key,
	`TTL` tinytext ,
	`login_name` text
);