package cmd

import (
	"bytes"
	_ "fmt"
	_ "gitnotes/internal/commands"
	_ "io"
	_ "strings"
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)

// TODO: write working tests for it
func TestAddCommand(t *testing.T) {
	input := bytes.NewBuffer([]byte("LOsdfd sfsdfl\n\n"))
	
	output := bytes.Buffer{}
	
	resolveGitRef = func(ref string) (string, error) {
	return "9934f4", nil
	}

	addCmnd := addCmd
	addCmnd.SetArgs([]string{"HEAD"})
	addCmnd.SetIn(input)
	addCmnd.SetOut(&output)
	addCmnd.SetErr(&output)
	
	err := addCmnd.Execute()

	assert.NoError(t, err)


	fmt.Println("== OUTPUT ==")
	fmt.Println(output.String())

	assert.Equal(t, "✅ Note added for 9934f4\n", output.String())

}


func TestAddCommand_Success(t *testing.T) {
	input := bytes.NewBuffer([]byte("This is a test note\n"))
	output := bytes.Buffer{}

	resolveGitRef = func(ref string) (string, error) {
		return "9934f4", nil
	}

	addCmnd := addCmd
	addCmnd.SetArgs([]string{"HEAD"})
	addCmnd.SetIn(input)
	addCmnd.SetOut(&output)
	addCmnd.SetErr(&output)

	err := addCmnd.Execute()

	assert.NoError(t, err)
	assert.Contains(t, output.String(), "✅ Note added for 9934f4")
}

func TestAddCommand_InvalidGitRef(t *testing.T) {
	input := bytes.NewBuffer([]byte("This is a test note\n"))
	output := bytes.Buffer{}

	resolveGitRef = func(ref string) (string, error) {
		return "", errors.New("invalid git reference")
	}

	addCmnd := addCmd
	addCmnd.SetArgs([]string{"INVALID_REF"})
	addCmnd.SetIn(input)
	addCmnd.SetOut(&output)
	addCmnd.SetErr(&output)

	err := addCmnd.Execute()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "could not resolve reference: INVALID_REF")
}

func TestAddCommand_FailedToReadInput(t *testing.T) {
	input := bytes.NewBuffer([]byte{})
	output := bytes.Buffer{}

	resolveGitRef = func(ref string) (string, error) {
		return "9934f4", nil
	}

	addCmnd := addCmd
	addCmnd.SetArgs([]string{"HEAD"})
	addCmnd.SetIn(input)
	addCmnd.SetOut(&output)
	addCmnd.SetErr(&output)

	err := addCmnd.Execute()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to read description")
}

func TestAddCommand_FailedToSaveNote(t *testing.T) {
	input := bytes.NewBuffer([]byte("This is a test note\n"))
	output := bytes.Buffer{}

	resolveGitRef = func(ref string) (string, error) {
		return "9934f4", nil
	}

	storage.SaveNotes = func(notes []storage.NotesMap) error {
		return errors.New("failed to save note")
	}

	addCmnd := addCmd
	addCmnd.SetArgs([]string{"HEAD"})
	addCmnd.SetIn(input)
	addCmnd.SetOut(&output)
	addCmnd.SetErr(&output)

	err := addCmnd.Execute()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to save note")
}
