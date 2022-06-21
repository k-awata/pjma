package pjma

import (
	"bytes"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Project struct {
	root string
	path string
	code string
	dirs []string
}

// NewProject returns project from directory path
func NewProject(root string, path string) (*Project, error) {
	p := &Project{
		root: os.ExpandEnv(root),
		path: os.ExpandEnv(path),
	}
	if err := p.loadEvars(); err != nil {
		return nil, err
	}
	return p, nil
}

// Code returns project code
func (p *Project) Code() string {
	return p.code
}

// DumpEvars returns bat commands setting project evars
func (p *Project) DumpEvars() string {
	var buf bytes.Buffer
	for _, d := range p.dirs {
		r := ""
		if p.root != "" {
			r = `%` + PROJECTS_DIR + `%\`
		} else if !filepath.IsAbs(d) {
			r = `%cd%\`
		}
		buf.WriteString("set " + filepath.Base(d) + "=" + r + d + "\n")
	}
	buf.WriteString("set " + p.code + "000id=" + filepath.Base(p.path) + "\n")
	return buf.String()
}

func (p *Project) loadEvars() error {
	dirs, err := os.ReadDir(p.path)
	if err != nil {
		return err
	}
	// Fetch project code
	err = p.fetchCode(dirs)
	if err != nil {
		return err
	}
	// Fetch sub directories
	path, err := filepath.Rel(p.root, p.path)
	if err != nil {
		path = p.path
		p.root = ""
	}
	p.dirs = []string{}
	for _, f := range dirs {
		n := f.Name()
		if f.IsDir() && strings.HasPrefix(n, p.code) {
			p.dirs = append(p.dirs, filepath.Join(path, n))
		}
	}
	return nil
}

func (p *Project) fetchCode(f []fs.DirEntry) error {
	const dir000 = "000"
	for _, f := range f {
		n := f.Name()
		if f.IsDir() && strings.HasSuffix(n, dir000) {
			p.code = strings.TrimSuffix(n, dir000)
			return nil
		}
	}
	return errors.New(dir000 + "directory is not found in " + p.path)
}
