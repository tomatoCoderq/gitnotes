/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"gitnotes/internal/storage"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show a note with of specific GIT commit or branch",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error{
		ref := args[0]
	
		standardRef, err := resolveGitRef(ref)
		if err != nil {
			return fmt.Errorf("could not resolve reference: %s", ref)
		}

		note, err := storage.FindByRef(standardRef)
		if err != nil {
			return fmt.Errorf("failed to find note by ref: %v", err)
		}

		cmd.Printf("Note:\n---\nTitle: %s\nDesc: %s\nCreate_at: %s\n", note.Title, note.Content, note.CreatedAt.String())
		return nil

	},
}

func init() {
	rootCmd.AddCommand(showCmd)


	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
