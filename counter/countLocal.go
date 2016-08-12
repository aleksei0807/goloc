package counter

import (
	"os"
	"fmt"
	"github.com/kirillDanshin/myutils"
	"strings"
)

func CountLocal(path string) error {
	file, err := os.Open(path)
	if os.IsNotExist(err) {
		return fmt.Errorf("Can't open path that doesn't exists")
	}
	if err != nil {
		return err
	}

	fileInfo, err := file.Stat()

	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		fi, err := file.Readdir(0)
		if err != nil {
			return err
		}
		for _, v := range fi {
			if v.IsDir() {
				CountLocal(
					myutils.Concat(path, "/", v.Name()),
				)
				continue
			}
			ProcessFile(v)
		}
	} else {
		ProcessFile(fileInfo)
	}

	return nil
}



func ProcessFile(fi os.FileInfo) error {
	idx := strings.LastIndex(fi.Name(), ".")
	if idx > 0 && len(fi.Name()) > idx {
		fmt.Printf("File=[%#+v] FileExtension=[%#+v]\n", fi.Name(), fi.Name()[idx:])
	} else {
		return nil
	}

	return nil
}
