/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"gitnotes/internal/storage"
	_"maps"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error{
		limit, err := cmd.Flags().GetInt("limit")
		if err != nil {
			return fmt.Errorf("failing while getting limit: %v", err)
		}

		notes, err := storage.LoadNotes(limit)
		if err != nil {
			return fmt.Errorf("failing at getting all notes: %v", err)
		}
		

		for _, note := range notes {
			for ref, value := range note {
				cmd.Printf("'%s'\n", ref)
				cmd.Printf("---\n")
				cmd.Printf("Title: %s\n", value.Title)
				cmd.Printf("Content: %s\n", value.Content)
				cmd.Printf("Created: %s\n\n", value.CreatedAt.Format("2006-01-02 15:04"))
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
