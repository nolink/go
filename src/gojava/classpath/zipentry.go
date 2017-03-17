
package classpath

import "io/ioutil"
import "path/filepath"
import "errors"
import "archive/zip"

type ZipEntry struct{
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (self *ZipEntry) String() string{
	return self.absPath
}

func (self *ZipEntry) ReadClass(className string) ([]byte, Entry, error) {
	zipReader, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer zipReader.Close()
	for _, f := range zipReader.File {
		if f.Name == className {
			fstream, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer fstream.Close()
			data, err := ioutil.ReadAll(fstream)
			if err != nil {
				return nil, nil, err
			}
			return data, self, err
		}
	}
	return nil, nil, errors.New("class not found: "+ className)
}