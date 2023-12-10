package cmd

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "use log to help you manage your saved colors",
	Long: ` You can use log to help you manage your saved colors, available commands:
$ clr-cli log list(aliases: l) to list all saved colors
$ clr-cli log clear(aliases: c) to clear all saved colors`,
	//Args:    cobra.MaximumNArgs(0),
	Example: "clr-cli log clear or clr-cli log list",
	Run:     logCmdFunc,
}

var clearCmd = &cobra.Command{
	Use:     "clear",
	Aliases: []string{"c"},
	Short:   "use clear or c to clear all saved colors",
	Long: ` You can use clear or c to clear all saved colors
$ clr-cli log clear(aliases: c)`,
	//Args:    cobra.MaximumNArgs(0),
	Example: "clr-cli log clear or clr-cli log c",
	Run: func(cmd *cobra.Command, args []string) {
		ColorData.Clear()
		err := ColorData.Save(ColorFile)
		if err != nil {
			cmd.Println("Error: can't clear the colors")
			return
		}
	},
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "use list or l to list all saved colors",
	Long: ` You can use list or l to list all saved colors
$ clr-cli log list(aliases: l)`,
	//Args:    cobra.MaximumNArgs(0),
	Example: "clr-cli log list or clr-cli log l",
	Run: func(cmd *cobra.Command, args []string) {
		ColorData.Print()
	},
}

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d"},
	Short:   "use delete or d to delete a saved color",
	Long: ` You can use delete or d to delete a saved color
$ clr-cli log delete(aliases: d) 1`,
	Args:    cobra.MinimumNArgs(1),
	Example: "clr-cli log delete [index] or clr-cli log d [index]",
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		err = ColorData.Delete(index)
		if err != nil {
			return
		}
		err = ColorData.Save(ColorFile)
		if err != nil {
			cmd.Println("Error: can't delete the color")
			return
		}
	},
}

var tokenCmd = &cobra.Command{
	Use:     "token",
	Aliases: []string{"t", "tk"},
	Short:   "use token to get your token",
	Long: ` You can use token to get your token
$ clr-cli log token`,
	Example: "clr-cli log token or clr-cli log t or clr-cli log tk",
	Run:     generateToken,
}

func generateToken(cmd *cobra.Command, args []string) {
	if len(*ColorData) == 0 {
		fmt.Println("Error: You don't have enough color data to generate a token")
		return
	}
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		cmd.Println("error while generating token")
		return
	}
	token := string(base64.StdEncoding.EncodeToString(b))
	// check and add color data to data collection
	DataCollection.Check(ColorData)
	TokenData = make(map[string][]Colors)
	TokenData[token] = *DataCollection
	postToken(TokenData)
	//for key, value := range TokenData {
	//	fmt.Printf("key[%s] value[%s]\n", key, value)
	//}
	log.Printf("\u001B[1m\u001B[48;5;128m[GENERATED TOKEN]\u001B[0m| \u001B[1m%-16s\n", token)
}

func postToken(tokenData map[string][]Colors) {
	postURL := os.Getenv("POST_URL")
	// restructuring the data
	var jsonStr []byte
	for token, colors := range tokenData {
		newData := map[string]interface{}{
			"colors_collection": colors,
			"cli_token":         token,
		}
		jsonStr, _ = json.Marshal(newData)
	}
	req, err := http.NewRequest("POST", postURL, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	log.Printf("\u001B[1m\u001B[48;5;22m[  SENT STATUS  ]\u001B[0m| \u001B[1m %-16s\u001B[0m", resp.Status)
}

func logCmdFunc(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Println("Invalid Flag, check the command description")
		return
	}
}
