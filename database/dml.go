package database

const AddNewUser = `insert into users (firstName, lastName, age, gender, login, password, role, remove) values(($1),($2),($3),($4),($5),($6),($7),($8));`

const AddNewCurrency = `insert into currencies (nameCurr) values ($1);`

const AddNewAcc = `insert into accounts (userId, numberAccount, amount, currency, remove) values (($1),($2),($3),($4),($5));`

const AddNewAtm = `insert into atms (address) values($1);`
