package utils

import (
	"encoding/json"
	"fmt"
)

type PackageJson struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func GetVersionJson() string {
	var pkg PackageJson
	err := json.Unmarshal(GetFileBytes("../../package.json"), &pkg)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("{ \"version\": \"%v\" }", pkg.Version)
}

func GetVersionMut(version *string) string {
	if *version == DEFAULT_VERSION {
		*version = GetVersionJson()
	}
	return *version
}
