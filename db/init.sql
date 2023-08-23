CREATE DATABASE notes;    

\c notes;     

create table users (
	user_id int primary key generated always as identity,
	username text unique,
	user_password text
);

create table notes (
	note_id int primary key generated always as identity,
	user_id int references users(user_id),
	body text
);

