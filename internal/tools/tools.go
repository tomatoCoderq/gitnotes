package tools

import (
	"os"
	"path/filepath"
)


// TODO: Use UserHomeDIR to store my json file
func GetHomePath() string {
	path, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	
	jsonNotesPath := filepath.Join(path, ".gitnotes", "gitnotes.json")
	return jsonNotesPath
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
