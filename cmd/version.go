package cmd

import (
	"fmt"

	"github.com/fatih/color"

	"github.com/spf13/cobra"
	goVersion "go.hein.dev/go-version"
)

const banner = `CAMEO`

var (
	shortened  = false
	version    = "unset"
	commit     = "unknown"
	date       = "unknown"
	output     = "yaml"
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "version will output the current build information",
		Long: `Print the version information. For example:
	sthings version`,

		Run: func(_ *cobra.Command, _ []string) {
			resp := goVersion.FuncWithOutput(shortened, version, commit, date, output)
			color.White(banner)
			fmt.Print(resp)
		},
	}
)