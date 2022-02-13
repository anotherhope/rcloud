package message

// Response is a representation of answer from server
type Response struct {
	Data []byte
}

// ToString transform Response to string
func (r *Response) ToString() string {
	return string(r.Data)
}
