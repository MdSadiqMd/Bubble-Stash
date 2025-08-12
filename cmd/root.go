package main

import (
	"path/filepath"

	homeDir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "secret",
	Short: "Secret is a API Key and other secrets manager",
}

var encodingKey string

func init() {
	RootCmd.PersistentFlags().StringVarP(&encodingKey, "key", "k", "", "key used in encoding and decoding secrets")
}

func secretsPath() string {
	home, _ := homeDir.Dir()
	return filepath.Join(home, ".secrets")
}
