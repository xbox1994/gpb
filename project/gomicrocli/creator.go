package gomicrocli

import (
	"archive/zip"
	"fmt"
	"github.com/rakyll/statik/fs"
	"github.com/xbox1994/gpb/common/util"
	"io/ioutil"
	"os"
	"path/filepath"
	_ "statik"
	"strconv"
	"strings"
	"unsafe"
)

func CreateProject(projectName string, outputPath string) (err error) {
	projectPath := projectName
	if outputPath != "" {
		projectPath = outputPath + "/" + projectName
	}

	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}
	err = fs.Walk(statikFS, "/", func(path string, info os.FileInfo, err error) error {
		ptrTof := unsafe.Pointer(&info)
		addressLength := 8
		if strconv.IntSize == 32 {
			addressLength = 4
		}
		ptrTof = unsafe.Pointer(uintptr(ptrTof) + uintptr(addressLength))
		ptrToy := (**zip.FileHeader)(ptrTof)
		_path := (*ptrToy).Name

		if info.IsDir() && info.Name() == "/" {
			return nil
		}
		if info.IsDir() && info.Name() != "/" {
			return filepath.SkipDir
		}

		read, err := fs.ReadFile(statikFS, "/"+_path)
		if err != nil {
			panic(err)
		}

		if strings.Contains(_path, "/") {
			os.MkdirAll(projectPath+"/"+_path[0:strings.LastIndex(_path, "/")], 0755)
		}
		err = ioutil.WriteFile(projectPath+"/"+_path, read, 0644)
		if err != nil {
			panic(err)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(projectPath+"/.gitignore", []byte(projectName), 0644)
	if err != nil {
		panic(err)
	}

	err = filepath.Walk(projectPath, util.Replace("microTemplate", projectName))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Create go micro project %s success", projectName)
	return err
}
