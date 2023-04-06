package pjma

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Launcher struct {
	pjdir string
	bat   string
	title string
	mod   string
	tty   string
	pj    string
	up    string
	mdb   string
	mac   string
}

// NewLauncher returns a new launcher
func NewLauncher(pjdir string, bat string, title string) *Launcher {
	return &Launcher{
		pjdir: os.ExpandEnv(pjdir),
		bat:   os.ExpandEnv(bat),
		title: title,
	}
}

// SetModule sets a module name
func (l *Launcher) SetModule(m string) *Launcher {
	l.mod = m
	return l
}

// SetTty sets whether launching bat with cli mode
func (l *Launcher) SetTty(t bool) *Launcher {
	if t {
		l.tty = "TTY"
	} else {
		l.tty = ""
	}
	return l
}

// SetProject sets a project code
func (l *Launcher) SetProject(p string) *Launcher {
	l.pj = p
	return l
}

// SetUser sets a username and password
func (l *Launcher) SetUser(u string, p string) *Launcher {
	if u != "" && p != "" {
		l.up = u + "/" + p
	} else {
		l.up = ""
	}
	return l
}

// SetMdb sets a MDB name
func (l *Launcher) SetMdb(m string) *Launcher {
	if m != "" && !strings.HasPrefix(m, "/") {
		l.mdb = "/" + m
	} else {
		l.mdb = m
	}
	return l
}

// SetMacro sets a macro command or macro filename
func (l *Launcher) SetMacro(m string) *Launcher {
	fs, err := os.Stat(m)
	if err == nil && !fs.IsDir() {
		if filepath.IsAbs(m) {
			l.mac = "$m'" + m + "'"
		} else {
			l.mac = `$m'%cd%\` + m + "'"
		}
	} else {
		l.mac = m
	}
	return l
}

// MakeBat returns contents of bat file
func (l *Launcher) MakeBat() string {
	var buf bytes.Buffer
	buf.WriteString("@echo off\n")
	if filepath.IsAbs(l.pjdir) {
		buf.WriteString("set " + PROJECTS_DIR + "=" + l.pjdir + "\n")
	} else {
		buf.WriteString("set " + PROJECTS_DIR + `=%cd%\` + l.pjdir + "\n")
	}
	buf.WriteString(`start "` + l.title + `" /wait cmd /c "` + l.bat + `"`)
	if l.mod != "" {
		buf.WriteString(" " + l.mod)
	}
	if l.tty != "" {
		buf.WriteString(" " + l.tty)
	}
	if l.pj != "" && l.up != "" {
		buf.WriteString(" " + l.pj + " " + l.up)
		if l.mdb != "" {
			buf.WriteString(" " + l.mdb)
		}
		if l.mac != "" {
			buf.WriteString(" " + l.mac)
		}
	}
	buf.WriteString("\n")
	return buf.String()
}

// Run runs launcher bat file
func (l *Launcher) Run(encode string) error {
	f, err := os.CreateTemp("", "*.bat")
	if err != nil {
		return err
	}
	defer func() {
		f.Close()
		os.Remove(f.Name())
	}()
	bat, err := EncodeForBatch(l.MakeBat(), encode)
	if err != nil {
		return err
	}
	if _, err := f.WriteString(bat); err != nil {
		return err
	}
	if err := exec.Command("cmd", "/c", f.Name()).Run(); err != nil {
		return err
	}
	return nil
}
