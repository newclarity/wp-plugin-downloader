package util

import (
	"archive/zip"
	"github.com/gearboxworks/go-status"
	"io"
	"os"
	"path/filepath"
	"strings"
	"wp-plugin-downloader/global"
	"wp-plugin-downloader/only"
)

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func Close(r io.Closer, err error) {
	if err == nil {
		_ = r.Close()
	}
}

// Unzip will decompress a zip archive, moving all files and folders
// within the zip file (parameter 1) to an output directory (parameter 2).
func Unzip(zipfilepath string, outputDir string) (filepaths global.Strings, sts status.Status) {
	r, err := zip.OpenReader(zipfilepath)
	defer Close(r, err)
	for range only.Once {
		if err != nil {
			sts = status.Wrap(err, &status.Args{}).
				SetMessage("unable to open .ZIP file '%s': %s",
					zipfilepath,
					err.Error(),
				)
			break
		}

		for _, f := range r.File {

			// Store filename/path for returning and using later on
			fp := filepath.Join(outputDir, f.Name)

			// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
			if !strings.HasPrefix(fp, filepath.Clean(outputDir)+string(os.PathSeparator)) {
				sts = status.Fail(&status.Args{}).
					SetMessage("illegal file path '%s'", fp)
				break
			}

			filepaths = append(filepaths, fp)

			if f.FileInfo().IsDir() {
				// Make Folder
				err = os.MkdirAll(fp, os.ModePerm)
				if err != nil {
					sts = status.Wrap(err, &status.Args{}).
						SetMessage("unable to make directory '%s': %s",
							fp,
							err.Error(),
						)
					break
				}
				continue
			}

			err = os.MkdirAll(filepath.Dir(fp), os.ModePerm)
			if err != nil {
				sts = status.Wrap(err, &status.Args{}).
					SetMessage("unable to make directory '%s': %s",
						filepath.Dir(fp),
						err.Error(),
					)
				break
			}

			outFile, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				sts = status.Wrap(err, &status.Args{}).
					SetMessage("unable to make directory '%s': %s",
						fp,
						err.Error(),
					)
				break
			}

			rc, err := f.Open()
			if err != nil {
				sts = status.Wrap(err, &status.Args{}).
					SetMessage("unable to open file '%s': %s",
						f.Name,
						err.Error(),
					)
				break
			}

			_, err = io.Copy(outFile, rc)

			// Close the file without defer to close before next iteration of loop
			Close(outFile, nil)
			Close(rc, nil)

		}
	}
	return filepaths, sts
}
