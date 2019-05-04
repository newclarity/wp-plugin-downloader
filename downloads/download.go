package downloads

import (
	"fmt"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/is"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"wp-plugin-downloader/global"
	"wp-plugin-downloader/only"
	"wp-plugin-downloader/svn"
	"wp-plugin-downloader/util"
)

func GetMissingPluginMarkerFilepath(plugin string, args *global.Args) string {
	return fmt.Sprintf("%s/%s/%s.missing",
		args.Basedir,
		global.MissingSubdirectory,
		plugin,
	)
}

func EnsureDirs(dir string, subdirs global.Strings) (sts status.Status) {
	for range only.Once {
		err := os.Mkdir(dir, 0777)
		if err != nil && !os.IsExist(err) {
			sts = status.Wrap(err, &status.Args{
				Message: fmt.Sprintf("unable to make directory '%s'", dir),
			})
			break
		}
		for _, sd := range subdirs {
			sts = EnsureDirs(
				filepath.FromSlash(fmt.Sprintf("%s/%s", dir, sd)),
				global.Strings{},
			)
			if is.Error(sts) {
				break
			}
		}
	}
	return sts
}

func Download(args *global.Args) (sts status.Status) {
	for range only.Once {
		sts = EnsureDirs(args.Basedir, global.Subdirectories)
		if is.Error(sts) {
			break
		}
		repo := svn.NewRepo(args)
		args.LastRev, sts = repo.GetLatestRevision()
		if is.Error(sts) {
			panic(sts.Message())
		}
		fmt.Printf("Most recent SVN revision: %d\n", args.LastRev)
		args.FirstRev, sts = repo.GetLastSyncedRevision()
		if is.Error(sts) {
			panic(sts.Message())
		}
		if args.FirstRev == 0 {
			fmt.Printf("You have not yet performed a successful sync. Settle in. This will take a while.\n")
		} else {
			fmt.Printf("Most recent SVN revision: %d\n", args.FirstRev)
		}
		var plugins global.Strings
		plugins, sts = repo.GetSubdirList(args)

		item := MakeDownloadableItem(args)
		c := util.NewHttpClient()

		start := time.Now()

		for _, p := range plugins {
			mfp := GetMissingPluginMarkerFilepath(p, args)
			if util.FileExists(mfp) {
				fmt.Printf("Skipping %s - previously found missing.\n", p)
				continue
			}
			item.SetPlugin(p)
			var b []byte
			fmt.Printf("Downloading '%s'...\n", p)
			b, sts = c.Download(item.DownloadUrl())
			if is.Error(sts) {
				if sts.HttpStatus() == http.StatusNotFound {
					_ = ioutil.WriteFile(mfp, []byte(""), os.ModePerm)
					fmt.Printf("Error downloading '%s': 404 Not Found\n", p)
					continue
				}
				fmt.Printf("Error downloading '%s': %s\n", p, sts.Message())
				continue
			}
			dlfp := item.DownloadFilepath()
			err := ioutil.WriteFile(dlfp, b, os.ModePerm)
			if err != nil {
				fmt.Printf("Error writing '%s': %s\n", dlfp, err.Error())
				continue
			}
			uzd := item.UnzipDir()
			err = os.RemoveAll(uzd)
			if err != nil {
				fmt.Printf("Error removing plugin directory '%s': %s\n", uzd, err.Error())
				continue
			}

			_, sts = util.Unzip(dlfp, item.Directory())
			if is.Error(sts) {
				fmt.Printf("Error unzipping '%s' to '%s: %s\n", dlfp, uzd, sts.Message())
				continue
			}

		}
		elapsed := time.Since(start)
		log.Printf("Downloading %d plugins took %s\n", len(plugins), elapsed.String())
		repo.PutLastSyncedRevision(args.LastRev)

	}
	return sts
}
