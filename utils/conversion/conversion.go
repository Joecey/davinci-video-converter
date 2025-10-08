package conversion

import (
	"os"
	"time"

	"github.com/schollz/progressbar/v3"
)

func Convert (originPath string, targetPath string)(error){
	files,_ := os.ReadDir(originPath)

	numFiles := len(files)
	bar := progressbar.Default(int64(numFiles))
	for i := 0; i < numFiles; i++ {
		bar.Add(1)
		time.Sleep(40 * time.Millisecond)
	}
	return nil
}	