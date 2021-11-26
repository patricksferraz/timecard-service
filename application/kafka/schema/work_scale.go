package schema

type WorkScale struct {
	Base      `json:",inline" valid:"required"`
	CompanyID string `json:"company_id" valid:"uuid"`
}

func NewWorkScale() *WorkScale {
	return &WorkScale{}
}
