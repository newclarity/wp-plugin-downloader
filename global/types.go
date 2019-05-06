package global

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

type Strings []string

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

type Args struct {
	What           string
	Plugin         string
	ListDomain     string
	DownloadDomain string
	Basedir        string
	FirstRev       int
	LastRev        int
}
