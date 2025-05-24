package storage

import (
	"encoding/json"
	_"fmt"
	"io"
	"maps"
	"os"
	"path/filepath"
	_ "reflect"
	"slices"

	"gitnotes/internal/models"
)

type NotesMap map[string]models.Note

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// TODO: Use UserHomeDIR to store my json file
func GetHomePath() string {
	path, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	jsonNotesPath := filepath.Join(path, ".gitnotes", "gitnotes.json")
	return jsonNotesPath
}

func LoadNotes(limit int) ([]NotesMap, error) {
	// Function return all notes stored in gitnotes.json
	var notes []NotesMap

	if exists := fileExists(GetHomePath()); !exists {
		if _, err := os.Create(GetHomePath()); err != nil {
			return nil, err
		}
	}

	reader, err := os.Open(GetHomePath())
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	err = json.NewDecoder(reader).Decode(&notes)

	if err == io.EOF {
		return []NotesMap{}, nil
	}
	if err != nil {
		return nil, err
	}

	if limit >= len(notes) {
		return notes, nil
	}

	return notes[:limit], nil
}

func SaveNotes(notes []NotesMap) error {
	// Save to gitnotes.json slice of note
	notesLoaded, err := LoadNotes(100)
	if err != nil {
		return err
	}

	concatNotes := slices.Concat(notesLoaded, notes)

	file, err := os.OpenFile(GetHomePath(), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0o666)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(concatNotes)
	if err != nil {
		return err
	}

	return nil
}

func FindByRef(ref string) (models.Note, error) {
	// Find all notes in gitnotes.json which satisfy ref
	notes, err := LoadNotes(100)
	if err != nil {
		return models.Note{}, err
	}

	for _, note := range notes {
		for key := range maps.Keys(note) {
			if key == ref {
				return note[key], nil
			}
		}
	}

	return models.Note{}, nil

	// for _, note := range notes {
	// 	for value := range maps.Values(note) {
	// 		if value.Title == ref || value.Content == ref {
	// 			return note, nil
	// 		}
	// 	}
	// }
	// return NotesMap{}, nil
}

// func RemoveNotes(reader io.Reader, notesToRemove []NotesMap) error{
// 	notes, err := LoadNotes(reader)
// 	if err != nil {
// 		return err
// 	}

// 	var notesToKeep []NotesMap

// 	for _, note := range notes {
// 		for _, noteToRemove := range notesToRemove {
// 			if !reflect.DeepEqual(note, noteToRemove) {
// 				notesToKeep = append(notesToKeep, note)
// 			}
// 		}
// 	}

// 	file, err := os.OpenFile("gitnotes.json", os.O_WRONLY | os.O_TRUNC, 0666)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	encoder := json.NewEncoder(file)
// 	err = encoder.Encode(notesToKeep)

// 	if err != nil {
// 		return err
// 	}

// 	return nil

// }
