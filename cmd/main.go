package main

import (
	"companyEmployee/cmd/delivery"
	"fmt"
	"log"
	"runtime"

	"github.com/gin-gonic/gin"
)

const (
	projectName = "companyEmployee"
)

type person struct {
	firstName string
}

// 建立 person 的 function receiver
func (p person) updateName(newFirstName string) {
	fmt.Printf("Before update: %+v\n", p)
	p.firstName = newFirstName
	fmt.Printf("After update: %+v\n", p)
}

func (p person) print() {
	fmt.Printf("Current person is: %+v\n", p)
}

type Person struct {
	Name string
	Age  int
}

const alphabetStr string = "abcdefghijklmnopqrstuvwxyz"

func add1(a *int) int {
	*a = *a + 1 // 我們改變了 a 的值
	return *a   //回傳一個新值
}

type testInt func(int) bool // 宣告了一個函式型別

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

func main() {
	router := gin.Default()
	delivery.InitRouter(router)
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start the server: ", err)
	}
}

// func InitRouter(engine *gin.Engine) {
// 	engine.Use(
// 		cors.New(cors.Config{
// 			AllowOrigins:     []string{"*"},
// 			AllowMethods:     []string{"GET", "POST", "PUT", "OPTIONS"},
// 			AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With", "X-Forwarded-For"},
// 			AllowCredentials: true,
// 			MaxAge:           12 * time.Hour,
// 		}),
// 	)
// }

// func listPeople(c *gin.Context) {
// 	name := c.Query("name")
// 	results := handler.PeopleHandler(name)
// 	if results != 123 {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.IndentedJSON(http.StatusOK, results)
// 	}
// }

// func listCompany(c *gin.Context) {
// 	companyList := handler.CompanyHandler()
// 	if companyList == nil || len(companyList) == 0 {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.IndentedJSON(http.StatusOK, companyList)
// 	}
// }

//	func loggEndpointURL(c *gin.Context) {
//		c.Next()
//		log.Printf("EndPoint URL is %v", c.Request.URL)
//	}
//

type AddCompanyRequest struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Pet         *MyPet `json:"pet"`
	RequestUUID string `json:"uuid"`
}

type MyPet struct {
	Name string
	Cate string
	Age  int
}

func getInfo(pet *MyPet) (*MyPet, error) {
	return pet, nil
}

// func wrapper(c *gin.Context) {
// 	var body AddCompanyRequest
// 	if err := c.ShouldBindJSON(&body); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	addCompany(c, body.Pet)
// }

// func addCompany(c *gin.Context, pet *MyPet) (*MyPet, error) {
// 	result, err := getInfo(pet)
// 	if err != nil {
// 		response.FailResponse(c, http.StatusBadRequest, err, "company exist already")
// 		return nil, err
// 	}

// 	check, err := handler.AddCompanyHandler(c, result.Cate, result.Name)
// 	if err != nil {
// 		response.FailResponse(c, http.StatusBadRequest, err, "fail to add")
// 		return nil, nil
// 	} else if !check {
// 		response.FailResponse(c, http.StatusBadRequest, err, "company exist already")
// 		return nil, nil
// 	}
// 	response.SuccResponse(c, http.StatusCreated, "add done", result)
// 	return result, nil
// }

// func addPeople(c *gin.Context) {
// 	var body struct {
// 		Name        string `json:"name" binding:"required,lt=100"`
// 		CompanyCode string `json:"company_code" binding:"required,lt=100"`
// 		Age         int    `json:"age" binding:"gt=16"`
// 		Gender      string `json:"gender" binding:"oneof=M F"`
// 	}

// 	err := c.BindJSON(&body)
// 	if err != nil {
// 		response.FailResponse(c, http.StatusBadRequest, err, "BindJSON fail")
// 		return
// 	}

// 	// /// 也可換成以下 ------------------------------------
// 	// data, err := io.ReadAll(c.Request.Body)
// 	// if err != nil {
// 	// 	response.FailResponse(c, http.StatusBadRequest, err, "BindJSON fail")
// 	// 	return
// 	// }

// 	// defer c.Request.Body.Close()

// 	// // var result map[string]string
// 	// err = json.Unmarshal(data, &body)
// 	// if err != nil {
// 	// 	response.FailResponse(c, http.StatusBadRequest, err, "BindJSON fail")
// 	// 	return
// 	// }
// 	/// -------------------------------------------------------------------------------

// 	// check company
// 	company, err := handler.CheckCompanyHandler(body.CompanyCode)
// 	if err != nil {
// 		response.FailResponse(c, http.StatusBadRequest, err, "fail to check company")
// 		return
// 	}

// 	// 檢查Company 是否存在
// 	if company == 0 {
// 		response.FailResponse(c, http.StatusNotFound, err, "company not foud")
// 		return
// 	}

// 	_, err = handler.AddPeopleHandler(body.Name, body.CompanyCode, body.Age, body.Gender)
// 	if err != nil {
// 		response.FailResponse(c, http.StatusBadRequest, err, "fail to add")
// 		return
// 	}

// 	response.SuccResponse(c, http.StatusCreated, "add done", body)
// 	return
// }

// func checkList(c *gin.Context) {
// 	var param struct {
// 		CompanyCode string `form:"company_code" binding:"required,lt=100"`
// 		Page        int    `form:"page" binding:"gte=0"`
// 		Size        int    `form:"size" binding:"gte=0"`
// 	}

// 	err := c.BindQuery(&param)
// 	if err != nil {
// 		response.FailResponse(c, http.StatusBadRequest, err, "BindJSON fail")
// 		return
// 	}

// 	companyCheck, err := handler.CheckCompanyHandler(param.CompanyCode)
// 	if err != nil {
// 		response.FailResponse(c, http.StatusBadRequest, err, "fail to check company")
// 		return
// 	}

// 	// 檢查Company 是否存在
// 	if companyCheck == 0 {
// 		response.FailResponse(c, http.StatusNotFound, err, "CompanyCode not found")
// 		return
// 	}

// 	peopleCheck, err := handler.CheckPeopleHandler(param.CompanyCode)
// 	if err != nil {
// 		response.FailResponse(c, http.StatusBadRequest, err, "fail to check people")
// 		return
// 	}

// 	// 檢查person 是否存在
// 	if peopleCheck == 0 {
// 		response.FailResponse(c, http.StatusNotFound, err, "Person not found")
// 		return
// 	}

// 	// 開始撈出清單 (page or size 只填入任一或都不填 一律回應全部)
// 	peopleList, err := handler.GetPeopleByCompanyHandler(param.CompanyCode, param.Page, param.Size)
// 	if err != nil {
// 		response.FailResponse(c, http.StatusBadRequest, err, "fail to get people list")
// 		return
// 	}

// 	response.SuccResponse(c, http.StatusOK, "ok", peopleList)
// 	return

// }

// func updatePeople(c *gin.Context) {
// 	var body struct {
// 		Name string `json:"name" binding:"required"`
// 		Age  int    `json:"age" binding:"gt=16"`
// 	}

// 	err := c.BindJSON(&body)
// 	if err != nil {
// 		response.FailResponse(c, http.StatusBadRequest, err, "failed to bind JSON")
// 		return
// 	}

// 	err = handler.UpdatePeopleHandler(c, body.Age, body.Name)
// 	if err != nil {
// 		response.FailResponse(c, http.StatusNotFound, err, "update db err")
// 		return
// 	}

// 	response.SuccResponse(c, http.StatusOK, "update done", body)
// 	return
// }

// 分支控制 : 單分支 v.s 多分支
// sample
// if ...{
// 	xxx
// } else if ...{
// 	xxx
// } else {
// 	xxx
// }
// 最後的 else 可省略
