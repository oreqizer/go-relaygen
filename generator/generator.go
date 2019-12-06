package generator

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"

	"github.com/pkg/errors"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/imports"
)

type templateData struct {
	Name    string
	Package string
	Import  string
	ID      string
}

func Generate(name string, id string, wd string) error {
	genPkg := getPackage(wd)
	if genPkg == nil {
		return fmt.Errorf("unable to find package info for " + wd)
	}

	data := templateData{
		Name:    name,
		Package: genPkg.Name,
		Import:  genPkg.PkgPath,
		ID:      id,
	}

	if err := writeTemplate(source, filepath.Join(wd, strings.ToLower(data.Name)+"_relaygen.go"), data); err != nil {
		return err
	}

	if err := writeTemplate(test, filepath.Join(wd, strings.ToLower(data.Name)+"_relaygen_test.go"), data); err != nil {
		return err
	}

	return nil
}

// Utils
// -----

func lcFirst(s string) string {
	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}

func getPackage(dir string) *packages.Package {
	p, _ := packages.Load(&packages.Config{
		Dir: dir,
	}, ".")

	if len(p) != 1 {
		return nil
	}

	return p[0]
}

func writeTemplate(tpl *template.Template, filepath string, data templateData) error {
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		return errors.Wrap(err, "generating code")
	}

	src, err := imports.Process(filepath, buf.Bytes(), nil)
	if err != nil {
		return errors.Wrap(err, "unable to gofmt")
	}

	if err := ioutil.WriteFile(filepath, src, 0644); err != nil {
		return errors.Wrap(err, "writing output")
	}

	return nil
}
