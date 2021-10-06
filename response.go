package response

import (
	"encoding/json"
	"net/http"
)

// Response is a object for sending back data to a client.
type Response struct {
	Data   interface{} `json:"data"`
	Errors []Error     `json:"errors"`
}

// Error is an object for error metadata.
type Error struct {
	Message string `json:"message"`
}

// CreateResponse creates a new response object.
func CreateResponse() *Response {
	return &Response{}
}

// SetData sets the data field of a response object.
func (r *Response) SetData(data interface{}) {
	r.Data = data
}

// AppendError appends to the errors field of a response object.
func (r *Response) AppendError(err error) {
	ErrorObject := Error{Message: err.Error()}
	r.Errors = append(r.Errors, ErrorObject)
}

// SendWithStatusCode writes the current response object as JSON with a
// given status code as an http response.
func (r *Response) SendWithStatusCode(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	encoder := json.NewEncoder(w)
	encoder.Encode(r)
}

// SendWithStatusCode writes the current response object as JSON with a
// given status code as an http response.
func (r *Response) SendJSONWithStatusCode(w http.ResponseWriter, JSON []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(JSON)
}

// SendDataWithStatusCode is a shorthand method that sets data on a request and writes
// it in a single call. This eliminates repetitive code.
func (r *Response) SendDataWithStatusCode(w http.ResponseWriter, data interface{}, code int) {
	r.SetData(data)
	r.SendWithStatusCode(w, code)
}

// SendErrorWithStatusCode is a shorthand method that apppends an error to a request and writes
// it in a single call. This eliminates repetitive code.
func (r *Response) SendErrorWithStatusCode(w http.ResponseWriter, err error, code int) {
	r.AppendError(err)
	r.SendWithStatusCode(w, code)
}
