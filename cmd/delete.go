/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
    "encoding/json"
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
	"github.com/spf13/cobra"
	sthingsCli "codehub.sva.de/Lab/stuttgart-things/dev/sthingsCli"
)

type NameDelete struct {
	DelName    string    `json:"name"`
}

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
		delete_multiple_names()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func delete_multiple_names(){
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
        select_name = sthingsCli.AskMultiSelectQuestion("Which names would you like to delete?", select_name)

        for _, name := range select_name{
            delete_name_api(httpurl, name)
        }
    }
}
func delete_name_api(httpurl string, name string){
    delname := NameDelete{name}
    jsonReq, err := json.Marshal(delname)
    req, err := http.NewRequest(http.MethodDelete, httpurl, bytes.NewBuffer(jsonReq))
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
    }

    defer resp.Body.Close()
    bodyBytes, _ := ioutil.ReadAll(resp.Body)

    // Convert response body to string
    bodyString := string(bodyBytes)
    fmt.Println(bodyString)
}
