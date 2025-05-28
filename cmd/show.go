/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/tomatoCoderq/gitnotes/internal/storage"
	"github.com/spf13/cobra"
)

const (
	longDescriptionShow = "\033[1mShow a single note by its Git reference or title.\033[0m\n\n" +
		"Use this command to view the full content of a note. You can locate a note using its Git reference\n" +
		"(e.g., commit hash, tag, branch) or a unique title.\n\n" +
		"\033[1mExamples:\033[0m\n" +
		"  \033[32mgitnotes show --p ref abc123\033[0m\n" +
		"  \033[32mgitnotes show --p title  \"Fix login bug\"\033[0m\n\n" +
		"\033[1mFlags:\033[0m\n" +
		"  \033[1;34m--ref\033[0m      Git commit hash, tag, or branch name.\n" +
		"  \033[1;34m--title\033[0m    Title of the note."
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show a note with of specific GIT commit or branch",
	Long: longDescriptionShow,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var notesMap storage.NotesMap
		parameter, err := cmd.Flags().GetString("parameter")
		if err != nil {
			return fmt.Errorf("failed at parsing parameter %v", err)
		}

		ref := args[0]

		switch parameter {
		case "ref":
			ref, err := resolveGitRef(ref)
			if err != nil {
				return fmt.Errorf("could not resolve reference: %s", ref)
			}
			notesMap, err = storage.FindByRef(db, ref)
			if err != nil {
				return fmt.Errorf("failed to find note by ref: %v", err)
			}
		case "title":
			// notesMap, err = storage.Find(db, ref)
		case "tag":
			// notesMap, err = storage.FindByTag(db, ref)
		}

		if err != nil {
			return fmt.Errorf("failed to find note by ref: %v", err)
		}

		
		for key, notes := range notesMap {
			for _, note := range notes {
				printFullNote(cmd, note, key)
			}
		}

		return nil
	},
}

func init() {
	showCmd.Flags().StringP("parameter", "p", "ref", "Define parameter to find notes: title, ref, tag")

	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
