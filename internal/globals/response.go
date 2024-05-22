package globals

type GeneralResponse struct {
	Status     string `json:"status"`
	Message    string `json:"message"`
	ResponseId string `json:"responseId"`
	Data       any    `json:"data"`
}

type GeneralError struct {
	Status     string `json:"status"`
	Error      string `json:"message"`
	ResponseId string `json:"responseId"`
}
