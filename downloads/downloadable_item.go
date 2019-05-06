package downloads

import "wp-plugin-downloader/global"

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

type DownloadableItem interface {
	Directory() string
	DownloadFilepath() string
	UnzipDir() string
	DownloadUrl() string
	SetPlugin(string)
}

func MakeDownloadableItem(args *global.Args) (item DownloadableItem) {
	switch args.What {
	case global.DownloadZip:
		item = NewZip(args)
	case global.DownloadReadme:
		item = NewReadme(args)
	}
	return item
}
