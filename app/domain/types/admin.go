package types

type RequestCreatAdmin struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

type ResponseCreateAdmin struct {
	Message string `json:"message"`
}

type RequestGetAdmin struct {
	ID int `uri:"id" binding:"required"`
}

type ResponseGetAdmin struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
}

type RequestUpdateAdmin struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	ID       int    `json:"id"`
}

type ResponseUpdateAdmin struct {
	Message string `json:"message"`
}

type ResponseDeleteAdmin struct {
	Message string `json:"message"`
}

type RequestLoginAdmin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseLoginAdmin struct {
	Token string `json:"token"`
}
