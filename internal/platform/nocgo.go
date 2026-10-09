//go:build !windows && !cgo

package platform

import (
	"cdman2/internal/game"
	"fmt"
)

func Run(*game.Machine, func() error) error {
	return fmt.Errorf("native Linux/macOS window requires CGO_ENABLED=1 and SDL2; Windows and cdman-headless do not require cgo")
}
func ShowError(s string) { fmt.Println(s) }
