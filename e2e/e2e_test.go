package e2e

import (
	"bytes"
	"github.com/andreyvit/diff"
	"log"
	"os/exec"
	"testing"
)

func TestE2E(t *testing.T) {
	t.Run("Generate With env=develp", func(t *testing.T) {
		cmd := exec.Command("./unienv", "unify", "--dry-run", "--env", "develop")
		var out bytes.Buffer
		cmd.Stdout = &out

		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		expect := `ENV1=xxx
ENV2=bar
ENV3=baz`

		if act := out.String(); act != expect {
			t.Errorf("must be same %v", diff.LineDiff(act, expect))
		}
	})
}
