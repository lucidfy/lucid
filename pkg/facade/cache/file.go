package cache

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"os"

	"github.com/golang-module/carbon"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/crypt"
	"github.com/lucidfy/lucid/pkg/facade/path"
	"github.com/lucidfy/lucid/pkg/functions/php"
	"github.com/lucidfy/lucid/pkg/helpers"
)

type FileCache struct {
	ResponseWriter http.ResponseWriter
	HttpRequest    *http.Request
	FilePath       string
	FileMode       os.FileMode
}

func File(args ...interface{}) *FileCache {
	var fp string
	var fm fs.FileMode

	if len(args) > 0 {
		fp = args[0].(string)
	} else {
		fp = path.Load().StoragePath("cache.json")
	}

	if len(args) > 1 {
		fm = args[1].(fs.FileMode)
	} else {
		fm = 0744
	}

	return &FileCache{
		FilePath: fp,
		FileMode: fm,
	}
}

func (s *FileCache) initializeFile(filepath string) string {
	if !php.FileExists(filepath) {
		s.updateContent(map[string]interface{}{
			"initialized_at": carbon.Now().ToString(),
		})
	}
	return filepath
}

func (s *FileCache) updateContent(content interface{}) *errors.AppError {
	filepath := s.FilePath
	if app_err := php.FilePutContents(filepath, content, s.FileMode); app_err != nil {
		return app_err
	}
	return nil
}

func (s *FileCache) Put(name string, value interface{}) (bool, *errors.AppError) {
	filepath := s.initializeFile(s.FilePath)
	content := *php.JsonDecode(string(*php.FileGetContents(filepath)))
	value, err := crypt.Encrypt(helpers.Stringify(value))
	if err != nil {
		return false, err
	}
	content[name] = value
	if err := s.updateContent(content); err != nil {
		return false, err
	}
	return true, nil
}

func (s *FileCache) Get(name string) (interface{}, *errors.AppError) {
	filepath := s.initializeFile(s.FilePath)
	content := *php.JsonDecode(string(*php.FileGetContents(filepath)))

	if content[name] != nil {
		value, err := crypt.Decrypt(content[name].(string))
		if err != nil {
			return nil, err
		}
		return value, nil
	}

	return nil, errors.InternalServerError("s.SessionKey error", fmt.Errorf("session [%s] does not exists", name))
}

func (s *FileCache) GetAs(name string, m interface{}) {
	value, _ := s.Get(name)
	hs := helpers.Stringify(value)

	if json.Valid([]byte(hs)) {
		json.Unmarshal([]byte(hs), m)
	}
}

func (s *FileCache) Forget(name string) (interface{}, *errors.AppError) {
	filepath := s.initializeFile(s.FilePath)
	content := *php.JsonDecode(string(*php.FileGetContents(filepath)))
	delete(content, name)
	if err := s.updateContent(content); err != nil {
		return false, err
	}
	return true, nil
}
