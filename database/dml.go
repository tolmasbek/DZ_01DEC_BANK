package database

const AddNewUser = `insert into users (firstName, lastName, age, gender, login, password, role, remove) 
					values(($1),($2),($3),($4),($5),($6),($7),($8));`

const AddNewCurrency = `insert into currencies (nameCurr) 	
						values ($1);`

const AddNewAcc = `insert into accounts (userId, numberAccount, amount, currency, pin, remove) 
					values (($1),($2),($3),($4),($5),($6));`

const AddNewAtm = `insert into atms (address) 
					values($1);`

const AddNewTransaction = `insert into transactions(accNumbSender, accNumbAddressee, translatedSum, date, time)
							values(($1),($2),($3),($4),($5));`

const AddNewServices = `insert into services(NameSrv, AccNumb, SumRemoved, SumFilled, date, time)
							values(($1),($2),($3),($4),($5),($6));`

const ShowHistoryTran = `Select *from transactions where accNumbSender = ($1);`

const ShowDataAcc = `Select id, numberAccount, amount, pin From accounts Where pin=($1)`

const ChekDataAcc = `Select accounts.numberAccount, accounts.pin From accounts Where numberAccount = ($1) and pin = ($2)`

const GetLoginPass = `Select users.login, users.password, users.role From users Where (login=($1) and password=($2))and role=($3)`

const TransferFromAcc = `update accounts set amount = amount - ($1) where numberAccount = ($2)`

const TransferToAcc = `update accounts set amount = amount + ($1) where numberAccount = ($2)`

const ChekNumAcc = `Select numberAccount From accounts Where numberAccount = ($1)`

const ChekAmountAcc = `Select Amount From accounts Where numberAccount=($1)`
