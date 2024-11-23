package pkg

type ErrorDetail struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
}

type Response struct {
	StatusCode int          `json:"-"`
	Success    bool         `json:"success"`
	Message    string       `json:"message,omitempty"`
	Data       interface{}  `json:"data,omitempty"`
	Error      *ErrorDetail `json:"error,omitempty"`
	RefCode    string       `json:"ref_code,omitempty"`
}

type ValidationErrorMeta struct {
	Code    string      `json:"code"`
	Type    string      `json:"type"`
	Field   string      `json:"field"`
	Minimum int         `json:"minimum,omitempty"`
	Maximum int         `json:"maximum,omitempty"`
	Exact   int         `json:"exact,omitempty"`
	Items   interface{} `json:"items,omitempty"`
	Message string      `json:"message"`
}
