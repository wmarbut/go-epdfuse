// Copyright 2017 Whit Marbut. All rights reserved.
// License information may be found in the LICENSE file.

package epdfuse

import (
	"github.com/fogleman/gg"
	"image"
)

func (epd *EpdFuse) buildTextImage(text string) (image.Image, error) {

	context := gg.NewContext(epd.Width, epd.Height)
	context.SetRGB(1, 1, 1)
	context.Clear()
	context.SetRGB(0, 0, 0)
	context.DrawStringWrapped(text, 0, 0, 0, 0, float64(epd.Width), 1.4, gg.AlignLeft)
	return context.Image(), nil

}
