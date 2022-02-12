package action

import (
	"strings"

	"github.com/anotherhope/rcloud/app/repositories"
)

func Do(query string) []byte {
	queryParts := strings.Split(strings.TrimSuffix(query, "\n"), ":")
	action := queryParts[0]
	arguments := queryParts[1:]

	switch action {
	case "getStatus":
		if repository := repositories.Get(arguments[0]); repository != nil {
			return []byte(repository.GetStatus())
		}
	}

	return []byte{}
}
