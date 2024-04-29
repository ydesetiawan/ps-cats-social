package response

type WebResponse struct {
	Status     int         `json:"status"`
	Message    string      `json:"message"`
	Error      interface{} `json:"error"`
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination,omitempty"`
}
