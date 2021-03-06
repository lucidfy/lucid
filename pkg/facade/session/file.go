package session

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/golang-module/carbon"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/crypt"
	"github.com/lucidfy/lucid/pkg/facade/path"
	"github.com/lucidfy/lucid/pkg/functions/php"
	"github.com/lucidfy/lucid/pkg/helpers"
)

type FileSession struct {
	SessionKey interface{}
	FileMode   os.FileMode
}

func File(sessionKey string) *FileSession {
	s := FileSession{
		SessionKey: sessionKey,
		FileMode:   0644,
	}
	return &s
}

func (s *FileSession) getFile() string {
	return path.Load().SessionPath(s.SessionKey.(string))
}

func (s *FileSession) initializeFile(filepath string) string {
	if !php.FileExists(filepath) {
		s.updateContent(map[string]interface{}{
			"created_at": carbon.Now().ToString(),
		})
	}
	return filepath
}

func (s *FileSession) updateContent(content interface{}) *errors.AppError {
	filepath := s.getFile()
	if app_err := php.FilePutContents(filepath, content, s.FileMode); app_err != nil {
		return app_err
	}
	return nil
}

func (s *FileSession) Put(name string, value interface{}) (bool, *errors.AppError) {
	if s.SessionKey.(string) == "" {
		return false, errors.InternalServerError("s.SessionKey error", fmt.Errorf("session [%s] does not exists", name))
	}

	filepath := s.initializeFile(s.getFile())
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

func (s *FileSession) Get(name string) (interface{}, *errors.AppError) {
	if s.SessionKey.(string) == "" {
		return nil, errors.InternalServerError("s.SessionKey error", fmt.Errorf("session [%s] does not exists", name))
	}

	filepath := s.getFile()

	// check if the file exists, or else return nil
	// we wont be initializing the file as it can flood
	// our session files, session is meant for users
	// that requires stateful data based on their session key
	if !php.FileExists(filepath) {
		return nil, nil
	}

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

func (s *FileSession) Flush(name string) (interface{}, *errors.AppError) {
	if s.SessionKey.(string) == "" {
		return nil, errors.InternalServerError("s.SessionKey error", fmt.Errorf("session [%s] does not exists", name))
	}

	filepath := s.initializeFile(s.getFile())
	content := *php.JsonDecode(string(*php.FileGetContents(filepath)))
	delete(content, name)
	if err := s.updateContent(content); err != nil {
		return false, err
	}
	return true, nil
}

func (s *FileSession) PutFlash(name string, value interface{}) {
	name = "flash-" + name
	s.Put(name, value)
}

func (s *FileSession) GetFlash(name string) interface{} {
	name = "flash-" + name
	value, err := s.Get(name)
	if err != nil || value == nil {
		return nil
	}
	s.Flush(name)
	return value
}

// PutFlashMap sets a session flash based on json format
// make sure the values you're providing is set as map[string]interface{}
// therefore, we can stringify it into json format
func (s *FileSession) PutFlashMap(name string, value interface{}) {
	j, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	s.PutFlash(name, string(j))
}

// GetFlashMap this pulls a session flash from PutFlashMap, in which
// it will reverse the json into a map
func (s *FileSession) GetFlashMap(name string) *map[string]interface{} {
	ret := &map[string]interface{}{}
	flash := s.GetFlash(name)
	if flash != nil {
		flashStr := flash.(string)
		json.Unmarshal([]byte(flashStr), ret)
	}
	return ret
}
