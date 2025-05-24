/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/tomatoCoderq/gitnotes/internal/storage"
	"github.com/tomatoCoderq/gitnotes/internal/tools"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove added note",
	Long: "\033[1mRemove a note associated with a Git reference.\033[0m\n\n" +
		"This command permanently deletes the note from the notes storage (e.g., gitnotes.json).\n" +
		"Use it when a note is outdated or added by mistake.\n\n" +
		"\033[1mExamples:\033[0m\n" +
		"  \033[32mgitnotes remove abc123\033[0m\n" +
		"  \033[32mgitnotes remove feature/login\033[0m\n\n" +
		"\033[1mArguments:\033[0m\n" +
		"  \033[1;34mref\033[0m   Git commit hash, tag, or branch name to remove the note for.\n\n" +
		"\033[1mNote:\033[0m\n" +
		"  \033[31mThis action is irreversible.\033[0m Make sure you intend to delete the note.",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		force, err := cmd.Flags().GetBool("force")
		if err != nil {
			return fmt.Errorf("failed during parsing force: %v", err)
		}

		standardRef, err := tools.ResolveGitRef(args[0])
		if err != nil {
			return fmt.Errorf("failed during resolving reference: %v", err)
		}

		if !force {
			cmd.Println("Are you sure [Y/n]:")
			reader := bufio.NewReader(cmd.InOrStdin())

			answer, err := reader.ReadString('\n')
			answer = strings.TrimSpace(answer)

			if err != nil {
				return fmt.Errorf("failed to read description: %v", err)
			}

			if strings.Compare(strings.ToLower(answer), "y") == 0 || strings.Compare(strings.ToLower(answer), "yes") == 0 {
				err = storage.RemoveNotesByReference(standardRef, tools.GetHomePath())
				if err != nil {
					return fmt.Errorf("failed during removing note: %v", err)
				}

				cmd.Printf("\033[1mSuccesfully removed note \033[32m%s\033[0m\033[0m\n", standardRef)
				return nil
			} else {
				return nil
			}
		} else {
			err = storage.RemoveNotesByReference(standardRef, tools.GetHomePath())
			if err != nil {
				return fmt.Errorf("failed during removing note: %v", err)
			}

			cmd.Printf("\033[1mSuccesfully removed note \033[32m%s\033[0m\033[0m\n", standardRef)
			return nil
		}
	},
}

func init() {
	rmCmd.Flags().BoolP("force", "f", false, "Delete without additional approval")

	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
