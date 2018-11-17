package e2e

import (
	"bytes"
	"github.com/k0kubun/pp"
	"os/exec"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestE2E(t *testing.T) {
	cases := []struct {
		name   string
		in     string
		expect []string
	}{
		{
			name:   "develop",
			in:     "develop",
			expect: []string{"ENV1=xxx", "ENV2=bar", "ENV3=baz", ""},
		},
		{
			name:   "staging",
			in:     "staging",
			expect: []string{"ENV1=111", "ENV2=222", "ENV3=333", "ENV4=444", ""},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cmd := exec.Command("./unienv", "unify", "--dry-run", "--env", c.name)
			var out bytes.Buffer
			cmd.Stdout = &out

			err := cmd.Run()
			if err != nil {
				t.Fatal(err)
			}
			acts := strings.Split(out.String(), "\n")
			sort.Strings(c.expect)
			sort.Strings(acts)
			if !reflect.DeepEqual(c.expect, acts) {
				pp.Println(acts)
				t.Errorf("must be same %+v but actual %+v", c.expect, acts)
			}
		})
	}
}
