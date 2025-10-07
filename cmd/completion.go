package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate shell completion script",
	Long: `Generate shell completion script for DevPreflight.

To load completions:

Bash:
  $ source <(devpreflight completion bash)
  
  # To load completions for each session, execute once:
  # Linux:
  $ devpreflight completion bash > /etc/bash_completion.d/devpreflight
  # macOS:
  $ devpreflight completion bash > /usr/local/etc/bash_completion.d/devpreflight

Zsh:
  # If shell completion is not already enabled in your environment,
  # you will need to enable it. You can execute the following once:
  $ echo "autoload -U compinit; compinit" >> ~/.zshrc
  
  # To load completions for each session, execute once:
  $ devpreflight completion zsh > "${fpath[1]}/_devpreflight"
  
  # You will need to start a new shell for this setup to take effect.

Fish:
  $ devpreflight completion fish | source
  
  # To load completions for each session, execute once:
  $ devpreflight completion fish > ~/.config/fish/completions/devpreflight.fish

PowerShell:
  PS> devpreflight completion powershell | Out-String | Invoke-Expression
  
  # To load completions for every new session, run:
  PS> devpreflight completion powershell > devpreflight.ps1
  # and source this file from your PowerShell profile.`,
	Example: `  # Generate bash completion
  devpreflight completion bash > /etc/bash_completion.d/devpreflight

  # Generate zsh completion
  devpreflight completion zsh > "${fpath[1]}/_devpreflight"

  # Generate fish completion
  devpreflight completion fish > ~/.config/fish/completions/devpreflight.fish`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
