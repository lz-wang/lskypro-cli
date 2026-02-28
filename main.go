package main

import (
	"fmt"
	"os"

	"github.com/lzwang/lskypro-cli/internal/cmd"
)

var version = "dev"

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		os.Exit(1)
	}
}
