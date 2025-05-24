/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	_ "maps"
	"math/rand/v2"

	"github.com/tomatoCoderq/gitnotes/internal/storage"
	"github.com/tomatoCoderq/gitnotes/internal/tools"

	"github.com/spf13/cobra"
)

func IntRange(min, max int) int {
	return rand.IntN(max-min) + min
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use: "list",
	Short: "\033[1mDisplay all notes in the system.\033[0m\n\n" +
		"This command outputs a list of all stored notes along with their associated references and titles.\n\n" +
		"\033[1mExamples:\033[0m\n" +
		"  \033[32mgitnotes list\033[0m\n" +
		"  \033[32mgitnotes list --limit 10\033[0m\n\n" +
		"\033[1mFlags:\033[0m\n" +
		"  \033[1;34m--limit\033[0m   (Optional) Maximum number of notes to display. If number is greater than actual number of notes then all notes are printed",
	Args: cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		limit, err := cmd.Flags().GetInt("limit")
		if err != nil {
			return fmt.Errorf("failing while getting limit: %v", err)
		}

		notes, err := storage.LoadNotes(limit, tools.GetHomePath())
		if err != nil {
			return fmt.Errorf("failing at getting all notes: %v", err)
		}

		for _, note := range notes {
			for ref, value := range note {

				randomColor := IntRange(30, 37)

				cmd.Printf("\n\033[1;%dm'%s'\033[0m\n", randomColor, ref)
				cmd.Printf("---\n")
				cmd.Printf("\033[3mTitle:\033[0m %s\n", value.Title)
				cmd.Printf("\033[3mDescription:\033[0m %s\n", value.Content)
				cmd.Printf("\033[3mCreated:\033[0m %s\n", value.CreatedAt.Format("2006-January-02 15:04"))
				if value.Tag != "DEFAULT" {
					cmd.Printf("\033[3mTag:\033[0m %s\n\n", value.Tag)
				}
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
