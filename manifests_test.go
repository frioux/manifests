package manifests

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestCheckSimple(t *testing.T) {
	var found int
	_, err := Check(map[string]interface{}{
		"foo": 1,
	}, []string{"foo"}, func(i interface{}) error {

		val, ok := i.(int)
		if !ok {
			return errors.New("not an int (" + fmt.Sprintf("%#v", i) + ")")
		}

		found = val

		return nil
	})

	assert.NoError(t, err)
	assert.Equal(t, 1, found)
}

func TestCheckList(t *testing.T) {
	var found []int

	_, err := Check(map[string]interface{}{
		"foo": []interface{}{
			1, 2, 3,
		},
	}, []string{"foo", ""}, func(i interface{}) error {

		val, ok := i.(int)
		if !ok {
			return errors.New("not an int (" + fmt.Sprintf("%#v", i) + ")")
		}

		found = append(found, val)

		return nil
	})

	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3}, found)
}

func TestCheckListMap(t *testing.T) {
	var found []int

	_, err := Check(map[string]interface{}{
		"foo": []interface{}{
			map[string]interface{}{"a": 1},
			map[string]interface{}{"a": 2},
			map[string]interface{}{"a": 3},
			map[string]interface{}{"b": 1},
		},
	}, []string{"foo", "", "a"}, func(i interface{}) error {

		val, ok := i.(int)
		if !ok {
			return errors.New("not an int (" + fmt.Sprintf("%#v", i) + ")")
		}

		found = append(found, val)

		return nil
	})

	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3}, found)
}

func TestCheckNotFound(t *testing.T) {
	ok, err := Check(map[string]interface{}{
		"foo": []interface{}{1, 2, 3},
	}, []string{"bar", "baz", "biff"}, func(i interface{}) error { return nil })

	assert.NoError(t, err)
	assert.False(t, ok)

	ok, err = Check(map[string]interface{}{
		"foo": []interface{}{1, 2, 3},
	}, []string{"bar", ""}, func(i interface{}) error { return nil })

	assert.NoError(t, err)
	assert.False(t, ok)
}

func TestCheckBool(t *testing.T) {
	errTooBig := errors.New("too big")
	_, err := Check(map[string]interface{}{
		"foo": []interface{}{
			map[string]interface{}{"a": 1},
			map[string]interface{}{"a": 2},
			map[string]interface{}{"a": 3},
			map[string]interface{}{"b": 1},
		},
	}, []string{"foo", "", "a"}, func(i interface{}) error {

		val, ok := i.(int)
		if !ok {
			return errors.New("not an int (" + fmt.Sprintf("%#v", i) + ")")
		}

		if val > 1 {
			return errTooBig
		}

		return nil
	})

	assert.Equal(t, err, errTooBig)
}

func TestResourceType(t *testing.T) {
	found, err := ResourceType(
		map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, "io.k8s.api.apps.v1.Deployment", found)

	found, err = ResourceType(
		map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, "io.k8s.api.core.v1.Service", found)

	_, err = ResourceType(
		map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "HERP",
		},
	)

	assert.Equal(t, errNoSuchType, err)
}
