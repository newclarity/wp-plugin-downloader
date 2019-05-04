package downloads

import (
	"fmt"
	"path/filepath"
	"strings"
	"wp-plugin-downloader/global"
)

var NilReadme = (*Readme)(nil)
var _ DownloadableItem = NilReadme

type Readme struct {
	plugin  string
	basedir string
}

func (me *Readme) SetPlugin(plugin string) {
	me.plugin = plugin
}

func NewReadme(args *global.Args) *Readme {
	return &Readme{
		plugin:  args.Plugin,
		basedir: strings.TrimRight(filepath.ToSlash(args.Basedir), "/"),
	}
}

func (me *Readme) Directory() string {
	dir := fmt.Sprintf("%s/%s", me.basedir, global.ReadmeSubdirectory)
	return filepath.FromSlash(dir)
}

func (me *Readme) DownloadFilepath() string {
	dir := fmt.Sprintf("%s/readmes/%s.readme", me.basedir, me.plugin)
	return filepath.FromSlash(dir)
}

func (me *Readme) UnzipDir() string {
	dir := fmt.Sprintf("%s/%s", me.Directory(), me.plugin)
	return filepath.FromSlash(dir)
}

func (me *Readme) DownloadUrl() string {
	return fmt.Sprintf("http://plugins.svn.wordpress.org/%s/trunk/readme.txt",
		me.plugin,
	)

}
