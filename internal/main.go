package main

import (
	"errors"
	"net/http"
	"sqltest/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/api/company", addCompany)
	router.POST("/api/people", addPeople) // pagenation (offset/limite) + 撈出同公司的員工list -----> 如果 offset/limite 沒填呢？
	router.GET("/api/list", checkList)
	router.PUT("/api/update", updatePeople)
	router.Run("localhost:8080")

	/// try
	// router.GET("/api/company", listCompany)
	// router.GET("/api/listlist", listlist)
}

func listCompany(c *gin.Context) {
	companyList := models.CompanyHandler()
	if companyList == nil || len(companyList) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, companyList)
	}
}

func addCompany(c *gin.Context) {
	// var com models.Company
	var param struct {
		CODE string `json:"code"`
		NAME string `json:"name"`
	}

	if err := c.BindJSON(&param); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.AddCompanyHandler(param)
		c.IndentedJSON(http.StatusCreated, param)
	}
}

func addPeople(c *gin.Context) {
	var p models.People

	if err := c.BindJSON(&p); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.AddPeopleHandler(p)
		c.IndentedJSON(http.StatusCreated, p)
	}
}

func checkList(c *gin.Context) {
	CompanyCode := strings.Trim(c.Query("company_code"), " ") // 控除前後空白
	Page := c.Query("page")
	Size := c.Query("size")

	// 檢查CompanyCode 是否為字串
	_, err := strconv.Atoi(CompanyCode)
	if err == nil {
		err := errors.New("CompanyCode must be string")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 檢查是否有填CompanyCode 參數
	if len(CompanyCode) == 0 {
		err := errors.New("CompanyCode is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 檢查CompanyCode 是否存在
	companyCheck := models.CheckCompanyHandler(CompanyCode)

	if !companyCheck {
		err := errors.New("CompanyCode not found")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 開始撈出清單 (page or pagesize 指填入任一或是都不填 一律回應全部清單)
	peopleList, count := models.GetPeopleByCompanyHandler(CompanyCode, Page, Size)

	// 當填入的參數都正確時 卻沒找到相對應資料
	if peopleList != nil && count == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "people not found",
		})
		return
	}
	// 當填入 page or pagesize 參數不正確時
	if peopleList == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "page or pagesize invalid",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"list": peopleList,
	})

}

func updatePeople(c *gin.Context) {
	var body struct {
		NAME string `json:"name"`
		AGE  int    `json:"age"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.UpdatePeopleHandler(body.AGE, body.NAME)
		c.IndentedJSON(http.StatusCreated, body)
	}
}
