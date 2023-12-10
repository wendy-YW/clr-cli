package cmd

import (
	"errors"
	"fmt"
	"github.com/crazy3lf/colorconv"
	"github.com/spf13/cobra"
)

var toRgbCmd = &cobra.Command{
	Use:     "toRgb",
	Aliases: []string{"2rgb", "2r", "torgb", "toRGB", "2rgb", "2RGB"},
	Short:   "use toRgb or other Aliases with -hex or -x to get rgb.",
	Long: ` You can use toRgb or other Aliases with -hex or -x to get rgb.
$ clr-cli toRgb(or aliases) -hex(shorthand: -x) FF0000 `,
	Args:    cobra.MinimumNArgs(2),
	Example: "clr-cli toRgb(or aliases) -hex(shorthand: -x) FF0000",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !(args[0] == "-hex" || args[0] == "-x" || args[0] == "-hsl" || args[0] == "-l") {
			fmt.Println("Invalid Flag, check the command description")
			return errors.New("invalid flag")
		} else if args[0] == "-hex" || args[0] == "-x" {
			hex2Rgb(cmd, args)
		}
		return nil
	},
}

func hex2Rgb(cmd *cobra.Command, args []string) {
	if !(args[0] == "-hex" || args[0] == "-x") {
		fmt.Println("Invalid Flag, check the command description")
		return
	}
	if len(args[1]) == 6 {
		r, g, b, err := colorconv.HexToRGB(args[1])
		if err != nil {
			fmt.Println("Error: hex is invalid")
			return
		}
		ColorData.Add(fmt.Sprintf("rgb(%d,%d,%d)", r, g, b), "#"+args[1])
		fmt.Printf("[INPUT]      \u001b[1m\u001b[41;1m R \u001b[42;1m G \u001b[44;1m B \u001B[0m| #%-16s | \033[38;2;%d;%d;%dm\u2588\u2588 \033[0m\n", args[1], r, g, b)
		fmt.Printf("[CONVERTED]  \u001b[1m\u001b[40;1m H \u001b[40;1m E \u001b[40;1m X \u001B[0m| \u001B[1m%-16s  | \u001B[38;2;%d;%d;%dm◀◀◀ \u001B[0m\n", fmt.Sprintf("rgb(%d,%d,%d)", r, g, b), r, g, b)
		err = ColorData.Save(ColorFile)
		if err != nil {
			fmt.Println("Error: can't save the color")
			return
		}
	} else {
		fmt.Println("Error: hex is invalid, check the command description, please dont use 0x prefix")
	}
}
