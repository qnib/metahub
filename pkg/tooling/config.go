package tooling

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// MachineType stores the name and the list of features
type MachineType struct {
	DisplayName string   `yaml:"name"`
	Features    []string `yaml:"features"`
}

// Config holds a list of MachineTypes
type Config struct {
	//TODO:  Password should be a hash
	User     string        `yaml:"user"`
	Password string        `yaml:"password"`
	Types    []MachineType `yaml:"types"`
}

// ParseConfig takes a bytearray and unmarshals it
func ParseConfig(data []byte) (cfg Config, err error) {
	err = yaml.Unmarshal(data, &cfg)
	return
}

// CreateConfigFromFile parses a file and returns a workshop
func CreateConfigFromFile(fpath string) (cfg Config, err error) {
	log.Printf("Reading file: %s", fpath)
	yData, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	cfg, err = ParseConfig(yData)
	return
}
