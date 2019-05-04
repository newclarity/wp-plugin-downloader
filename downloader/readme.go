package downloader

import (
	"fmt"
	"path/filepath"
	"strings"
)

var NilReadme = (*Readme)(nil)
var _ DownloadableItem = NilReadme

type Readme struct {
	plugin  string
	basedir string
}

func NewReadme(basedir string, plugin string) *Readme {
	return &Readme{
		plugin:  plugin,
		basedir: strings.TrimRight(filepath.ToSlash(basedir),"/"),
	}
}

func (me *Readme) Directory() string {
	dir := fmt.Sprintf("%s/%s",me.basedir,ReadmeSubdirectory)
	return filepath.FromSlash(dir)
}

func (me *Readme) DownloadFilepath() string {
	dir := fmt.Sprintf("%s/readmes/%s.readme", me.basedir, me.plugin)
	return filepath.FromSlash(dir)
}

func (me *Readme) UrlFormat() string {
	return fmt.Sprintf("http://plugins.svn.wordpress.org/%s/trunk/readme.txt",
		me.plugin,
	)

}
