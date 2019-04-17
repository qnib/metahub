package clouddatastore

var machineTypeEntityKind = "MachineType"

type machineTypeModel struct {
	DisplayName string   `datastore:"name,noindex"`
	Features    []string `datastore:"features,noindex"`
	Password    string   `datastore:"password,noindex"`
}
