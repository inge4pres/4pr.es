create database short;
create user shortener identified by 'passwd';
grant all on short.* to shortener@'localhost' identified by 'passwd';

create table urls ( 
	id serial primary key, 
	url varchar(500) not null, 
	short varchar(30) not null,
	created timestamp not null default now()
  );