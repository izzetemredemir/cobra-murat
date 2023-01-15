package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Youtube'dan bir şarkı çalar",
	Long:  `Youtube'dan bir şarkı çalar`,
	Run: func(cmd *cobra.Command, args []string) {

		url := "https://www.youtube.com/watch?v=vGDFWcUTM8s"

		// Open the URL in a browser
		browser := exec.Command("open", url)
		err := browser.Start()
		if err != nil {
			fmt.Println("An error occurred: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(playCmd)

}
