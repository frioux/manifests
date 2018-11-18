package manifests

import (
	"fmt"

	"github.com/pkg/errors"
)

// Inject sets a value deep in a complex data structure
// XXX this doesn't support arrays yet and should in the same fashion as
// Check.  Also needs tests.
func Inject(h interface{}, path []string, value interface{}) error {
	v, ok := h.(map[string]interface{})
	if !ok {
		return errors.New("Cannot munge non-hash resource " + fmt.Sprintf("%#v", h))
	}

	def := map[string]interface{}{}
	for _, k := range path[:len(path)-1] {
		def = map[string]interface{}{}
		raw, ok := v[k]
		if ok && raw != nil {
			v, ok = raw.(map[string]interface{})
			if !ok {
				return errors.New("Cannot munge non-hash " + fmt.Sprintf("%#v", raw))
			}
			def = v
		} else {
			v[k] = def
		}
	}

	def[path[len(path)-1]] = value

	return nil
}
