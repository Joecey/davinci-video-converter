package main

// run the main function
import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/Joecey/davinci-video-converter/utils/conversion"

	"github.com/Joecey/davinci-video-converter/utils/validation"
	"github.com/charmbracelet/huh"
)

// Define variables for answers
var (
    targetFolder	string 
    originFolder	string 
    mp4ToMov		bool
)

func main(){
	form := huh.NewForm(
		huh.NewGroup(huh.NewNote().
		Title("DaVinci Video Converter for Linux").
		Description("Convert your videos for DaVinci Studio Free on Linux").
		Next(true).
		NextLabel("Begin"),
	),
	
	// gather details about conversion
		huh.NewGroup(
			huh.NewInput().
				Title("Where are your '.mp4' videos located"). 
				Value(&originFolder).
				// Validating fields is easy. The form will mark erroneous fields
				// and display error messages accordingly.
				Validate(func(str string) error {
					pathHasVideos := validation.ValidateMP4Folder(str)

					if !pathHasVideos{
						return errors.New("sorry, this path has no .mp4 files")
					}
					return nil
				}),
		
	
			huh.NewInput().
				Title("Where are your '.mov' videos located"). 
				Value(&targetFolder).
				// Validating fields is easy. The form will mark erroneous fields
				// and display error messages accordingly.
				Validate(func(str string) error {
					if strings.TrimSpace(str) == strings.TrimSpace(originFolder) {
						return errors.New("please select a different folder for your new .mov files")
					}
					return nil
				}),
		
			huh.NewConfirm().
				Title("Are you converting from .mp4 to .mov").
				Value(&mp4ToMov),
		),
	)


	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	if len(targetFolder) == 0  {
		fmt.Println("targetFolder was not defined")
	} else if len(originFolder) == 0 {
		fmt.Println("originFolder was not defined")
	} else {
		fmt.Println("Converting your videos to .mov 📷")
		conversion.Convert(strings.TrimSpace(originFolder), strings.TrimSpace(targetFolder))
	}



}