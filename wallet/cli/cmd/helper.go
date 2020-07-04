package cmd

import (
	"fmt"
	"github.com/incwallet/wallet/debugtool"
)

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
var privateSeeds = []string{
	"12MZ4QiFoETNbdLKgRQWPMQMqsceWPKo71Jma9NzwvLTabpcDhn",
	"158ZGK5EHmoyrHEd8aA2HCaqbNQ4r45ZsnwL4Zh8mH8dueWHWs",
	"1mYRSzV7yigD7qNpuQwnyKeVMQcnenjSxAB1L8MEpDuT3RRbZc",
	"1G5Q9uGSxekPSgC1w1ZFaDJ8RxeYrekk2FtFLF33QCKNbg2V88",
	"1mYRSzV7yigD7qNpuQwnyKeVMQcnenjSxAB1L8MEpDuT3RRbZc",
	"12MZ4QiFoETNbdLKgRQWPMQMqsceWPKo71Jma9NzwvLTabpcDhn",
	"1G5Q9uGSxekPSgC1w1ZFaDJ8RxeYrekk2FtFLF33QCKNbg2V88",
	"1cQCTV1m33LxBKpNW2SisbuJfp5VcBSEau7PE5aD16gGLAN7eq",
	"1DRtHReKFQqPQzd689EsjivNgUScPBUwvbw8azgxaLRUBtmFL2",
	"12YCyPu6KyBToSaaQkw7hzbWkJnUi78DLkfvWokgi4dCtkbVusC",
}


func ConvertCoinVersion(tool *debugtool.DebugTool, privKey string) {
	fmt.Println("CONVERT COIN")
	b, _ := tool.SwitchCoinVersion(privKey)
	fmt.Println(string(b))
	fmt.Println("END CONVERT COIN")
}

func GetPRVOutPutCoin(tool *debugtool.DebugTool, privkey string) {
	fmt.Println("========== GET PRV OUTPUT COIN ==========")
	b, _ := tool.GetListOutputCoins(privkey)
	fmt.Println(string(b))
	fmt.Println("========== END GET PRV OUTPUT COIN ==========")
}

func GetPRVBalance(tool *debugtool.DebugTool, privkey string) {
	fmt.Println("========== GET PRV BALANCE ==========")
	b, _ := tool.GetBalanceByPrivatekey(privkey)
	fmt.Println(string(b))
	fmt.Println("========== END GET PRV BALANCE ==========")
}

func GetRawMempool(tool *debugtool.DebugTool) {
	fmt.Println("========== GET RAW MEMPOOL ==========")
	b, _ := tool.GetRawMempool()
	fmt.Println(string(b))
	fmt.Println("========== END GET RAW MEMPOOL ==========")
}

func GetTxByHash(tool *debugtool.DebugTool, txHash string) {
	fmt.Println("========== GET TX BY HASH ==========")
	b, _ := tool.GetTransactionByHash(txHash)
	fmt.Println(string(b))
	fmt.Println("========== END GET TX BY HASH ==========")
}

func Staking(tool *debugtool.DebugTool, privKey string, privSeed string) {
	fmt.Println("========== STAKING  ==========")
	b, _ := tool.Stake(privKey, privSeed)
	fmt.Println(string(b))
	fmt.Println("========== END STAKING ==========")
}

func UnStaking(tool *debugtool.DebugTool, privKey string, privSeed string) {
	fmt.Println("========== UN-STAKING  ==========")
	b, _ := tool.Unstake(privKey, privSeed)
	fmt.Println(string(b))
	fmt.Println("========== END UN-STAKING ==========")
}

func WithdrawReward(tool *debugtool.DebugTool, privKey string, tokenID string) {
	fmt.Println("========== WITHDRAW REWARD  ==========")
	b, _ := tool.WithdrawReward(privKey, tokenID)
	fmt.Println(string(b))
	fmt.Println("========== END WITHDRAW REWARD  ==========")
}

func SwitchPort(newPort string) *debugtool.DebugTool {
	tool := new(debugtool.DebugTool).InitLocal(newPort)
	return tool
}

func TransferPRV(tool *debugtool.DebugTool, fromPrivKey, toPrivKey, amount string) {
	fmt.Println("========== TRANSFER PRV  ==========")
	b, _ := tool.CreateAndSendTransactionFromAToB(fromPrivKey, toPrivKey, amount)
	fmt.Println(string(b))
	fmt.Println("========== END TRANSFER PRV  ==========")
}
