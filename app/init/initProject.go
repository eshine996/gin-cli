package init

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
)

func GetGoVersionSkipMinor() string {
	strArray := strings.Split(runtime.Version()[2:], `.`)
	return strArray[0] + `.` + strArray[1]
}

func NewProject(projectName string) (err error) {
	var bfileMap []byte
	if bfileMap, err = base64.StdEncoding.DecodeString(projectTemplate); err != nil {
		return
	}

	fileMap := make(map[string][]byte)
	if err = json.Unmarshal(bfileMap, &fileMap); err != nil {
		return
	}

	for fPath, content := range fileMap {
		if err = os.MkdirAll(path.Join(projectName, path.Dir(fPath)), 0777); err != nil {
			return
		}

		content = bytes.Replace(content, []byte("gin-scaffolding"), []byte(projectName), -1)
		if path.Base(fPath) == "go.mod" {
			content = bytes.Replace(content, []byte("go 1.18"), []byte(fmt.Sprintf("go %s", GetGoVersionSkipMinor())), -1)
		}

		f, _ := os.Create(path.Join(projectName, fPath))
		if _, err = f.Write(content); err != nil {
			return
		}

		if err = f.Close(); err != nil {
			return
		}
	}
	return
}
