package cah

import (
	"os"
	"path"
)

var AppDir string = path.Join(os.Getenv("GOPATH"), "src", "github.com", "j4rv", "cah", "cahApp")
var FrontendDir string = path.Join(AppDir, "frontend")
var PublicDir string = path.Join(FrontendDir, "build")
