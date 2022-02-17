package update

import (
	"io"
	"net/http"
	"os"
)

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	os.Remove(filepath)

	// Create the file
	out, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0700)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
