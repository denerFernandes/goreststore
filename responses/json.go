package responses

import (
	"encoding/json"
	"net/http"
)

// Default struct for JSON return
type Response struct {
	HttpStatus   int         `json:"http_status"`
	HttpMethod   string      `json:"http_method"`
	Message      string      `json:"message"`
	ErrorMessage string      `json:"error_message"`
	Data         interface{} `json:"data"`
}

// JSON function used for return success response
func JSON(w http.ResponseWriter, r *http.Request, httpStatus int, message string, data interface{}) {
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(Response{
		HttpMethod:   r.Method,
		HttpStatus:   httpStatus,
		Message:      message,
		ErrorMessage: "NO_ERROR",
		Data:         data,
	})
}

// ERROR function used for return fail response
func ERROR(w http.ResponseWriter, r *http.Request, httpStatus int, err error) {
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(Response{
		HttpMethod:   r.Method,
		HttpStatus:   httpStatus,
		Message:      err.Error(),
		ErrorMessage: "ERROR",
		Data:         nil,
	})
}
