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

const (
	longDescriptionRm = "\033[1mRemove a note associated with a Git reference.\033[0m\n\n" +
		"This command permanently deletes the note from the notes storage (e.g., gitnotes.json).\n" +
		"Use it when a note is outdated or added by mistake.\n\n" +
		"\033[1mExamples:\033[0m\n" +
		"  \033[32mgitnotes remove abc123\033[0m\n" +
		"  \033[32mgitnotes remove feature/login\033[0m\n\n" +
		"\033[1mArguments:\033[0m\n" +
		"  \033[1;34mref\033[0m   Git commit hash, tag, or branch name to remove the note for.\n\n" +
		"\033[1mNote:\033[0m\n" +
		"  \033[31mThis action is irreversible.\033[0m Make sure you intend to delete the note."
)

func removeByParameter(ref, pureRef, parameter string, cmd *cobra.Command) error{
	if strings.ToLower(parameter) == "ref" {
		if err := storage.RemoveNotesByReferencBold(db, ref); err != nil {
			return fmt.Errorf("failed during removing note: %v", err)
		}
	}

	if strings.ToLower(parameter) == "title" {
		if err := storage.RemoveNotesByTitleBold(db, ref); err != nil {
			return fmt.Errorf("failed during removing note: %v", err)
		}
	}

	cmd.Printf("\033[1mSuccesfully removed note \033[32m%s\033[0m\033[0m\n", pureRef)
	return nil
}


// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove added note",
	Long: longDescriptionRm,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Parsing flags
		force, err := cmd.Flags().GetBool("force")
		if err != nil {
			return fmt.Errorf("failed during parsing force: %v", err)
		}

		parameter, err := cmd.Flags().GetString("parameter")
		if err != nil {
			return fmt.Errorf("failed during parsing parameter: %v", err)
		}

		var ref string

		ref = args[0]
		if strings.ToLower(parameter) == "ref" {
			ref, err = tools.ResolveGitRef(ref)
			if err != nil {
				return fmt.Errorf("failed during resolving reference: %v", err)
			}
		}
 
		pureRef, err := storage.GetRefFromNoteFields(db, ref)
		if err != nil {
			return fmt.Errorf("failed during getting ref: %v", err)
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
				return removeByParameter(ref, pureRef, parameter, cmd)
			} else {
				return nil
			}
		}
	
		return removeByParameter(ref, pureRef, parameter, cmd)		
	},
}

func init() {
	rmCmd.Flags().BoolP("force", "f", false, "Delete without additional approval")
	rmCmd.Flags().StringP("parameter", "p", "ref", "Define parameter to delete by: ref, title")

	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
