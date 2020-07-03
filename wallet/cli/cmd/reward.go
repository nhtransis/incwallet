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

// rewardCmd represents the reward command
var rewardCmd = &cobra.Command{
	Use:   "reward",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.ParseInt(args[0], 10, 32)
		if err != nil {
			panic(err)
		}
		tokenID := "0000000000000000000000000000000000000000000000000000000000000004"
		if args[1] != "" && len(args[1]) > 0 {
			tokenID = args[1]
		}
		WithdrawReward(tool, privateKeys[index], tokenID)
	},
}
func WithdrawReward(tool *debugtool.DebugTool, privKey string, tokenID string) {
	fmt.Println("========== WITHDRAW REWARD  ==========")
	b, _ := tool.WithdrawReward(privKey, tokenID)
	fmt.Println(string(b))
	fmt.Println("========== END WITHDRAW REWARD  ==========")
}
func init() {
	rootCmd.AddCommand(rewardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rewardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rewardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
