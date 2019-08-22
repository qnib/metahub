package static

import (
	"metahub/pkg/storage"
)

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
	// Dummy Manifests
	mfUbuntuAMD64, _    = storage.NewManifest("ubuntu", "latest", "linux", "amd64")
	mfUbuntuPPC, _      = storage.NewManifest("ubuntu", "latest", "linux", "ppc64le")
	mfQFeatGeneric, _   = storage.NewManifest("qnib/plain-featuretest", "generic", "linux", "amd64")
	mfQFeatBroadwell, _ = storage.NewManifest("qnib/plain-featuretest", "cpu-broadwell", "linux", "amd64", "cpu:broadwell")
	mfQFeatSkylake, _   = storage.NewManifest("qnib/plain-featuretest", "cpu-skylake", "linux", "amd64", "cpu:skylake")
	// Dummy ManifestLists
	mlUbuntu, _ = storage.NewManifestList("ubuntu", "latest", mfUbuntuAMD64, mfUbuntuPPC)
	mlQBench, _ = storage.NewManifestList("qnib/bench", "latest", mfQFeatGeneric, mfQFeatBroadwell, mfQFeatSkylake)
)

func getDummyManifestLists() []storage.ManifestList {
	return []storage.ManifestList{mlQBench, mlUbuntu}
}

func getMachineTypes() []storage.MachineType {
	return []storage.MachineType{
		mType0, mType1, mType2, mType3, mType4,
	}
}
