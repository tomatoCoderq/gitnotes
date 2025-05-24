package storage

import (
	"gitnotes/internal/models"
	"reflect"
	"slices"
	"testing"
	"time"

)

// TODO: Add tests with emulated file system
func TestLoadNotes(t *testing.T) {
	// Test LoadNotes function

	t.Run("Empty File", func(t *testing.T) {	
		t.Helper()
		notes, err := LoadNotes()
		if err != nil {
			t.Errorf("LoadNotes() error = %v", err)
		}
		if len(notes) != 0 {
			t.Logf("LoadNotes() returned more than 0 notes")
		}
	})


	t.Run("Non empty", func(t *testing.T) {
		timestamp := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
	
		valid_notes := []NotesMap{
			{"91919" : {Title: "Note 0404", Content: "Content 1", CreatedAt: timestamp}},
			{"91912" : {Title: "Note 0401", Content: "Content 1", CreatedAt: timestamp}},
		}
	
		notes, err := LoadNotes()
		if err != nil {
			t.Errorf("LoadNotes() error = %v", err)
		}
		if len(notes) == 0 {
			t.Errorf("LoadNotes() returned no notes")
		}
		if !reflect.DeepEqual(valid_notes, notes) {
			t.Errorf("Notes are not the same! Got %v Want %v", notes, valid_notes)
		}
	})


}

func TestSaveNotes(t *testing.T) {
	// Test LoadNotes function

	timestamp := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
	
	notes_to_save := []NotesMap{
		{"20000" : {Title: "Note 0404", Content: "Content 1", CreatedAt: timestamp}},
	}

	valid_notes, err := LoadNotes()

	if err != nil {
		t.Errorf("LoadNotes() error = %v", err)
	}

	err = SaveNotes(notes_to_save)

	if err != nil {
		t.Errorf("SaveNotes() error = %v", err)
	}

	read_notes, err := LoadNotes()

	if err != nil {
		t.Errorf("LoadNotes() readnotes error = %v", err)
	}

	if !reflect.DeepEqual(slices.Concat(valid_notes, notes_to_save), read_notes) {
		t.Errorf("Notes are not the same! Got %v Want %v", read_notes, slices.Concat(valid_notes, notes_to_save))
	}
}

func TestFindByRef(t *testing.T) {
	t.Run("No found", func(t *testing.T) {
		note, err := FindByRef("Note 123")
		if err != nil {
			t.Errorf("Find By red error: %v", err)
		}
		if !reflect.DeepEqual(note, models.Note{}) {
			t.Errorf("Should be empty")
		}
	})


	t.Run("Found ref", func(t *testing.T) {
		note, err := FindByRef("20000")
		if err != nil {
			t.Errorf("Find By red error: %v", err)
		}

		if !reflect.DeepEqual(note, models.Note{Title: "Note 0404", Content: "Content 1", CreatedAt: time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)}) {
			t.Errorf("Notes are not the same! Got %v Want %v", note, models.Note{Title: "Note 1", Content: "Content 1", CreatedAt: time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)})
		}

	})

}


// func TestRemoveNotes(t *testing.T) {
// 	// Test LoadNotes function

// 	t.Run("Standard remove", func(t *testing.T) {
// 		timestamp := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
	
// 		notes_to_remove := []models.Note{
// 			{Title: "Note 0404", Content: "Content 1", CreatedAt: timestamp},
// 		}
	
// 		reader1, _ := os.Open("gitnotes.json")
// 		valid_notes, err := LoadNotes(reader1)
	
// 		if err != nil {
// 			t.Errorf("LoadNotes() error = %v", err)
// 		}
	
// 		reader1.Seek(0, io.SeekStart)
// 		err = RemoveNotes(reader1, notes_to_remove)
	
// 		if err != nil {
// 			t.Errorf("SaveNotes() error = %v", err)
// 		}
// 		reader2, _ := os.Open("gitnotes.json")
	
// 		read_notes, err := LoadNotes(reader2)
	
// 		if err != nil {
// 			t.Errorf("LoadNotes() readnotes error = %v", err)
// 		}
	
// 		validNotesDeleted := make([]models.Note, 0)

// 		for _, note := range valid_notes {
// 			for _, noteToRemove := range notes_to_remove {
// 				if !reflect.DeepEqual(note, noteToRemove) {
// 					validNotesDeleted = append(validNotesDeleted, note)
// 				}
// 			}
// 		}

// 		if !slices.Equal(validNotesDeleted, read_notes) {
// 			t.Errorf("Notes are not the same! Got %v Want %v", read_notes, validNotesDeleted)
// 		}
// 	})


// }