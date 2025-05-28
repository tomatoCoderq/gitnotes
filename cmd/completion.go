package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(completionCmd)
}

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate shell completion script",
	Long: `To load completions:

Bash:

$ source <(gitnotes completion bash)

# To load completions for each session, execute once:
Linux:
  $ gitnotes completion bash > /etc/bash_completion.d/gitnotes
MacOS:
  $ gitnotes completion bash > /usr/local/etc/bash_completion.d/gitnotes

Zsh:

# If shell completion is not already enabled in your environment,
# you will need to enable it. You can execute the following once:

$ echo "autoload -U compinit; compinit" >> ~/.zshrc

$ gitnotes completion zsh > "${fpath[1]}/gitnotes"

Restart: 
$ source ~/.zshrc

Fish:

$ gitnotes completion fish | source

# To load completions for each session, execute once:
$ gitnotes completion fish > ~/.config/fish/completions/gitnotes.fish
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			rootCmd.GenBashCompletion(os.Stdout)
		case "zsh":
			rootCmd.GenZshCompletion(os.Stdout)
		case "fish":
			rootCmd.GenFishCompletion(os.Stdout, true)
		case "powershell":
			rootCmd.GenPowerShellCompletion(os.Stdout)
		}
	},
}
