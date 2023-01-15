package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// kobraCmd represents the kobra command
var kobraCmd = &cobra.Command{
	Use:   "kobra",
	Short: "Kobrayı gösterir",
	Long:  `Kobrayı Gösterir`,
	Run: func(cmd *cobra.Command, args []string) {

		file, err := os.Open("cmd/cobra.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		// Dosyadaki tüm verileri oku
		fileStat, err := file.Stat()
		if err != nil {
			fmt.Println(err)
			return
		}

		data := make([]byte, fileStat.Size())
		count, err := io.ReadFull(file, data)
		if err != nil {
			fmt.Println(err)
			return
		}

		str := string(data[:count])
		fmt.Println(str)
	},
}

func init() {
	rootCmd.AddCommand(kobraCmd)

}
