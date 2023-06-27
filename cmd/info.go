package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"linkcli/app"
)

var infoCmd = &cobra.Command{
	Use:   "get",
	Short: "To get the information you need",
	Long:  `This command is what helps netcli know what you want to do`,
	Run: func(cmd *cobra.Command, args []string) {

		url, _ := cmd.Flags().GetString("url")
		num, _ := cmd.Flags().GetString("profile")
		if url != "" {
			data, err := app.GetProperties(url, num)
			fmt.Println(PrettyPrint(data))
			if err != nil {
				fmt.Println(err)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	infoCmd.PersistentFlags().String("url", "", "The link you want to profile.")
	infoCmd.PersistentFlags().String("profile", "1", "The number of times you want to profile the link.")
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
