/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/incwallet/wallet/debugtool"
	"strconv"

	"github.com/spf13/cobra"
)

// outcoinCmd represents the outcoin command
var outcoinCmd = &cobra.Command{
	Use:   "outcoin",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			index, err := strconv.ParseInt(args[0], 10, 32)
			if err != nil {
				panic(err)
			}
			GetPRVOutPutCoin(tool, privateKeys[index])

		} else {
			GetPRVOutPutCoin(tool,privateKeys[0])
		}
	},
}
func GetPRVOutPutCoin(tool *debugtool.DebugTool, privkey string) {
	fmt.Println("========== GET PRV OUTPUT COIN ==========")
	b, _ := tool.GetListOutputCoins(privkey)
	fmt.Println(string(b))
	fmt.Println("========== END GET PRV OUTPUT COIN ==========")
}
func init() {
	rootCmd.AddCommand(outcoinCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// outcoinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// outcoinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
