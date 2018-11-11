package monoenv

import (
	"gopkg.in/ini.v1"
	"io"
	"io/ioutil"
)

//go:generate generr -t noDefaultSection -i
type noDefaultSection interface {
	NoDefaultSection()
}

type ConfigMap map[string]map[string]string

func NewConfigMap() ConfigMap {
	return map[string]map[string]string{}
}

func Parse(r io.Reader) (ConfigMap, error) {
	in, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	cfg, err := ini.Load(
		[]byte(in),
		"env",
		ioutil.NopCloser(r),
	)
	if err != nil {
		return nil, err
	}

	defaultSection := cfg.Section("default")
	if defaultSection == nil {
		return nil, &NoDefaultSection{}
	}

	m := NewConfigMap()
	m["default"] = defaultSection.KeysHash()

	return m, nil
}
