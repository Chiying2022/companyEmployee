package main

import (
	"companyEmployee/cmd/handler"
	"errors"
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
	router.Run("localhost:8080")
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
		CODE string `json:"code" binding:"required"`
		NAME string `json:"name"`
	}

	// 消除空格
	err := c.BindJSON(&body)
	companyCode := strings.ReplaceAll(body.CODE, " ", "")
	name := strings.ReplaceAll(body.NAME, " ", "")

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

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		handler.AddCompanyHandler(companyCode, name)
		c.IndentedJSON(http.StatusCreated, body)
	}
}

func addPeople(c *gin.Context) {
	var body struct {
		NAME        string `json:"name" binding:"required"`
		COMPANYCODE string `json:"company_code" binding:"required"`
		AGE         int    `json:"age"`
		GENDER      string `json:"gender"`
	}

	// 消除空格
	err := c.BindJSON(&body)
	companyCode := strings.ReplaceAll(body.COMPANYCODE, " ", "")
	name := strings.ReplaceAll(body.NAME, " ", "")

	// 排除為空字串
	if companyCode == "" || name == "" {
		err := errors.New("request missed name or companyCode")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 檢查CompanyCode 是否存在
	companyCheck := handler.CheckCompanyHandler(body.COMPANYCODE)
	if companyCheck {
		err := errors.New("CompanyCode not found")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		handler.AddPeopleHandler(name, companyCode, body.AGE, body.GENDER)
		c.IndentedJSON(http.StatusCreated, body)
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
		// err := errors.New("CompanyCode is required")
		response := map[string]string{
			"error": "CompanyCode is required",
		}
		// err := fmt.Errorf("error: %s", "CompanyCode is required")
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
		NAME string `json:"name" binding:"required"`
		AGE  int    `json:"age"`
	}

	// 消除空格
	err := c.BindJSON(&body)
	name := strings.ReplaceAll(body.NAME, " ", "")

	// 排除為空字串
	if name == "" {
		err = errors.New("request missed name")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// BindJSON fail
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	// 名字去掉空白後 update + 沒有找到這個人跳錯
	results := handler.UpdatePeopleHandler(body.AGE, name) // UpdatePeopleHandlerTwo -> db.NamedExec
	if results != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "fail, person not found",
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, body)
}
