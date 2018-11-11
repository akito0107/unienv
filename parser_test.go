package monoenv_test

import (
	"bytes"
	"github.com/akito0107/monoenv"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	for _, testcase := range []struct {
		name string
		in   string
		out  monoenv.ConfigMap
	}{
		{
			name: "simple case",
			in: `[default]
ENV1=test
`,
			out: map[string]map[string]string{
				"default": {
					"ENV1": "test",
				},
			},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {

			buf := bytes.NewBufferString(testcase.in)

			act, err := monoenv.Parse(buf)

			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(testcase.out, act) {
				t.Error(act)
			}
		})
	}
}
