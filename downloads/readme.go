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
