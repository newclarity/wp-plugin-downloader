package downloader

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


