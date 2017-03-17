

package classpath

import "path/filepath"
import "io/ioutil"

type DirEntry struct{
	absPath string
}

func newDirEntry(path string) *DirEntry{
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

func (self *DirEntry) ReadClass(className string) ([]byte, Entry, error) {
	filaName := filepath.Join(self.absPath, className)
	data, err := ioutil.ReadFile(filaName)
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absPath
}