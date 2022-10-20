package pkg

type Response struct {
	Code      int    `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
	Reference string `json:"reference,omitempty"`
}
