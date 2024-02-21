package install

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"
	"github.com/yusufcanb/tlama/pkg/config"
	"log"
)

func GetCommand() *cli.Command {
	return &cli.Command{
		Name:    "install",
		Aliases: []string{"i"},
		Usage:   "Install LLM to your system.",
		Action: func(c *cli.Context) error {
			cfg := c.App.Metadata["config"].(*config.TlamaConfig)
			ollama := cfg.GetOllamaApi()

			model := initialModel(&initialModelArgs{alreadyInstalled: ollama.IsInstalled()})
			program := tea.NewProgram(model)
			_, err := program.Run()
			if err != nil {
				log.Fatalf("could not run program: %s", err)
			}
			defer program.Quit()

			if model.questions[len(model.questions)-1].answer == false {
				fmt.Println("\nAbort...")
				return nil
			}

			err = ollama.Install()
			if err != nil {
				log.Fatalf("could not install LLM: %s", err.Error())
			}

			fmt.Println("\n\nInstallation complete.")
			return nil
		},
	}
}