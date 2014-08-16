package template

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Template interface {
	String() string
	GetContent() (string, bool)
	appendContent(string)
	Do(string) error
}

type TemplateFunc struct {
	new func(opt string) Template
}

type TemplateAuthor struct {
	author string
}

func (this TemplateAuthor) String() string {
	return "author: " + this.author
}

func (this TemplateAuthor) GetContent() (string, bool) {
	return "", false
}

func (this *TemplateAuthor) appendContent(s string) {

}

func (this TemplateAuthor) Do(baseDir string) error {
	return nil
}

func TemplateAuthorFunc() TemplateFunc {
	return TemplateFunc{func(opt string) Template {
		return &TemplateAuthor{author: opt}
	}}
}

type TemplateDir struct {
	name       string
	permission uint32
}

func (this TemplateDir) String() string {
	return fmt.Sprintf("dir: %s(%04o)", this.name, this.permission)
}

func (this TemplateDir) GetContent() (string, bool) {
	return "", false
}

func (this *TemplateDir) appendContent(s string) {

}

func (this TemplateDir) Do(baseDir string) error {
	return os.MkdirAll(baseDir+"/"+this.name, os.FileMode(this.permission))
}

func TemplateDirFunc() TemplateFunc {
	return TemplateFunc{func(opt string) Template {
		a := strings.SplitN(opt, " ", 2)
		permission, err := strconv.ParseInt(a[1], 8, 32)
		if err != nil {
			panic("")
		}
		return &TemplateDir{name: a[0], permission: uint32(permission)}
	}}
}

type TemplateFile struct {
	name       string
	permission uint32
	content    string
}

func (this TemplateFile) String() string {
	return fmt.Sprintf("file: %s(%o)", this.name, this.permission)
}

func (this TemplateFile) GetContent() (string, bool) {
	if this.content == "" {
		return "", false
	} else {
		return this.content, true
	}
}

func (this *TemplateFile) appendContent(s string) {
	this.content += s + "\n"
}

func (this TemplateFile) Do(baseDir string) error {
	return ioutil.WriteFile(baseDir+"/"+this.name, []byte(this.content), os.FileMode(this.permission))
}

func TemplateFileFunc() TemplateFunc {
	return TemplateFunc{func(opt string) Template {
		a := strings.SplitN(opt, " ", 2)
		permission, err := strconv.ParseInt(a[1], 8, 32)
		if err != nil {
			panic("")
		}
		return &TemplateFile{name: a[0], permission: uint32(permission)}
	}}
}

func Parse(s string) ([]Template, error) {
	cmds := map[string]TemplateFunc{
		"author": TemplateAuthorFunc(),
		"dir":    TemplateDirFunc(),
		"file":   TemplateFileFunc(),
	}
	result := make([]Template, 0)

	patternDecl := regexp.MustCompile(`^([^\s]+)\s*:\s*(.+)\s*`)
	patternContent := regexp.MustCompile(`^([\s]+)(.*)`)
	var latest Template
	for _, line := range strings.Split(strings.Trim(s, " \t\r\n"), "\n") {
		if matches := patternDecl.FindAllStringSubmatch(line, -1); len(matches) == 1 && len(matches[0]) == 3 {
			cmd := matches[0][1]
			opt := matches[0][2]
			tmplFunc, ok := cmds[cmd]
			if !ok {
				return result, fmt.Errorf("Illegal command.")
			}
			tmpl := tmplFunc.new(opt)
			latest = tmpl
			result = append(result, tmpl)
		} else if matches := patternContent.FindAllStringSubmatch(line, -1); len(matches) == 1 && len(matches[0]) == 3 {
			if latest != nil {
				latest.appendContent(matches[0][2])
			}
		}

	}
	return result, nil
}

func Do(templates []Template, baseDir string) error {
	var err error
	for _, template := range templates {
		err = template.Do(baseDir)
		if err != nil {
			return err
		}
	}
	return nil
}
