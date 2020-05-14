package utils

import (
	"encoding/json"
	"fmt"
)

type PackageJson struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func GetVersionJson() (string, error) {
	var pkg PackageJson
	err := json.Unmarshal(GetFileBytes("../../package.json"), &pkg)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("{ \"version\": \"%v\" }", pkg.Version), nil
}
