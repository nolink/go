
package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func NewWildEntry(path string) MultiEntry {
	baseDir := path[:len(path) - 1]
	multiEntry := []Entry{}
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := NewZipEntry(path)
			multiEntry = append(multiEntry, jarEntry)
		}
		return nil
	}
	filepath.Walk(baseDir, walkFn)
	return multiEntry
}