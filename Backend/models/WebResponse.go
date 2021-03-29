package models

type WebResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"Message"`
	Data    interface{} `json:"data"`
}

func ConstructWebResponse(code int, message string, data interface{}) WebResponse {
	var webResponse WebResponse
	webResponse.Code = code
	webResponse.Message = message
	webResponse.Data = data

	return webResponse
}
