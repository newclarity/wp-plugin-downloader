package cmd

import (
	"github.com/spf13/cobra"
)

//
// Name: WP Plugin Downloader
//
// Copyright (C) 2019 NewClarity Consulting LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
//

var NoCache bool

var RootCmd = &cobra.Command{
	Use:   "wp plugin downloader",
	Short: "Download all WordPress plugins to local storage",
}

func init() {
	pf := RootCmd.PersistentFlags()
	pf.BoolVarP(&NoCache, "no-cache", "", false, "Disable caching")
}
