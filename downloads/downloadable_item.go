package downloads

import "wp-plugin-downloader/global"

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
