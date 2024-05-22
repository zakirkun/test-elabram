package types

type RequestCreatEmploye struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FullName  string `json:"full_name"`
	IDCOmpany int    `json:"company_id"`
	Msisdn    string `json:"msisdn"`
	Position  string `json:"position"`
	Foto      string `json:"foto"`
}

type ResponseCreateEmploye struct {
	Message string `json:"message"`
}

type RequestGetEmploye struct {
	ID int `uri:"id" binding:"required"`
}

type ResponseGetEmploye struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	FullName  string `json:"full_name"`
	IDCOmpany int    `json:"company_id"`
	Msisdn    string `json:"msisdn"`
	Position  string `json:"position"`
	Foto      string `json:"foto"`
}

type RequestUpdateEmploye struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Position string `json:"position"`
	Foto     string `json:"foto"`
	ID       int    `json:"id"`
}

type ResponseUpdateEmploye struct {
	Message string `json:"message"`
}

type ResponseDeleteEmploye struct {
	Message string `json:"message"`
}

type RequestLoginEmploye struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseLoginEmploye struct {
	Token string `json:"token"`
}
