package sample_handler

import (
	"net/http"

	"github.com/lucidfy/lucid/pkg/engines"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/logger"
	"github.com/lucidfy/lucid/pkg/facade/routes"
	"github.com/lucidfy/lucid/pkg/helpers"
	"github.com/lucidfy/lucid/pkg/storage"
)

var StorageRoute = routes.Routing{
	Path:    "/samples/storage",
	Name:    "",
	Method:  routes.Method{"POST"},
	Handler: sample_file_storage,
}

func sample_file_storage(T engines.EngineContract) *errors.AppError {
	engine := T.(engines.NetHttpEngine)
	req := engine.Request
	res := engine.Response

	files, app_err := req.GetFiles()

	if app_err != nil {
		return res.Json(helpers.MP{
			"error": app_err.Error,
		}, http.StatusOK)
	}

	images := files["files"]
	logger.Info("Image Length", len(images))

	// initialize local storage
	store := storage.NewLocalStorage()

	for _, image := range images {
		app_err := store.Put(image.Filename, image)
		if app_err != nil {
			continue
		}

		go logger.Info("Storage Size: ", store.Size(image.Filename))
		go logger.Info("File Exist: ", store.Exists(image.Filename))
		go logger.Info("File Missing: ", store.Missing(image.Filename))
		store.Delete(image.Filename)
	}

	if app_err != nil {
		return res.Json(helpers.MP{
			"error": app_err.Error,
		}, http.StatusOK)
	}

	// prepare the data
	data := helpers.MP{
		"file": len(images),
	}

	return res.Json(data, http.StatusOK)
}
