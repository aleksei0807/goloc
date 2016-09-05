package counter

import (
	"os"
	"fmt"
	"strings"
	"log"
	"github.com/kirillDanshin/myutils"
	"github.com/spf13/afero"
)

type Counter struct {
	fs *afero.Afero
	Exclude []string
}

func New() (*Counter, error) {
	return &Counter{
		fs: &afero.Afero{
			Fs: afero.NewOsFs(),
		},
	}, nil
}

func (c *Counter) isExcluded(path string) bool {
	for _, v := range c.Exclude {
		if v == path {
			return true
		}
	}
	return false
}

func (c *Counter) Local(path string) error {
	file, err := c.fs.Open(path)
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
			exPath := v.Name()
			if path != "./" {
				exPath = myutils.Concat(path, "/", exPath)
			}
			if c.isExcluded(exPath) {
				log.Printf("%s excluded", exPath)
				continue
			}
			if v.IsDir() {
				c.Local(
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
	if idx >= 0 && len(fi.Name()) > idx {
		fmt.Printf("File=[%#+v] FileExtension=[%#+v]\n", fi.Name(), fi.Name()[idx:])
	} else {
		return nil
	}

	return nil
}
