package global

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

const SvnDownloadDomain = "plugins.trac.wordpress.org"
const SvnListDomain = "svn.wp-plugins.org"

type Subdirectory = string

const (
	PluginSubdirectory    Subdirectory = "plugins"
	ReadmeSubdirectory    Subdirectory = "readmes"
	ZipSubdirectory       Subdirectory = "zips"
	ChangelogSubdirectory Subdirectory = "changelogs"
	MissingSubdirectory   Subdirectory = "missing"
)

var Subdirectories = []string{
	PluginSubdirectory,
	ReadmeSubdirectory,
	ZipSubdirectory,
	ChangelogSubdirectory,
	MissingSubdirectory,
}

const (
	WhatToDownloadFlag = "what"
	DownloadDirFlag    = "dir"
	SvnDownloadFlag    = "download-url"
	SvnListFlag        = "list-url"
)

type DownloadWhat = string

const (
	DownloadZip    DownloadWhat = "zip"
	DownloadReadme DownloadWhat = "readme"
)

var ValidDownloadWhat = Strings{
	DownloadZip,
	DownloadReadme,
}
