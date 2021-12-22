package pjma

import (
	"bytes"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func MakeEvars() error {
	// Create file
	pjdir := viper.GetString("projects_dir")
	evars, err := os.Create(pjdir + `\custom_evars.bat`)
	if err != nil {
		return err
	}
	defer evars.Close()

	// Output multiple paths evars
	for _, v := range [3]string{"caf_uic_path", "pmllib", "pmlui"} {
		_, err := evars.WriteString(dirsToEvar(v, viper.GetStringSlice(v)))
		if err != nil {
			return err
		}
	}
	if _, err := evars.WriteString(dirsToEvar("pdmsui", viper.GetStringSlice("pmlui"))); err != nil {
		return err
	}

	// Output project evars defined by pjma.yml
	for _, v := range viper.GetStringSlice("extrapj") {
		output, err := prjDirToEvars(v)
		if err != nil {
			return err
		}
		if _, err := evars.WriteString(output); err != nil {
			return err
		}
	}

	// Output local project evars
	output, err := walkPrjDirToEvars(pjdir)
	if err != nil {
		return err
	}
	if _, err := evars.WriteString(output + "\r\n"); err != nil {
		return err
	}

	// Output additional commands
	for _, v := range viper.GetStringSlice("extracmd") {
		_, err := evars.WriteString(v + "\r\n")
		if err != nil {
			return err
		}
	}

	return nil
}

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
	buf.WriteString("set projects_dir=" + pjdir + "\\\r\n")
	buf.WriteString(`start "" cmd /c "` + viper.GetString(viper.GetString("appname")) + `"`)
	buf.WriteString(" " + viper.GetString("context.module"))
	if viper.GetBool("context.tty") {
		buf.WriteString(" TTY")
	}
	buf.WriteString(" " + viper.GetString("context.project"))
	buf.WriteString(" " + viper.GetString("context.user"))
	mdb := viper.GetString("context.mdb")
	if mdb != "" && !strings.HasPrefix(mdb, "/") {
		mdb = "/" + mdb
	}
	buf.WriteString(" " + mdb)
	macro := viper.GetString("context.macro")
	if macro != "" {
		buf.WriteString(" $m'%%projects_dir%%..\\" + macro + "'")
	}
	return strings.TrimSpace(buf.String()), nil
}

func dirsToEvar(name string, dirs []string) string {
	if len(dirs) == 0 {
		return ""
	}
	var buf bytes.Buffer
	for _, d := range dirs {
		if strings.HasPrefix(d, `./`) || strings.HasPrefix(d, `.\`) {
			buf.WriteString("%projects_dir%.")
		}
		buf.WriteString(d + ";")
	}
	return "set " + name + "=" + buf.String() + "%" + name + "%\r\n"
}

func prjDirToEvars(path string) (string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	// Search 000 directory
	pjcode := ""
	for _, f := range files {
		prj000 := f.Name()
		if f.IsDir() && strings.HasSuffix(prj000, "000") {
			pjcode = strings.TrimSuffix(prj000, "000")
			break
		}
	}
	if pjcode == "" {
		return "", nil
	}

	var buf bytes.Buffer
	buf.WriteString("\r\n")
	for _, f := range files {
		n := f.Name()
		if f.IsDir() && strings.HasPrefix(n, pjcode) {
			buf.WriteString("set " + n + "=" + filepath.Join(path, n) + "\r\n")
		}
	}
	buf.WriteString("set " + pjcode + "000id=" + filepath.Base(path) + "\r\n")
	return buf.String(), nil
}

func walkPrjDirToEvars(root string) (string, error) {
	var buf bytes.Buffer
	err := fs.WalkDir(os.DirFS(root), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !(d.IsDir() && strings.HasSuffix(path, "000")) {
			return nil
		}

		pjcode := strings.TrimSuffix(d.Name(), "000")
		pjdir := filepath.Dir(path)
		files, err := os.ReadDir(filepath.Join(root, pjdir))
		if err != nil {
			return err
		}
		buf.WriteString("\r\n")
		for _, f := range files {
			n := f.Name()
			if f.IsDir() && strings.HasPrefix(n, pjcode) {
				buf.WriteString("set " + n + "=%projects_dir%" + filepath.Join(pjdir, n) + "\r\n")
			}
		}
		buf.WriteString("set " + pjcode + "000id=" + filepath.Base(pjdir) + "\r\n")

		return nil
	})
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
