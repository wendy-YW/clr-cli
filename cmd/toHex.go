package cmd

import (
	"errors"
	"fmt"
	"github.com/crazy3lf/colorconv"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

var toHexCmd = &cobra.Command{
	Use:     "toHex",
	Aliases: []string{"2hx", "2x", "tohex", "toHEX", "2hex", "2HEX"},
	Short:   "use toHex or other Aliases with -rgb or -r to get hex.",
	Long: ` You can use toHex or other Aliases with -rgb or -r to get hex.
$ clr-cli toHex(or aliases) -rgb(shorthand: -r) 255,0,0 `,
	Args:    cobra.MinimumNArgs(2),
	Example: "clr-cli toHex(or aliases) -rgb(shorthand: -r) 255,0,0",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !(args[0] == "-rgb" || args[0] == "-r" || args[0] == "-hsl" || args[0] == "-l") {
			fmt.Println("Invalid Flag, check the command description")
			return errors.New("invalid flag")
		} else if args[0] == "-rgb" || args[0] == "-r" {
			rgb2Hex(cmd, args)
		}
		return nil
	},
}

func rgb2Hex(cmd *cobra.Command, args []string) {
	//fmt.Printf("Your \u001b[1m\u001b[41;1m R \u001b[42;1m G \u001b[44;1m B \u001b[0m: %s\n ", args[1])
	//fmt.Println("\033[38;2;0;0;255m Test \u001B[0m")
	// 058D
	if !(args[0] == "-rgb" || args[0] == "-r") {
		fmt.Println("Invalid Flag, check the command description")
		return
	}
	if len(strings.Split(args[1], ",")) == 3 {
		r64, err := strconv.ParseUint(strings.Split(args[1], ",")[0], 10, 8)
		if err != nil {
			fmt.Println("Error: r is invalid")
			return
		}
		r := uint8(r64)
		g64, err := strconv.ParseUint(strings.Split(args[1], ",")[1], 10, 8)
		if err != nil {
			fmt.Println("Error: g is invalid")
			return
		}
		g := uint8(g64)
		b64, err := strconv.ParseUint(strings.Split(args[1], ",")[2], 10, 8)
		if err != nil {
			fmt.Println("Error: b is invalid")
			return
		}
		b := uint8(b64)
		// store the data to json
		ColorData.Add(args[1], "#"+colorconv.RGBToHex(r, g, b))
		fmt.Printf("[INPUT]     \u001b[1m\u001b[41;1m R \u001b[42;1m G \u001b[44;1m B \u001B[0m| %-17s | \033[38;2;%d;%d;%dm\u2588\u2588 \033[0m\n", args[1], r, g, b)
		fmt.Printf("[CONVERTED] \u001B[1m\u001B[40;1m H \u001B[40;1m E \u001B[40;1m X \u001B[0m| \u001B[1m#%-16s | \u001B[38;2;%d;%d;%dm\u25c0\u25c0\u25c0 \u001B[0m\n", colorconv.RGBToHex(r, g, b), r, g, b)
		err = ColorData.Save(ColorFile)
		if err != nil {
			fmt.Println("Error: can't save the color")
			return
		}
	} else {
		fmt.Println("Error: rgb is invalid, check the command description")
	}
}
