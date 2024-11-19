package goutil

import (
	"embed"
	"os"
	"path"
	"text/template"
)

func IsDir(dir string, filename string) bool {
	stat, err := os.Lstat(path.Join(dir, filename))

	return err == nil && stat.IsDir()
}

func IsFile(dir string, filename string) bool {
	stat, err := os.Lstat(path.Join(dir, filename))

	return err == nil && stat.Mode().IsRegular()
}

func MustReadFile(fs embed.FS, filename string) string {
	contents, err := fs.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(contents)
}

func MustParseTemplate(fs embed.FS, filename string) *template.Template {
	tpl, err := template.New("").Funcs(template.FuncMap{
		"CamelToSnake": CamelToSnake,
	}).Parse(MustReadFile(fs, filename))
	if err != nil {
		panic(err)
	}

	return tpl
}
