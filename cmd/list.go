/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	_ "maps"
	"math/rand/v2"

	"github.com/tomatoCoderq/gitnotes/internal/models"
	"github.com/tomatoCoderq/gitnotes/internal/storage"
	_ "github.com/tomatoCoderq/gitnotes/internal/tools"

	"github.com/spf13/cobra"
)

const (
	longDescriptionList = "\033[1mShow a single note by its Git reference or title.\033[0m\n\n" +
		"Use this command to view the full content of a note. You can locate a note using its Git reference\n" +
		"(e.g., commit hash, tag, branch) or a unique title.\n\n" +
		"\033[1mExamples:\033[0m\n" +
		"  \033[32mgitnotes show --p ref abc123\033[0m\n" +
		"  \033[32mgitnotes show --p title  \"Fix login bug\"\033[0m\n\n" +
		"\033[1mFlags:\033[0m\n" +
		"  \033[1;34m--ref\033[0m      Git commit hash, tag, or branch name.\n" +
		"  \033[1;34m--title\033[0m    Title of the note."
)


func IntRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func printFullNote(cmd *cobra.Command, note models.Note, key string) {
	// Generating some color for ref
	randomColor := IntRange(30, 37)

	// Printing
	cmd.Printf("\n\033[1;%dm'%s'\033[0m\n", randomColor, key)
	cmd.Printf("---\n")
	cmd.Printf("\033[3mTitle:\033[0m %s\n", note.Title)
	cmd.Printf("\033[3mDescription:\033[0m %s\n", note. Content)
	cmd.Printf("\033[3mCreated:\033[0m %s\n", note.CreatedAt.Format("2006-January-02 15:04"))
	if note.Tag != "DEFAULT" {
		cmd.Printf("\033[3mTag:\033[0m %s\n", note.Tag)
	}
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use: "list",
	Short: longDescriptionList,
	Args: cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Parsing flags
		_, err := cmd.Flags().GetInt("limit")
		if err != nil {
			return fmt.Errorf("failing while getting limit: %v", err)
		}

		notes, err := storage.LoadAllNotesBolt(db)
		if err != nil {
			return fmt.Errorf("failing at getting all notes: %v", err)
		}

		cmd.Printf("Found notes: %d", len(notes))

		for key, notes := range notes {
			for _, note := range notes {
				printFullNote(cmd, note, key)
			}
		}

		return nil
	},
}

func init() {
	listCmd.Flags().IntP("limit", "l", 10, "Limit number of notes")
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
