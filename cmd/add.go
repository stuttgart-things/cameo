/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"fmt"
	"os"
	"io/ioutil"
	"net/http"
	"log"
	"encoding/json"
	"github.com/spf13/cobra"
	sthingsCli "github.com/stuttgart-things/sthingsCli"
)

type NewNames struct {
	NewName    string   `json:"Name"`
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		new_name_inquiry() 
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func new_name_inquiry(){
	selection := "yes"
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if host == "" || port == ""{
		fmt.Println("Connection could not be created for lack of HOST or PORT env variable.")
	} else {
		httpurl := fmt.Sprintf("%s:%s/names", host, port)
		for selection == "yes" {
			var new_name = sthingsCli.AskSingleInputQuestion("Write desired Name", "")
			var response = add_name_api(httpurl, new_name)
			fmt.Println(response, "\n")
			selection = sthingsCli.AskSingleSelectQuestion("Add New Name:", []string{"yes", "no"} )
		}
	}
}

func add_name_api(httpurl string, new_name string) string {
    names := NewNames{new_name}
    jsonReq, err := json.Marshal(names)
    resp, err := http.Post(httpurl, "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
    if err != nil {
        log.Fatalln(err)
    }

    defer resp.Body.Close()
    bodyBytes, _ := ioutil.ReadAll(resp.Body)

    // Convert response body to string
    bodyString := string(bodyBytes)

	return bodyString

}