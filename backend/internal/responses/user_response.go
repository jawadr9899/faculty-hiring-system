package responses

type DefaultResponse struct {
	Status  int  `json:"status"`
	Success bool `json:"success"`
	Message any  `json:"message"`
}
