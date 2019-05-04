package global

const SvnDownloadDomain = "plugins.trac.wordpress.org"
const SvnListDomain = "svn.wp-plugins.org"

type Subdirectory = string

const (
	PluginSubdirectory    Subdirectory = "plugins"
	ReadmeSubdirectory    Subdirectory = "readmes"
	ZipSubdirectory       Subdirectory = "zips"
	ChangelogSubdirectory Subdirectory = "changelogs"
	MissingSubdirectory   Subdirectory = "missing"
	TempSubdirectory      Subdirectory = "temp"
)

var Subdirectories = []string{
	PluginSubdirectory,
	ReadmeSubdirectory,
	ZipSubdirectory,
	ChangelogSubdirectory,
	MissingSubdirectory,
	TempSubdirectory,
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
