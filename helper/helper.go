package helper

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func APIResponse(message string, status string, data interface{}) Response {

	jsonResponse := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return jsonResponse
}
