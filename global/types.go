package global

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
