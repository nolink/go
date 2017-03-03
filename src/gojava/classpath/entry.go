
package classpath

import "os"
import "strings"

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	ReadClass(className string) ([]byte, Entry, error)
	String() string
}

func NewEntry(path string) {
	
}
