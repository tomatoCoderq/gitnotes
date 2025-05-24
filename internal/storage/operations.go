package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"maps"
	"os"
	"slices"
	"sync"

	"github.com/tomatoCoderq/gitnotes/internal/models"
	"github.com/tomatoCoderq/gitnotes/internal/tools"
)

type NotesMap map[string]models.Note

var loadNotesFunc = LoadNotes

func LoadNotesRaw(limit int, pathfile string) ([]NotesMap, error) {
	data, err := os.ReadFile(pathfile)
	if err != nil {
		return []NotesMap{}, err
	}

	var raw []json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return []NotesMap{{"s": models.Note{Title: "lol"}}}, err
	}

	var (
		wg    sync.WaitGroup // for waiting all goroutines
		mu    sync.Mutex     // for shared variables
		notes []NotesMap
	)

	for _, item := range raw {
		wg.Add(1)
		go func(item json.RawMessage) {
			defer wg.Done()

			var note NotesMap
			if err := json.Unmarshal(item, &note); err != nil {
				fmt.Println("Decode error:", err) // log or collect if needed
				return
			}

			mu.Lock()
			notes = append(notes, note)
			mu.Unlock()
		}(item)
	}

	wg.Wait()
	return notes, nil
}

func LoadNotes(limit int, filename string) ([]NotesMap, error) {
	// Function return all notes stored in gitnotes.json
	var notes []NotesMap

	if exists := tools.FileExists(filename); !exists {
		if _, err := os.Create(filename); err != nil {
			return nil, err
		}
	}

	reader, err := os.Open(filename)
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
	notesLoaded, err := LoadNotes(100, tools.GetHomePath())
	if err != nil {
		return err
	}

	concatNotes := slices.Concat(notesLoaded, notes)

	file, err := os.OpenFile(tools.GetHomePath(), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0o666)
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

// Find all notes in gitnotes.json which satisfy ref
func FindByParameter(parameter string, value string) ([]models.Note, error) {
	foundNotes := []models.Note{}
	notes, err := loadNotesFunc(100, tools.GetHomePath())
	if err != nil {
		return []models.Note{}, err
	}

	switch parameter {
	case "title":
		for _, notesmap := range notes {
			for note := range maps.Values(notesmap) {
				if note.Title == value {
					foundNotes = append(foundNotes, note)
				}
			}
		}
	case "ref":
		for _, notesmap := range notes {
			for key := range maps.Keys(notesmap) {
				if key == value {
					foundNotes = append(foundNotes, notesmap[key])
				}
			}
		}
	case "tag":
		for _, notesmap := range notes {
			for note := range maps.Values(notesmap) {
				if note.Tag == value {
					foundNotes = append(foundNotes, note)
				}
			}
		}
	}

	return foundNotes, nil
}

func RemoveNotesByReference(reference string, filename string) error {
	notes, err := loadNotesFunc(100, tools.GetHomePath())
	if err != nil {
		return err
	}

	var notesToKeep []NotesMap

	for _, note := range notes {
		for key := range maps.Keys(note) {
			if key != reference {
				notesToKeep = append(notesToKeep, note)
			}
		}
	}

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0o666)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(notesToKeep)
	if err != nil {
		return err
	}
	return nil
}

func RemoveNotesByTitle(title string, filename string) error {
	notes, err := loadNotesFunc(100, tools.GetHomePath())
	if err != nil {
		return err
	}

	var notesToKeep []NotesMap

	for _, note := range notes {
		for value := range maps.Values(note) {
			if value.Title != title {
				notesToKeep = append(notesToKeep, note)
			}
		}
	}

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0o666)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(notesToKeep)
	if err != nil {
		return err
	}

	return nil
}
