package main

import (
	"fmt"

	"github.com/anotherhope/rcloud/app/repositories"
	"github.com/spf13/viper"
)

type Config struct {
	Repositories *repositories.Repositories
}

func LoadConfig() *Config {
	conf := &Config{
		Repositories: &repositories.Repositories{},
	}

	err := viper.Unmarshal(conf)
	fmt.Println(err, conf)

	conf.Repositories.Add(&repositories.Directory{
		Name:        "Toto",
		Source:      "/toto",
		Destination: "/tata",
	})

	fmt.Println(conf.Save())

	/*
		data := &Config{
			Repositories: map[string]*repositories.Directory{
				"Hello": {
					Source:      "/toto",
					Destination: "/tata",
				},
			},
		}

		/*
			data := viper.Get("repositories")

			if data == nil {
				data := &Repositories{
					list: map[string]*repositories.Directory{
						"Hello": {
							Source:      "/toto",
							Destination: "/tata",
						},
					},
				}
				data.Save()
			} else {
				fmt.Println("data", data)
			}
	*/
	return nil
}

func (c *Config) Save() error {
	for key, repository := range c.Repositories.List() {
		viper.Set(key, repository)
	}

	return viper.WriteConfig()
}

/*
	s := Student{"Chetan", "Kumar", "Bangalore", 7777777777}
	v := reflect.ValueOf(s)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}
*/
