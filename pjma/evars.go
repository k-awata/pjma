package pjma

import (
	"bytes"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const PROJECTS_DIR = "projects_dir"
const CUSTOM_EVARS = "custom_evars.bat"

type Evars struct {
	pjdir string
	refpj []Project
	jenv  map[string][]string
	acmd  string
}

// NewProject returns evars
func NewEvars(pjdir string) *Evars {
	return &Evars{
		pjdir: os.ExpandEnv(pjdir),
		jenv:  map[string][]string{},
	}
}

// AddReferProject appends a new reference project
func (e *Evars) AddReferProject(p Project) {
	e.refpj = append(e.refpj, p)
}

// AddReferProject appends new reference project directories
func (e *Evars) AddReferProjectDirs(dirs []string) error {
	for _, v := range dirs {
		pj, err := NewProject("", v)
		if err != nil {
			return err
		}
		e.AddReferProject(*pj)
	}
	return nil
}

// AddJoinEnv appends environment variables to join to existing variables
func (e *Evars) AddJoinEnv(m map[string][]string) {
	for k, v := range m {
		e.jenv[k] = append(e.jenv[k], v...)
	}
}

// AddAfterCmd appends commands
func (e *Evars) AddAfterCmd(s string) {
	if s != "" {
		e.acmd = e.acmd + s + "\n"
	}
}

// Save saves custom_evars.bat to projects_dir
func (e *Evars) Save() error {
	if e.pjdir == "" {
		return errors.New("projects_dir is not specified")
	}
	// Make directory if it doesn't exist
	if fs, err := os.Stat(e.pjdir); err != nil {
		if err := os.MkdirAll(e.pjdir, os.ModePerm); err != nil {
			return err
		}
	} else if !fs.IsDir() {
		return errors.New("projects_dir is not a directory")
	}
	// Make commands
	l, err := e.makeProjectsDir()
	if err != nil {
		return err
	}
	f := e.makeReferProjects()
	j := e.makeJoinEnv()
	a := e.acmd
	// Create custom evars
	file, err := os.Create(filepath.Join(e.pjdir, CUSTOM_EVARS))
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.WriteString(strings.ReplaceAll(f+l+j+a, "\n", "\r\n")); err != nil {
		return err
	}
	return nil
}

func (e *Evars) makeReferProjects() string {
	var buf bytes.Buffer
	for _, v := range e.refpj {
		buf.WriteString(v.DumpEvars() + "\n")
	}
	return buf.String()
}

func (e *Evars) makeProjectsDir() (string, error) {
	var buf bytes.Buffer
	uniq := map[string]struct{}{}
	err := filepath.WalkDir(e.pjdir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			return nil
		}
		pj, err := NewProject(e.pjdir, path)
		if err != nil {
			return nil
		}
		if _, ok := uniq[pj.Code()]; ok {
			return errors.New("project code " + pj.Code() + " is duplicate in projects_dir")
		}
		uniq[pj.Code()] = struct{}{}
		buf.WriteString(pj.DumpEvars() + "\n")
		return nil
	})
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (e *Evars) makeJoinEnv() string {
	var buf bytes.Buffer
	for _, k := range SortStringKeys(e.jenv) {
		buf.WriteString("set " + k + "=")
		for _, v := range e.jenv[k] {
			p := os.ExpandEnv(v)
			if filepath.IsAbs(p) {
				buf.WriteString(p + ";")
			} else {
				buf.WriteString(`%cd%\` + p + ";")
			}
		}
		buf.WriteString("%" + k + "%\n")
	}
	buf.WriteString("\n")
	return buf.String()
}
