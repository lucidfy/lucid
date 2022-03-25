package php

import (
	"io/fs"
	"io/ioutil"
	"os"
	"reflect"
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

// InArray
// php equivalent for https://www.php.net/manual/en/function.in-array.php
func InArray(val interface{}, arr interface{}) int {
	values := reflect.ValueOf(arr)
	if reflect.TypeOf(arr).Kind() == reflect.Slice || values.Len() > 0 {
		for i := 0; i < values.Len(); i++ {
			if reflect.DeepEqual(val, values.Index(i).Interface()) {
				return i
			}
		}
	}
	return -1
}
