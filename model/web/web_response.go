package web

type WebResponse struct {
	Code   int         `json:"name"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
