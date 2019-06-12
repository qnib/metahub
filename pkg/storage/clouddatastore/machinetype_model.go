package clouddatastore

var machineTypeEntityKind = "MachineType"

type machineTypeModel struct {
	DisplayName string   `datastore:"name,noindex"`
	Features    []string `datastore:"features,noindex"`
	Login       string   `datastore:"login"`
	Password    string   `datastore:"password,noindex"`
}

