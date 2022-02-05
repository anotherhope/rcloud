package repositories

import (
	"fmt"
	"strings"

	"github.com/anotherhope/rcloud/app/config"
)

func parent(name string) bool {
	repo := config.Get().Repositories

	for _, d := range repo {
		if strings.HasPrefix(d.Destination, name) {
			return true
		}
	}

	return false
}

func sub(name string) bool {
	repo := config.Get().Repositories

	for _, d := range repo {
		if strings.HasPrefix(name, d.Destination) {
			return true
		}
	}

	return false
}

func same(name string) bool {
	repo := config.Get().Repositories

	for _, d := range repo {
		if name == d.Destination {
			return true
		}
	}

	return false
}

func Add(d *config.Directory) error {

	if same(d.Destination) {
		return fmt.Errorf("destination path already exist as a sync folder ")
	} else if parent(d.Destination) {
		return fmt.Errorf("destination path is parent directory of a sync folder ")
	} else if sub(d.Destination) {
		return fmt.Errorf("destination path is sub directory of a sync folder ")
	}

	config.Set("repositories", append(config.Get().Repositories, d))
	return nil
}

func Del(n string) error {
	repo := config.Get().Repositories

	for k, v := range repo {
		if v.Name == n {
			config.Set("repositories", append(repo[:k], repo[k+1:]...))
			return nil
		}
	}

	return fmt.Errorf("repository not exists")
}

func List() []*config.Directory {
	return config.Get().Repositories
}
