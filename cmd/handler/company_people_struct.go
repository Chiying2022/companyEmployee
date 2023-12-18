package handler

type Company struct {
	CODE string `json:"code"`
	NAME string `json:"name"`
}

type People struct {
	NAME        string `json:"name"`
	COMPANYCODE string `json:"company_code"`
	AGE         int    `json:"age"`
	GENDER      string `json:"gender"`
}

type Peoplelist struct {
	ID   string `json:"id"`
	NAME string `json:"name"`
}
