

package classpath

import (
	"os" 
	"path/filepath"
)

type Classpath struct{
	bootClasspath Entry
	extClasspath Entry
	userClasspath Entry
}

func Parse(jreOption string, cpOption string) *Classpath{
	cp := &Classpath{}
	cp.ParseBootAndExtClasspath(jreOption)
	cp.ParseUserClasspath(cpOption)
	return cp
}

func (self *Classpath) ParseBootAndExtClasspath(jreOption string) {
	jreDir := GetJreDir(jreOption)
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = NewWildEntry(jreLibPath)
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = NewWildEntry(jreExtPath)
}

func (self *Classpath) ParseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = NewEntry(cpOption)
}

func GetJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre path")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err){
			return false
		}
	}
	return true
}

func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.ReadClass(className); err == nil {
		return data, entry ,err
	}
	if data, entry, err := self.extClasspath.ReadClass(className); err == nil {
		return data, entry, err
	}
	return self.userClasspath.ReadClass(className)
}

func (self *Classpath) String() string{
	return self.userClasspath.String()
}


