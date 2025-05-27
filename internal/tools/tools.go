package tools

import (
	"os"
	"path/filepath"
)

const (
	fileName = "gitnotes.db"
)

func GetHomePath() string {
	path, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	jsonNotesPath := filepath.Join(path, ".gitnotes", fileName)
	return jsonNotesPath
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
