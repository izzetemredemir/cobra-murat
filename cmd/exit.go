package cmd

import (
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/izzetemredemir/cobramurat/util"
	"github.com/spf13/cobra"
)

// exitCmd represents the exit command
var exitCmd = &cobra.Command{
	Use:   "56",
	Short: "Ortalık çok karıştığı durumda bu komutu çalıştırmak gerekir",
	Long:  `Ortalık 56 olunca bu komutu çağırıp programı kapatabilirsin. `,
	RunE: func(cmd *cobra.Command, args []string) error {
		color.Red("Ortalık çok karıştı her yer 56")
		indicator := util.NewLoadingIndicator("Son durumlar kontrol ediliyor", 5)
		indicator.Start()
		time.Sleep((3 * time.Second))
		indicator.SetStep("56", 2)
		time.Sleep((3 * time.Second))
		indicator.SetStep("56", 4)

		indicator.Stop()

		color.Red("Her yer 56 program kapatılıyor")

		os.Exit(0)
		return nil
	},
}

func init() {

	rootCmd.AddCommand(exitCmd)
}
