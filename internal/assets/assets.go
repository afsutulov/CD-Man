// Package assets embeds the resources needed to run the game.
package assets

import "embed"

//go:embed state.dat
var Initial []byte

//go:embed files/*
var resources embed.FS

func Files() (map[string][]byte, error) {
	entries, err := resources.ReadDir("files")
	if err != nil {
		return nil, err
	}
	files := make(map[string][]byte, len(entries))
	for _, e := range entries {
		b, err := resources.ReadFile("files/" + e.Name())
		if err != nil {
			return nil, err
		}
		files[e.Name()] = b
	}
	return files, nil
}
