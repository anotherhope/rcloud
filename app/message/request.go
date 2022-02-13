package message

import "strings"

// Request is a representation of a query from client socket
type Request struct {
	Method string
	Args   []string
}

// ToBytes transform Request to bytes
func (r *Request) ToBytes() []byte {
	return []byte(r.Method + ":" + strings.Join(r.Args, ",") + "\n")
}

// ReqStatus is a formated request to get status of a repository
func ReqStatus(repositoryName string) *Request {
	return &Request{
		Method: "getStatus",
		Args:   []string{repositoryName},
	}
}
