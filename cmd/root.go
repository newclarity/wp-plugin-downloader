package cmd

import (
	"github.com/spf13/cobra"
)

var NoCache bool

var RootCmd = &cobra.Command{
	Use:   "wp plugin downloader",
	Short: "Download all WordPress plugins to local storage",
}

func init() {
	pf := RootCmd.PersistentFlags()
	pf.BoolVarP(&NoCache, "no-cache", "", false, "Disable caching")
}
