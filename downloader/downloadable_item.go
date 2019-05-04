package downloader


type DownloadableItem interface {
	Directory() string
	DownloadFilepath() string
	UrlFormat() string
}
