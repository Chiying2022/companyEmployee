package handler

import (
	"companyEmployee/cmd/response"
	"companyEmployee/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type A struct{}

const (
	a = "今天天氣很好"
	b = len(a)
)

// 批次公司 is_active 確認 把沒有 ａctive 的印出
func ActiveBatchCheck(ctx *gin.Context) {
	var params struct {
		Company []model.Company `json:"companies" binding:"required"`
	}

	var num map[string]int
	num = make(map[string]int)
	num["one"] = 1  //賦值
	num["ten"] = 10 //賦值
	num["three"] = 3

	err := ctx.BindJSON(&params)
	if err != nil {
		response.FailResponse(ctx, http.StatusBadRequest, err, "failed to bind JSON")
		return
	}

	company, err := BatchCompanyHandler(params.Company)
	if err != nil {
		response.FailResponse(ctx, http.StatusNotFound, err, "company not found")
		return
	}

	results := batchCheckActive(company)

	response.SuccResponse(ctx, http.StatusOK, results, num)
	return
}

func batchCheckActive(company []model.Company) (name []map[string]string) {
	for _, eachCompany := range company {
		if eachCompany.IsActive != 1 {
			name = append(name, map[string]string{
				"CODE": eachCompany.Code,
				"NAME": eachCompany.Name,
			})
		}
	}
	return
}

// 使用 map 判斷此公司有沒有 active, 用名字先撈出 然後判斷是否是 active
func ActiveCheck(c *gin.Context) {
	var params struct {
		Name string `form:"name" binding:"required"`
	}

	var test = a
	fmt.Println(b)
	fmt.Println(test)

	// m := []int{
	// 	1:    1,
	// 	0x02: 2,
	// 	'c':  3,
	// }

	// m := map[int]int{
	// 	1:    1,
	// 	0x02: 2,
	// 	'c':  3,
	// }
	// mm := make(map[int]int)

	// fmt.Println(len(m))
	// fmt.Println(len(m))

	err := c.BindQuery(&params)
	if err != nil {
		response.FailResponse(c, http.StatusBadRequest, err, "failed to bind JSON")
		return
	}

	company, err := CheckCompanyHandler(params.Name)
	if err != nil {
		response.FailResponse(c, http.StatusNotFound, err, "company not found")
		return
	}

	if !checkActive(company) {
		response.FailResponse(c, http.StatusNotFound, err, "company not active")
		return
	}

	response.SuccResponse(c, http.StatusOK, company, nil)
	return
}

func checkActive(company *model.Company) bool {
	return company.IsActive == 1
}
