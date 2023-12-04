package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
)

const dbuser = "root"
const dbname = "central"

func BankHandler() []Bank {
	db, err := sql.Open("mysql", dbuser+":safesync"+"@tcp(10.1.103.111:3306)/"+dbname)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()

	execSQL := `
		SELECT
			ba.account,
			ba.client_type
		FROM bank_analytics ba
		WHERE ba.id = 37203
		`

	results, err := db.Query(execSQL)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	bankTry := []Bank{}
	for results.Next() {
		var bank Bank
		err = results.Scan(&bank.ACCOUNT, &bank.CLIENT)

		if err != nil {
			panic(err.Error())
		}
		bankTry = append(bankTry, bank)
	}
	return bankTry
}
