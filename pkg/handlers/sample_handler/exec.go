package sample_handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/lucidfy/lucid/pkg/engines"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/routes"
)

var ExecSampleRoute = routes.Routing{
	Path:                 "/exec-sample",
	Name:                 "exec-sample",
	Method:               routes.Method{"GET", "POST", "DELETE", "PUT", "PATCH"},
	Handler:              ExecSample,
	WithGlobalMiddleware: false,
}

// ExecSample is a sample way to run a command via handler
// as an example below, it will execute a php file containing
// all the helpful variables.
func ExecSample(T engines.EngineContract) *errors.AppError {
	engine := T.(engines.NetHttpEngine)
	w := engine.ResponseWriter
	r := engine.HttpRequest

	jsonErr := func(err error) *errors.AppError {
		return &errors.AppError{
			Code:    http.StatusInternalServerError,
			Message: "Json Marshal Error: " + err.Error(),
			Error:   err,
		}
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return &errors.AppError{
			Code:    http.StatusInternalServerError,
			Message: "ioutil.ReadAll(): " + err.Error(),
			Error:   err,
		}
	}

	data := map[string]interface{}{
		"vars":  mux.Vars(r),
		"form":  r.Form,
		"query": r.URL.Query(),
		"HttpRequest": map[string]interface{}{
			"Method": r.Method,
			"URL": map[string]interface{}{
				"Scheme":      r.URL.Scheme,
				"Opaque":      r.URL.Opaque,
				"Host":        r.URL.Host,
				"Path":        r.URL.Path,
				"RawPath":     r.URL.RawPath,
				"ForceQuery":  r.URL.ForceQuery,
				"RawQuery":    r.URL.RawQuery,
				"Fragment":    r.URL.Fragment,
				"RawFragment": r.URL.RawFragment,
			},
			"Proto":            r.Proto,
			"ProtoMajor":       r.ProtoMajor,
			"ProtoMinor":       r.ProtoMinor,
			"Header":           r.Header,
			"Body":             body,
			"ContentLength":    r.ContentLength,
			"TransferEncoding": r.TransferEncoding,
			"Host":             r.Host,
			"Form":             r.Form,
			"PostForm":         r.PostForm,
			"RemoteAddr":       r.RemoteAddr,
			"RequestURI":       r.RequestURI,
		},
	}

	inputs, err := json.Marshal(data)
	if err != nil {
		return jsonErr(err)
	}

	app := "php"
	args := []string{
		"-r",
		fmt.Sprintf("echo \"<pre>\"; var_dump(json_decode('%s'));", string(inputs)),
	}

	// similar above, it's just that you need to create a file
	// and capture the variable using $argv
	// args := []string{"index.php", string(inputs)}

	stdout, exitCode, err := RunCommand(app, args...)
	if err != nil && exitCode == 999 {
		return &errors.AppError{
			Code:    http.StatusInternalServerError,
			Message: "Internal Error: " + err.Error(),
			Error:   err,
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, string(stdout))
	return nil
}

func RunCommand(name string, args ...string) (stdout string, exitCode int, err error) {
	log.Println("run command:", name, args)
	var outbuf, errbuf bytes.Buffer
	cmd := exec.Command(name, args...)
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	err = cmd.Run()
	stdout = outbuf.String()
	stderr := errbuf.String()

	if err != nil {
		// try to get the exit code
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
		} else {
			// This will happen (in OSX) if `name` is not available in $PATH,
			// in this situation, exit code could not be get, and stderr will be
			// empty string very likely, so we use the default fail code, and format err
			// to string and set to stderr
			log.Printf("Could not get exit code for failed program: %v, %v", name, args)
			exitCode = 999 // our custom exit code
			if stderr == "" {
				stderr = err.Error()
			}
		}
	} else {
		// success, exitCode should be 0 if go is ok
		ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
		exitCode = ws.ExitStatus()
	}
	log.Printf("command result, stdout: %v, stderr: %v, exitCode: %d", stdout, stderr, exitCode)
	return
}
