package storage

import (
	"encoding/json"
	"errors"

	"github.com/tomatoCoderq/gitnotes/internal/models"
	bolt "go.etcd.io/bbolt"
)

const (
	notesBucket     = "gitnotes"
	errBuckNotFound = "bucket not found"
	errNoteNotFound = "note not found"
)

type NotesMap map[string][]models.Note

type NoteStruct struct {
	Ref  string
	Note models.Note
}

func GetRefFromNoteFields(db *bolt.DB, value string) (string, error) {
	notesMap, err := LoadAllNotesBolt(db)
	if err != nil {
		return "", err
	}
	for key, notes := range notesMap {
		for _, note := range notes {
			if note.Title == value {
				return key, nil
			}
		}
	}
	return "", nil
}

func LoadNoteBold(db *bolt.DB, ref string) ([]models.Note, error) {
	var notes []models.Note
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(notesBucket))
		if b == nil {
			return errors.New(errBuckNotFound)
		}

		data := b.Get([]byte(ref))
		if data == nil {
			return errors.New(errNoteNotFound)
		}
		return json.Unmarshal(data, &notes)
	})
	return notes, err
}

func LoadAllNotesBolt(db *bolt.DB) (NotesMap, error) {
	notesMap := make(NotesMap, 0)

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(notesBucket))
		if b == nil {
			return errors.New(errBuckNotFound)
		}
		return b.ForEach(func(k, v []byte) error {
			var notes []models.Note
			if err := json.Unmarshal(v, &notes); err != nil {
				return err
			}
			notesMap[string(k)] = notes
			return nil
		})
	})
	return notesMap, err
}

func SaveNoteBold(db *bolt.DB, ref string, note models.Note) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(notesBucket))
		if err != nil {
			return err
		}

		var existing []models.Note
		data := b.Get([]byte(ref))
		if data != nil {
			if err := json.Unmarshal(data, &existing); err != nil {
				return err
			}
		}

		allNotes := append(existing, note)

		data, err = json.Marshal(allNotes)
		if err != nil {
			return err
		}

		return b.Put([]byte(ref), data)
	})
}

func SaveNotesBold(db *bolt.DB, ref string, notes []models.Note) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(notesBucket))
		if err != nil {
			return err
		}

		var existing []models.Note
		data := b.Get([]byte(ref))
		if data != nil {
			if err := json.Unmarshal(data, &existing); err != nil {
				return err
			}
		}

		allNotes := append(existing, notes...)

		data, err = json.Marshal(allNotes)
		if err != nil {
			return err
		}
		return b.Put([]byte(ref), data)
	})
}

func FindByRef(db *bolt.DB, ref string) (NotesMap, error) {
	var notes []models.Note
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(notesBucket))
		if b == nil {
			return errors.New(errBuckNotFound)
		}

		data := b.Get([]byte(ref))

		err := json.Unmarshal(data, &notes)
		if err != nil {
			return err
		}

		return nil
	})
	return NotesMap{ref: notes}, err
}

func FindByValue(db *bolt.DB, value string, parameter string) ([]NoteStruct, error) {
	var (
		notes      []models.Note
		foundNotes []NoteStruct
	)

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(notesBucket))
		if b == nil {
			return errors.New(errBuckNotFound)
		}

		b.ForEach(func(k, v []byte) error {
			if err := json.Unmarshal(v, &notes); err != nil {
				return err
			}

			for _, note := range notes {
				if (parameter == "tag" && note.Tag == value) ||
				(parameter == "title" && note.Title == value){
					foundNotes = append(foundNotes, NoteStruct{string(k), note})
				}
			}
			return nil
		})
		return nil
	})
	return foundNotes, err
}

func RemoveNotesByReferencBold(db *bolt.DB, reference string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(notesBucket))
		if err != nil {
			return err
		}
		return b.Delete([]byte(reference))
	})
}

func RemoveNotesByTitleBold(db *bolt.DB, title string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(notesBucket))
		if err != nil {
			return err
		}

		b.ForEach(func(k, v []byte) error {
			var (
				notes []models.Note
				remainingNotes []models.Note
			)

			if err := json.Unmarshal(v, &notes); err != nil {
				return err
			}

			for _, note := range notes {
				if note.Title != title {
					remainingNotes = append(remainingNotes, note)
				}
			}
			if len(remainingNotes) == 0 {
				return b.Delete(k) // Remove the bucket if no notes remain
			}
			data, err := json.Marshal(remainingNotes)
			if err != nil {
				return err
			}
			return b.Put(k, data)
		})
		return nil
	})
}
