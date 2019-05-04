package cmd

const (
	WhatToDownloadFlag = "what"
	DownloadDirFlag    = "dir"
	SvnDomainFlag      = "svn"
)

type DownloadWhat = string

const (
	DownloadAll    DownloadWhat = "all"
	DownloadReadme DownloadWhat = "readme"
)

type Strings []string

var ValidDownloadWhat = Strings{
	DownloadAll,
	DownloadReadme,
}

func (me Strings) Contains(s string) (contains bool) {
	for _, ss := range me {
		if ss != s {
			continue
		}
		contains = true
		break
	}
	return contains
}

const DefaultSvnDomain = "plugins.trac.wordpress.org"