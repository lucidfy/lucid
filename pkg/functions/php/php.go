package php

import (
	"encoding/json"
	"errors"
	"io/fs"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/daison12006013/lucid/pkg/facade/logger"
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
func FilePutContents(filename string, data interface{}, mode os.FileMode) error {
	var content string
	switch reflect.TypeOf(data).Kind() {
	case reflect.Map:
		content = string(JsonEncode(data))

	case reflect.String:
		content = data.(string)
	}

	return ioutil.WriteFile(filename, []byte(content), mode)
}

// FileGetContents
// php equivalent for https://www.php.net/manual/en/function.file-get-contents.php
func FileGetContents(filename string) *[]byte {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.Error("Error getting file contents!", err)
		return nil
	}
	return &content
}

// FileExists
// php equivalent for https://www.php.net/manual/en/function.file-exists.php
func FileExists(filename string) bool {
	if _, err := os.Stat(filename); errors.Is(err, fs.ErrNotExist) {
		return false
	}
	return true
}

// JsonEncode
// php equivalent for https://www.php.net/manual/en/function.json-encode.php
func JsonEncode(value interface{}) []byte {
	j, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return j
}

// JsonDecode
// php equivalent for https://www.php.net/manual/en/function.json-decode.php
func JsonDecode(j interface{}) *map[string]interface{} {
	data := &map[string]interface{}{}
	json.Unmarshal([]byte(j.(string)), data)
	return data
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
