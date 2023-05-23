/*
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
	"github.com/spf13/cobra"
	"log"
	"encoding/json"
	sthingsCli "github.com/stuttgart-things/sthingsCli"
)

type NameUpdate struct {
	OldName    string    `json:"OldName"`
	NewName    string    `json:"NewName"`
}


// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		update_name()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func update_name() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if host == "" || port == ""{
		fmt.Println("Connection could not be created for lack of HOST or PORT env variable.")
	} else {
		httpurl := fmt.Sprintf("%s:%s/names", host, port)
		var select_name []string

		names := extract_names(httpurl)
		for _, name := range names {
			select_name = append(select_name, name.Name)
		}
		oldname := sthingsCli.AskSingleSelectQuestion("Select name to update :", select_name)
		newname := sthingsCli.AskSingleInputQuestion("Write new name :", oldname)

		data := NameUpdate{oldname, newname}
		jsonReq, err := json.Marshal(data)
		req, err := http.NewRequest(http.MethodPut, httpurl, bytes.NewBuffer(jsonReq))
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}

	defer resp.Body.Close()
    bodyBytes, _ := ioutil.ReadAll(resp.Body)

    // Convert response body to string
    bodyString := string(bodyBytes)
    fmt.Println(bodyString)
	}
}
