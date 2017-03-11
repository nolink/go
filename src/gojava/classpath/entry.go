
package classpath

import(
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	ReadClass(className string) ([]byte, Entry, error)
	String() string
}

func NewEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator){
		return NewMultiEntry(path)
	}
	if strings.HasSuffix(path, "*"){
		return NewWildEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") || 
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
			return NewZipEntry(path)
	}
	return NewDirEntry(path)
}
