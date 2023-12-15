package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const dbuser = "root"
const dbname = "central"

var db *sqlx.DB

const (
	username = "root"
	password = "safesync"
	host     = "10.1.103.111"
	port     = 3306
	dbName   = "central"
)

func PeopleHandler(name string) int {
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, dbName))
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database")

	var people []People
	err = db.Select(&people, "SELECT name, age FROM employee WHERE name = ?", name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(people) // 輸出搜尋結果
	return 123

}

func CompanyHandler() []Company {
	db, err := sql.Open("mysql", dbuser+":safesync"+"@tcp(10.1.103.111:3306)/"+dbname)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()

	execSQL := `
		SELECT
			company.id,
			company.code
		FROM company
		WHERE company.code = "AF"
		`

	results, err := db.Query(execSQL)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	companyTry := []Company{}
	for results.Next() {
		var com Company
		err = results.Scan(&com.CODE)

		if err != nil {
			panic(err.Error())
		}
		companyTry = append(companyTry, com)
	}
	return companyTry
}

func AddCompanyHandler(companyCode string, name string) {
	db, err := sql.Open("mysql", dbuser+":safesync"+"@tcp(10.1.103.111:3306)/"+dbname)
	if err != nil {
		fmt.Println("Err", err.Error())
	}

	defer db.Close()

	execSQL := `
		INSERT INTO company (
			code,
			name
		) VALUES (
			?,
			?
		)
	`
	_, err = db.Exec(execSQL, companyCode, name)

	if err != nil {
		fmt.Println("Err", err.Error())
	}
}

func AddPeopleHandler(NAME string, COMPANYCODE string, AGE int, GENDER string) {
	db, err := sql.Open("mysql", dbuser+":safesync"+"@tcp(10.1.103.111:3306)/"+dbname)
	if err != nil {
		fmt.Println("Err", err.Error())
	}

	defer db.Close()

	execSQL := `
		INSERT INTO employee (
			name,
			company_code,
			age, 
			gender
		) VALUES (
			?,
			?,
			?,
			?
		)
	`
	_, err = db.Exec(execSQL, NAME, COMPANYCODE, AGE, GENDER)

	if err != nil {
		fmt.Println("Err", err.Error())
	}

}

func GetPeopleByCompanyHandler(companyCode string, page string, size string) ([]interface{}, int) {

	db, err := sql.Open("mysql", dbuser+":safesync"+"@tcp(10.1.103.111:3306)/"+dbname)
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil, 0
	}
	defer db.Close()

	// 將傳入page參數轉成數字
	var intPage, intSize = 0, 0
	if page != "" {
		intPage, err = strconv.Atoi(page)
		if err != nil {
			fmt.Println("Error converting page to integer:", err)
			return nil, 0
		}
	}

	// 將傳入size參數轉成數字
	if size != "" {
		intSize, err = strconv.Atoi(size)
		if err != nil {
			fmt.Println("Error converting page to integer:", err)
			return nil, 0
		}
	}

	// intPage , intSize 任一沒填等同沒有分頁功能 即印出全部
	if companyCode != "" && (intPage == 0 || intSize == 0) {
		execSQL := `
			SELECT
				employee.id,
				employee.name
			FROM employee
			WHERE company_code = ?
			ORDER BY id ASC
	`
		results, err := db.Query(execSQL, companyCode)
		if err != nil {
			return nil, 0
		}

		defer results.Close()

		var peopleList []interface{}
		for results.Next() {
			var person Peoplelist
			err := results.Scan(&person.ID, &person.NAME)
			if err != nil {
				fmt.Println("Err", err.Error())
				return nil, 0
			}
			peopleList = append(peopleList, &person)
		}

		return peopleList, len(peopleList)
	}

	// 當有填入 page + size 開始計算 offset
	var offset = 0
	if intPage != 0 && intSize != 0 {
		offset = (intPage - 1) * intSize // offset 設定從0開始
		// 預防填入數字為 -1
		if offset < 0 {
			offset = 0
		}
	}

	execSQL := `
		SELECT
			employee.id,
			employee.name
		FROM employee
		WHERE employee.company_code = ?
		ORDER BY id ASC 
		LIMIT ?  OFFSET ?
	`

	results, err := db.Query(execSQL, companyCode, size, offset)

	defer results.Close()

	var peopleList []interface{}
	for results.Next() {
		var person Peoplelist
		err := results.Scan(&person.ID, &person.NAME)
		if err != nil {
			fmt.Println("Err", err.Error())
			return nil, 0
		}
		peopleList = append(peopleList, &person)
	}

	countQuery := `
		SELECT COUNT(*)
		FROM employee
		WHERE company_code = ?
	`
	var count int
	err = db.QueryRow(countQuery, companyCode).Scan(&count)
	if err != nil {
		fmt.Println("Error counting rows:", err)
		return nil, 0
	}

	// 計算頁數/筆數
	var totalPage, remainder int
	totalPage = count / intSize
	remainder = totalPage % intSize
	if remainder > 0 {
		totalPage++
	}

	// 判斷頁數/筆數 有無錯誤
	if intPage > totalPage {
		fmt.Println("Error: page or pagesize invalide")
		return nil, 0
	}

	return peopleList, len(peopleList)

}

func UpdatePeopleHandler(age int, name string) (err error) {
	db, err := sql.Open("mysql", dbuser+":safesync"+"@tcp(10.1.103.111:3306)/"+dbname)
	if err != nil {
		return fmt.Errorf("error opening database: %s", err)
	}

	defer db.Close()

	execSQL := `
		UPDATE employee
		SET
			age = ?
		WHERE name = ?
	`
	results, err := db.Exec(execSQL, age, name)

	if err != nil {
		return fmt.Errorf("error updating person: %s", err)
	}

	rowsAffected, _ := results.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("person not found")
	}

	fmt.Println("Update successful")
	return

}

func countResults(companyCode string) []interface{} {
	db, err := sql.Open("mysql", dbuser+":safesync"+"@tcp(10.1.103.111:3306)/"+dbname)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	defer db.Close()

	execSQL := `
		SELECT
			employee.id,
			employee.name
		FROM employee
		WHERE employee.company_code = ?
		ORDER BY id ASC
	`
	results, err := db.Query(execSQL, companyCode)

	defer results.Close()

	var peopleList []interface{}
	for results.Next() {
		var person Peoplelist
		err := results.Scan(&person.ID, &person.NAME)
		if err != nil {
			fmt.Println("Err", err.Error())
			return nil
		}
		peopleList = append(peopleList, &person)
	}

	return peopleList
}

func CheckCompanyHandler(companyCode string) bool {
	db, err := sql.Open("mysql", dbuser+":safesync"+"@tcp(10.1.103.111:3306)/"+dbname)

	if err != nil {
		fmt.Println("Err", err.Error())
		return false
	}
	defer db.Close()

	execSQL := `
		SELECT
			company.id,
			company.name
		FROM company
		WHERE code = ?
	`
	results, err := db.Query(execSQL, companyCode)
	if err != nil {
		fmt.Println("Err", err.Error())
		return false
	}

	defer results.Close()

	if !results.Next() {
		return false
	}
	return true
}

func UpdatePeopleHandlerTwo(age int, name string) (err error) {
	db, err := sqlx.Open("mysql", dbuser+":safesync"+"@tcp(10.1.103.111:3306)/"+dbname)
	if err != nil {
		return fmt.Errorf("error opening database: %s", err)
	}

	defer db.Close()

	execSQL := `
		UPDATE employee
		SET
			age = :age
		WHERE name = :name
	`

	results, err := db.NamedExec(execSQL, map[string]interface{}{
		"name": name,
		"age":  age,
	})

	if err != nil {
		return fmt.Errorf("error updating person: %s", err)
	}

	rowsAffected, _ := results.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("person not found")
	}

	fmt.Println("Update successful")
	return

}

// SELECT id, title FROM books ORDER BY id ASC LIMIT 10 OFFSET 10;
// offset := (page - 1) * pageSize
// totalPage := (totalCount + pageSize - 1) / pageSize
