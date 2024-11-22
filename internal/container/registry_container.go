package container

import "github.com/umardev500/pos-api/pkg"

func NewRegistryContainer(db *pkg.PGX, v pkg.Validator) []pkg.Container {
	return []pkg.Container{
		NewAuthContainer(db, v),
	}
}
