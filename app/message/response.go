package message

type Response struct {
	Data []byte
}

func (r *Response) ToString() string {
	return string(r.Data)
}
