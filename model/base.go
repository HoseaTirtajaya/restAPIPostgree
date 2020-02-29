package model

//ResponseAPI
type ResponseAPI struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
