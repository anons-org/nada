package Utils

import (
	"path/filepath"
	"runtime"
)

var (
	_, Path, _, _ = runtime.Caller(0)

	// Root folder of this project
	RootPath = filepath.Join(filepath.Dir(Path), "../..")
)
