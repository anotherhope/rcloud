package repositories

type Repositories struct {
	collections map[string]*Directory `mapstructure:"Repositories"`
}

func (r *Repositories) Add(d *Directory) {
	r.collections[d.Name] = d
}

func (r *Repositories) List() map[string]*Directory {
	return r.collections
}
