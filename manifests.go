package manifests

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/frioux/yaml"
	"github.com/pkg/errors"
)

//go:generate ./destiny.pl paths_generated.go io.k8s.api.core.v1.Container io.k8s.api.core.v1.PodSpec io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta

var paths map[string]map[string][][]string

// Walk takes a needle type, a list of resources, and a function to verify and
// possibly transform the resource of needleType within the original resource
func Walk(needleType string, r interface{}, f func(r interface{}) error) error {
	resourcePaths, ok := paths[needleType]
	if !ok {
		return errors.New("No paths defined, update go:generate")
	}

	rtype, err := ResourceType(r)
	if err != nil {
		return errors.Wrap(err, "manifests.ResourceType")
	}
	rpaths, ok := resourcePaths[rtype]
	if !ok {
		return nil
	}

	for _, p := range rpaths {
		_, err := Check(r, p, f)
		if err != nil {
			return errors.Wrap(err, "findHash")
		}
	}

	return nil
}

// ToString returns a list of resources as a multi doc YAML string.
func ToString(r []interface{}) (string, error) {
	buf := &bytes.Buffer{}

	e := yaml.NewEncoder(buf)
	for _, doc := range r {
		err := e.Encode(doc)
		if err != nil {
			return "", errors.Wrap(err, "yaml.Encode")
		}
	}

	return buf.String(), nil
}

// Check descends within h via path and runs f.  If no data was found the
// boolean result is false.
func Check(h interface{}, path []string, f func(interface{}) error) (bool, error) {
	if len(path) == 0 {
		return true, f(h)
	}

	if path[0] == "" {
		v, ok := h.([]interface{})
		if !ok {
			return true, errors.New(`"" passed for non-slice ` + fmt.Sprintf("%#v", h))
		}
		for _, val := range v {
			_, err := Check(val, path[1:], f)
			if err != nil {
				return true, err
			}
		}
		return true, nil
	}

	v, ok := h.(map[string]interface{})
	if !ok {
		return true, errors.New("Cannot search non-hash" + fmt.Sprintf("%#v", h))
	}

	inner, ok := v[path[0]]
	if !ok {
		return false, nil
	}

	return Check(inner, path[1:], f)
}

type ktype struct{ group, kind, version string }

var types map[ktype]string

var errNoSuchType = errors.New("no such type")

// ResourceType resolves the kind and apiVersion to a single ResourceType string
func ResourceType(h interface{}) (string, error) {
	v, ok := h.(map[string]interface{})
	if !ok {
		return "", errors.New("Cannot get type for non-hash resource " + fmt.Sprintf("%#v", h))
	}

	apiVersionRaw, ok := v["apiVersion"]
	if !ok {
		return "", errors.New("apiVersion is required")
	}

	apiVersion, ok := apiVersionRaw.(string)
	if !ok {
		return "", errors.New("apiVersion must be a string " + fmt.Sprintf("%#v", apiVersionRaw))
	}

	kindRaw, ok := v["kind"]
	if !ok {
		return "", errors.New("kind is required")
	}

	kind, ok := kindRaw.(string)
	if !ok {
		return "", errors.New("kind must be a string " + fmt.Sprintf("%#v", kindRaw))
	}

	split := strings.Split(apiVersion, "/")
	if len(split) == 2 {
		found, kok := types[ktype{group: split[0], version: split[1], kind: kind}]
		if !kok {
			return "", errNoSuchType
		}
		return found, nil
	}

	found, kok := types[ktype{version: apiVersion, kind: kind}]
	if !kok {
		return "", errNoSuchType
	}
	return found, nil
}
