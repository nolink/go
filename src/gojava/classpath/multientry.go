
package classpath

import (
	"strings"
	"errors"
)

type MultiEntry []Entry

func newMultiEntry(pathList string) MultiEntry{
	multiEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator){
		entry := newEntry(path)
		multiEntry = append(multiEntry, entry)
	}
	return multiEntry
}

func (self MultiEntry) ReadClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, from, err := entry.ReadClass(className)
		if err == nil {
			return data, from, err
		}
	}
	return nil, nil, errors.New("class not found: "+ className)
}

func (self MultiEntry) String() string{
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}