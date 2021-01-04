package tooling

import (
	"fmt"
	"testing"
)

var testConfig1 = `
- name: broadwell
  features:
  - cpu:broadwell
- name: skylake
  features:
  - cpu:skylake
`

func Test_ParseMAchineTypes_testConfig1(t *testing.T) {
	mts, err := ParseMAchineTypes([]byte(testConfig1))
	if err != nil {
		t.Error(err.Error())
	}
	if len(mts) != 2 {
		t.Error("Length of testConfig1 should be 2")
	}
	for _, mt := range mts {
		fmt.Println(mt)
	}
}
