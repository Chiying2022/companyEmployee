package models

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
)

const dbuser = "root"
const dbname = "central"

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
	insert, err := db.Query(execSQL, companyCode, name)

	if err != nil {
		fmt.Println("Err", err.Error())
	}
	defer insert.Close()
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
	insert, err := db.Query(execSQL, NAME, COMPANYCODE, AGE, GENDER)

	if err != nil {
		fmt.Println("Err", err.Error())
	}
	defer insert.Close()
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

func UpdatePeopleHandler(age int, name string) {
	db, err := sql.Open("mysql", dbuser+":safesync"+"@tcp(10.1.103.111:3306)/"+dbname)
	if err != nil {
		fmt.Println("Err", err.Error())
	}

	defer db.Close()

	execSQL := `
		UPDATE employee
		SET
			age = ?
		WHERE name = ?
	`
	_, err = db.Exec(execSQL, age, name)

	if err != nil {
		fmt.Println("Err", err.Error())
		return
	}

	fmt.Println("Update successful")

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

// 需抓取 params -> 已抓到
// 如果搜尋結果是多筆? -> ok
// 如何只取 response 只需要的資訊 -> 解決到可能不聰明
// 未加分頁 ->
// SELECT id, title FROM books ORDER BY id ASC LIMIT 10 OFFSET 10;
// offset := (page - 1) * pageSize
// totalPage := (totalCount + pageSize - 1) / pageSize
