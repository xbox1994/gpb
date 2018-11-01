package template

import (
	"fmt"
	"os"
	"strings"
)

const (
	Create = 1 << iota
	Modify = 1 << iota
	Del    = 1 << iota
)

type Template struct {
	Mode     int64
	FilePath string
	Content  string
	StdOut   func(template *Template, args ...string) error
}

var AvailableTemplates = []*Template{}

func getPackagePath(absPath string) string {
	goSrcPath := fmt.Sprintf("%s%s", os.Getenv("GoPath"), `src`)
	packagePath := strings.Replace(absPath, goSrcPath+"\\", "", -1)
	return packagePath
}
