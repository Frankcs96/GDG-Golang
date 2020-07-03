package main

import "github.com/fcsuarez96/http/user"

type Database struct {
	Accounts map[string]string
}

func NewDatabase() *Database {

	accounts := make(map[string]string)
	accounts["user1"] = "1234"
	accounts["user2"] = "1234"
	accounts["user3"] = "1234"

	return &Database{
		Accounts: accounts,
	}
}

func (db Database) CheckUser(user user.User) bool {

	if val, ok := db.Accounts[user.Username]; ok {

		if val == user.Password {
			return true
		}
	}

	return false
}
