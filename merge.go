package unienv

//go:generate bin/generr -t noEnvName -i
type noEnvName interface {
	NoEnvName() (envname string)
}

func Merge(envname string, m ConfigMap) (map[string]string, error) {
	merged := m["default"]

	envConf, ok := m[envname]
	if !ok {
		return nil, &NoEnvName{Envname: envname}
	}

	for k, v := range envConf {
		merged[k] = v
	}

	return merged, nil
}
