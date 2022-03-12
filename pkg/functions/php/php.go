package php

import (
	"io/fs"
	"io/ioutil"
	"os"
)

// Mkdir
// php equivalent https://www.php.net/manual/en/function.mkdir.php
func Mkdir(filename string, mode fs.FileMode, recursive bool) error {
	if recursive {
		return os.MkdirAll(filename, mode)
	}
	return os.Mkdir(filename, mode)
}

// FilePutContents
// php equivalent for https://www.php.net/manual/en/function.file-put-contents.php
func FilePutContents(filename string, data string, mode os.FileMode) error {
	return ioutil.WriteFile(filename, []byte(data), mode)
}
