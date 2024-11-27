package container

import "github.com/umardev500/pos-api/pkg"

func NewRegistryContainer(db *pkg.GormDB, v pkg.Validator) []pkg.Container {
	return []pkg.Container{
		NewAuthContainer(db, v),
		NewUserContainer(db, v),
	}
}
