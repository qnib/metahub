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

// MachineTypes holds a list of MachineTypes
type MachineTypes []MachineType

// ParseMAchineTypes takes a bytearray and unmarshals it
func ParseMAchineTypes(data []byte) (mts MachineTypes, err error) {
	err = yaml.Unmarshal(data, &mts)
	return
}

// CreateMachineTypesFromFile parses a file and returns a workshop
func CreateMachineTypesFromFile(fpath string) (mts MachineTypes, err error) {
	log.Printf("Reading file: %s", fpath)
	yData, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	mts, err = ParseMAchineTypes(yData)
	return
}
