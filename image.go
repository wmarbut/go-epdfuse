// Copyright 2017 Whit Marbut. All rights reserved.
// License information may be found in the LICENSE file.

package epdfuse

import (
	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
	"image"
)

type ScalePlan byte

const (
	SCALE_NO         ScalePlan = '0'
	SCALE_UP         ScalePlan = '1'
	SCALE_DOWN       ScalePlan = '2'
	SCALE_PLACE_ONLY ScalePlan = '3'
)

type Axis byte

const (
	AXIS_X Axis = '0'
	AXIS_Y Axis = '1'
)

func (epd *EpdFuse) scaleAndPlaceImage(img image.Image) image.Image {
	plan := epd.detectScalePlacePlan(img)

	if plan == SCALE_DOWN || plan == SCALE_UP {
		img = epd.scale(img)
		img = epd.placeImage(img)
	} else if plan == SCALE_PLACE_ONLY {
		img = epd.placeImage(img)
	}

	return img
}

func (epd *EpdFuse) detectScalePlacePlan(img image.Image) ScalePlan {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	if width > epd.Width || height > epd.Height {
		return SCALE_DOWN
	} else if width < epd.Width && height < epd.Height {
		return SCALE_UP
	} else if width < epd.Width || height < epd.Height {
		return SCALE_PLACE_ONLY
	} else {
		return SCALE_NO
	}
}

func (epd *EpdFuse) placeImage(img image.Image) image.Image {
	context := gg.NewContext(epd.Width, epd.Height)
	mW := int(epd.Width / 2)
	mH := int(epd.Height / 2)

	context.DrawImageAnchored(img, mW, mH, 0.5, 0.5)

	return context.Image()
}

func (epd *EpdFuse) scale(img image.Image) image.Image {
	cstrAxis := epd.constrainingAxis(img)
	sFactor := epd.scaleFactor(img, cstrAxis)
	nW := uint(float64(img.Bounds().Dx()) * sFactor)
	nH := uint(float64(img.Bounds().Dy()) * sFactor)
	img = resize.Resize(nW, nH, img, resize.Lanczos2)

	return img
}

func (epd *EpdFuse) constrainingAxis(img image.Image) Axis {
	origRatio := float64(epd.Width) / float64(epd.Height)
	ratio := float64(img.Bounds().Dx()) / float64(img.Bounds().Dy())
	if ratio > origRatio {
		return AXIS_X
	}
	return AXIS_Y
}

func (epd *EpdFuse) scaleFactor(img image.Image, axis Axis) float64 {
	if axis == AXIS_X {
		return float64(epd.Width) / float64(img.Bounds().Dx())
	}
	return float64(epd.Height) / float64(img.Bounds().Dy())
}
