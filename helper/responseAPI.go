package helper

type ResponseAPI struct {
	Acknowledge int         `json:"acknowledge"`
	Status      int         `json:"status"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
}

func APIResponse(acknowledge int, status int, data interface{}, message string) ResponseAPI {
	response := ResponseAPI{
		Acknowledge: acknowledge,
		Status:      status,
		Data:        data,
		Message:     message,
	}

	return response
}
