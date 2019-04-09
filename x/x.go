package x

import (
	"github.com/daved/multimod/a"
	"github.com/daved/multimod/b"
	"github.com/daved/multimod/d"
	"github.com/daved/multimod/e"
	"github.com/daved/multimod/f/v2"
)

// X contains multimod funcs.
type X map[string]func() string

// Make returns an X formed using various functions from various modules from
// the multimod repo.
func Make() X {
	return X{
		"a": a.Version1,
		"b": b.Version1,
		"d": d.Version,
		"e": e.Version1,
		"f": f.Version2,
	}
}
