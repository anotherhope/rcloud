package message

import "strings"

type Request struct {
	Method string
	Args   []string
}

func (r *Request) ToBytes() []byte {
	return []byte(r.Method + ":" + strings.Join(r.Args, ",") + "\n")
}

func ReqStatus(repositoryName string) *Request {
	return &Request{
		Method: "getStatus",
		Args:   []string{repositoryName},
	}
}
