package database

const CreateTableUsers = `Create table if not exists users (
	id 			integer 	primary key	autoincrement,
	firstName 	text 		not null,
	lastName 	text 		not null,
	age 		integer 	not null,
	gender 		text 		not null,
	login 		text 		not null 	unique,
	password 	text 		not null,
	role		text		not null,
	remove 		boolean 	not null 	default false
);`

const CreateTableAccounts = `Create table if not exists accounts(
	id 				integer primary key autoincrement,
	userId 			integer references users(id) not null,
	numberAccount 	integer not null,
	amount 			integer not null,
	currency 		integer references currency(id),
	pin				integer not null unique,
	remove 			boolean not null default false 
);`

const CreateTableATMs = `Create table if not exists atms (
	id 			integer 	primary key	autoincrement,
	address 	text 		not null,
	status		boolean 	not null default true
);`

const CreateTableCurrencies = `create table if not exists currencies(
	id 			integer primary key autoincrement,
	nameCurr 	text 
);`
