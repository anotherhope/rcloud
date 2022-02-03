package repositories

type Directory struct {
	Name        string `mapstructure:"Name"`
	Source      string `mapstructure:"Source"`
	Destination string `mapstructure:"Destination"`
}
