package current_path

import "os"

type BaseWidget struct {
}

func (w *BaseWidget) Init() {

}

func (w *BaseWidget) Run() {

}

func (w *BaseWidget) Rend(view string, data interface{}) {

}

func (w *BaseWidget) ViewPath() string {
	/*
		_, filename, line, _ := runtime.Caller(0)
		_ = line
		return path.Dir(filename)
	*/
	var currentDir string
	currentDir, err := os.Getwd()
	_ = err
	return currentDir
}
