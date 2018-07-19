// Copyright 2017 Whit Marbut. All rights reserved.
// License information may be found in the LICENSE file.

package epdfuse

import (
	"bytes"
	"github.com/wmarbut/goxbm"
	"image"
	"io"
	"os"
	"path"
)

const DISPLAY_PATH = "/BE/display"
const EPD_COMMAND_PATH = "/command"
const EPD_DEFAULT_PATH = "/dev/epd"
const EPD_DEFAULT_WIDTH = 200
const EPD_DEFAULT_HEIGHT = 96

type EpdCommand byte

const (
	// EPD Command to trigger a display update
	COMMAND_UPDATE EpdCommand = 'U'
	// EPD Command to clear display
	COMMAND_CLEAR EpdCommand = 'C'
	// EPD Command to trigger a partial display update
	COMMAND_PARTIAL EpdCommand = 'P'
)

// Configures EPD Fuse
type EpdFuse struct {
	EpdPath string
	Width   int
	Height  int
}

// Create a new EpdFuse to interact with your display
func NewEpdFuse() EpdFuse {
	return EpdFuse{
		EPD_DEFAULT_PATH,
		EPD_DEFAULT_WIDTH,
		EPD_DEFAULT_HEIGHT,
	}
}

// Create a new EpdFuse to interact with your display
func NewCustomEpdFuse(path string, width, height int) EpdFuse {
	return EpdFuse{
		path,
		width,
		height,
	}
}

// Write wrapped text to your PaPiRus display
func (epd *EpdFuse) WriteText(text string) error {
	img, err := epd.buildTextImage(text)
	if err != nil {
		return err
	}
	return epd.WriteImage(img)
}

// Write image to your PaPiRus display
func (epd *EpdFuse) WriteImage(img image.Image) error {
	img = epd.scaleAndPlaceImage(img)
	err := epd.writeDisplay(goxbm.ToRawXBMBytes(img))
	if err != nil {
		return err
	}
	return epd.Update()
}

// Write partial image to your PaPiRus display
func (epd *EpdFuse) WriteImagePartial(img image.Image) error {
	img = epd.scaleAndPlaceImage(img)
	err := epd.writeDisplay(goxbm.ToRawXBMBytes(img))
	if err != nil {
		return err
	}
	return epd.PartialUpdate()
}

// Force an update of the PaPiRus display
func (epd *EpdFuse) Update() error {
	return epd.update(COMMAND_UPDATE)
}

// Clear the PaPiRus display
func (epd *EpdFuse) Clear() error {
	return epd.update(COMMAND_CLEAR)
}

// Perform a partial update
func (epd *EpdFuse) PartialUpdate() error {
	return epd.update(COMMAND_PARTIAL)
}

func (epd *EpdFuse) update(command EpdCommand) error {
	cmdPath := path.Join(epd.EpdPath, EPD_COMMAND_PATH)
	f, err := os.OpenFile(cmdPath, os.O_WRONLY|os.O_APPEND, 0222)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte{byte(command)})
	if err != nil {
		return err
	}
	err = f.Sync()
	if err != nil {
		return err
	}
	return nil
}

func (epd *EpdFuse) writeDisplay(data []byte) error {
	displayPath := path.Join(epd.EpdPath, DISPLAY_PATH)
	f, err := os.OpenFile(displayPath, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	dataReader := bytes.NewReader(data)
	_, err = io.Copy(f, dataReader)
	if err != nil {
		return err
	}
	err = f.Sync()
	if err != nil {
		return err
	}
	return nil
}
