package Helper

import (
	"bytes"
	"html/template"
	"log"
	"path"
	"runtime"
)

func StaticPath(Staticpath string) string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(filename), Staticpath)
}

func GetTemplateEmail(TemplateFile string, templateData map[string]interface{}) (data string, err error) {
	path := StaticPath("")
	t, err := template.ParseFiles(path + TemplateFile)
	if err != nil {
		log.Println(err.Error())
		return
	}
	var tpl bytes.Buffer
	err = t.Execute(&tpl, templateData)
	return tpl.String(), err
}
