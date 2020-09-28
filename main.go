package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var nodeRpc string

func init() {
	rootCmd.Flags().StringVarP(&nodeRpc, "nodeRpc", "n", "tcp://127.0.0.1:26657", "monitor the node by the address")
}

var rootCmd = &cobra.Command{
	Use:   "node-monitor",
	Short: "monitor halle chain",
	Run: func(cmd *cobra.Command, args []string) {
		cli, err := NewClient(nodeRpc)
		if err != nil {
			panic(err)
		}
		cli.Monitor()
	},
}
