package core

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Imager(cmd *cobra.Command, args []string) {

	input, _ := cmd.Flags().GetString("input")
	output, _ := cmd.Flags().GetString("output")
	width, _ := cmd.Flags().GetInt("width")
	height, _ := cmd.Flags().GetInt("height")

	fmt.Printf("Input: %s\nOutput: %s\nWidth: %d\nHeight: %d\n", input, output, width, height)

}
