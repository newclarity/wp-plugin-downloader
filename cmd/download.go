package cmd

import (
	"fmt"
	"github.com/gearboxworks/go-status/is"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"wp-plugin-downloader/downloads"
	"wp-plugin-downloader/global"
	"wp-plugin-downloader/only"
)

var DownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download all WordPress plugins to local storage",
	Run: func(cmd *cobra.Command, args []string) {
		for range only.Once {
			fflags := cmd.Flags()
			what, err := fflags.GetString(global.WhatToDownloadFlag)
			if err != nil {
				fmt.Printf("could not access value of 'what' option: %s\n", err.Error())
			}
			if !global.ValidDownloadWhat.Contains(what) {
				fmt.Printf("'%s' is invalid for 'what'; must be one of: %s\n",
					what,
					strings.Join(global.ValidDownloadWhat, ", "),
				)
			}
			sts := downloads.EnsureDirs(global.DownloadDir, global.Subdirectories)
			if is.Error(sts) {
				fmt.Println(sts.Message())
				break
			}
			sts = downloads.Download(&global.Args{
				DownloadDomain: global.DownloadDomain,
				ListDomain:     global.ListDomain,
				Basedir:        global.DownloadDir,
				What:           global.WhatToDownload,
			})
			if is.Error(sts) {
				fmt.Println(sts.Message())
			}
		}
	},
}

func init() {
	fs := DownloadCmd.Flags()
	fs.StringVarP(&global.WhatToDownload, global.WhatToDownloadFlag, "", global.DownloadZip, "What to download?")
	fs.StringVarP(&global.DownloadDir, global.DownloadDirFlag, "", DefaultDownloadDir(), "Where to download?")
	fs.StringVarP(&global.DownloadDomain, global.SvnDownloadFlag, "", global.SvnDownloadDomain, "SVN domain to download from?")
	fs.StringVarP(&global.ListDomain, global.SvnListFlag, "", global.SvnListDomain, "SVN domain to download from?")
	RootCmd.AddCommand(DownloadCmd)
}

func DefaultDownloadDir() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	return fmt.Sprintf("%s/.wp-plugin-downloads", dir)
}
