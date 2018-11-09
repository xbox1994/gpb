package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Replace(old, new string) filepath.WalkFunc {
	visit := func(path string, fi os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if !!fi.IsDir() {
			return nil
		}

		matched, err := filepath.Match("*", fi.Name())

		if err != nil {
			panic(err)
			return err
		}

		if matched {
			read, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}
			fmt.Println(path)

			newContents := strings.Replace(string(read), old, new, -1)

			err = ioutil.WriteFile(path, []byte(newContents), 0)
			if err != nil {
				panic(err)
			}

		}

		return nil
	}
	return visit
}
