package views

type Response struct {
	StatusCode   int         `json:"status_code"`
	ResponseBody interface{} `json:"response_body"`
}
