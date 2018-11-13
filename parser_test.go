package unienv_test

import (
	"bytes"
	"github.com/akito0107/unienv"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	for _, testcase := range []struct {
		name string
		in   string
		out  unienv.ConfigMap
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
		{
			name: "multi section",
			in: `[default]
ENV1=test1
ENV2=test2

[othersection]
ENV1=s1
ENV2=s2
`,
			out: map[string]map[string]string{
				"default": {
					"ENV1": "test1",
					"ENV2": "test2",
				},
				"othersection": {
					"ENV1": "s1",
					"ENV2": "s2",
				},
			},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {

			buf := bytes.NewBufferString(testcase.in)

			act, err := unienv.Parse(buf)

			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(testcase.out, act) {
				t.Error(act)
			}
		})
	}
}
