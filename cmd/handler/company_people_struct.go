package handler

type Company struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type People struct {
	Name        string `json:"name"`
	CompanyCode string `json:"company_code"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
}

type Peoplelist struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
