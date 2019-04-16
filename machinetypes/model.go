package machinetypes

var machineTypeEntityKind = "MachineType"

type machineType struct {
	DisplayName string   `datastore:"name,noindex"`
	Features    []string `datastore:"features,noindex"`
	Password    string   `datastore:"password,noindex"`
}

/*func getPassword(accountName string, id int64) {
	return fmt.Sprintf("%s-%d", accountName, id)
}

func getID*/
