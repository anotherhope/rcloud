package internal

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

func parent(name string) bool {
	for _, d := range App.Repositories {
		if strings.HasPrefix(d.Destination, name) {
			return true
		}
	}

	return false
}

func sub(name string) bool {
	for _, d := range App.Repositories {
		if strings.HasPrefix(name, d.Destination) {
			return true
		}
	}

	return false
}

func same(name string) bool {
	for _, d := range App.Repositories {
		if name == d.Destination {
			return true
		}
	}

	return false
}

func exists(d *Directory) bool {
	for _, repository := range App.Repositories {
		if repository.Name == d.Name {
			return true
		}
	}

	return false
}

// Add repository in configuration file
func Add(d *Directory) error {
	var exitMessage error = nil
	if same(d.Destination) {
		exitMessage = fmt.Errorf("destination path already exist as a sync folder ")
	} else if parent(d.Destination) {
		exitMessage = fmt.Errorf("destination path is parent directory of a sync folder ")
	} else if sub(d.Destination) {
		exitMessage = fmt.Errorf("destination path is sub directory of a sync folder ")
	} else if exists(d) {
		exitMessage = fmt.Errorf("sorry repository already exists")
	}

	if exitMessage != nil {
		return exitMessage
	}

	App.Set("repositories",
		append(App.Repositories, d),
	)

	return exitMessage
}

// Del repository in configuration file
func Del(n string) error {
	for k, v := range App.Repositories {
		if strings.HasPrefix(v.Name, n) {
			App.Set("repositories", append(
				App.Repositories[:k],
				App.Repositories[k+1:]...,
			))
			return nil
		}
	}

	return fmt.Errorf("repository not exists")
}

// List repositories in configuration file
func List() []*Directory {
	return App.Repositories
}

// Get repository by name
func Get(repositoryName string) *Directory {
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
