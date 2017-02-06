package compiler

import (
	"bytes"
	"io/ioutil"
	"text/template"

	"github.com/moul/funcmap"
)

func ParseFile(path string) (config string, err error) {
	var configSlice []byte
	var buf bytes.Buffer

	if configSlice, err = ioutil.ReadFile(path); err != nil {
		return
	}

	tmpl, err := template.New("confFile").Funcs(funcmap.FuncMap).Parse(string(configSlice))

	err = tmpl.Execute(&buf, nil)
	if err != nil {
		return
	}

	config = buf.String()
	return
}
