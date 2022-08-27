package php

import (
	"encoding/json"
	e "errors"
	"io/fs"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/logger"
)

// Mkdir
// php equivalent https://www.php.net/manual/en/function.mkdir.php
func Mkdir(filename string, mode fs.FileMode, recursive bool) *errors.AppError {
	if recursive {
		err := os.MkdirAll(filename, mode)
		return errors.InternalServerError("os.MkdirAll() error", err)
	}
	err := os.Mkdir(filename, mode)
	return errors.InternalServerError("os.Mkdir() error", err)
}

// FilePutContents
// php equivalent for https://www.php.net/manual/en/function.file-put-contents.php
func FilePutContents(filename string, data interface{}, mode os.FileMode) *errors.AppError {
	var content string
	switch reflect.TypeOf(data).Kind() {
	case reflect.Map:
		content = string(JsonEncode(data))

	case reflect.String:
		content = data.(string)
	}

	err := ioutil.WriteFile(filename, []byte(content), mode)
	return errors.InternalServerError("ioutil.WriteFile() error", err)
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
	if _, err := os.Stat(filename); e.Is(err, fs.ErrNotExist) {
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

// Strtr
// php equivalent for https://www.php.net/manual/en/function.strtr.php
func Strtr(str string, replace map[string]string) string {
	if len(replace) == 0 || len(str) == 0 {
		return str
	}
	for old, new := range replace {
		str = strings.ReplaceAll(str, old, new)
	}
	return str
}
