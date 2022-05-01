package php

import (
	"testing"

	"github.com/daison12006013/lucid/pkg/env"
	"github.com/daison12006013/lucid/pkg/facade/path"
)

func init() {
	env.LoadEnv()
}

func TestFileGetContents(t *testing.T) {
	content := FileGetContents(path.Load().BasePath("stubs/handler/handler.stub"))
	if len(content.(string)) == 0 {
		t.Errorf("FileGetContents is not working!")
	}
}

func TestFilePutContents(t *testing.T) {
	if err := FilePutContents("/tmp/lucid-test.log", "Hello World!", 0775); err != nil {
		t.Errorf("Error FilePutContents using map: %s", err)
	}

	m := map[string]interface{}{
		"TestFilePutContents": "It worked!",
		"attributes": map[string]interface{}{
			"age": 31,
			"sex": "male",
		},
	}

	if err := FilePutContents("/tmp/lucid-test.log", m, 0775); err != nil {
		t.Errorf("Error FilePutContents using map: %s", err)
	}
}

func TestFileExists(t *testing.T) {
	exists := FileExists(path.Load().BasePath("stubs/handler/handler.stub"))
	if !exists {
		t.Errorf("FileExists cant find the file!")
	}
}

func TestJsonDecode(t *testing.T) {
	ret := *JsonDecode("{\"name\": \"John Doe\", \"attributes\": {\"age\": 31}}")
	if ret["name"] == nil {
		t.Errorf("JsonDecode is not working!")
	}
	if ret["attributes"].(map[string]interface{})["age"] == nil {
		t.Errorf("JsonDecode is not working!")
	}
}
