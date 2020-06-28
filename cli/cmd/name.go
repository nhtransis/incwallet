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
	"encoding/json"
	"fmt"
	"github.com/incwallet/wallet/debugtool"
	"github.com/spf13/cobra"
	"io/ioutil"
)

// nameCmd represents the name command
var nameCmd = &cobra.Command{
	Use:   "name",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		test(args)
	},
}
var privateKeys = []string{
"112t8roafGgHL1rhAP9632Yef3sx5k8xgp8cwK4MCJsCL1UWcxXvpzg97N4dwvcD735iKf31Q2ZgrAvKfVjeSUEvnzKJyyJD3GqqSZdxN4or",
"112t8rnZDRztVgPjbYQiXS7mJgaTzn66NvHD7Vus2SrhSAY611AzADsPFzKjKQCKWTgbkgYrCPo9atvSMoCf9KT23Sc7Js9RKhzbNJkxpJU6",
"112t8rne7fpTVvSgZcSgyFV23FYEv3sbRRJZzPscRcTo8DsdZwstgn6UyHbnKHmyLJrSkvF13fzkZ4e8YD5A2wg8jzUZx6Yscdr4NuUUQDAt",
"112t8rnXoBXrThDTACHx2rbEq7nBgrzcZhVZV4fvNEcGJetQ13spZRMuW5ncvsKA1KvtkauZuK2jV8pxEZLpiuHtKX3FkKv2uC5ZeRC8L6we",
"112t8rnXWRThUTJQgoyH6evV8w19dFZfKWpCh8rZpfymW9JTgKPEVQS44nDRPpsooJiGStHxu81m3HA84t9DBVobz8hgBKRMcz2hddPWNX9N",
"112t8rnXfotCdLe7Gb8z13Qr2QqAneRYQPYN8faYTp8WLZCbd2JY2iSdB2qgxhhrkQ2PNZxrxZAj8944TjEEdjzPbhUTKUspxhA4vTFDW6aV",
"112t8rnX2FzscSBNqNAMuQfeSQhMPamSAKDX9f7X7Nft6JR4hfxNh5WJ8r9jeAmbmzTytW8hvpeXVn9FwLD8fjuvn4k7oHtXeh2YxnMDUzZv",
"112t8rnXoEWG5H8x1odKxSj6sbLXowTBsVVkAxNWr5WnsbSTDkRiVrSdPy8QfMujntKRYBqywKMJCyhMpdr93T3XiUD5QJR1QFtTpYKpjBEx",
"112t8rnqijhT2AqiS8NBVgifb86sqjfwQwf4MHLMAxK3gr1mwxMaeUWQtR1MfxHscrKQ2MsyQMvJ3LEu49LEcZzTzoJCkCiewQ9p6wN5SrG1",
"112t8rnud1R3of9rPkdKHWy8mQ5gMpXuBjLGhVrNurvHC93fF6qfiaEC8Nf7AHRbgrn1KF33akoNMNqUEUdSU7caXYvRL4uT58fhCuDV2Qs8",
"112t8rne7fpTVvSgZcSgyFV23FYEv3sbRRJZzPscRcTo8DsdZwstgn6UyHbnKHmyLJrSkvF13fzkZ4e8YD5A2wg8jzUZx6Yscdr4NuUUQDAt",
}
func init() {


	rootCmd.AddCommand(nameCmd)


	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func test(args [] string){
	tool := new(debugtool.DebugTool).InitDevNet("9334")
	//index, _ := strconv.ParseInt(args[1], 10, 32)
	GetPRVBalance(tool, privateKeys[0])
}
func hexToName(args []string) {
	var hexMap map[string]string

	// read the color.min.json
	content, err := ioutil.ReadFile("colornames.min.json")

	if err != nil {
		fmt.Printf("Error while reading the file %v", err)
	}
	_ = json.Unmarshal(content, &hexMap)

	name, ok := hexMap[args[0]]
	if ok {
		fmt.Printf("Name: %s, Hex: %s\n", name, args[0])
	} else {

		fmt.Println("Color name not found")
	}
}
func GetPRVBalance(tool *debugtool.DebugTool, privkey string) {
	fmt.Println("========== GET PRV BALANCE ==========")
	b, _ := tool.GetBalanceByPrivatekey(privkey)
	fmt.Println(string(b))
	fmt.Println("========== END GET PRV BALANCE ==========")
}