/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/tomatoCoderq/gitnotes/internal/models"
	"github.com/tomatoCoderq/gitnotes/internal/storage"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show a note with of specific GIT commit or branch",
	Long: "\033[1mShow a single note by its Git reference or title.\033[0m\n\n" +
		"Use this command to view the full content of a note. You can locate a note using its Git reference\n" +
		"(e.g., commit hash, tag, branch) or a unique title.\n\n" +
		"\033[1mExamples:\033[0m\n" +
		"  \033[32mgitnotes show --p ref abc123\033[0m\n" +
		"  \033[32mgitnotes show --p title  \"Fix login bug\"\033[0m\n\n" +
		"\033[1mFlags:\033[0m\n" +
		"  \033[1;34m--ref\033[0m      Git commit hash, tag, or branch name.\n" +
		"  \033[1;34m--title\033[0m    Title of the note.",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		parameter, err := cmd.Flags().GetString("parameter")
		if err != nil {
			return fmt.Errorf("failed at parsing parameter %v", err)
		}

		ref := args[0]

		var notes []models.Note

		switch parameter {
		case "ref":
			ref, err := resolveGitRef(ref)
			if err != nil {
				return fmt.Errorf("could not resolve reference: %s", ref)
			}
			notes, err = storage.FindByParameter("ref", ref)
		case "title":
			notes, err = storage.FindByParameter("title", ref)
		case "tag":
			notes, err = storage.FindByParameter("tag", ref)
		}

		if err != nil {
			return fmt.Errorf("failed to find note by ref: %v", err)
		}

		randomColor := IntRange(30, 37)

		for _, note := range notes {
			cmd.Printf("\033[1;%dm'%s'\033[0m\n", randomColor, ref)
			cmd.Printf("---\n")
			cmd.Printf("\033[3mTitle:\033[0m %s\n", note.Title)
			cmd.Printf("\033[3mDescription:\033[0m %s\n", note.Content)
			if note.Tag != "DEFAULT" {
				cmd.Printf("\033[3mCreated:\033[0m %s\n", note.CreatedAt.Format("2006-January-02 15:04"))
				cmd.Printf("\033[3mTag:\033[0m %s\n\n", note.Tag)
			} else {
				cmd.Printf("\033[3mCreated:\033[0m %s\n\n", note.CreatedAt.Format("2006-January-02 15:04"))
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
