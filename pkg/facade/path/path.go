package path

import (
	"os"
	"path/filepath"
	"runtime"
)

type PathStruct struct {
	BASE_PATH        string
	CONSOLE_PATH     string
	HANDLERS_PATH    string
	MIDDLEWARES_PATH string
	MODELS_PATH      string
	DATABASE_PATH    string
	TRANSLATION_PATH string
	VIEW_PATH        string
	ROUTES_PATH      string
	STORAGE_PATH     string
}

func Load() *PathStruct {
	basePath, err := RootPath()
	if err != nil {
		panic(err)
	}

	p := &PathStruct{
		BASE_PATH:        *basePath,
		CONSOLE_PATH:     PathTo(os.Getenv("CONSOLE_PATH")),
		HANDLERS_PATH:    PathTo(os.Getenv("HANDLERS_PATH")),
		MIDDLEWARES_PATH: PathTo(os.Getenv("MIDDLEWARES_PATH")),
		MODELS_PATH:      PathTo(os.Getenv("MODELS_PATH")),
		DATABASE_PATH:    PathTo(os.Getenv("DATABASE_PATH")),
		TRANSLATION_PATH: PathTo(os.Getenv("TRANSLATION_PATH")),
		VIEW_PATH:        PathTo(os.Getenv("VIEW_PATH")),
		ROUTES_PATH:      PathTo(os.Getenv("ROUTES_PATH")),
		STORAGE_PATH:     PathTo(os.Getenv("STORAGE_PATH")),
	}
	return p
}

func (p *PathStruct) BasePath(str string) string {
	return append(p.BASE_PATH, str)
}
func (p *PathStruct) ConsolePath(str string) string {
	return append(p.CONSOLE_PATH, str)
}
func (p *PathStruct) HandlersPath(str string) string {
	return append(p.HANDLERS_PATH, str)
}
func (p *PathStruct) MiddlewaresPath(str string) string {
	return append(p.MIDDLEWARES_PATH, str)
}
func (p *PathStruct) ModelsPath(str string) string {
	return append(p.MODELS_PATH, str)
}
func (p *PathStruct) DatabasePath(str string) string {
	return append(p.DATABASE_PATH, str)
}
func (p *PathStruct) TranslationPath(str string) string {
	return append(p.TRANSLATION_PATH, str)
}
func (p *PathStruct) ViewPath(str string) string {
	return append(p.VIEW_PATH, str)
}
func (p *PathStruct) RoutesPath(str string) string {
	return append(p.ROUTES_PATH, str)
}
func (p *PathStruct) StoragePath(str string) string {
	return append(p.STORAGE_PATH, str)
}

func append(path string, str string) string {
	if str != "" {
		str = "/" + str
	}
	return path + str
}

func RootPath() (*string, error) {
	if len(os.Getenv("LUCID_ROOT")) != 0 {
		projectpath := os.Getenv("LUCID_ROOT")
		return &projectpath, nil
	}

	// if LUCID_ROOT isn't present
	// let's try to look up the runtime caller
	_, callerFile, _, _ := runtime.Caller(0)
	path := filepath.Dir(callerFile)
	projectpath, err := filepath.Abs(path + "/../../../")

	if err != nil {
		return nil, err
	}

	return &projectpath, nil
}

func PathTo(path string) string {
	basepath, err := RootPath()
	if err != nil {
		panic(err)
	}
	return *basepath + path
}
