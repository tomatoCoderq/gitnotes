package cmd

import (
	"bytes"
	"fmt"
	"testing"

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
