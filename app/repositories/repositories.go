package repositories

import (
	"fmt"

	"github.com/anotherhope/rcloud/app/config"
)

func Contains(name string) bool {
	for _, d := range config.Get().Repositories {
		if name == d.Destination {
			return true
		}
	}
	return false
}

func Add(d *config.Directory) error {
	if Contains(d.Destination) {
		return fmt.Errorf("value already existe %s", d.Name)
	}

	config.Set("repositories", append(config.Get().Repositories, d))
	return nil
}

func Del(n string) {
	repo := config.Get().Repositories

	for k, v := range repo {
		if v.Name == n {
			config.Set("repositories", append(repo[:k], repo[k+1:]...))
			break
		}
	}

}

func List() []*config.Directory {
	return config.Get().Repositories
}
