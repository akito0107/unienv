package unienv_test

import (
	"github.com/akito0107/unienv"
	"log"
	"reflect"
	"testing"
)

func TestMerge_MergeOriginal(t *testing.T) {
	src := map[string]map[string]string{
		"default": {
			"PARAM1": "123",
			"PARAM2": "1234",
			"PARAM3": "12345",
		},
		"develop": {
			"PARAM1": "1234",
		},
	}

	merged, err := unienv.Merge("develop", src)
	if err != nil {
		log.Fatal(err)
	}
	expect := map[string]string{
		"PARAM1": "1234",
		"PARAM2": "1234",
		"PARAM3": "12345",
	}

	if !reflect.DeepEqual(merged, expect) {
		t.Errorf("must be same %v, but actual %v", expect, merged)
	}
}
