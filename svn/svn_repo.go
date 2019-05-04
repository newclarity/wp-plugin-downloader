package svn

import (
	"fmt"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/is"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
	"wp-plugin-downloader/global"
	"wp-plugin-downloader/only"
	"wp-plugin-downloader/util"
)

const LogUrlFormat = "http://%s/log"
const LatestRevisionUrlFormat = LogUrlFormat + "/?format=changelog&stop_rev=HEAD"
const ChangeLogUrlFormat = LogUrlFormat + "/?verbose=on&mode=follow_copy&format=changelog&rev=%d&limit=%d"

type Repos []*Repo
type Repo struct {
	listdomain     string
	downloaddomain string
	basedir        string
}

func NewRepo(args *global.Args) *Repo {
	return &Repo{
		listdomain:     args.ListDomain,
		downloaddomain: args.DownloadDomain,
		basedir:        args.Basedir,
	}
}

func (me *Repo) GetLatestRevisionUrl() string {
	return fmt.Sprintf(LatestRevisionUrlFormat, me.downloaddomain)
}

func (me *Repo) GetChangeLogUrl(args *global.Args) string {
	return fmt.Sprintf(ChangeLogUrlFormat,
		me.downloaddomain,
		args.LastRev,
		args.LastRev-args.FirstRev,
	)
}

func (me *Repo) GetLatestRevision() (rev int, sts status.Status) {
	for range only.Once {
		c := util.NewHttpClient()
		var body []byte
		body, sts = c.Download(me.GetLatestRevisionUrl())
		if is.Error(sts) {
			break
		}
		r, _ := regexp.Compile("\\[([0-9]+)]")
		match := r.FindStringSubmatch(string(body))
		if len(match) < 2 {
			sts = status.Fail(&status.Args{}).SetMessage("could not determine most recent revision")
		}
		var err error
		rev, err = strconv.Atoi(match[1])
		if err != nil {
			sts = status.Wrap(err, &status.Args{}).SetMessage("could not convert latest revision to int").
				SetSuccess(true) // no need to fail if we can't read this file
			break
		}
	}
	return rev, sts
}

func (me *Repo) PutLastSyncedRevision(lastrev int) {
	fp := me.GetLastSyncedRevisionFilepath()
	err := ioutil.WriteFile(fp, []byte(strconv.Itoa(lastrev)), os.ModePerm)
	if err != nil {
		fmt.Printf("Error writing last synced revision %d: %s", lastrev, err.Error())
	}
}

func (me *Repo) GetLastSyncedRevision() (rev int, sts status.Status) {
	for range only.Once {
		rev = 0
		fp := me.GetLastSyncedRevisionFilepath()
		if !util.FileExists(fp) {
			break
		}
		b, err := ioutil.ReadFile(fp)
		if err != err {
			sts = status.Wrap(err, &status.Args{}).SetMessage("could not read last revision file").
				SetSuccess(true) // no need to fail if we can't read this file
			break
		}
		rev, err = strconv.Atoi(string(b))
		if err != nil {
			sts = status.Wrap(err, &status.Args{}).SetMessage("could not convert last revision to int").
				SetSuccess(true) // no need to fail if we can't read this file
			break
		}
	}
	return rev, sts
}

func (me *Repo) GetLastSyncedRevisionFilepath() (fp string) {
	return filepath.FromSlash(fmt.Sprintf("%s/last-revision.txt", me.basedir))
}

func (me *Repo) GetListUrl() string {
	return fmt.Sprintf("http://%s/", me.listdomain)
}

func (me *Repo) GetSubdirList(args *global.Args) (sds global.Strings, sts status.Status) {
	for range only.Once {
		if args.FirstRev == 0 {
			sds, sts = me.GetFullSubdirList(args)
			break
		}
		sds, sts = me.GetPartialSubdirList(args)
		break
	}
	return sds, sts
}

func (me *Repo) GetFullSubdirList(args *global.Args) (sds global.Strings, sts status.Status) {
	for range only.Once {
		c := util.NewHttpClient()
		var body []byte
		body, sts = c.Download(me.GetListUrl())
		if is.Error(sts) {
			break
		}
		r, _ := regexp.Compile(`<li><a href="(.+?)/">.+?/</a></li>`)
		matches := r.FindAllStringSubmatch(string(body), -1)
		sds = make(global.Strings, len(matches))
		for i, m := range matches {
			sds[i] = m[1]
		}
		break
	}
	return sds, sts
}

func (me *Repo) GetChangeLogFilepath() string {
	return filepath.FromSlash(fmt.Sprintf("%s/%s/%s-changes.log",
		me.basedir,
		global.ChangelogSubdirectory,
		time.Now().Format(time.RFC3339),
	))
}

func (me *Repo) GetPartialSubdirList(args *global.Args) (sds global.Strings, sts status.Status) {
	for range only.Once {
		c := util.NewHttpClient()
		var body []byte
		body, sts = c.Download(me.GetChangeLogUrl(args))
		if is.Error(sts) {
			break
		}
		fp := me.GetChangeLogFilepath()
		err := ioutil.WriteFile(fp, body, os.ModePerm)
		if err != nil {
			sts = status.Wrap(err, &status.Args{}).SetMessage("cannot write to '%s': %s", fp, err.Error())
			break
		}
		r, _ := regexp.Compile(`\*\s+([^/]+).+?\((added|modified|deleted|moved|copied)\)`)
		matches := r.FindAllStringSubmatch(string(body), -1)
		idx := make(map[string]bool, 0)
		for _, m := range matches {
			idx[m[1]] = true
		}
		i := 0
		sds = make(global.Strings, len(idx))
		for sd := range idx {
			sds[i] = sd
			i++
		}
	}
	return sds, sts
}
