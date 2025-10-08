package conversion

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/schollz/progressbar/v3"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func Convert (originPath string, targetPath string)(error){
	files,err := os.ReadDir(originPath)
	if err != nil {
        return err
    }

	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		err = os.MkdirAll(targetPath, 0755)
		if err != nil {
			return err
		}
	}

	numFiles := len(files)
	bar := progressbar.Default(int64(numFiles))

	for _, file := range files {
        if file.IsDir() {
            continue
        }

		if strings.ToLower(filepath.Ext(file.Name())) != ".mp4"{
			continue
		}

        inputFile := filepath.Join(originPath, file.Name())
        // Change extension to .mov for output
        outputFileName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())) + ".mov"
        outputFile := filepath.Join(targetPath, outputFileName)

        err := ffmpeg.Input(inputFile).
            Output(outputFile, ffmpeg.KwArgs{
                "c:v": "mjpeg",
                "q:v": "2",
                "c:a": "pcm_s16be",
                "q:a": "0",
                "f":   "mov",
				"loglevel": "quiet", 
            }).
            OverWriteOutput().
			WithOutput().
			Run()

        if err != nil {
            return err
        }

        bar.Add(1)
        time.Sleep(40 * time.Millisecond) // Keep for smooth bar update
    }
    return nil
}	