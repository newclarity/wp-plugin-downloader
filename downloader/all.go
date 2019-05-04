package downloader

import (
	"fmt"
	"path/filepath"
	"strings"
)

var NilAll = (*All)(nil)
var _ DownloadableItem = NilAll

type All struct {
	plugin  string
	basedir string
}

func NewAll(basedir string, plugin string) *All {
	return &All{
		plugin:  plugin,
		basedir: strings.TrimRight(filepath.ToSlash(basedir),"/"),
	}
}

func (me *All) Directory() string {
	dir := fmt.Sprintf("%s/%s",me.basedir,PluginSubdirectory)
	return filepath.FromSlash(dir)
}

func (me *All) DownloadFilepath() string {
	dir := fmt.Sprintf("%s/zips/%s.zip", me.basedir, me.plugin)
	return filepath.FromSlash(dir)
}

func (me *All) UrlFormat() string {
	return fmt.Sprintf("http://downloads.wordpress.org/plugin/%s.latest-stable.zip?nostats=1",
		me.plugin,
	)
}
