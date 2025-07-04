package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func FindModuleRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("cannot determine current directory: %w", err)
	}

	dir =filepath.Clean(dir)
	for {
		if fi, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil && !fi.IsDir() {
			return dir, nil
		}
		d := filepath.Dir(dir)
		if d == dir {
			break
		}
		dir = d
	}

	return "", fmt.Errorf("cannot find module root")
}
