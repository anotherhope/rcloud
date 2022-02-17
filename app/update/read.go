package update

import (
	"bytes"
	"net/http"
	"strings"
)

func Read(url string) (string, error) {
	// Get the data
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)

	return strings.TrimSpace(buf.String()), nil
}
