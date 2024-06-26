package profiles

import (
	"os"

	fuzz "github.com/AdaLogics/go-fuzz-headers"
)

// FuzzProfilesValidate fuzzes the ProfilesValidate function.
func FuzzProfilesValidate(data []byte) int {
	_ = Validate(data)
	return 1
}

// FuzzRenderProto fuzzes the RenderProto function.
func FuzzRenderProto(data []byte) int {
	f := fuzz.NewConsumer(data)
	protodata, err := f.GetBytes()
	if err != nil {
		return 0
	}
	namespace, err := f.GetString()
	if err != nil {
		return 0
	}
	name, err := f.GetString()
	if err != nil {
		return 0
	}
	clusterDomain, err := f.GetString()
	if err != nil {
		return 0
	}
	protofile, err := os.Create("protofile")
	if err != nil {
		return 0
	}
	defer protofile.Close()
	defer os.Remove(protofile.Name())

	_, err = protofile.Write(protodata)
	if err != nil {
		return 0
	}
	_, err = RenderProto(protofile.Name(), namespace, name, clusterDomain)
	if err != nil {
		return 0
	}
	return 1
}
