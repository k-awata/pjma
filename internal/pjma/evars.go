package pjma

import (
	"bytes"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func MakeEvars() error {
	// Create file
	pjdir := viper.GetString("projects_dir")
	if _, err := os.Stat(pjdir); err != nil {
		os.Mkdir(pjdir, 0777)
	}
	evars, err := os.Create(pjdir + `\custom_evars.bat`)
	if err != nil {
		return err
	}
	defer evars.Close()

	// Output local project evars
	output, err := walkPrjDirToEvars(pjdir)
	if err != nil {
		return err
	}
	if _, err := evars.WriteString(output); err != nil {
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

	// Output multiple paths evars
	for _, v := range [3]string{"caf_uic_path", "pmllib", "pmlui"} {
		_, err := evars.WriteString(dirsToEvar(v, viper.GetStringSlice(v)))
		if err != nil {
			return err
		}
	}
	if _, err := evars.WriteString(dirsToEvar("pdmsui", viper.GetStringSlice("pmlui")) + "\r\n"); err != nil {
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

func GetProjectCode(path string) (string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}
	for _, f := range files {
		prj000 := f.Name()
		if f.IsDir() && strings.HasSuffix(prj000, "000") {
			return strings.TrimSuffix(prj000, "000"), nil
		}
	}
	return "", errors.New("000 directory is not found in " + path)
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
		for _, f := range files {
			n := f.Name()
			if f.IsDir() && strings.HasPrefix(n, pjcode) {
				buf.WriteString("set " + n + "=%projects_dir%" + filepath.Join(pjdir, n) + "\r\n")
			}
		}
		buf.WriteString("set " + pjcode + "000id=" + filepath.Base(pjdir) + "\r\n")
		buf.WriteString("\r\n")

		return nil
	})
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func prjDirToEvars(path string) (string, error) {
	pjdir, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	pjcode, err := GetProjectCode(pjdir)
	if err != nil {
		return "", err
	}
	files, err := os.ReadDir(pjdir)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	for _, f := range files {
		n := f.Name()
		if f.IsDir() && strings.HasPrefix(n, pjcode) {
			buf.WriteString("set " + n + "=" + filepath.Join(pjdir, n) + "\r\n")
		}
	}
	buf.WriteString("set " + pjcode + "000id=" + filepath.Base(pjdir) + "\r\n")
	buf.WriteString("\r\n")
	return buf.String(), nil
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
