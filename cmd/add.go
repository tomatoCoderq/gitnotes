/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	_ "os"
	"strings"
	"time"

	"gitnotes/internal/commands"
	"gitnotes/internal/models"
	"gitnotes/internal/storage"

	cc "github.com/ivanpirog/coloredcobra"


	"github.com/spf13/cobra"

)

var resolveGitRef = commands.ResolveGitRef

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a note to specific GIT commit or branch",
	Args: cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		ref := args[0]
	
		standardRef, err := resolveGitRef(ref)
		if err != nil {
			return fmt.Errorf("could not resolve reference: %s", ref)
		}

		cmd.Println("Write your note name and description:")
		reader := bufio.NewReader(cmd.InOrStdin())
		message, err := reader.ReadString('\n')

		cmd.Println("Is input redirected?", cmd.InOrStdin() != os.Stdin)
		
		if err != nil {
			return fmt.Errorf("failed to read description: %v", err)
		}

		message = strings.TrimSpace(message)


		note := models.Note{
			Title: "title",
			Content: message,
			CreatedAt: time.Now(),
		}

		if err = storage.SaveNotes([]storage.NotesMap{{standardRef:note}}); err != nil {
			return fmt.Errorf("failed to save note: %v", err)
		}
		cmd.Println("✅ Note added for", standardRef)
		return nil
	},
}

func init() {

	cc.Init(&cc.Config{
        RootCmd:       rootCmd,
        Headings:      cc.HiCyan + cc.Bold + cc.Underline,
        Commands:      cc.HiYellow + cc.Bold,
        Example:       cc.Italic,
        ExecName:      cc.Bold,
        Flags:         cc.Bold,
    })

	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
