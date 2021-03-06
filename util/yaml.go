package util

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func ReadYAML(f string) ([]byte, error) {
	m := make(map[interface{}]interface{})

	filename, _ := filepath.Abs(f)
	fb, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(fb, &m)
	if err != nil {
		return nil, err
	}

	yb, err := yaml.Marshal(&m)
	if err != nil {
		return nil, err
	}

	return yb, nil
}
