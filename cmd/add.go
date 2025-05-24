/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	_ "os"
	"strings"
	"time"

	"github.com/tomatoCoderq/gitnotes/internal/tools"
	"github.com/tomatoCoderq/gitnotes/internal/models"
	"github.com/tomatoCoderq/gitnotes/internal/storage"

	"github.com/spf13/cobra"

)

var resolveGitRef = tools.ResolveGitRef

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a note related to git commit or branch",
	Long: "\033[1mYou can use this command to annotate commits, tags, or branches with structured notes.\033[0m\n\n" +
	"\033[1mExamples:\033[0m\n" +
	"  \033[32mgitnotes add a ... \"Fix Bug\" ... \"This commit fixes...\"\033[0m\n" +
	"  \033[32mgitnotes add ab ... \"Start Login\" ... \"Initial login...\"\033[0m\n\n" +
	"\033[1mArguments:\033[0m\n" +
	"  \033[1;34mref\033[0m      Git commit hash...\n" +
	"  \033[1;34mtitle\033[0m    A short title...\n" +
	"  \033[1;34mcontent\033[0m  Detailed content...\n\n" +
	"\033[1mNote:\033[0m\n" +
	"  If a note with the same reference already exists, \033[31mit will not be overwritten.\033[0m",  
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		tag, err := cmd.Flags().GetString("tag")
		if err != nil {
			return fmt.Errorf("failed during parsing tag: %v", err)
		}

		valid_tags := map[string]bool {
			"TODO":true, "BUG":true, "INFO":true, "CRITICAL":true, "DEFAULT": true,
		}

		if !valid_tags[tag] && tag != "" {
			return fmt.Errorf("tag is invalid. Should be `TODO`, `BUG`, `INFO`, or `CRITICAL`")
		}

		ref := args[0]
	
		standardRef, err := resolveGitRef(ref)
		if err != nil {
			return fmt.Errorf("could not resolve reference: %s", ref)
		}

		cmd.Println("\033[1mWrite \033[42mtitle\033[0m \033[1mfor your note:\033[0m")
		reader := bufio.NewReader(cmd.InOrStdin())

		title, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read description: %v", err)
		}
		title = strings.TrimSpace(title)

		cmd.Println("\033[1mWrite \033[42mdescription\033[0m \033[1mfor your note:\033[0m")

		message, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read description: %v", err)
		}

		message = strings.TrimSpace(message)

		note := models.Note{
			Title: title,
			Content: message,
			CreatedAt: time.Now(),
			Tag: tag,
		}

		if err = storage.SaveNotes([]storage.NotesMap{{standardRef:note}}); err != nil {
			return fmt.Errorf("failed to save note: %v", err)
		}
		cmd.Println("Note \033[1madded\033[0m for", standardRef, "succesfully!")
		return nil
	},
}

func init() {
	addCmd.Flags().StringP("tag", "t", "DEFAULT", "Include tag to the specific note: `TODO`, `BUG`, `INFO`, `CRITICAL`")

	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
