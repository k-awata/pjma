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
	buf, err := MakeLaunch(true)
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

func MakeLaunch(abs bool) (string, error) {
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
	appname := viper.GetString("context.bat")
	buf.WriteString("echo Running app " + appname + "...\r\n")
	buf.WriteString(`start "" /wait cmd /c "` + viper.GetString("apps."+appname) + `"`)
	module := strings.TrimSpace(viper.GetString("context.module"))
	if module != "" {
		buf.WriteString(" " + module)
	}
	if viper.GetBool("context.tty") {
		buf.WriteString(" TTY")
	}
	proj := strings.TrimSpace(viper.GetString("context.project"))
	user := strings.TrimSpace(viper.GetString("context.user"))
	if proj != "" && user != "" {
		buf.WriteString(" " + proj)
		buf.WriteString(" " + user)
		mdb := strings.TrimSpace(viper.GetString("context.mdb"))
		if mdb != "" {
			if !strings.HasPrefix(mdb, "/") {
				mdb = "/" + mdb
			}
			buf.WriteString(" " + mdb)
		}
		macro := strings.TrimSpace(viper.GetString("context.macro"))
		if macro != "" {
			if _, err := os.Stat(macro); err == nil {
				buf.WriteString(" $m'%%projects_dir%%..\\" + macro + "'")
			} else {
				buf.WriteString(" " + macro)
			}
		}
	}
	buf.WriteString("\r\n")
	return buf.String(), nil
}
