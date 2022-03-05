package repositories

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

var Repositories = make([]*Repository, 0)

// List repositories in configuration file
func List() []*Repository {
	return Repositories
}

// GetRepository repository by name
func GetRepository(repositoryName string) *Repository {
	for _, repository := range List() {
		if repository.Name == repositoryName {
			return repository
		}
	}

	return nil
}

// IsValid detect if remote repository is configured and valid
func IsValid(path string, isRemote bool) (string, error) {
	if strings.Contains(path, ":") {
		remote := strings.Split(path, ":")[0]
		output, err := exec.Command("rclone", "listremotes").Output()
		if err != nil {
			return path, err
		}

		availableChoices := strings.Split(string(output), "\n")
		if i := sort.SearchStrings(availableChoices, remote+":"); i < 0 {
			return path, fmt.Errorf("rclone remote not available")
		}

		return path, nil
	}

	return filepath.Abs(path)
}
