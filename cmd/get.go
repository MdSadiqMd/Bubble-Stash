package main

import (
	"fmt"

	"github.com/MdSadiqMd/Bubble-Stash/internal/vault"
	"github.com/spf13/cobra"
)

var getKeyFlag string

func init() {
	getCmd.Flags().StringVar(&getKeyFlag, "key", "", "Key to retrieve")
	getCmd.MarkFlagRequired("key")
	RootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the secret associated with the key",
	Run: func(cmd *cobra.Command, args []string) {
		v := vault.File(encodingKey, secretsPath())
		value, err := v.Get(getKeyFlag)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s=%s\n", getKeyFlag, value)
	},
}
