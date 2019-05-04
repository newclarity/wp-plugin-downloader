package downloads

import (
	"fmt"
	"path/filepath"
	"strings"
	"wp-plugin-downloader/global"
)

var NilZip = (*Zip)(nil)
var _ DownloadableItem = NilZip

type Zip struct {
	plugin  string
	basedir string
}

func (me *Zip) SetPlugin(plugin string) {
	me.plugin = plugin
}

func NewZip(args *global.Args) *Zip {
	return &Zip{
		plugin:  args.Plugin,
		basedir: strings.TrimRight(filepath.ToSlash(args.Basedir), "/"),
	}
}

func (me *Zip) Directory() string {
	dir := fmt.Sprintf("%s/%s",
		me.basedir,
		global.PluginSubdirectory,
	)
	return filepath.FromSlash(dir)
}

func (me *Zip) DownloadFilepath() string {
	dir := fmt.Sprintf("%s/%s/%s.zip",
		me.basedir,
		global.ZipSubdirectory,
		me.plugin,
	)
	return filepath.FromSlash(dir)
}

func (me *Zip) UnzipDir() string {
	dir := fmt.Sprintf("%s/%s", me.Directory(), me.plugin)
	return filepath.FromSlash(dir)
}

func (me *Zip) DownloadUrl() string {
	return fmt.Sprintf("http://downloads.wordpress.org/plugin/%s.latest-stable.zip?nostats=1",
		me.plugin,
	)
}
