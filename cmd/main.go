package main

import (
	"companyEmployee/cmd/handler"
	"companyEmployee/cmd/response"
	"companyEmployee/model"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
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
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)
	// 17 , -5 , 12
	// go say("world") //開一個新的 Goroutines 執行
	// say("hello")    //當前 Goroutines 執行

	// 定義 a 為空介面
	// var a interface{}
	// var i int = 5
	// s := "Hello world"
	// a 可以儲存任意型別的數值
	// a = i
	// a = s

	// fmt.Println(a)

	// slice := []int{1, 2, 3, 4, 5, 7}
	// fmt.Println("slice = ", slice)
	// odd := filter(slice, isOdd) // 函式當做值來傳遞了
	// fmt.Println("Odd elements of slice are: ", odd)
	// even := filter(slice, isEven) // 函式當做值來傳遞了
	// fmt.Println("Even elements of slice are: ", even)
	// x := 3
	// fmt.Println("x = ", x)    // 應該輸出 "x = 3"
	// x1 := add1(&x)            //呼叫 add1(x)
	// fmt.Println("x+1 = ", x1) // 應該輸出"x+1 = 4"
	// fmt.Println("x = ", x)    // 應該輸出"x = 4"

	// c := [...]int{4, 5, 6, 7, 8}
	// fmt.Println(c)
	// fmt.Println(len(c))

	// numbers := make(map[string]int)
	// var num map[string]int
	// num = make(map[string]int)
	// num["one"] = 1  //賦值
	// num["ten"] = 10 //賦值
	// num["three"] = 3

	// fmt.Println(num)

	// {"one" : 1, "ten":10, "three":3}
	// err := handler.initDB()
	// if err != nil {
	// 	log.Fatalf("Failed to initialize the database: %v", err)
	// }
	// a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// b := []int{}
	// c := []map[string]int{}
	// var sliceOfMaps []map[string]int

	// for _, result := range a {
	// 	if result%2 == 0 {
	// 		c = append(c, map[string]int{"偶數": result})
	// 	}
	// }

	// fmt.Println("")
	// fmt.Println(c)
	// fmt.Println("")

	// router := gin.New()
	// mid1 := func(c *gin.Context) {
	// 	fmt.Println("mid1 start")
	// 	// c.Abort()
	// 	c.Next()
	// 	fmt.Println("mid1 end")
	// }

	// mid2 := func(c *gin.Context) {
	// 	fmt.Println("mid2 start")
	// 	// c.Abort()
	// 	c.Next()
	// 	// c.Abort()
	// 	fmt.Println("mid2 end")
	// }

	// mid3 := func(c *gin.Context) {
	// 	fmt.Println("mid3 start")
	// 	// c.Abort()
	// 	// c.Next()
	// 	fmt.Println("mid3 end")
	// }

	// mid4 := func(c *gin.Context) {
	// 	fmt.Println("mid4 start")
	// 	// 加上條件判斷，如果不符合條件，則中止請求
	// 	if !middleware.CheckCondition(c) {
	// 		c.Abort()
	// 		return
	// 	}
	// 	c.Next()
	// 	fmt.Println("mid4 end")
	// }

	// BasicAuth := gin.BasicAuth(gin.Accounts{
	// 	"lorem": "ipsum",
	// })

	// router.Use(mid1, mid2, mid3)
	// router.GET("/", BasicAuth, mid4, func(c *gin.Context) {
	// 	fmt.Println("process get request")
	// 	c.JSON(http.StatusOK, "hello")
	// })

	// alphabetMap := make(map[string]bool)
	// for _, r := range alphabetStr {
	// 	c := string(r)
	// 	alphabetMap[c] = true
	// }
	// fmt.Println(alphabetMap["x"])
	// alphabetMap["x"] = false
	// fmt.Println(alphabetMap["x"])

	router := gin.Default()

	// 定義一個切片
	// weekdays := []string{"一 ", "二 ", "三 ", "四 ", "五 ", "六 ", "日 "}
	// for _, weekday := range weekdays {
	// 	fmt.Print(weekday)
	// }

	// fmt.Println("")
	// fmt.Println("")

	// weekdays := []string{"一 ", "二 ", "三 ", "四 ", "五 ", "六 ", "日 "}
	// for _, weekday := range weekdays {
	// 	fmt.Print(weekday)
	// }

	// 使用 for range 遍歷切片
	// for _, num := range numbers {
	// 如果數字
	// 打印當前數字
	// fmt.Println(num)

	// 如果數字等於 8，結束循環
	// if num == 8 {
	// 	break
	// }
	// fmt.Println("循環結束")

	// router.Use(BasicAuth)

	// authorized := router.Group("/", BasicAuth)

	// authorized.GET("/welcome", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "welcome to middleware",
	// 	})
	// })

	// router.POST("/api/company", wrapper)
	// router.POST("/api/people", addPeople)
	// router.GET("/api/list", checkList)
	// router.PUT("/api/update", updatePeople)
	// router.GET("/api/active", activeCheck)
	// router.GET("/api/batch_active", activeBatchCheck)

	err := router.Run(":8080")
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

//	func loggEndpointURL(c *gin.Context) {
//		c.Next()
//		log.Printf("EndPoint URL is %v", c.Request.URL)
//	}
//

// 批次公司 is_active 確認 把沒有 ａctive 的印出
func activeBatchCheck(c *gin.Context) {
	var params struct {
		Company []model.Company `json:"companies" binding:"required"`
	}

	var num map[string]int
	num = make(map[string]int)
	num["one"] = 1  //賦值
	num["ten"] = 10 //賦值
	num["three"] = 3

	err := c.BindJSON(&params)
	if err != nil {
		response.FailResponse(c, http.StatusBadRequest, err, "failed to bind JSON")
		return
	}

	company, err := handler.BatchCompanyHandler(params.Company)
	if err != nil {
		response.FailResponse(c, http.StatusNotFound, err, "company not found")
		return
	}

	results := batchCheckActive(company)

	response.SuccResponse(c, http.StatusOK, results, num)
	return
}

// 使用 map 判斷此公司有沒有 active, 用名字先撈出 然後判斷是否是 active
func activeCheck(c *gin.Context) {
	var params struct {
		Name string `form:"name" binding:"required"`
	}

	err := c.BindQuery(&params)
	if err != nil {
		response.FailResponse(c, http.StatusBadRequest, err, "failed to bind JSON")
		return
	}

	company, err := handler.TestHandler(params.Name)
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

func checkActive(company *model.Company) bool {
	return company.IsActive == 1
}

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
