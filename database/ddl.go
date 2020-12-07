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

const CreateTableTransactions = `create table if not exists transactions(
	id 					integer primary key autoincrement,
	accNumbSender  		integer not null,
	accNumbAddressee  	integer not null,
	translatedSum		integer not null,
	date				text not null,
	time				text not null
);`

const CreateTableSrvcs = `create table if not exists services(
	id 			integer primary key autoincrement,
	NameSrv  	text not null,
	AccNumb  	integer not null,
	SumRemoved  integer,
	SumFilled	integer,
	date		text not null,
	time		text not null
);`

