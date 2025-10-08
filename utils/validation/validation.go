package validation

import (
	"os"
)

func ValidateMP4Folder (folderPath string)(bool){
	files,_ := os.ReadDir(folderPath)
	return len(files) > 0
}