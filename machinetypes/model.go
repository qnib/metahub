package machinetypes

var machineTypeEntityKind = "MachineType"

type machineType struct {
	DisplayName string   `datastore:"name,noindex"`
	Features    []string `datastore:"features,noindex"`
	Login       string   `datastore:"login"`
	Password    string   `datastore:"password,noindex"`
}
