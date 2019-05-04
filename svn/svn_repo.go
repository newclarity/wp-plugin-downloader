package svn

import (
	"fmt"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/is"
	"github.com/newclarity/wp-plugin-downloader/only"
	"github.com/newclarity/wp-plugin-downloader/util"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strconv"
)

const UrlFormat = "http://%s/log/?format=changelog&stop_rev=HEAD"

type Repo struct {
	domain string
	basedir string
}

func NewSvn(domain string, basedir string) *Repo {
	return &Repo{
		domain: domain,
		basedir: basedir,
	}
}

func (me *Repo) GetLatestRevision() (rev int,sts status.Status) {
	for range only.Once {
		url := fmt.Sprintf(UrlFormat,me.domain)
		c := NewHttpClient()
		sts = c.GET(url)
		if is.Error(sts) {
			break
		}
		var body []byte
		body,sts = c.GetBody()
		if is.Error(sts) {
			break
		}
		r,_ := regexp.Compile("\\[([0-9]+)]")
		match := r.FindStringSubmatch(string(body))
		if len(match)<2 {
			sts = status.Fail(&status.Args{}).SetMessage("could not determine most recent revision")
		}
		var err error
		rev,err = strconv.Atoi(match[1])
		if err != nil {
			sts = status.Wrap(err,&status.Args{}).SetMessage("could not convert latest revision to int").
				SetSuccess(true) // no need to fail if we can't read this file
			break
		}
	}
	return rev,sts
}

func (me *Repo) GetLastSyncedRevision() (rev int,sts status.Status) {
	for range only.Once {
		rev = 0
		fp := me.GetLastSyncedRevisionFilepath()
		if !util.FileExists(fp) {
			break
		}
		b,err := ioutil.ReadFile(fp)
		if err != err {
			sts = status.Wrap(err,&status.Args{}).SetMessage("could not read last revision file").
				SetSuccess(true) // no need to fail if we can't read this file
			break
		}
		rev,err = strconv.Atoi(string(b))
		if err != nil {
			sts = status.Wrap(err,&status.Args{}).SetMessage("could not convert last revision to int").
				SetSuccess(true) // no need to fail if we can't read this file
			break
		}
	}
	return rev,sts
}

func (me *Repo) GetLastSyncedRevisionFilepath() (fp string) {
	return filepath.FromSlash(fmt.Sprintf("%s/.last-revision", me.basedir))
}

