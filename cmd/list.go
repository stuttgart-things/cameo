/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd
// package main

import (
	"fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "encoding/json"
	"github.com/spf13/cobra"
	sthingsCli "codehub.sva.de/Lab/stuttgart-things/dev/sthingsCli"
    "github.com/jedib0t/go-pretty/v6/table"
)

type Names struct {
    Name string `mapstructure:"name"`
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
        list_command()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func list_command(){
    host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if host == "" || port == ""{
		fmt.Println("Connection could not be created for lack of HOST or PORT env variable.")
	} else {
        httpurl := fmt.Sprintf("%s:%s/names", host, port)
        operation := sthingsCli.AskSingleSelectQuestion("Desired operation:", []string{"All Names", "Random Names"})
        if operation == "All Names"{
            names := extract_names(httpurl)
            create_table(names)
        } else if operation == "Random Names"{
            name_amount := sthingsCli.AskSingleSelectQuestion("How many names should be displayed? :", []string{"1", "2", "3", "4", "5", "10"})
            for name_amount != "Exit"{
                names := extract_names(httpurl + "/" + name_amount)
                create_table(names)
                var more_names = sthingsCli.AskSingleSelectQuestion("Would you like to fetch more random names? :", []string{"Yes", "No"})
                if more_names == "Yes" {
                    name_amount = sthingsCli.AskSingleSelectQuestion("How many names should be displayed? :", []string{"1", "2", "3", "4", "5", "10"})
                } else {
                    name_amount = "Exit"
                }
            }
        }
    }
}


func extract_names(httpurl string) []Names {
    var names []Names

    response, err := http.Get(httpurl)
    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    json.Unmarshal(responseData, &names)

    return names

}

func create_table(names []Names){
    tw := table.NewWriter()

    for i, name := range names {
        if i==0{
            var header table.Row
            header = sthingsCli.CreateTableHeader(name)
            tw.AppendHeader(header)
        }
        tw.AppendRow(sthingsCli.CreateTableRows(name))
        tw.AppendSeparator()
        tw.SetStyle(table.StyleColoredBright)
        tw.SetOutputMirror(os.Stdout)
    }
    fmt.Println("\n")
    tw.Render()
    fmt.Println("\n")
}

