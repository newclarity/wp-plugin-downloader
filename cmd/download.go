package cmd

import (
	"fmt"
	"github.com/newclarity/wp-plugin-downloader/options"
	"strings"

	"github.com/gearboxworks/go-status/is"
	"github.com/newclarity/wp-plugin-downloader/downloader"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var DownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download all WordPress plugins to local storage",
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags()
		what, err := fflags.GetString(WhatToDownloadFlag)
		if err != nil {
			fmt.Printf("could not access value of 'what' option: %s\n", err.Error())
		}
		if !ValidDownloadWhat.Contains(what) {
			fmt.Printf("'%s' is invalid for 'what'; must be one of: %s\n",
				what,
				strings.Join(ValidDownloadWhat,", "),
			)
		}
		downloader.EnsureDir(options.DownloadDir)
		sts := downloader.Download(what)
		if is.Error(sts) {
			fmt.Println(sts.Message())
		}
	},
}

func init() {
	fs := DownloadCmd.Flags()
	fs.StringVarP(&options.WhatToDownload, WhatToDownloadFlag, "", DownloadAll, "What to download?")
	fs.StringVarP(&options.DownloadDir, DownloadDirFlag, "", DefaultDownloadDir(), "Where to download?")
	fs.StringVarP(&options.SvnDomain, SvnDomainFlag, "", DefaultSvnDomain, "SVN domain to download from?")
	RootCmd.AddCommand(DownloadCmd)
}

func subdirFilepath(sd string) string {
	return filepath.FromSlash(fmt.Sprintf("%s/%s", options.DownloadDir, sd))
}

func DefaultDownloadDir() string {
	dir, err := os.Getwd()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	return fmt.Sprintf("%s/downloads", dir)
}

/*
#!/usr/bin/php
<?php

foreach( array( 'plugins','readmes','zips','changelogs','missing','temp' ) as $subdir )
  if ( ! is_dir( "{$download_dir}{$subdir}" ) )
    mkdir( "{$download_dir}{$subdir}" );

$wget_dir = rtrim( getenv( 'wp_plugin_downloader_wget_dir' ), '/' );
if ( empty( $wget_dir ) )
  $wget_dir = '/usr/local/bin';

switch ( $type ) {
	case 'readme':
		$directory = "{$download_dir}readmes";
    	$download = "{$download_dir}readmes/%s.readme";
		$url = 'http://plugins.svn.wordpress.org/%s/trunk/readme.txt';
		break;
	case 'all':
		$directory = "{$download_dir}plugins";
		$download = "{$download_dir}zips/%s.zip";
		$url = 'http://downloads.wordpress.org/plugin/%s.latest-stable.zip?nostats=1';
		break;
	default:
		echo $cmd . ": invalid command\r\n";
		echo 'Usage: php ' . $cmd . " [command]\r\n\r\n";
		echo "Available commands:\r\n";
		echo "  all - Downloads full plugin zips\r\n";
		echo "  readme - Downloads plugin readmes only\r\n";
		die();
}
*/