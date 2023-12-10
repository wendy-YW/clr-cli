/*
Copyright © 2023 wyw
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	ColorFile = ".color.json"
)

var ColorData = &Colors{}
var DataCollection = &ColorsCollection{}
var TokenData map[string][]Colors

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "clr-cli",
	Short: "A command-line interface for color conversion",
	Long: `Color-Converter-CLI is a command-line interface (CLI) tool designed for color conversion. 
This application is a powerful tool that allows users to convert colors between different color models such as RGB, HEX.
For example, if you have a color in RGB format like rgb(255, 0, 0) and you want to convert it to HEX format, you can simply use our tool like this:
$ clr-cli toHex -r 255,0,0 or $ clr-cli 2x -rgb 255,0,0
And it will output the color in HEX format: #FF0000

Color-Converter-CLI also allows you to save the colors you have converted to a file and list them later.
$ clr-cli log list or $ clr-cli log l to list all the colors you have converted
You can also generate a token for your saved colors and they will be sent to our server, so you can see your saved colors from color converter website.
$ clr-cli log token or $ clr-cli log t to get your token

Color-Converter-CLI is built using Go and Cobra, a CLI library for Go that empowers applications.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	//read the file
	if err := ColorData.Read(ColorFile); err != nil {
		fmt.Print(fmt.Errorf("can't read the file"))
	}
	// rgb2hex command
	var rgb string
	rootCmd.AddCommand(toHexCmd)
	toHexCmd.Flags().StringVarP(&rgb, "rgb", "r", "", "RGB color")
	err := toHexCmd.MarkFlagRequired("rgb")
	// below isnt working
	toHexCmd.DisableFlagParsing = true
	if err != nil {
		fmt.Println("Error: rgb is invalid, check the command description ", err)
		return
	}
	// hex2rgb command
	var hex string
	rootCmd.AddCommand(toRgbCmd)
	toRgbCmd.Flags().StringVarP(&hex, "hex", "x", "", "HEX color")
	toRgbCmd.DisableFlagParsing = true
	// below isnt working
	err = toRgbCmd.MarkFlagRequired("hex")
	if err != nil {
		return
	}
	// log command
	rootCmd.AddCommand(logCmd, versionCmd)
	logCmd.AddCommand(clearCmd, listCmd, deleteCmd, tokenCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
