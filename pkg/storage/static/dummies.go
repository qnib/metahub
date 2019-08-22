package static

import "metahub/pkg/storage"

// Consts for protoype
const (
	user        = "qnib"
	accountName = user
)

// Dummy MachineTypes
var (
	mType0 = storage.MachineType{
		ID:          0,
		DisplayName: "type0",
		Features:    []string{},
		Login:       user + "-type0",
		Password:    user + "-type0",
	}
	mType1 = storage.MachineType{
		ID:          1,
		DisplayName: "type1",
		Features:    []string{"cpu:broadwell"},
		Login:       user + "-type1",
		Password:    user + "-type1",
	}
	mType2 = storage.MachineType{
		ID:          2,
		DisplayName: "type2",
		Features:    []string{"cpu:skylake"},
		Login:       user + "-type2",
		Password:    user + "-type2",
	}
	mType3 = storage.MachineType{
		ID:          3,
		DisplayName: "type3",
		Features:    []string{"cpu:coffelake"},
		Login:       user + "-type3",
		Password:    user + "-type3",
	}
	mType4 = storage.MachineType{
		ID:          4,
		DisplayName: "type4",
		Features:    []string{"cpu:broadwell", "nvcap:5.2"},
		Login:       user + "-type4",
		Password:    user + "-type4",
	}
)

func getMachineTypes() []storage.MachineType {
	return []storage.MachineType{
		mType0, mType1, mType2, mType3, mType4,
	}

}
