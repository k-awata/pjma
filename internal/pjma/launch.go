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
	buf, err := genLaunch(true)
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
	buf, err := genLaunch(viper.GetBool("absbat"))
	if err != nil {
		return err
	}
	if _, err := launch.WriteString(buf); err != nil {
		return err
	}
	return nil
}

func genLaunch(abs bool) (string, error) {
	pjrel := viper.GetString("projects_dir")
	pjdir, err := filepath.Abs(pjrel)
	if err != nil {
		return "", err
	}
	if !abs && !filepath.IsAbs(pjrel) {
		pjdir = "%cd%\\" + pjrel
	}
	var buf bytes.Buffer
	buf.WriteString("@echo off\r\n")
	buf.WriteString("set projects_dir=" + pjdir + "\\\r\n")
	buf.WriteString("cd /d \"%temp%\"\r\n")
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
	return strings.TrimSpace(buf.String()) + "\r\n", nil
}
