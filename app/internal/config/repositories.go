package config

import (
	"fmt"
	"strings"

	"github.com/anotherhope/rcloud/app/internal/repositories"
)

func parent(name string) bool {
	for _, d := range repositories.Repositories {
		if strings.HasPrefix(d.Destination, name) {
			return true
		}
	}

	return false
}

func sub(name string) bool {
	for _, d := range repositories.Repositories {
		if strings.HasPrefix(name, d.Destination) {
			return true
		}
	}

	return false
}

func same(name string) bool {
	for _, d := range repositories.Repositories {
		if name == d.Destination {
			return true
		}
	}

	return false
}

func exists(d *repositories.Repository) bool {
	for _, repository := range repositories.Repositories {
		if repository.Name == d.Name {
			return true
		}
	}

	return false
}

// Add repository in configuration file
func Add(d *repositories.Repository) error {
	var exitMessage error = nil
	if same(d.Destination) {
		exitMessage = fmt.Errorf("destination path already exist as a sync folder ")
	} else if parent(d.Destination) {
		exitMessage = fmt.Errorf("destination path is parent repository of a sync folder ")
	} else if sub(d.Destination) {
		exitMessage = fmt.Errorf("destination path is sub repository of a sync folder ")
	} else if exists(d) {
		exitMessage = fmt.Errorf("sorry repository already exists")
	}

	if exitMessage != nil {
		return exitMessage
	}

	App.Set("repositories",
		append(repositories.Repositories, d),
	)

	return exitMessage
}

// Del repository in configuration file
func Del(n string) error {
	for k, v := range repositories.Repositories {
		if strings.HasPrefix(v.Name, n) {
			App.Set("repositories", append(
				repositories.Repositories[:k],
				repositories.Repositories[k+1:]...,
			))
			return nil
		}
	}

	return fmt.Errorf("repository not exists")
}

/*
// List repositories in configuration file
func List() []*repositories.Repository {
	return repositories.List()
}

// GetRepository repository by name
func GetRepository(repositoryName string) *repositories.Repository {
	return repositories.GetRepository(repositoryName)
}

// IsValid detect if remote repository is configured and valid
func IsValid(path string, isRemote bool) (string, error) {
	return repositories.IsValid(path, isRemote)
}
*/
