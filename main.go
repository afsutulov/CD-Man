package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"cdman2/internal/assets"
	"cdman2/internal/game"
	"cdman2/internal/platform"
)

func run() error {
	files, err := assets.Files()
	if err != nil {
		return err
	}
	base, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	saveDir := filepath.Join(base, "CDMan-Go")
	for _, name := range []string{"HIGHSC.CDM", "DEMO.CDM"} {
		b, err := os.ReadFile(filepath.Join(saveDir, name))
		if err == nil {
			files[name] = b
		} else if !os.IsNotExist(err) {
			return err
		}
	}
	m, err := game.New(assets.Initial, files, time.Now())
	if err != nil {
		return err
	}
	save := func() error {
		for name, b := range m.DrainWrites() {
			if name != "HIGHSC.CDM" && name != "DEMO.CDM" {
				return fmt.Errorf("unexpected save name %q", name)
			}
			if err := os.MkdirAll(saveDir, 0700); err != nil {
				return err
			}
			f, err := os.CreateTemp(saveDir, "save-*")
			if err != nil {
				return err
			}
			tmp := f.Name()
			_, writeErr := f.Write(b)
			closeErr := f.Close()
			if writeErr != nil {
				os.Remove(tmp)
				return writeErr
			}
			if closeErr != nil {
				os.Remove(tmp)
				return closeErr
			}
			if err := os.Rename(tmp, filepath.Join(saveDir, name)); err != nil {
				os.Remove(tmp)
				return err
			}
		}
		return nil
	}
	if err := platform.Run(m, save); err != nil {
		return err
	}
	return save()
}
func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, strings.TrimSpace(err.Error()))
		platform.ShowError(err.Error())
		os.Exit(1)
	}
}
