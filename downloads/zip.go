package downloads

import (
	"fmt"
	"path/filepath"
	"strings"
	"wp-plugin-downloader/global"
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
