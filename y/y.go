package y

import (
	"github.com/daved/multimod/f"
)

// Y contains multimod funcs.
type Y map[string]func() string

// Make returns an Y formed using various functions from various modules from
// the multimod repo.
func Make() Y {
	return Y{
		"f": f.Version,
	}
}
