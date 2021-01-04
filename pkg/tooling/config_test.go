package tooling

import (
	"fmt"
	"testing"
)

var testConfig1 = `
user: user
password: hash
types:
  - name: broadwell
    features:
    - cpu:broadwell
  - name: skylake
    features:
    - cpu:skylake
`

func Test_ParseConfig_testConfig1(t *testing.T) {
	cfg, err := ParseConfig([]byte(testConfig1))
	if err != nil {
		t.Error(err.Error())
	}
	if cfg.User != "user" {
		t.Errorf("User should be 'user'; got '%s'", cfg.User)
	}
	if len(cfg.Types) != 2 {
		t.Error("Length of testConfig1 should be 2")
	}
	for _, mt := range cfg.Types {
		fmt.Println(mt)
	}
}
