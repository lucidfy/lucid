package session

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-module/carbon"
	"github.com/lucidfy/lucid/pkg/facade/cookie"
	"github.com/lucidfy/lucid/pkg/facade/crypt"
	"github.com/lucidfy/lucid/pkg/facade/path"
	"github.com/lucidfy/lucid/pkg/functions/php"
)

type FileSession struct {
	SessionKey     interface{}
	ResponseWriter http.ResponseWriter
	HttpRequest    *http.Request
}

func File(w http.ResponseWriter, r *http.Request) *FileSession {
	coo := cookie.New(w, r)
	sessionKey, err := coo.Get(os.Getenv("SESSION_NAME"))
	if err != nil && errors.Is(err, http.ErrNoCookie) {
		return &FileSession{}
	}
	s := FileSession{
		SessionKey: sessionKey,
	}
	return &s
}

func (s *FileSession) getSessionFile() string {
	return path.Load().StoragePath("framework/sessions/" + s.SessionKey.(string))
}

func (s *FileSession) initializeFile(filepath string) string {
	if !php.FileExists(filepath) {
		s.updateContent(map[string]interface{}{
			"created_at": carbon.Now().ToString(),
		})
	}
	return filepath
}

func (s *FileSession) updateContent(content interface{}) error {
	filepath := s.getSessionFile()
	if err := php.FilePutContents(filepath, content, 0644); err != nil {
		return err
	}
	return nil
}

func (s *FileSession) Set(name string, value interface{}) (bool, error) {
	if s.SessionKey == nil {
		return false, fmt.Errorf("session [%s] does not exists", name)
	}

	filepath := s.initializeFile(s.getSessionFile())
	content := *php.JsonDecode(string(*php.FileGetContents(filepath)))
	value, err := crypt.Encrypt(value)
	if err != nil {
		return false, err
	}
	content[name] = value
	if err := s.updateContent(content); err != nil {
		return false, err
	}
	return true, nil
}

func (s *FileSession) Get(name string) (interface{}, error) {
	if s.SessionKey == nil {
		return nil, fmt.Errorf("session [%s] does not exists", name)
	}

	filepath := s.initializeFile(s.getSessionFile())
	content := *php.JsonDecode(string(*php.FileGetContents(filepath)))

	if content[name] != nil {
		value, err := crypt.Decrypt(content[name].(string))
		if err != nil {
			return nil, err
		}
		return value, nil
	}

	return nil, fmt.Errorf("session [%s] does not exists", name)
}

func (s *FileSession) Destroy(name string) (interface{}, error) {
	if s.SessionKey == nil {
		return nil, fmt.Errorf("session [%s] does not exists", name)
	}

	filepath := s.initializeFile(s.getSessionFile())
	content := *php.JsonDecode(string(*php.FileGetContents(filepath)))
	delete(content, name)
	if err := s.updateContent(content); err != nil {
		return false, err
	}
	return true, nil
}

func (s *FileSession) SetFlash(name string, value interface{}) {
	name = "flash-" + name
	s.Set(name, value)
}

func (s *FileSession) GetFlash(name string) interface{} {
	name = "flash-" + name
	value, err := s.Get(name)
	if err != nil || value == nil {
		return nil
	}
	s.Destroy(name)
	return value
}

// SetFlashMap sets a session flash based on json format
// make sure the values you're providing is set as map[string]interface{}
// therefore, we can stringify it into json format
func (s *FileSession) SetFlashMap(name string, value interface{}) {
	j, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	s.SetFlash(name, string(j))
}

// GetFlashMap this pulls a session flash from SetFlashMap, in which
// it will reverse the json into a map
func (s *FileSession) GetFlashMap(name string) *map[string]interface{} {
	ret := &map[string]interface{}{}
	flash := s.GetFlash(name)
	if flash != nil {
		// flashStr := (*flash.(*interface{})).(string)
		flashStr := flash.(string)
		json.Unmarshal([]byte(flashStr), ret)
	}
	return ret
}
