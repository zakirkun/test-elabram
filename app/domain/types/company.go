package types

type RequestCreateCompany struct {
	Name        string `json:"company_name"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
}

type ResponseCreateCompany struct {
	Message string `json:"message"`
}

type ResponseGetCompany struct {
	Name        string `json:"company_name"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
}
