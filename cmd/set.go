package main

import (
	"fmt"

	"github.com/MdSadiqMd/Bubble-Stash/internal/vault"
	"github.com/spf13/cobra"
)

var keyFlag string
var valueFlag string

func init() {
	setCmd.Flags().StringVar(&keyFlag, "key", "", "Key to set")
	setCmd.Flags().StringVar(&valueFlag, "value", "", "Value to set")
	setCmd.MarkFlagRequired("key")
	setCmd.MarkFlagRequired("value")
	RootCmd.AddCommand(setCmd)
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a secret in your storage",
	Run: func(cmd *cobra.Command, args []string) {
		v := vault.File(encodingKey, secretsPath())
		if err := v.Set(keyFlag, valueFlag); err != nil {
			panic(err)
		}
		fmt.Printf("%s = %s\n", keyFlag, valueFlag)
	},
}
