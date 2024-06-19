package handler

import (
	"companyEmployee/model"

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

// func AddCompanyHandler(ctx context.Context, companyCode string, name string) (bool, error) {
// 	db, err := getDB()
// 	if err != nil {
// 		return false, fmt.Errorf("failed to connect to database: %v", err)
// 	}

// 	query := `
// 		SELECT COUNT(*)
// 		FROM company
// 		WHERE code = ?
// 	`

// 	var count int
// 	err = db.QueryRowContext(ctx, query, companyCode).Scan(&count)
// 	if err != nil {
// 		return false, fmt.Errorf("error querying database: %v", err)
// 	}
// 	if count > 0 {
// 		logger.FromCtx(ctx).Errorf("company code %s exists already", companyCode)
// 		return false, errors.New("公司代碼不可重複")

// 	}

// 	execSQL := `
// 		INSERT INTO company (
// 			code,
// 			name
// 		) VALUES (
// 			?,
// 			?
// 		)
// 	`
// 	result, err := db.Exec(execSQL, companyCode, name)
// 	if err != nil {
// 		err = fmt.Errorf("fail to insert, err(%v)", result)
// 		return false, err
// 	}

// 	affected, _ := result.RowsAffected()
// 	if affected == 0 {
// 		return false, errors.Wrap(model.ErrAddCompanyFail, "新增公司失敗")
// 	} else if affected > 1 {
// 		return false, errors.Errorf("更新單狀態錯誤, 條數 %d != 1", affected)
// 	}

// 	logger.FromCtx(ctx).Info("company added successfully")
// 	return true, nil
// }

// func AddPeopleHandler(name string, companyCode string, age int, gender string) (b bool, err error) {
// 	db, _ := getDB()
// 	execSQL := `
// 		INSERT INTO employee (
// 			name,
// 			company_code,
// 			age,
// 			gender
// 		) VALUES (
// 			?,
// 			?,
// 			?,
// 			?
// 		)
// 	`
// 	_, err1 := db.Exec(execSQL, name, companyCode, age, gender)
// 	if err1 != nil {
// 		return false, err1
// 	}

// 	return true, nil
// }

// func GetPeopleByCompanyHandler(companyCode string, page int, size int) ([]interface{}, error) {
// 	db, _ := getDB()

// 	// page , size 任一沒填等同沒有分頁功能 即印出全部
// 	if page == 0 || size == 0 {
// 		execSQL := `
// 			SELECT
// 				employee.id,
// 				employee.name
// 			FROM employee
// 			WHERE company_code = ?
// 			ORDER BY id ASC
// 	`
// 		results, err := db.Query(execSQL, companyCode)
// 		if err != nil {
// 			err = fmt.Errorf("query people fail, err(%v)", err)
// 			return nil, err
// 		}

// 		defer results.Close()

// 		var peopleList []interface{}
// 		for results.Next() {
// 			var person Peoplelist
// 			err := results.Scan(&person.Id, &person.Name)
// 			if err != nil {
// 				return nil, errors.WithMessagef(err, "get people list fail %s", companyCode)
// 			}
// 			peopleList = append(peopleList, &person)
// 		}

// 		return peopleList, err
// 	}

// 	// 當有填入 page + size 開始計算 offset
// 	var offset = 0
// 	if page != 0 && size != 0 {
// 		offset = (page - 1) * size
// 	}

// 	execSQL := `
// 		SELECT
// 			employee.id,
// 			employee.name
// 		FROM employee
// 		WHERE employee.company_code = ?
// 		ORDER BY id ASC
// 		LIMIT ?  OFFSET ?
// 	`
// 	results, err := db.Query(execSQL, companyCode, size, offset)
// 	if err != nil {
// 		err = fmt.Errorf("fail to query people with page and size, err(%v)", err)
// 		return nil, err
// 	}

// 	defer results.Close()

// 	var peopleList []interface{}
// 	for results.Next() {
// 		var person Peoplelist
// 		err := results.Scan(&person.Id, &person.Name)
// 		if err != nil {
// 			err = fmt.Errorf("fail to get people with page and size, err(%v)", err)
// 			return nil, err
// 		}
// 		peopleList = append(peopleList, &person)
// 	}

// 	countQuery := `
// 		SELECT COUNT(*)
// 		FROM employee
// 		WHERE company_code = ?
// 	`
// 	var count int
// 	err = db.QueryRow(countQuery, companyCode).Scan(&count)
// 	if err != nil {
// 		err = fmt.Errorf("fail to count data, err(%v)", err)
// 		return nil, err
// 	}

// 	// 判斷筆數有無錯誤
// 	var totalPage, remainder int
// 	if size > count {
// 		err = errors.New("invalid size")
// 		return nil, err
// 	}

// 	totalPage = count / size
// 	remainder = count % size
// 	if remainder > 0 {
// 		totalPage++
// 	}

// 	// 判斷頁數 有無錯誤
// 	if page > totalPage {
// 		err = errors.New("invalid page")
// 		return nil, err
// 	}

// 	return peopleList, nil

// }

// func UpdatePeopleHandler(c *gin.Context, age int, name string) (err error) {
// 	db, _ := getDB()

// 	execSQL := `
// 		SELECT count(*)
// 		FROM employee
// 		WHERE 1 = 1
// 		AND name = ?
// 	`
// 	var count int

// 	err = db.QueryRowContext(c, execSQL, name).Scan(&count)
// 	if err != nil {
// 		return errors.New("db error")
// 	}

// 	if count == 0 {
// 		return errors.New("user not found")
// 	}

// 	execSQL = `
// 		UPDATE employee
// 		SET age = ?
// 		WHERE name = ?
// 	`

// 	_, err = db.ExecContext(c, execSQL, age, name)
// 	if err != nil {
// 		return errors.New("db update fail")
// 	}

// 	// results, err := db.ExecContext(c, execSQL, age, name)
// 	// if err != nil {
// 	// 	err = fmt.Errorf("db error: %s", err)
// 	// 	return err
// 	// }

// 	// rowsAffected, _ := results.RowsAffected()
// 	// if rowsAffected == 0 {
// 	// 	err = errors.New("person not found")
// 	// 	return err
// 	// }

// 	// err = r.DB(ctx).GetContext(ctx, &depositAddress, query, networkCode, user.ID)
// 	// if errors.Is(err, sql.ErrNoRows) {
// 	// 	return "", nil
// 	// } else if err != nil {
// 	// 	return "", err
// 	// }

// 	return
// }

// func CheckCompanyHandler(companyCode string) (c int, err error) {
// 	db, _ := getDB()
// 	var count int
// 	execSQL := `
// 		SELECT count(*)
// 		FROM company
// 		WHERE code = ?
// 	`
// 	results, err := db.Query(execSQL, companyCode)
// 	if err != nil {
// 		err = fmt.Errorf("company query fail, err(%v)", err)
// 		return count, err
// 	}

// 	defer results.Close()

// 	if results.Next() {
// 		err = results.Scan(&count)
// 		if err != nil {
// 			err = fmt.Errorf("fail to check company, err(%v)", err)
// 			return count, err
// 		}
// 	}

// 	return count, err
// }

// func CheckPeopleHandler(companyCode string) (c int, err error) {
// 	db, _ := getDB()
// 	var count int
// 	execSQL := `
// 		SELECT count(*)
// 		FROM employee
// 		WHERE company_code = ?
// 	`
// 	results, err := db.Query(execSQL, companyCode)
// 	if err != nil {
// 		err = fmt.Errorf("people query fail, err(%v)", err)
// 		return count, err
// 	}

// 	defer results.Close()

// 	if results.Next() {
// 		err = results.Scan(&count)
// 		if err != nil {
// 			err = fmt.Errorf("fail to check people, err(%v)", err)
// 			return count, err
// 		}
// 	}

// 	return count, err
// }

func TestHandler(name string) (company *model.Company, err error) {
	dbx, _ := getDB()
	query := `
		SELECT
			id, 
			name, 
			is_active
		FROM company
		WHERE name = ?
	`
	company = &model.Company{}
	err = dbx.db.Get(company, query, name)
	if err != nil {
		return nil, err
	}
	return

}

func BatchCompanyHandler(batchCompany []model.Company) (boxes []model.Company, err error) {
	db, _ := getDB()
	query := `
		SELECT
			id,
			code,
			name,
			is_active
		FROM company
		WHERE code = ?
	`
	company := model.Company{}
	boxes = []model.Company{}

	for _, companyInfo := range batchCompany {
		err = db.db.Get(&company, query, companyInfo.Code)
		if err != nil {
			return nil, err
		}
		boxes = append(boxes, company)
	}
	return
}

// SELECT id, title FROM books ORDER BY id ASC LIMIT 10 OFFSET 10;
// offset := (page - 1) * pageSize
// totalPage := (totalCount + pageSize - 1) / pageSize
