package downloader

import (
	"fmt"
	"path/filepath"
	"strings"
)

var NilOther = (*Other)(nil)
var _ DownloadableItem = NilOther

type Other struct {
	plugin  string
	basedir string
}

func NewOther(basedir string, plugin string) *Other {
	return &Other{
		plugin:  plugin,
		basedir: strings.TrimRight(filepath.ToSlash(basedir),"/"),
	}
}

func (me *Other) Directory() string {
	dir := fmt.Sprintf("%s/%s",me.basedir,PluginSubdirectory)
	return filepath.FromSlash(dir)
}

func (me *Other) DownloadFilepath() string {
	dir := fmt.Sprintf("%s/zips/%s.zip", me.basedir, me.plugin)
	return filepath.FromSlash(dir)
}

func (me *Other) UrlFormat() string {
	return fmt.Sprintf("http://downloads.wordpress.org/plugin/%s.latest-stable.zip?nostats=1",
		me.plugin,
	)
}
