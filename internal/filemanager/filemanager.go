package filemanager

import (
	"path/filepath"
	"runtime"
)

func ProjectPath() (string, error) {
	_, callerFile, _, _ := runtime.Caller(0)
	path := filepath.Dir(callerFile)

	projectpath, err := filepath.Abs(path + "/../../")

	if err != nil {
		return "", err
	}

	return projectpath, nil
}

func PathTo(path string) string {
	basepath, err := ProjectPath()

	if err != nil {
		panic(err)
	}

	return basepath + path
}
