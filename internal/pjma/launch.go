package pjma

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func Launch() error {
	launch, err := os.CreateTemp("", "*.bat")
	if err != nil {
		return err
	}
	defer func() {
		launch.Close()
		os.Remove(launch.Name())
	}()
	buf, err := genLaunch("")
	if err != nil {
		return err
	}
	if _, err := launch.WriteString(buf); err != nil {
		return err
	}
	if err := exec.Command("cmd", "/c", launch.Name()).Run(); err != nil {
		return err
	}
	return nil
}

func MakeLaunch(path string) error {
	launch, err := os.Create(path)
	if err != nil {
		return err
	}
	defer launch.Close()
	dir := ""
	if !viper.GetBool("absbat") {
		dir = filepath.Dir(path)
	}
	buf, err := genLaunch(dir)
	if err != nil {
		return err
	}
	if _, err := launch.WriteString(buf); err != nil {
		return err
	}
	return nil
}

func genLaunch(path string) (string, error) {
	pjdir, err := filepath.Abs(viper.GetString("projects_dir"))
	if err != nil {
		return "", err
	}
	if path != "" {
		// Set relative path
		base, err := filepath.Abs(path)
		if err != nil {
			return "", err
		}
		pwd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		reldir, err := filepath.Rel(base, pwd)
		if err != nil {
			return "", err
		}
		pjdir = "%~dp0" + filepath.Join(reldir, viper.GetString("projects_dir"))
	}

	var buf bytes.Buffer
	buf.WriteString("@echo off\r\n")
	buf.WriteString("cd /d %temp%\r\n")
	buf.WriteString("set projects_dir=" + pjdir + "\\\r\n")
	buf.WriteString(`start "" /wait cmd /c "` + viper.GetString("apps."+viper.GetString("context.bat")) + `"`)
	buf.WriteString(" " + viper.GetString("context.module"))
	if viper.GetBool("context.tty") {
		buf.WriteString(" TTY")
	}
	buf.WriteString(" " + viper.GetString("context.project"))
	buf.WriteString(" " + viper.GetString("context.user"))
	mdb := viper.GetString("context.mdb")
	if mdb != "" {
		if !strings.HasPrefix(mdb, "/") {
			mdb = "/" + mdb
		}
		buf.WriteString(" " + mdb)
	}
	macro := viper.GetString("context.macro")
	if macro != "" {
		if _, err := os.Stat(macro); err == nil {
			buf.WriteString(" $m'%%projects_dir%%..\\" + macro + "'")
		} else {
			buf.WriteString(" " + macro)
		}
	}
	return strings.TrimSpace(buf.String()), nil
}
