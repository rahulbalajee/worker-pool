package streamer

import (
	"fmt"

	"github.com/xfrr/goffmpeg/transcoder"
)

type Encoder interface {
	EncodeToMP4(v *Video, baseFileName string) error
}

type VideoEncoder struct{}

func (ve *VideoEncoder) EncodeToMP4(v *Video, baseFileName string) error {
	// Create a transcoder
	trans := new(transcoder.Transcoder)

	// Build the output path
	outputPath := fmt.Sprintf("%s/%s.mp4", v.OutputDir, baseFileName)

	// Initialise the transcoder
	err := trans.Initialize(v.InputFile, outputPath)
	if err != nil {
		return err
	}

	// Set the codec
	trans.MediaFile().SetVideoCodec("libx264")

	// Start transcoding process
	done := trans.Run(false)

	err = <-done
	if err != nil {
		return err
	}

	return nil
}
