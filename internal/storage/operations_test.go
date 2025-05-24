package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"slices"
	"testing"
	"time"

	"github.com/tomatoCoderq/gitnotes/internal/models"
	"github.com/tomatoCoderq/gitnotes/internal/tools"

	"github.com/stretchr/testify/assert"
)

func TestLoadNotesRaw(t *testing.T) {
	tmpDir := t.TempDir()
	originalWD, _ := os.Getwd()
	defer os.Chdir(originalWD)

	err := os.Chdir(tmpDir)
	assert.NoError(t, err)

	// Create dummy notes data
	testData := []NotesMap{
		{"abc123": {Title: "Test 2", Content: "One"}},
		{"def456": {Title: "Test 1", Content: "Two"}},
	}

	// Write to gitnotes.json
	file, err := os.Create("gitnotes.json")
	assert.NoError(t, err)

	df, _ := os.Getwd()
	fmt.Println("CURR", df)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(testData)
	assert.NoError(t, err)
	file.Close()

	// Run the function
	notes, err := LoadNotesRaw(100, "gitnotes.json")
	assert.NoError(t, err)
	assert.Equal(t, 2, len(notes))
	assert.Equal(t, "Test 1", notes[0]["def456"].Title)
	assert.Equal(t, "Test 2", notes[1]["abc123"].Title)
}

func TestLoadNotes(t *testing.T) {
	// Test LoadNotes function
	tmpDir := t.TempDir()
	originalWD, _ := os.Getwd()
	defer os.Chdir(originalWD)

	err := os.Chdir(tmpDir)
	assert.NoError(t, err)

	// Create dummy notes data
	testData := []NotesMap{
		{"abc123": {Title: "Test 2", Content: "One"}},
		{"def456": {Title: "Test 1", Content: "Two"}},
	}

	// Write to gitnotes.json
	file, err := os.Create("gitnotes.json")
	assert.NoError(t, err)

	df, _ := os.Getwd()
	fmt.Println("CURR", df)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(testData)
	assert.NoError(t, err)
	file.Close()

	// Run the function
	notes, err := LoadNotes(100, "gitnotes.json")
	assert.NoError(t, err)
	assert.Equal(t, 2, len(notes))
	assert.Equal(t, "Test 1", notes[0]["def456"].Title)
	assert.Equal(t, "Test 2", notes[1]["abc123"].Title)
}

func BenchmarkLoadNotes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoadNotes(5, tools.GetHomePath())
	}
}

func BenchmarkLoadNotesRaw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoadNotesRaw(5, tools.GetHomePath())
	}
}

func TestSaveNotes(t *testing.T) {
	// Test LoadNotes function

	limit := 5

	timestamp := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)

	notes_to_save := []NotesMap{
		{"20000": {Title: "Note 0404", Content: "Content 1", CreatedAt: timestamp}},
	}

	valid_notes, err := LoadNotes(limit, tools.GetHomePath())
	if err != nil {
		t.Errorf("LoadNotes() error = %v", err)
	}

	err = SaveNotes(notes_to_save)
	if err != nil {
		t.Errorf("SaveNotes() error = %v", err)
	}

	read_notes, err := LoadNotes(limit, tools.GetHomePath())
	if err != nil {
		t.Errorf("LoadNotes() readnotes error = %v", err)
	}

	if !reflect.DeepEqual(slices.Concat(valid_notes, notes_to_save), read_notes) {
		t.Errorf("Notes are not the same! Got %v Want %v", read_notes, slices.Concat(valid_notes, notes_to_save))
	}
}

func TestFindBy(t *testing.T) {
	mockNote := NotesMap{
		"abc123": {Title: "Mock Title", Content: "Mock Content", CreatedAt: time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC), Tag: "BUG"},
		"abc":    {Title: "Mock Title2", Content: "Mock Content2", CreatedAt: time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC), Tag: "CRITICAL"},
	}

	loadNotesFunc = func(limit int, filename string) ([]NotesMap, error) {
		return []NotesMap{mockNote}, nil
	}
	t.Run("Not found by reference", func(t *testing.T) {
		notes, err := FindByParameter("Note 123", "reference")
		if err != nil {
			t.Errorf("Find By red error: %v", err)
		}

		assert.Equal(t, notes, []models.Note{})
	})

	t.Run("Not found by title", func(t *testing.T) {
		notes, err := FindByParameter("WRONG TITLE", "title")
		if err != nil {
			t.Errorf("Find By red error: %v", err)
		}

		assert.Equal(t, notes, []models.Note{})
	})

	t.Run("Not found by tag", func(t *testing.T) {
		notes, err := FindByParameter("WRONG TAG", "tag")
		if err != nil {
			t.Errorf("Find By red error: %v", err)
		}

		assert.Equal(t, notes, []models.Note{})
	})

	t.Run("Found by reference", func(t *testing.T) {
		notes, err := FindByParameter("ref", "abc123")
		if err != nil {
			t.Errorf("Find By red error: %v", err)
		}
		assert.Equal(t, []models.Note{mockNote["abc123"]}, notes)
	})

	t.Run("Found by title", func(t *testing.T) {
		notes, err := FindByParameter("title", "Mock Title")
		if err != nil {
			t.Errorf("Find By red error: %v", err)
		}
		assert.Equal(t, notes, []models.Note{mockNote["abc123"]})
	})

	t.Run("Found by tag", func(t *testing.T) {
		notes, err := FindByParameter("tag", "CRITICAL")
		if err != nil {
			t.Errorf("Find By red error: %v", err)
		}
		assert.Equal(t, notes, []models.Note{mockNote["abc"]})
	})
}

func TestRemoveNotesByReference(t *testing.T) {
	// Test LoadNotes function
	tmpFile, err := os.CreateTemp("", "test_gitnotes_*.json")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name()) // clean up

	mockNotes := []NotesMap{
		{"abc123": models.Note{Title: "Note1"}},
		{"xyz789": models.Note{Title: "Note2"}},
	}

	loadNotesFunc = func(limit int, filename string) ([]NotesMap, error) {
		return mockNotes, nil
	}
	defer func() { loadNotesFunc = LoadNotes }()

	err = RemoveNotesByReference("abc123", tmpFile.Name())
	assert.NoError(t, err)

	// Check file content
	content, err := os.ReadFile(tmpFile.Name())
	assert.NoError(t, err)

	var saved []NotesMap
	err = json.Unmarshal(content, &saved)
	assert.NoError(t, err)

	assert.Len(t, saved, 1)
	_, exists := saved[0]["xyz789"]
	assert.True(t, exists)
}

func TestRemoveNotesByTitle(t *testing.T) {
	// Test LoadNotes function
	tmpFile, err := os.CreateTemp("", "test_gitnotes_*.json")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name()) // clean up

	mockNotes := []NotesMap{
		{"abc123": models.Note{Title: "Note1"}},
		{"xyz789": models.Note{Title: "Note2"}},
	}

	loadNotesFunc = func(limit int, filename string) ([]NotesMap, error) {
		return mockNotes, nil
	}
	defer func() { loadNotesFunc = LoadNotes }()

	err = RemoveNotesByTitle("Note1", tmpFile.Name())
	assert.NoError(t, err)

	// Check file content
	content, err := os.ReadFile(tmpFile.Name())
	assert.NoError(t, err)

	var saved []NotesMap
	err = json.Unmarshal(content, &saved)
	assert.NoError(t, err)

	assert.Len(t, saved, 1)
	_, exists := saved[0]["xyz789"]
	assert.True(t, exists)
}
