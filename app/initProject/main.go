package initProject

import (
	"gin-cli/app/initProject/template"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

func GetGoVersionSkipMinor() string {
	strArray := strings.Split(runtime.Version()[2:], `.`)
	return strArray[0] + `.` + strArray[1]
}

func createFile(filePath, template string) (err error) {
	projectName := strings.Split(filePath, "/")[0]
	if err = os.MkdirAll(path.Dir(filePath), 0777); err != nil {
		return
	}

	var content string
	content = strings.Replace(template, "{{projectName}}", projectName, -1)
	content = strings.Replace(content, "{{goVersion}}", GetGoVersionSkipMinor(), -1)

	var f *os.File
	if f, err = os.Create(filePath); err != nil {
		return
	}
	defer f.Close()

	_, err = f.WriteString(content)
	return
}

type Template struct {
	FilePath string
	Template string
}

var templateList []*Template

func bind(filePath, template string) {
	templateList = append(templateList, &Template{
		FilePath: filePath,
		Template: template,
	})
}

func NewProject(projectName string) (err error) {
	//main
	bind(path.Join(projectName, "main.go"), template.MainTmp)
	bind(path.Join(projectName, "config.yaml"), template.ConfigTmp)
	bind(path.Join(projectName, "README.md"), template.ReadMeTmp)
	bind(path.Join(projectName, ".gitignore"), template.GitignoreTmp)
	bind(path.Join(projectName, "go.mod"), template.GoModTmp)
	//consts
	bind(path.Join(projectName, "consts", "consts.go"), template.ConstsTmp)
	//controller
	bind(path.Join(projectName, "controller", "base.go"), template.ControllerBaseTmp)
	bind(path.Join(projectName, "controller", "user", "login.go"), template.ControllerUserLoginTmp)
	//dao
	bind(path.Join(projectName, "dao", "dao.go"), template.DaoTmp)
	bind(path.Join(projectName, "dao", "internal", "base.go"), template.DaoInternalBaseTmp)
	bind(path.Join(projectName, "dao", "internal", "user.go"), template.DaoInternalUserTmp)
	//docs
	bind(path.Join(projectName, "docs", "docs.go"), template.DocsTmp)
	bind(path.Join(projectName, "docs", "swagger.json"), template.DocsSwaggerJson)
	bind(path.Join(projectName, "docs", "swagger.yaml"), template.DocsSwaggerYaml)
	//global
	bind(path.Join(projectName, "global", "global.go"), template.GlobalTmp)
	bind(path.Join(projectName, "global", "config.go"), template.GlobalConfigTmp)
	bind(path.Join(projectName, "global", "logger.go"), template.GlobalLoggerTmp)
	bind(path.Join(projectName, "global", "mysql.go"), template.GlobalMysqlTmp)
	bind(path.Join(projectName, "global", "redis.go"), template.GlobalRedisTmp)
	bind(path.Join(projectName, "global", "mongo.go"), template.GlobalMongoTmp)
	//middleware
	bind(path.Join(projectName, "middleware", "cors.go"), template.MiddlewareCorsTmp)
	bind(path.Join(projectName, "middleware", "auth.go"), template.MiddlewareAuthTmp)
	//model
	bind(path.Join(projectName, "model", "api", "user.go"), template.ModelApiUserTmp)
	bind(path.Join(projectName, "model", "entity", "user.go"), template.ModelEntityUserTmp)
	bind(path.Join(projectName, "model", "user.go"), template.ModelUserTmp)
	//pkg
	bind(path.Join(projectName, "pkg", "jwt.go"), template.PkgJwtTmp)
	//router
	bind(path.Join(projectName, "router", "router.go"), template.RouterTmp)
	//service
	bind(path.Join(projectName, "service", "user.go"), template.ServiceUserTmp)
	//test
	bind(path.Join(projectName, "test", "service", "config.yaml"), template.ConfigTmp)
	bind(path.Join(projectName, "test", "service", "user.go"), template.TestServiceUserTestTmp)
	//utils
	bind(path.Join(projectName, "utils", "utils.go"), template.ConfigTmp)

	//执行创建
	for _, temp := range templateList {
		if err = createFile(temp.FilePath, temp.Template); err != nil {
			log.Panicln(err)
			return
		}
	}
	return
}
