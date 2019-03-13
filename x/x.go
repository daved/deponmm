package x

import (
	"github.com/daved/multimod/a"
	"github.com/daved/multimod/b"
)

// X contains multimod funcs.
type X map[string]func() string

// Make returns an X formed using various functions from various modules from
// the multimod repo.
func Make() X {
	return X{
		"a": a.Version,
		"b": b.Version,
	}
}
