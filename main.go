package main

import (
	"fmt"
	"os"
	"wp-plugin-downloader/cmd"
)

//
// Name: WP Plugin Downloader
// Version 1.0
// Author: Mike Schinkel <mike@newclarity.net>
//

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
