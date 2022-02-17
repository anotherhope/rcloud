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
		if file, err := os.Open(d.Source + "/.gitignore"); err == nil {
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				if !strings.HasPrefix(line, "#") && len(line) > 0 {
					ignores = append(ignores, fmt.Sprintf("--exclude=\"%s\"", line))
				}
			}

			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}

			return ignores
		}
	}

	return ignores
}
