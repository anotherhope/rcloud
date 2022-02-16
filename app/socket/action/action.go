package action

import (
	"strings"

	"github.com/anotherhope/rcloud/app/interfaces"
)

// Do Execute a command from socket
func Do(query string) []byte {
	queryParts := strings.Split(strings.TrimSuffix(query, "\n"), ":")
	action := queryParts[0]
	arguments := queryParts[1:]

	switch action {
	case "getStatus":
		if repository := interfaces.Get(arguments[0]); repository != nil {
			return []byte(repository.GetStatus())
		}
	}

	return []byte{}
}
