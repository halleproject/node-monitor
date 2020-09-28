package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	Execute()
}

var nodeRpc string

func init() {
	rootCmd.Flags().StringVarP(&nodeRpc, "nodeRpc", "n", "tcp://192.168.3.200:26657", "monitor the node by the address")
}

var rootCmd = &cobra.Command{
	Use:   "node-monitor",
	Short: "monitor halle chain",
	Run: func(cmd *cobra.Command, args []string) {
		cli := NewClient(nodeRpc)
		cli.Monitor()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
