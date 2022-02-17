package update

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"strings"
)

func ReadRemote(url string) (string, error) {
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

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
