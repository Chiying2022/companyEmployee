package main

import (
	"companyEmployee/cmd/handler"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/api/company", addCompany)
	router.POST("/api/people", addPeople)
	router.GET("/api/list", checkList)
	router.PUT("/api/update", updatePeople)

	// ------------
	// router.GET("/api/listcompany", listCompany) // using for test
	// router.GET("/api/listpeople", listPeople)   // using for test
	err := router.Run()
	if err != nil {
		log.Fatal("Failed to start the server: ", err)
	}
}

func listPeople(c *gin.Context) {
	name := c.Query("name")
	results := handler.PeopleHandler(name)
	if results != 123 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, results)
	}
}

func listCompany(c *gin.Context) {
	companyList := handler.CompanyHandler()
	if companyList == nil || len(companyList) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, companyList)
	}
}

func addCompany(c *gin.Context) {
	var body struct {
		Code string `json:"code" binding:"required"`
		Name string `json:"name"`
	}

	err := c.BindJSON(&body)
	if err != nil {
		err := errors.New("failed to bind JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 消除空格
	companyCode := strings.ReplaceAll(body.Code, " ", "")
	name := strings.ReplaceAll(body.Name, " ", "")

	// 排除為空字串
	if companyCode == "" || name == "" {
		err := errors.New("request missed companyCode or name")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 檢查CompanyCode 是否重複
	companyCheck := handler.CheckCompanyHandler(companyCode)
	if companyCheck {
		err := errors.New("CompanyCode existed already")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := handler.AddCompanyHandler(companyCode, name)
	if !result {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to add company"})
		return
	}
	c.IndentedJSON(http.StatusCreated, body)
}

func addPeople(c *gin.Context) {
	var body struct {
		Name        string `json:"name" binding:"required"`
		CompanyCode string `json:"company_code" binding:"required"`
		Age         int    `json:"age"`
		Gender      string `json:"gender"`
	}

	err := c.BindJSON(&body)
	if err != nil {
		err := errors.New("failed to bind JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 消除空格
	companyCode := strings.ReplaceAll(body.CompanyCode, " ", "")
	name := strings.ReplaceAll(body.Name, " ", "")

	// 排除為空字串
	if companyCode == "" || name == "" {
		err := errors.New("request missed name or companyCode")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 檢查CompanyCode 是否存在
	companyCheck := handler.CheckCompanyHandler(body.CompanyCode)
	if companyCheck {
		err := errors.New("CompanyCode not found")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := handler.AddPeopleHandler(name, companyCode, body.Age, body.Gender)
	if !result {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to add people"})
		return
	}

	c.IndentedJSON(http.StatusCreated, body)

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
		response := map[string]string{
			"error": "CompanyCode is required",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// 檢查CompanyCode 是否存在
	companyCheck := handler.CheckCompanyHandler(CompanyCode)

	if !companyCheck {
		err := errors.New("CompanyCode not found")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 開始撈出清單 (page or pagesize 指填入任一或是都不填 一律回應全部清單)
	peopleList, count := handler.GetPeopleByCompanyHandler(CompanyCode, Page, Size)

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
		Name string `json:"name" binding:"required"`
		Age  int    `json:"age"`
	}

	err := c.BindJSON(&body)
	if err != nil {
		err := errors.New("failed to bind JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 消除空格
	name := strings.ReplaceAll(body.Name, " ", "")

	// 排除為空字串
	if name == "" {
		err = errors.New("request missed name")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 名字去掉空白後 update + 沒有找到這個人跳錯
	results := handler.UpdatePeopleHandler(body.Age, name)
	if results != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "fail, person not found",
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, body)
}
