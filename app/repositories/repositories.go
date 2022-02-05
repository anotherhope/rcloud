package repositories

import (
	"fmt"

	"github.com/anotherhope/rcloud/app/config"
)

func Add(d *config.Directory) error {
	if value, ok := config.Get().Repositories[d.Name]; ok {
		return fmt.Errorf("value already existe %s", value.Name)
	}

	config.Get().Repositories[d.Name] = d
	config.Set("repositories", config.Get().Repositories)
	return nil
}

func Del(n string) {
	delete(config.Get().Repositories, n)
	config.Set("repositories", config.Get().Repositories)
}

func List() map[string]*config.Directory {
	return config.Get().Repositories
}
