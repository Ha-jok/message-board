```mysql
create database messageboard;

use messageboard;

create table users(
username varchar(10) not null primary key,
password varchar(20) not null,
mibao varchar(30),
answer varchar(20)
)ENGINE=InnoDB CHARSET=utf8mb4;

insert into users(username,password,mibao,answer) values('ddz','yyds','ddz','nhnz');

insert into users(username,password,mibao,answer) values('zzz','cai','zzz','c');
```

