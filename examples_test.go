package manifests

import (
	"os"
	"strings"

	"github.com/frioux/yaml"
	"go.zr.org/common/go/errors"
)

func Example() {
	var r interface{}

	d := yaml.NewDecoder(strings.NewReader(`
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hello-world
  namespace: frew
spec:
  replicas: 2
  selector:
    matchLabels:
      app: hello
  template:
    metadata:
      labels:
        app: hello
    spec:
      containers:
      - name: hello
        securityContext:
          runAsUser: 1001
        image: "ubuntu:14.04"
        command: [ "bash", "-c", "env"]
      - name: world
        securityContext:
          runAsUser: 1001
        image: "ubuntu:14.04"
        command: [ "bash", "-c", "echo 'woo!'"]
`))

	err := d.Decode(&r)
	if err != nil {
		panic("Couldn't set up Walk example " + err.Error())
	}

	err = Walk("io.k8s.api.core.v1.PodSpec", r, func(i interface{}) error {
		h, ok := i.(map[string]interface{})
		if !ok {
			return errors.New("nonhash")
		}

		h["shareProcessNamespace"] = true

		hostname, ok := h["hostname"]
		if ok {
			return errors.New("hostname should not be set on pods", "hostname", hostname)
		}

		return nil
	})
	if err != nil {
		panic("Couldn't set up Walk example " + err.Error())
	}

	err = Walk("io.k8s.api.core.v1.Container", r, func(i interface{}) error {
		h, ok := i.(map[string]interface{})
		if !ok {
			return errors.New("nonhash")
		}

		if strings.HasSuffix(h["image"].(string), ":latest") {
			return errors.New("bad image")
		}

		return nil
	})
	if err != nil {
		panic("Couldn't set up Walk example " + err.Error())
	}

	e := yaml.NewEncoder(os.Stdout)
	err = e.Encode(r)
	if err != nil {
		panic("Couldn't set up Walk example " + err.Error())
	}
	// Output:
	// apiVersion: apps/v1
	// kind: Deployment
	// metadata:
	//   name: hello-world
	//   namespace: frew
	// spec:
	//   replicas: 2
	//   selector:
	//     matchLabels:
	//       app: hello
	//   template:
	//     metadata:
	//       labels:
	//         app: hello
	//     spec:
	//       containers:
	//       - command:
	//         - bash
	//         - -c
	//         - env
	//         image: ubuntu:14.04
	//         name: hello
	//         securityContext:
	//           runAsUser: 1001
	//       - command:
	//         - bash
	//         - -c
	//         - echo 'woo!'
	//         image: ubuntu:14.04
	//         name: world
	//         securityContext:
	//           runAsUser: 1001
	//       shareProcessNamespace: true
}
