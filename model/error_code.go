package model

type Code string

const (
	ErrAddCompanyFail Code = "add company fail"
	ErrNotActive      Code = "not active company"
)

func (c Code) Error() string {
	return string(c)
}
