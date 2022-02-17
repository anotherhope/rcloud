package rclone

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/anotherhope/rcloud/app/internal"
)

func gitIgnore(d *internal.Directory) []string {
	ignores := make([]string, 0)
	if d.IsLocal(d.Source) {
		if file, err := os.Open(d.Source + "/.gitignore"); err != nil {
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				fmt.Println(line)
				if !strings.HasPrefix(line, "#") {
					ignores = append(ignores, scanner.Text())
				}
			}

			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}

			return make([]string, 0)
		}
	}

	return ignores
}
