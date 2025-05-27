package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomatoCoderq/gitnotes/internal/models"
	bolt "go.etcd.io/bbolt"
)

func TestLoadNoteBold(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	// Prepare test data
	note1 := models.Note{Title: "Note 1", Content: "Content for note 1"}
	note2 := models.Note{Title: "Note 2", Content: "Content for note 2"}
	ref := "test-ref"

	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(notesBucket))
		if err != nil {
			return err
		}
		data, err := json.Marshal([]models.Note{note1, note2})
		if err !=nil {
			return err
		}
		return b.Put([]byte(ref), data) // Create the bucket for the reference
	})

	// Test LoadNoteBold
	loadedNotes, err := LoadNoteBold(db, ref)
	assert.NoError(t, err)

	// Verify the loaded notes
	expectedNotes := []models.Note{note1, note2}
	assert.Equal(t, expectedNotes, loadedNotes)
}

func TestLoadNoteBold_NoteNotFound(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	// Create the bucket but do not add any notes
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(notesBucket))
		return err
	})
	assert.NoError(t, err)
}

func TestSaveNotesBold(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	// Prepare test data
	notes := []models.Note{
		{Title: "Note 1", Content: "Content for note 1"},
		{Title: "Note 2", Content: "Content for note 2"},
	}
	ref := "test-ref"

	// Test SaveNotesBold

	err := SaveNotesBold(db, ref, notes)
	assert.NoError(t, err)

	// Verify the saved notes
	var loadedNotes []models.Note
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(notesBucket))
		if b == nil {
			return fmt.Errorf("bucket not found")
		}

		data := b.Get([]byte(ref))
		if data == nil {
			return fmt.Errorf("notes not found")
		}

		return json.Unmarshal(data, &loadedNotes)
	})
	assert.NoError(t, err)
	assert.Equal(t, notes, loadedNotes) // Only the last note is saved due to overwriting
}


func TestSaveNoteBold(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	note := models.Note{Title: "Test Note", Content: "This is a test note"}
	ref := "test-ref"

	err := SaveNoteBold(db, ref, note)
	assert.NoError(t, err)

	var storedNote []models.Note
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(notesBucket))
		data := b.Get([]byte(ref))
		return json.Unmarshal(data, &storedNote)
	})
	assert.NoError(t, err)
	assert.Equal(t, []models.Note{note}, storedNote)
}


func TestFindByTitle(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	note := models.Note{Title: "Unique Title", Content: "This is a test note"}
	ref := "test-ref"

	err := SaveNoteBold(db, ref, note)
	assert.NoError(t, err)

	foundNotes, err := FindByValue(db, "Unique Title", "title")
	assert.NoError(t, err)
	assert.Equal(t, []NoteStruct{{ref, note}}, foundNotes)
}



func setupTestDB(t *testing.T) (*bolt.DB, func()) {
	db, err := bolt.Open("test.db", 0o600, nil)
	if err != nil {
		t.Fatalf("Failed to open test database: %v", err)
	}

	cleanup := func() {
		db.Close()
		os.Remove("test.db")
	}

	return db, cleanup
}
